# LevelDB_Write

---

1、架构

![](https://youjiali1995.github.io/assets/images/leveldb/architecture.png)

    数据结构LSM：

    LSM-Tree：leveldb 是以 LSM-tree 为模型实现的。LSM-tree 将随机写转换为顺序写从而获得更高的写性能，
    大致思路如下：
        数据分为2部分存储，共同维护一个有序的空间：
        内存中维护最新的写数据 memtable，所有的写操作都在内存中进行，同时顺序写 WAL。
        磁盘上维护较老的数据 sstable。
        读操作先读内存中的数据，然后读磁盘上的数据。
        在合适的时机进行 compaction:
        当内存中的数据大小超过阈值时，刷新到磁盘上。
        在合适的时机清理、合并磁盘上的数据。

    考虑最基本的情况，磁盘上维护一个大的 sstable，当 memtable 大小超过阈值时，就与磁盘上的 sstable 合并。
    可以是 in-place 的修改，可能伴随着很多随机 I/O，或者顺序写生成新的 sstable， 但可能要将整个 sstable
    进行重写，会带来极大的写放大。现在只关注后面一种方式，一步步进行优化:
    
    不是每次 memtable 大小超过阈值就与 sstable 合并，而是将 memtable 转储为 sstable，当到了一定数量后批量合并，
    从而减少合并的次数。同时需要控制 sstable 的数量，因为读可能需要读每个 sstable，数量太多会降低读的性能。但每次
    合并仍有可能将整个 sstable 重写，为了降低写放大，将大的sstable划分为多个小的、不重叠的sstable，这样sstable
    就分为2层：level-0 是 memtable 直接转换而来，文件之间有重叠；level-1 由 level-0 合并而来，文件之间无重叠。
    当 level-0 的文件个数达到上限时，只要挑选 level-1 中和 level-0 文件有重叠的合并即可，读 level-1 也只要读一
    个文件。 但仍有可能 level-1 中所有的 sstable 都需要合并，所以又要限制 level-1的大小。当level-1的sstable
    达到上限时，使用类似的方法 compaction 为 level-2 的 sstable。level-2 达到上限，再compaction为level-3……
    leveldb 就是类似上面的思路分层管理 sstable，所以叫 leveldb。现在看一下它的实现：
    
    memtable 由 skiplist 实现，只有追加操作，每个操作先顺序写到 log 中。
    当 memtable 大小超过阈值时(默认为 4MB)，变为 immutable memtable，在后台线程 compaction 为 level-0 的
    sstable。sstable 的大小固定，默认为 2MB。
    共分为 7 层:
    level-0 由 memtable 直接转化而来，文件之间有重叠。当 level-0 的 sstable 数量到达阈值时，会将 level-0 中
    相互重叠的sstable和level-1中重叠的sstable合并为新的 level-1 的 sstable。
    其余 level 由低层 compaction而来，除最高层外，每层有大小限制，level-(N+1) 的大小限制是 level-N 的 10 倍，
    其中level-1 为 10MB。相同level 的文件之间无重叠。当 level-N 的 sstable 大小到达阈值时，会挑选一个文件(可
    能不止一个)和 level-(N+1)中有重叠的 sstable 合并为新的 level-(N+1) 的 sstable。
    
    所有的写操作，包括 compacion 都是顺序写。
    读的顺序如下，只要读到对应的 key 就会停止:
        1、memtable
        2、immutable memtable
        3、level-0 中文件有重叠，从新到旧读取
        4、其余 level 文件不重叠，每层最多只要读 1 个文件，按照层的顺序从低到高读取

2、并发控制：

    leveldb 使用 MVCC 实现并发控制，不支持事务，支持单个写操作和多个读操作同时执行:

    每个成功的写操作会更新 db 内部维护的顺序递增的 sequence，sequence 会追加到 key 后一同保存。
    读操作会使用当前最新的 sequence(或者使用传入的 snapshot)，只会读到小于等于自己的最大的 sequence 数据。
    因为 compaction 会改变磁盘上文件的组织，为了不影响正在进行的操作，leveldb 使用 VersionSet 维护不同版
    本的 sstable 组织，每当有文件删除或增加时，就会创建新的 Version 插入到 VersionSet。只有当一个 sstable 
    不再被任意 Version 使用时才会进行删除。

3、数据恢复：

    数据恢复
    leveldb 中每个 db 对应一个目录，目录名即 dbname。目录下有如下几种文件:
    
    LOG: 记录操作日志。
    *.log: 记录 memtable 的 WAL。
    *.ldb(*.sst): sstable 文件。
    MANIFEST-*: 记录元信息，如 sstable 的组织结构，每个 sstable 的 key range 等。
    CURRENT: 记录当前的 MANIFEST 文件名。
    LOCK: leveldb 不支持多个进程同时打开相同 db，会加文件锁。
    数据恢复主要就是通过 CURRENT 文件找到当前的 MANIFEST 文件读取之后进行清理和恢复操作。

---

## Write

1、流程：

    leveldb 的写操作很简单，先写 WAL，然后写到 memtable:
    每个写操作都有 sequence，会一同记录在 key 中，通过 sequence 实现 MVCC。
    删除 key 时是插入一个 tombstone。
    所有的操作顺序追加到 log 中。

2、写操作执行：

    leveldb不支持多个写操作同时执行，写操作会保存在一个deque队列中，只有队首的操作可以执行

```
Status DBImpl::Write(const WriteOptions& options, WriteBatch* my_batch) {
  Writer w(&mutex_);
  w.batch = my_batch;
  w.sync = options.sync;
  w.done = false;

  MutexLock l(&mutex_);
  writers_.push_back(&w);   // 出列
  while (!w.done && &w != writers_.front()) {
    w.cv.Wait();
  }
  if (w.done) {
    return w.status;
  }
```

    leveldb 会将写操作合并执行：队首的线程将 deque 中正在等待执行的写操作 batch 到当前线程中一起执行，
    执行后设置 done 通知其他线程完成。batch 能够提高写 WAL 的效率。

3、WAL

    leveldb 不是直接将 kv 插入到 memtable 中，而是先生成 WAL，然后解析 WAL 插入，目的是为了减少重复代码，
    复用了重启时用 WAL 恢复 memtable 的代码。 WAL 的格式如下：

![](https://youjiali1995.github.io/assets/images/leveldb/wal.png)

    WriteBatch 记录了当前 batch 的起始 Sequence，会追加在 key 后用于实现 MVCC，写操作合并就是修改 
    WriteBatch 的 Count 并追加 Record。

4、Log

    leveldb 使用 log::Writer 写 log，包括 WAL 和 MANIFEST。log::Writer 会在具体内容之外增加如 
    checksum 之类的保护，用于检测 corruption，它的格式如下：

![](https://youjiali1995.github.io/assets/images/leveldb/log.png)

    log::Writer 将文件划分为固定大小的 Block(32 KB)，Record 不能跨越 Block，每个 Block 开始一定
    是一个新的 Record，这么做的目的是当一个 Block 的数据发生 corruption 时不会影响到其他的 Block。
    显然，单条 log 的大小可能会超过 Block size，通过 Record 中的 Type 区分：

    kFullType：当前 Block 中是完整的一条 log。
    kFirstType、kMiddleType、kLastType：组装成一条完整的 log。

5、Memtable（skipList）

    memtable 是 leveldb 里实现最复杂的部分。任意有序的结构都可以实现 memtable，leveldb 使用 skiplist 实现，
    因为结构简单， 更容易实现 lock-free 的支持一写多读的。

    InternalKey:为了实现并发控制，每个写入的 key 会携带着 sequence，并且不能够修改之前写入的数据，而是插入新的数据。
    InternalKey 的格式如下：

![](https://youjiali1995.github.io/assets/images/leveldb/memtable_entry.png)

    5.1 InternalKey 的比较按照：
    UserKey 升序：使用用户传入的 comparator 比较，默认是 memcmp。
    Sequence 降序。
    ValueType 降序：每个写操作都有不同的 sequence，所以 ValueType 用不到。

    查找的时候会根据传入的 snapshot 或者最新的 sequence 构造出 InternalKey 查询，只要查找到第一个大于等于自己的即可

    5.2 skipList Node:
    leveldb 中 skiplist node 结构如下所示，只有 key 用于存储数据，所以 memtable 会构造如上所示的 Add/Delete 结构
    插入到 skiplist 中，通过 InternalKey 即可比较出大小， 后面追加的 value 不会影响先后顺序。

```
template<typename Key, class Comparator>
struct SkipList<Key,Comparator>::Node {
  Key const key;
 private:
  port::AtomicPointer next_[1];
};
```

    leveldb 在插入和读取 memtable 的时候是不加锁的，全依赖 skiplist 实现并发控制。leveldb 实现的是支持一写多读的、
    lock-free 的 skiplist。 想要知道原理必须先理解 atomic pointer 的实现。

    5.3 Atomic pointer:（leveldb自己实现）

```
#if defined(ARCH_CPU_X86_FAMILY) && defined(__GNUC__)
inline void MemoryBarrier() {
  __asm__ __volatile__("" : : : "memory");
}

class AtomicPointer {
 private:
  void* rep_;
 public:
  AtomicPointer() { }
  explicit AtomicPointer(void* p) : rep_(p) {}
  inline void* NoBarrier_Load() const { return rep_; }
  inline void NoBarrier_Store(void* v) { rep_ = v; }
  inline void* Acquire_Load() const {
    void* result = rep_; // read-acquire
    MemoryBarrier();
    return result;
  }
  inline void Release_Store(void* v) {
    MemoryBarrier();
    rep_ = v; // write-release
  }
};
```

    原子性: store/load 是原子的，不会被其他线程看到 half-write，这里的原子性不包括原子变量的可见性。
    可见性: Release_Store/Acquire_Load 提供了 release/acquire 语义。Release_Store 立即对接下
    来的 Acquire_Load 可见(原子变量的可见性)；Release_Store 之前的写命令对 Acquire_Load 之后的
    读命令可见。

    额外的补充：
        多核体系下，每个 CPU 有独占的 cache，使用 MESI 协议来保证 cache coherence。为了降低同步对
    性能的影响，每个 CPU 有 store buffer 和 invalidate queue 来缓冲相应的同步消息，对同步消息处理
    的时机和顺序是不确定的。
        CPU 在保证单线程程序的正确运行的前提下，为了提高性能会对命令做乱序处理。多线程情况下，因为store
    buffer 和 invalidate queue 的存在，其他线程的修改不会立即对另外的线程可见；受到 CPU 乱序和对同步
    消息处理的顺序影响，可见的顺序也不能保证。

    memory barrier 用于保证多核之间操作的执行顺序，包含4类:
        LoadLoad-barrier: memory barrier 前(后)的 load 不会乱序到 memory barrier 后(前)。
        LoadStore-barrier: memory barrier 前的 load(后的 store)不会乱序到 memory barrier 后(前)。
        StoreStore-barrier: memory barrier 前(后) 的 store 不会乱序到 memory barrier 后(前)。
        StoreLoad-barrier: memory barrier 前的 store(后的 load)不会乱序到 memory barrier 后(前)。

    比 memory barrier 更高一层的语义是 acquire/release，同样也要成对使用:
        acquire: 保证了在 acquire 之后的读写操作不会乱序到 acquire 前。
        release: 保证了在 release 之前的读写操作不会乱序到 release 后。
    acquire/release 同样只保证可见性顺序，对可见的时机不能保证，需要有方法判断 acquire 或者 release 完成。
    一般将 load(RMW)/store(RMW) 操作和 memory barrier 搭配使用构成 acquire/release 语义， 对 load/store
    操作和 memory barrier 的顺序有要求：
    
    acquire: load/RMW 在前，memory barrier 在后，load/RMW 操作构成 read-acquire，防止了 LoadLoad/LoadStore 乱序。
    release: store/RMW 在后，memory barrier 在前，store/RMW 操作构成 write-release，防止了 LoadStore/StoreStore 乱序。
    这种顺序很符合逻辑，只有这样才能够确保 store 操作完成，之前的 release 操作一定完成；load 操作不会在 memory barrier 之后执行。
    构成 acquire/release 语义的读写操作要操作相同的变量，同时需要是原子操作。
    除了 CPU 乱序，编译器也会在保证单线程程序正确性的前提下，对命令乱序处理，同样有 compiler barrier 保证编译的命令顺序。

    x86/64 是 strong memory model，load/store 自带 acquire/release 语义，只会出现操作不同地址的 StoreLoad 乱序。
    对于 naturally aligned native types 且大小不超过 memory bus 的变量读写操作是原子的 ，所以只要防止编译器乱序即可。
    所以 leveldb 的 AtomicPointer 在 gcc on x86 的实现中只使用了 compiler barrier。

    5.4 Lock-free

    skiplist 只需要支持如下场景的线程安全即可：
        一写多读：写操作不影响正在进行的读操作。
        写之后读：写操作之后的读要立刻读到最新的写入。
        写之后写：写操作之后继续写入要保证线程安全。

    skipList的结构：

```
template<typename Key, class Comparator>
class SkipList {
 private:
  enum { kMaxHeight = 12 };

  Comparator const compare_;
  Arena* const arena_;

  Node* const head_;

  port::AtomicPointer max_height_;

  Random rnd_;
};

template<typename Key, class Comparator>
struct SkipList<Key,Comparator>::Node {
  explicit Node(const Key& k) : key(k) { }

  Key const key;

  Node* Next(int n) {
    return reinterpret_cast<Node*>(next_[n].Acquire_Load());
  }
  void SetNext(int n, Node* x) {
    next_[n].Release_Store(x);
  }

  Node* NoBarrier_Next(int n) {
    return reinterpret_cast<Node*>(next_[n].NoBarrier_Load());
  }
  void NoBarrier_SetNext(int n, Node* x) {
    next_[n].NoBarrier_Store(x);
  }

 private:
  port::AtomicPointer next_[1];
};
```

    leveldb 的 skiplist 的结构没什么特别的，查找和插入算法和不支持并发的区别也不大:
    查找：从最高层开始向右、向下查找;
    插入：查找算法会保存查找过程中每层中向下的结点，也就是被插入节点每层的 prev 结点，
    最后将被插入结点插入在 prev[] 后。

    需要注意的是使用 AtomicPointer 的地方:

    Skiplist::AtomicPointer max_height_：
    插入：使用 NoBarrier_Store 设置。
    查找：使用 NoBarrier_Load 读取。
    Node::AtomicPointer next_[1]：
    插入：设置被插入结点的 next_[] 时用的是 NoBarrier_Store，设置 prev 结点的 next_[]时使用 Release_Store。
    查找：使用 AcquireLoad 读取 next_[]。
    前面的 AtomicPointer 分析可知：成对 Acquire_Load/Release_Store 能够保证在不同线程间立刻可见，且 Release_Store
    之前的修改对 Acquire_Load 之后立即可见。因为查找是按照从高层向低层、从小 到大的顺序遍历，而插入的时候是按照从低层到高层
    、用 Release_Store 设置 prev 结点的 next_[]，确保了当观察到新插入的结点时，后续的遍历一定是个完好的 skiplist。

![](https://youjiali1995.github.io/assets/images/leveldb/skiplist.png)

    查找操作只会有下面2种情形观察不到新插入的结点。
        在某一层观察到新插入的结点，且更低层也会观察到，也就是完好的 skiplist。
        当读写同时发生时，上述两种情况都有可能发生，但都不会影响正确的结果，因为不会查找正在插入的 key(Sequence 只有写操作
    完成才会更新)。 需要注意一点，max_height_ 只保证了原子性，没有保证对 max_height_ 可见性，也没有保证对 next_[] 的可
    见性，但都不会影响读的结果。

    假设插入增大了 max_height_:
        读操作观察到了 max_height_ 的更新，对应上面两种情况分别是:
            新增的 level 的 head_ 都指向 NULL，leveldb 保证了 skiplist 中 NULL 是最大的，所以会立刻向下层查找。
            在某一层观察到了新插入的 key，继续遍历。
        读操作未观察到 max_height_ 的更新，直接从低层开始遍历不影响 skiplist 的查找。

    而当写之后再读刚写入的 key 时，因为写已经完成，一定会观察到新插入的 key。写之后写我个人觉得有点问题，写操作使用了
    查找操作来获取需要设置的 prev 结点，但是遍历的时候不能保证获取到最新的 max_height_， 所以设置 prev[] 时高层会
    有问题。但是注释写到插入操作时有外部同步。

    临界区主要有这几个作用：
        保证只有一个线程能进入，临界区内的操作是原子的；临界区内的修改是可见的。
    认为在临界区内的修改才是可见的，这是错误的。可见性是由 acquire/release 保证的，lock 是 read-acquire，
    unlock 是 write-release，unlock 保证了之前的修改一定对 lock 之后可见，包括不在临界区内的。 这也就是 leveldb 
    的外部同步：leveldb 在写 WAL 和 memtable 时是不持有锁的，但是之前抢占队首和之后通知其他线程时是有 lock/unlcok 
    操作的，读的时候也要 lock 获取 Sequence, 在这里保证了可见性。