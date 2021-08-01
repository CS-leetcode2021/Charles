# BoltdbDB
---

1、简介

    boltdb 有如下性质：
        K/V 型存储，使用B+树索引（改进）。
        支持namespace，每对K/V存放在一个Bucket下，不同Bucket可以有相同的key，支持嵌套的Bucket。
        支持事务(ACID)，使用MVCC和COW，允许多个读事务和一个写事务并发执行，但是读事务有可能会阻塞写事务，适合读多写少的场景

2、存储：

    每个 db 对应一个文件，文件按照 page size(一般为 4096 Bytes) 划分为 page:
        前2个 page 保存 metadata；
        特殊的 page 保存 freelist，存放空闲 page 的 id；
        剩下的 page 构成一个 B+ 树结构。

    每个 Bucket 是一个完整的 B+ 树；

    B+ 树的每个结点对应一个或多个连续的 page；
    
    因为内存比磁盘小，一般会实现 page cache 缓存部分 page，比如使用 LRU 算法。boltdb 没有实现，而是使用 mmap() 创建共享、
    只读的文件映射并调用 madvise(MADV_RANDOM)，由操作系统 管理 page cache；
    
    没有 WAL，事务中所有操作都在内存中进行，只有 commit 时才会写到磁盘；
    
    commit 时会将 dirty page 写入新的 page，从而保证同时读的事务不受到影响。

3、B+Tree实现：

![](https://youjiali1995.github.io/assets/images/boltdb/B+tree.png)

    3.1 每个节点上的key和value的个数是相同的，而一般的val会比key多一个；
        Branch：每对 key/val 指向一个子节点，key 是子节点的起始 range，val 存放子节点的 page id；
        Leaf：每对 key/val 存放数据，没有指针指向 sibiling node；通常的 B+ 树多出的一个指针会指向 sibiling node。

    3.2 boltdb 中有 3 个结构和 B+ 树密切相关：
        page: 大小一般为 4096 bytes，对应文件里的每个 page，读写文件都是以 page 为单位。
        node: B+ 树的单个结点，访问结点时首先将 page 的内容转换为内存中的 node，每个 node 对应一个或多个连续的 page。
        Bucket: 每个 Bucket 都是一个完整的 B+ 树，所有操作都是针对 Bucket。
    
    3.3 一个典型的查找过程如下：
        首先找到 Bucket 的根节点，也就是 B+ 树的根节点的 page id；
        读取对应的 page，转化为内存中的 node；
        若是 branch node，则根据 key 查找合适的子节点的 page id;
        重复2、3直到找到 leaf node，返回 node 中对应的 val。


4、Page的实现：

```
type page struct {
	id    pgid      // page id
	flags uint16    // 区分不同类型的 page
	count    uint16 // page data 中的数据个数
	overflow uint32 // 若单个 page 大小不够，会分配多个 page
	ptr uintptr     // 存放 page data 的起始地址
}
```

    4.1 ptr 是保存数据的起始地址，不同类型page保存的数据格式也不同，共有4种page, 通过flags区分:
        meta page: 存放db的meta data。
        freelist page: 存放db的空闲page。
        branch page: 存放branch node的数据。
        leaf page: 存放leaf node的数据。
    
        磁盘上的格式和结构体相同，将数据写入文件会首先分配一个 page，然后设置 page header 和 page data:

![](https://youjiali1995.github.io/assets/images/boltdb/page.png)

        boltdb 直接将 page 结构体的二进制格式写入文件，避免了序列化和反序列化的开销。
        写的时候直接将需要写入的结构体转换为 byte 数组写入文件;读的时候直接将 byte 数组转换为对应结构体即可。
    
    4.2 page cache
        boltdb 没有实现 page cache，而是调用 mmap() 将整个文件映射进来，并调用 madvise(MADV_RANDOM) 
        由操作系统管理 page cache，后续对磁盘上文件的所有读操作直接读取。

5、Node的实现:

    5.1 Leaf Node:

![](https://youjiali1995.github.io/assets/images/boltdb/leafNode.png)

    node 的数据存放在 page.ptr 的位置:
        首先是所有的 leafPageElement:
```
// leafPageElement represents a node on a leaf page.
type leafPageElement struct {
        flags uint32 // 通过 flags 区分 subbucket 和普通 value
        pos   uint32 // key 距离 leafPageElement 的位移
    ksize uint32
    vsize uint32
}
```
        然后是所有的 key 和 value。
        通过 pos、ksize 和 vsize 获取对应的 key/value 的地址：
            &leafPageElement + pos == &key。
            &leafPageElement + pos + ksize == &val。

    5.2 Branch Node:

![](https://youjiali1995.github.io/assets/images/boltdb/branchNode.png)

    和leaf node的区别是：branch node的value是子节点的 page id，存放在branchPageElement里，
    而key的存储相同都是通过pos得到。

```
// branchPageElement represents a node on a branch page.
type branchPageElement struct {
	pos   uint32
	ksize uint32
	pgid  pgid
}
```
    将 item header 和 item 分开存储减少了查找的时间:可以进行二分查找，反之只能按照顺序查找。

    访问结点时，首先要将 page 反序列化后得到得到 node 结构体，代表 B+ 树中一个结点:

```
// node represents an in-memory, deserialized page.
type node struct {
	bucket     *Bucket
	isLeaf     bool // 区分 branch node 和 leaf node
	unbalanced bool
	spilled    bool
	key        []byte // 该 node 的起始 key
	pgid       pgid
	parent     *node
	children nodes
	inodes   inodes // 存放 node 的数据
}
```
    inodes 保存了该 node 的 K/V 数据:
```
// inode represents an internal node inside of a node.
// It can be used to point to elements in a page or point
// to an element which hasn't been added to a page yet.
type inode struct {
	flags uint32 // 用于 leaf node，区分是正常 value 还是 subbucket
	pgid  pgid // 用于 branch node, 子节点的 page id
	key   []byte
	value []byte
}
```

6、Bucket实现：

    操作都是对 Bucket 进行操作，Bucket 是一个 namespace，相当于一张表，一个 Bucket 代表一个完整的 B+ 树，
    结构如下：

```
// Bucket represents a collection of key/value pairs inside the database.
type Bucket struct {
	*bucket
	tx      *Tx                     // the associated transaction
	buckets map[string]*Bucket      // subbucket cache
	page *page                      // inline page reference
	rootNode *node                  // materialized node for the root page.
	nodes map[pgid]*node            // node cache
	// Sets the threshold for filling nodes when they split. By default,
	// the bucket will fill to 50% but it can be useful to increase this
	// amount if you know that your write workloads are mostly append-only.
	//
	// This is non-persisted across transactions so it must be set in every Tx.
	FillPercent float64
}

// bucket represents the on-file representation of a bucket.
// This is stored as the "value" of a bucket key. If the bucket is small enough,
// then its root page can be stored inline in the "value", after the bucket
// header. In the case of inline buckets, the "root" will be 0.
type bucket struct {
	root pgid // page id of the bucket's root-level page
	sequence uint64 // monotonically incrementing, used by NextSequence()
}
```

    6.1 subBucket:
        boltdb 支持嵌套的 Bucket，对于父 Bucket 而言，subbucket 只是特殊的 value 而已，设置
        leafPageElement.flags = bucketLeafFlag 标记，而 subbucket 本身是一个完整的 B+ 树:

![](https://youjiali1995.github.io/assets/images/boltdb/subbucketTree.png)

    subbucket 在父 bucket 的 leaf node 中存储格式如下:
        root: subbucket 的 root 结点的 page id。
        sequence: boltdb 支持每 bucket 的自增 id，即这里的 sequence。

```
type bucket struct {
  root pgid // page id of the bucket's root-level page
  sequence uint64 // monotonically incrementing, used by NextSequence()
}
```

![](https://youjiali1995.github.io/assets/images/boltdb/subbucketLeafNode.png)

    为了实现的统一，db 有一个 root Bucket，也就是整个文件构成的 B+ 树的 root 结点对应的 
    Bucket，之后创建的 Bucket 都是 root Bucket 的 subbucket。

    6.2 inline bucket：
        从上面看到，父 bucket 中只保存了 subbucket 的 bucket header，每个 subbucket 
        都会占据至少一个 page，若 subbucket 中的数据很少，会造成磁盘空间的浪费。inline 
        bucket 是将小的 subbucket 的 值完整放在父bucket的 leaf node上，从而减少 page
        的个数。inline bucket 的限制如下：

        没有 subbucket；
        整个 Bucket 的大小不能超过 page size/4。

    inline bucket 在父 bucket 的存放的格式如下，可以看出来是把完整的 page 格式追加到了 bucket header 后面，很好的复用了代码：

![](https://youjiali1995.github.io/assets/images/boltdb/inlineBucket.png)

7、事务：

    boltdb 支持完整的事务特性(ACID)，使用 MVCC 并发控制，允许多个读事务和一个写事务并发执行，但是读事务有可能会阻塞写事务。
    它的特点如下：
        Durability: 写事务提交时，会为该事务修改的数据(dirty page)分配新的 page，写入文件。
        Atomicity: 未提交的写事务操作都在内存中进行；提交的写事务会按照 B+ 树数据、freelist、metadata 的顺序写入文件，只有 
        metadata 写入成功，整个事务才算完成，只写入前两个数据对数据库无影响。
        Isolation: 每个读事务开始时会获取一个版本号，读事务涉及到的 page 不会被写事务覆盖；提交的写事务会更新数据库的版本号。

    写事务的流程如下：
        根据 db 初始化事务：拷贝一份 metadata，初始化 root bucket，自增 txid；
        从 root bucket 开始，遍历 B+ 树进行操作，所有的修改在内存中进行；
        提交写事务:
        平衡 B+ 树，在分裂的时候会给每个修改过的 node 分配新的 page；
        给 freelist 分配新的 page；
        将 B+ 树数据和 freelist 数据写入文件；
        将 metadata 写入文件。
    
    读事务的流程如下：
        根据 db 初始化事务：拷贝一份 metadata，初始化 root bucket；
        将当前读事务添加到 db.txs 中；
        从 root bucket 开始，遍历 B+ 树进行查找；
        结束时，将自身移出 db.txs。

    Durability：
        数据库的前两个 page 保存 metadata，从而保证重启时恢复 db 的信息：

```
type meta struct {
	magic    uint32
	version  uint32
	pageSize uint32
	flags    uint32
	root     bucket     // root bucket 的 page id
	freelist pgid       // freelist 的 page id
	pgid     pgid       // 文件的 page 个数
	txid     txid       // 当前 db 提交的最大的写事务 id
	checksum uint64
}
```
    
    文件格式如下：

![](https://youjiali1995.github.io/assets/images/boltdb/metaPage.png)

    在写事务 commit 时，会为脏 node 分配新的 page，同时将之前使用的 page 释放。
    freelist 中维护了当前文件中的空闲 page id，分配时会从 freelist.ids 中寻找合适的 page:

```
type freelist struct {
	ids     []pgid          // all free and available free page ids.
	pending map[txid][]pgid // mapping of soon-to-be free page ids by tx.
	cache   map[pgid]bool   // fast lookup of all free and pending page ids.
}
```
    freelist page的文件格式如下：

![](https://youjiali1995.github.io/assets/images/boltdb/freelistPage.png)

    这里的分配 page 不是真的分配文件中某一 page 来写，而是分配了一个 buffer 和起始的 page id，
    首先将 node 的信息写入这个 buffer，之后统一的写入 page id 对应的文件位置:

```go
// allocate returns a contiguous block of memory starting at a given page.
func (db *DB) allocate(count int) (*page, error) {
	// Allocate a temporary buffer for the page.
	var buf []byte
	if count == 1 {
		buf = db.pagePool.Get().([]byte) // db.pagePool 是 sync.pool，缓存了大小为 page size 的 buffer
	} else {
		buf = make([]byte, count*db.pageSize)
	}
	p := (*page)(unsafe.Pointer(&buf[0]))
	p.overflow = uint32(count - 1)

	// Use pages from the freelist if they are available.
	if p.id = db.freelist.allocate(count); p.id != 0 { // db.freelist.allocate() 查找合适的连续的 page，返回首 page id
		return p, nil
	}

	// Resize mmap() if we're at the end.
	p.id = db.rwtx.meta.pgid
	var minsz = int((p.id+pgid(count))+1) * db.pageSize
	if minsz >= db.datasz {
		if err := db.mmap(minsz); err != nil {
			return nil, fmt.Errorf("mmap allocate error: %s", err)
		}
	}

	// Move the page id high water mark.
	db.rwtx.meta.pgid += pgid(count)

	return p, nil
}
```

    Atomicity:事务要求完全执行或完全不执行:
        若事务未提交时出错，因为 boltdb 的操作都是在内存中进行，不会对数据库造成影响。
        若是在 commit 的过程中出错，如写入文件失败或机器崩溃，boltdb 写入文件的顺序也保证了不会造成影响:
        先写入 B+ 树数据和 freelist 数据；
        后写入 metadata。
        因为 db 的信息如 root bucket 的位置、freelist 的位置等都保存在 metadata 中，只有成功写入 metadata 
        事务才算成功。 如果第一步时出错，因为数据会写在新的 page 不会覆盖原来的数据，且此时的 metadata 不变，
        后面的事务仍会访问之前的完整一致的数据。

    关键就是要保证 metadata 写入出错也不会影响数据库:
        meta.checksum 用于检测 metadata 的 corruption。
        metadata 交替保存在文件前2个 page 中，当发现一个新写入的 metadata 出错时会使用另一个。

```go
// write writes the meta onto a page.
func (m *meta) write(p *page) {
  if m.root.root >= m.pgid {
      panic(fmt.Sprintf("root bucket pgid (%d) above high water mark (%d)", m.root.root, m.pgid))
  } else if m.freelist >= m.pgid {
      panic(fmt.Sprintf("freelist pgid (%d) above high water mark (%d)", m.freelist, m.pgid))
  }

  // Page id is either going to be 0 or 1 which we can determine by the transaction ID.
  p.id = pgid(m.txid % 2)
  p.flags |= metaPageFlag

  // Calculate the checksum.
  m.checksum = m.sum64()

  m.copy(p.meta())
}
// meta retrieves the current meta page reference.
func (db *DB) meta() *meta {
  // We have to return the meta with the highest txid which doesn't fail
  // validation. Otherwise, we can cause errors when in fact the database is
  // in a consistent state. metaA is the one with the higher txid.
  metaA := db.meta0
  metaB := db.meta1
  if db.meta1.txid > db.meta0.txid {
      metaA = db.meta1
      metaB = db.meta0
  }

  // Use higher meta page if valid. Otherwise fallback to previous, if valid.
  if err := metaA.validate(); err == nil {
      return metaA
  } else if err := metaB.validate(); err == nil {
      return metaB
  }

  // This should never be reached, because both meta1 and meta0 were validated
  // on mmap() and we do fsync() on every write.
  panic("bolt.DB.meta(): invalid meta pages")
}
```

    Isolation:
        boltdb 支持多个读事务与一个写事务同时执行，写事务提交时会释放旧的 page，分配新的 page，只要确保分配的新 page 
        不会是其他读事务使用到的就能实现 Isolation。

        在写事务提交时，释放的老 page 有可能还会被其他读事务访问到，不能立即用于下次分配，所以放在 freelist.pending 中， 
        只有确保没有读事务会用到时，才将相应的 pending page 放入 freelist.ids 中用于分配:
        freelist.pending: 维护了每个写事务释放的 page id。
        freelist.ids: 维护了可以用于分配的 page id。

        每个事务都有一个 txid，db.meta.txid 保存了最大的已提交的写事务 id:
            读事务: txid == db.meta.txid。
            写事务：txid == db.meta.txid + 1。

        当写事务成功提交时，会更新 metadata，也就更新了 db.meta.txid。
        txid 相当于版本号，只有低版本的读事务才有可能访问高版本写事务释放的 node，当没有读事务的 txid 比该写事务的 txid 
        小时，就能释放 pending page 用于分配。

    实现如下：
        db中维护了一个读事务的队列：
        创建读事务是会追加到db.txs中，当读事务rollback时会从中移除；
        在创建写事务时，会找到 db.txs 中最小的 txid，释放 freelist.pending 中所有 txid 小于它的 pending page:

```go
// Keep track of transaction until it closes.
    db.txs = append(db.txs, t)

    tx.db.removeTx(tx)
```

    不支持写写并发：
        boltdb 不支持多个写事务同时执行，在创建写事务的时候会加上锁。
        在现在的实现上支持写写并发会增加很多复杂度，不过 boltdb 提供了 db.Batch():
        当多个 goroutine 同时调用时，会将多个写事务合并为一个大的写事务执行。
        若其中某个事务执行失败，其余的事务会重新执行，所以要求事务是幂等的。


    