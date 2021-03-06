# 深入理解LSM-tree

---
[连接1-深入理解LSM](https://cloud.tencent.com/developer/article/1441835)

1、前言：

    Google的三驾马车：GFS（2003）、MapReduce(2004)、BigTable（2006）
    其中BigTable中使用的文件组织方式，名字叫Log Structured Merge Tree.
    
    在面对亿级别之上的海量数据的存储和检索的场景下，我们选择的数据库通常都是各种强力的Nosql
    Apache：Hbase、Cassanrda
    Google：LevelDB
    Facebook：RocksDB

    底层使用的数据结构，都是仿照BigTable中的文件组织方式来实现的

2、LSM-Tree

    是一种分层、有序，面相磁盘的数据结构，其核心思想充分利用了，磁盘批量的顺序写要远比随机写性能高出很多
    围绕这一原理进行设计和优化，以此让谢性能达到最优，正如我们普通的Log的写入方式，这种结构的写入，全部
    都是以Append的模式追加，不存在删除和修改。不过虽然大量提升了写能力，但是是建立在牺牲读性能的基础之
    上的，适用于多写少读的场景。
    
    故LSM被设计来提供比传统的B+树或者ISAM更好的写操作吞吐量，通过消去随机的本地更新操作来达到这个目标。
    这里面最典型的例子就属于Kakfa了，把磁盘顺序写发挥到了极致，故而在大数据领域成为了互联网公司标配的分
    布式消息中间件组件。

    优缺点：虽然这种结构的写是非常简单有效，但其缺点是对读取性能特别是随机读写很不友好，所以要添加日志进
    行优化

    1、数据是被整体访问的，大多数数据库的log，包括mysql的binlog
    2、数据是通过文件的偏移量offset访问的，比如Kafka

    `但要想支持更复杂和高效的读取，比如key查询和按range查询，就得需要做一步的设计`，这就是LSM结构，`除了
    利用磁盘顺序写之外，还划分了内存+磁盘多层合并结构的原因`。正是基于这种结构再加上不同的优化实现，才造
    就了在这之上的各种独具特点的NoSQL数据库，如Hbase，Cassandra，Leveldb，RocksDB，MongoDB，TiDB等。

3、SSTable和LSM

    Google的bigtable是一片闭源的高性能的KV系统，而LevelDB就是这个KV系统开源的单机版实现，是高度复刻的版本。

    在LSM-Tree里面，核心的数据结构就是SSTable，全称是Sorted String Table，SSTable的概念其实也是来自于
    Google 的 Bigtable 论文，论文中对 SSTable 的描述如下：

    An SSTable provides a persistent, ordered immutable map from keys to values, where both 
    keys and values are arbitrary byte strings. Operations are provided to look up the value 
    associated with a specified key, and to iterate over all key/value pairs in a specified 
    key range. Internally, each SSTable contains a sequence of blocks (typically each block 
    is 64KB in size, but this is configurable). A block index (stored at the end of the SSTable) 
    is used to locate blocks; the index is loaded into memory when the SSTable is opened. 
    A lookup can be performed with a single disk seek: we first find the appropriate block 
    by performing a binary search in the in-memory index, and then reading the appropriate 
    block from disk. Optionally, an SSTable can be completely mapped into memory, which allows 
    us to perform lookups and scans without touching disk.

    SSTable是一种持久化，有序且不可变的键值存储结构，它的Key和Value都是任意的字节数组，并且提供了按指定Key查找和
    指定范围的Key区间迭代遍历的功能，SSTable内部包含了一系列可配置大小的Block块，`大小是64kb`，关于这些Block块的
    Index存储在SSTable的尾部，用于帮助快速查找特定的Block，当一个SSTable被打开的时候，Index会被加载到内存，然后
    根据Key进行一个二分查找，找到key的offset之后，进行读取，如果内存足够大的话，可以直接把SSTable直接通过MMAP的技
    术映射到内存，进行更快的查找。
![](photo/3001.png)
    

4、LSM写数据
    
    LSM的分层结构：SSTable有一份在内存里面，其它的多级是在磁盘上。

![](photo/3002.png)  
    
    1、当收到一个写请求的时候，会先把该条数据记录在log里面。用作故障恢复
    2、当写完log之后，会把该条数据写入内存的SSTable里面，也成Memtable注意为了维持有序性，在内存里面可以采用红黑树
    或者跳跃表相关的数据结构。
    3、当Memtable超过一定的大小后，会在内存里面冻结，变成不可变的Memtable，同时为了不阻塞写操作需要更新生成一个新
    的memtable继续提供服务
    4、把内存里面不可变的Memtable给dump到硬盘上的SSTable层中，此过程称为Minor Compaction，这里需要注意在L0层
    的SSTable是没有进行合并的，其中可能会出现key range重叠，在层数大于0层之后的SSTable，不存在重叠key
    5、当每层的磁盘体积上超过一定的大小或者个数，也会周期的进行合并，此步骤称为Major Compaction，这个阶段会真正的
    消除掉被标记删除的数据以及多版本数据的合并，避免浪费空间，注意SSTable都是有序的，我们可以直接采用merge sort进行高效合并。

5、LSM读数据

    1、当收到一个读请求的时候，会直接现在内存里面查询，如果查询到就返回。
    2、如果没有查询到就会下沉，直到所有level层查询一遍得到结果。
    
    Q：如果分层过多，会导致查询效率低下？如何优化？

6、查询优化：
    
    1，压缩

    SSTable 是可以启用压缩功能的，并且这种压缩不是将整个 SSTable 一起压缩，而是根据 locality 将数据分组，每个组分别压缩，
    这样的好处当读取数据的时候，我们不需要解压缩整个文件而是解压缩部分 Group 就可以读取。

    2，缓存

    因为SSTable在写入磁盘后，除了Compaction之外，是不会变化的，所以我可以将Scan的Block进行缓存，从而提高检索的效率

    3，索引，Bloom filters

    正常情况下，一个读操作是需要读取所有的 SSTable 将结果合并后返回的，但是对于某些 key 而言，有些 SSTable 是根本不包含对
    应数据的，因此，我们可以对每一个 SSTable 添加 Bloom Filter，因为布隆过滤器在判断一个SSTable不存在某个key的时候，那么
    就一定不会存在，利用这个特性可以减少不必要的磁盘扫描。

    4，合并

    这个在前面的写入流程中已经介绍过，通过定期合并瘦身， 可以有效的清除无效数据，缩短读取路径，提高磁盘利用空间。但Compaction
    操作是非常消耗CPU和磁盘IO的，尤其是在业务高峰期，如果发生了Major Compaction，则会降低整个系统的吞吐量，这也是一些NoSQL
    数据库，比如Hbase里面常常会禁用Major Compaction，并在凌晨业务低峰期进行合并的原因、

7、为什么LSM不直接顺序写入磁盘，而是需要在内存中缓冲一下？

    单条写的性能肯定没有批量写来的块，这个原理其实在Kafka里面也是一样的，虽然kafka给我们的感觉是写入后就落地，但其实并不是，
    本身是 可以根据条数或者时间比如200ms刷入磁盘一次，这样能大大提升写入效率。此外在LSM中，在磁盘缓冲的另一个好处是，针对新增
    的数据，可以直接查询返回，能够避免一定的IO操作。

8、B+Tree VS LSM-Tree

    传统关系型数据采用的底层数据结构是B+树，那么同样是面向磁盘存储的数据结构LSM-Tree相比B+树有什么异同之处呢？

    LSM-Tree的设计思路是，将数据拆分为几百M大小的Segments，并是顺序写入。// 一个SSTable是64kb大小

    B+Tree则是将数据拆分为固定大小的Block或Page, 一般是4KB大小，（Innodb是默认是16kb大小=4个4k），和磁盘
    一个扇区的大小对应，Page是读写的最小单位。

    在数据的更新和删除方面，B+Tree可以做到原地更新和删除，这种方式对数据库事务支持更加友好，因为一个key只会出现
    一个Page页里面，但由于LSM-Tree只能追加写，并且在L0层key的rang会重叠，所以对事务支持较弱，只能在Segment 
    Compaction的时候进行真正地更新和删除。

    因此LSM-Tree的优点是支持高吞吐的写（可认为是O（1）），这个特点在分布式系统上更为看重，当然针对读取普通的
    LSM-Tree结构，读取是O（N）的复杂度，在使用索引或者缓存优化后的也可以达到O（logN）的复杂度。

    而B+tree的优点是支持高效的读（稳定的OlogN），但是在大规模的写请求下（复杂度O(LogN)），效率会变得比较低，
    因为随着insert的操作，为了维护B+树结构，节点会不断的分裂和合并。操作磁盘的随机读写概率会变大，故导致性能降低。

    还有一点需要提到的是基于LSM-Tree分层存储能够做到写的高吞吐，带来的副作用是整个系统必须频繁的进行compaction，
    写入量越大，Compaction的过程越频繁。而compaction是一个compare & merge的过程，非常消耗CPU和存储IO，在高
    吞吐的写入情形下，大量的compaction操作占用大量系统资源，必然带来整个系统性能断崖式下跌，对应用系统产生巨大影响，
    当然我们可以禁用自动Major Compaction，在每天系统低峰期定期触发合并，来避免这个问题。

    阿里为了优化这个问题，在X-DB引入了异构硬件设备FPGA来代替CPU完成compaction操作，使系统整体性能维持在高水位并避免抖动，
    是存储引擎得以服务业务苛刻要求的关键。