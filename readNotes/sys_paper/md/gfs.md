# Google GFS
---
[连接1-GFS论文详解](https://spongecaptain.cool/post/paper/googlefilesystem/)

## GFS是什么？

    Google的文件系统，是一种分布式文件系统，由Google公司开发，运行在Linux上，未开源，所以由雅虎牵头的根据Google三篇论文做出来的Hadoop并没有
    达到官方的速度，在大数据领域只有Google在领跑

    GFS 的地位相当高，早些年基于 GFS 系统的 MapReduce 以及 BigTable 框架实际上随着时代都逐渐被取代了，并且 MapReduce 以及 BigTable
    作为 Google 大数据三驾马车的后两篇论文，也没有 GFS 那样写得详细。

    下图是基于GFS文件系统发展出来的文件系统分支：

![](../photo/070221.png)

---
## GFS介绍

1、GFS系统使用背景：
    
    1.1、分布式组件经常发生错误，应当将此视为常态而不是意外
    2.1、文件通常是大文件、而不是小文件
    3.1、大部分文件通过append新数据的方式实现修改，而不是直接重写现有数据
    4.1、协同设计应用以及文件系统可以提高系统整体灵活性最终使系统受益

2、论文的前提和假设：
    
    2.1、分布式系统的各个组件是廉价的商业主机，而不是专业服务器，不得不频繁的自我检测，发现故障、自我容错、恢复
    2.2、文件数量处于几百万的规模。每个文件通常是100MB（GB），同时也支持小文件
    2.3、读工作模式主要由两种读方式构：大规模的串行读和小部分的随机读
            大规模的串行读：顺序读取数百及以上个KB大小的数据
            小规模的随机读：以任意偏移量读取几个KB大小的数据
    2.4、写工作负载主要是大规模的、连续的写操作，这些操作讲数据追加到文件末尾，写操作的规模通常和大规模串行读的规模类似
    2.5、系统需要支持并发写，即支持数百台并发地追加数据到一个文件，操作的原子性和同步开销是主要的指标
    2.6、高持续带宽比低延迟更重要

![](../photo/070222.png)

3、接口-为完全实现的POSIX

    GFS作为一个分布式文件系统，对外提供了一个传统的单机文件系统接口，但是由于效率和实用性的角度，并没有实现标准的文件系统POSIX API

    文件通过目录进行分层管理，通过路径名来定位，支持文件的 create，delete，open，close，read 以及 write 操作。

    此外 GFS 还支持如下两个特性：
    
        Snapshot 快照：快照指的是以低成本方式创建文件和目录树（directory tree）的副本；
        Record Append 记录追加：记录追加指的是 GFS 允许多个客户机并发安全地向同一文件追加数据，同时保证每个客户追加操作的原子性；

---
## GFS系统架构与设计

1、GFS集群的系统架构

    一个GFS Cluster分为两个组件：
        单个master节点
        多个chunkserver节点

    一个GFS集群同时可以被多个client节点访问
    一个GFS集群的架构可以使用下图表示：
![](https://spongecaptain.cool/images/img_paper/image-20200719162238223.png)
    
    这是个典型的Naster+Worker结构：存在一个Master 来管理任务、分配任务，而 Worker 是真正干活的节点。在这里干的活自然是数据的存储和读取。

2、大文件的分块存储：

    ？：大文件分块存储和 MySQL 的水平扩展、垂直扩展的理念是一样的，或者说类似于 Redis 的主从节点的设计。不过，如果要讨论最基本的原理，
    那便是：将串行通为并行。

    Chunk Size 是整个分布式文件系统的最重要的参数之一，GFS 以 64 MB 为固定的 Chunk Size 大小，这远远大于典型的单机文件系统 chunk 的大小。
    Linux 系统每页4k大小，Chunl内部是按照64kb进行存储的。

    
    分布式系统由于不可避免的故障，因此我们需要使用 replication 机制，每一个 chunk 都存在着若干个副本（它们不一定完全一样 ，因为 GFS 并不是一
    个强一致性文件管理系统），我们称这些 chunk 的副本为 replica（复数形式为 replicas）。每个 chunk 或者 replica 都作为普通的 Linux 文件
    存储在 chunkserver 上。
    
3、大的Chunk有什么优点：

    3.1、减少了 Client 与 Master 服务器交互的次数
    3.2、减少了 GFS Client 与 GFS chunkserver 进行交互的数据开销，这是因为数据的读取具有连续读取的倾向，即读到 offset 的字节数据后，下一
         次读取有较大的概率读紧挨着 offset 数据的后续数据，chunk 的大尺寸相当于提供了一层缓存，减少了网络 I/O 的开销；
    3.3、减少了存储在主服务器上的元数据的大小。这允许我们将元数据保存在内存中。

    缺点：小数据量（比如仅仅占据一个 chunk 的文件，文件至少占据一个 chunk）的文件很多时，当很多 GFS Client 同时将 record 存储到该文件时就会
        造成局部的 hot spots 热点。

    事实上，GFS 并没有很好地支持彻底的小文件系统，应用场景还是大文件存储。

![](https://spongecaptain.cool/images/img_paper/image-20200723110430786.png)

    从上面图示可以看出，每一个文件都被划分为多个 chunk，这里先单独将文件存储拎出来，并不考虑 replication 机制，
    那么一个文件是以如下的形式拆分管理的：
    这里此文件是一个相当大的文件（GB），如果文件足够小，那么仅仅会对应一个 chunk

![](https://spongecaptain.cool/images/img_paper/FileToChunkInGFS.png)

    File分块存储的特点：
        1、每个chunk的大小固定，默认值是64MB
        2、chunk_1~chunk_n在逻辑上是连续的，但是在磁盘上存储上并不一定是连续的
        3、如果chunk涉及replication，GFS在默认情况下会为每一个chunk多创建2gereplicas，一式三份，分布在多个磁盘上

    但是replication机制带来了额外的复杂性，eg：查找的复杂性。
        解决方式：每一个chunk在创建的时候使用一个不可变的全局唯一64位的ID 身份符来标记，由Master节点在chunk被创建时进行分配

4、单Master节点的设计：

    GFS Master节点负责的工作室系统级别的控制：
        1、chunk的版本控制
        2、孤块的回收
        3、两个chunkservers上的偏移
        4、append record 原子性的确保
        5、metadata的维护和管理

    Master 周期性通过 HeartBeat 机制和每一个 chunkserver 进行通信，进行指令的发送以及状态信息的接收。

    但是需要注意是，GFS 中 Master 节点基于 Master-Worker 模式设计（有点类似于中间件的名称服务器），并没有把 Master 节点作为代理节点，
    换句话说，GFS 中 Master 只是个大老板，或者大管家，但是其本身不干活。这意味着 GFS Client 读写操作的数据 I/O 传输直接与 chunkserver 进行。

    同时，GFS 实际上仅仅拥有一个 Master 服务器，这极大地简化了设计难度，做出快速的决策。
    
5、GFS Client文件数据的步骤：

![](https://spongecaptain.cool/images/img_paper/image-20200719162238223-6596397.png)

    1、GFS Client 首先对要读取的字节相对偏移量在 chunk size 固定的背景下计算出 chunk index；一个chunk size为64B
    2、给 GFS Master 发送 file name 以及 chunk index，即文件名和 chunk index(索引)；
    3、GFS Master 接收到查询请求后，将 filename 以及 chunk index 映射为 chunk handle 以及 chunk locations，并返回给 GFS Client；
    4、GFS Client 接收到响应后以 key 为 file name + chunk index，value 为 chunk handle + chunk locations 的键值对形式缓存此次查询信息 ；
    5、GFS Client 向其中一个 replicas (最有可能是最近的副本)发送请求，去请求中指定 chunk handle 以及块中的字节范围；

        如果你足够细心的话，你会注意到从 GFS Master 返回的是 chunk locations，表示 replicas，即多个 replica 的地址。
        注意，如果缓存没有过期，那么 GFS Client 与同一个 replica 对同一 chunk 的读并不需要 clinet-master 进行通信；
        同时，客户端通常会对多个 chunk 进行合理的聚合，可以一次向 master 查询多个 chunk + index 的 metadata，以及一次向 chunkserver 
        读取多个 chunk 的数据；

    注：此处的句柄就是chunk的全局身份标识符

6、Master节点在内存中的设计：

![](../photo/070223.png)

    Table1：基于Hashmap实现，根据filename进行快速的查找，O（1）
    Table2：基于B+树实现，因为可能涉及到大字节的数据，需要一次性返回多个chunk数据，同时也可能存在范围查找

    GFS 系统中，为了加快响应客户端关于 metadata 数据的请求，因此会将 metadata 存储于内存中，但是因为内存是易失性存储体，因此还需要持久化操作。具体来说：
        客户端向 Master 节点请求的 metadata 数据直接存储于 Master 的内存中，避免每次请求都需要进行磁盘 I/O；
        Master 节点使用日志 + checkpoint 的方式来确保数据的持久化；
    

    Q：为什么有 Table1 与 Table2 的持久化机制了，还需要日志和 checkpoint？

    这是因为 Table1 与 Table2 的数据结构无论是 HashMap 还是 B+Tree，如果选择新加入节点后马上进行持久化，那么就会面临随机 I/O 的问题，因为
    它们本质上都基于节点实现，而节点并不基于连续地址进行存储。出于效率的考虑，这两张表并不会在每一个写操作执行时就执行持久化机制，而是定期执行。
    不过定期执行就存在因为掉电、故障后数据丢失问题，因此需要引入日志系统。这通常被称为持久化 snapshot of memory。

    由于日志仅仅就是追加数据，日志的追加操作属于顺序磁盘 I/O，因此每一条写操作生效前都可以提前把日志记录到磁盘上，再进行真正的写操作，因此数据总是
    能够安全地确保持久化。主机重启时重新执行一遍日志即可。

    日志系统非常冗长，如果每次启动 Master 节点时都执行全部的日志记录一次，那么效率就会很低。另一方面，GFS 本身就必然会持久化 Table1、Table2，
    我们应当利用这个特点。checkpoint 就是来解决这个问题的。每次持久化 Table1、Table2 成功后，都会在日志系统上打上一个 checkpoint，用于说明
    下一次启动 master 之后可以先读取持久化了的 Table1、Table2 磁盘数据，然后从日志系统的 checkpoint 向后执行。这样一来 Master 启动时的效率
    就不会很低了，因为并不会将日志从头到尾执行一遍，而仅仅是 checkpoint 到尾执行一遍。另一面，checkpoint 也赋予了日志系统删除陈旧的日志的能力，
    用来节约磁盘空间（checkpoint 前的字节数据理论上都可以删除）。

7、Metadata-元数据

    Master存储三类最重要的metadata数据：1、File、chunk、namespace；2、file到chunk的映射Map；3、每一个chunk replica的存储位置。
    所有的元数据都保存在主服务器的内存中。

    前两种类型(名称空间和文件到块的映射)数据同时也会通过 logging metuations 存储到主机本地磁盘上、复制到远程机器上来持久化。使用日志允许
    我们简单、可靠地更新 Master 的状态，并且不会在主崩溃时出现不一致的风险。另一方面，Master 并不会持久化 chunk replica 的位置信息。其
    通过启动时以及 chunkserver 加入集群时发起 chunkserver 的每个 chunk 数据块位置的询问。

    内存内的数据结构

    元数据保存在 Master 的内存中使得 Master 要对元数据做出变更变得极为容易；同时，这也使得 Master 能够更加高效地周期扫描集群的元数据，这
    扫描主要用户实现垃圾收集、出现 chunkserver 时重新创建其一个副本以及 chunkserver 之间的负载平衡。
    
    唯一的不足在于这使得整个集群所能拥有的 Chunk 数量受限于 Master 的内存大小，不过这是一个小问题，因为对于一个 64MB 大小的 Chunk，Master
    只需要维持不到 64 字节的元数据（虽然前者数据存储在磁盘上，后者数据存储在内存中）。
    
    如果需要支持更大的文件系统，通过给 Master 节点内存扩容代价是不大的，但是这个不大的代价可以带来在内存中存储元数据所获得的简单性、可靠性、性
    能和灵活性。
    
8、Chunk Locations：
    
    Master 服务器不持久化 chunkservers 拥有哪个 chunk 副本的持久记录，而是通过询问实现：
    Master 在启动时通过轮询获取每一个 chunkserver 上所有 chunk replica 的初始值；
    Master 在启动后通过心跳机制来检测每一个 chunkserver 上所有 chunk replica 的变化，以保证其拥有最新的数据。

9、Operation Log

    Operation Log 包含关键metadata 数据更改的历史记录，它是数据库的核心，它不仅是元数据的唯一持久记录，还充当定义并发操作顺序的逻辑时间线、
    文件和块，以及它们的版本都是唯一的，永远由它们创建的逻辑时间来标识。

    由于 Operation Log 日志非常重要，所以我们必须可靠地存储它，并且在元数据的更改被持久化之前，不能使更改对客户端可见。否则，即使块本身存活
    下来，我们也会丢失整个文件系统或最近的客户端操作。因此，我们将它复制到多个远程机器上，只有在本地和远程将相应的日志记录刷新到磁盘之后才响应
    客户机操作。在刷新之前，Master 批处理多个日志记录，从而减少刷新和复制对总体系统吞吐量的影响。

    Master 通过重播操作日志恢复其文件系统状态。为了最小化启动时间，我们必须保持日志较小。每当日志增长超过一定的大小时，Master 检查其状态，
    以便通过从本地磁盘加载最新的 checkpoint 并在此之后仅重放有限数量的日志记录来恢复。checkpoint 是一种紧凑的类似 b 树的形式，可以直接
    映射到内存中并用于名称空间查找，而无需进行额外的解析。这进一步加快了恢复并提高了可用性。
    
    `Master 恢复仅仅需要最新的完整检查点和后续日志文件`。旧的检查点和日志文件可以自由删除，但通常还是会保留了一些近期的日志，以防止意外。

10、Consistency Model-一致性模型

    Guarantees by GFS：

    文件名 namespace 命名空间的变化（比如，文件的创建）全权由 Master 节点在内存中进行管理，这个过程通过 namespace lock 确保操作的原
    子性以及并发正确性，Mater 节点的 operation log 定义了这些操作在系统中的全局顺序；
    
    在数据修改后，文件区域的状态取决于很多个条件，比如修改类型、修改的成功与否、是否存在并发修改，下表总结了文件区域的状态（来自于论文）：
![](https://spongecaptain.cool/images/img_paper/image-20200719211636393.png)

    `GFS对file region状态的概念定义：`

    consistent：所有的GFS Client将总是看到完全相同的数据。
    
    defined：当一个文件数据修改之后，如果file region还是保持 consistent状态，并且所有的client能够看到全部修改的内容。
    
    consistent but undefined：从定义上看，就是所有的Client能够看到相同的数据，但是并不能及时反映并发修改的任意修改。
            这里通常是指写操作冲突了，GFS并不能保证多个客户端的并发写请求的执行顺序，但是保证最终数据的一致性，也就是说它的执行次序并不
            能充分保证，但是最终所有的客户端查询时能读到相同的结果

    inconsistent：因为处于inconsistent状态，因此一定也处于undefined状态，造成此状态的操作也被认为是failed的，不同的Client在读到的
    数据不一致，同一个client在不同的时刻读取的文件数据也不一致
   

    表格将数据的修改分为两种情况：
        1、write：修改File中原有数据，具体来说就是在指定文件的偏移地址下写入数据
            GFS 没有为这类覆写操作提供完好的一致性保证：如果多个的 Client 并发地写入同一块文件区域，操作完成后这块区域的数据可能由各次
            写入的数据碎片所组成，此时的状态最好也就是 consistant but undefined 状态
        2、Record Append：即在原有File末尾Append（追加）数据，这种操作被GFS系统确保为原子操作，这是GFS系统最重要的优化之一。GFS 中
           的 append 操作并不简单地在原文件的末尾对应的 offset 处开始写入数据（这是通常意义下的append操作），而是通过选择一个offset，
           这一点在下面会详细说到。最后该被选择的 offset 会返回给 Client，代表此次 record 的起始数据偏移量。由于 GFS 对于 Record 
           Append 采用的是 at least once 的消息通信模型，在绝对确保此次写操作成功的情况下，可能造成在重复写数据。
        
    在一系列成功的修改操作以后，被修改的文件区域的状态是 defined 并包含最后一次修改的写内容。GFS 通过以下两种方式实现这一目标：
    在所有写操作相关的 replicas 上以同一顺序采用给 chunk 进行修改；
    使用 chunk version numbers（也就是版本号）去检测 replicas 上的数据是否已经 stale（过时），这种情况可能是由于 
    chunkserver 出现暂时的宕机(down)；
        一旦一个 replicas 被判定为过时，那么 GFS 就不会基于此 replicas 进行任何修改操作，客户机再向 Master 节点请求元数据时，
        也会自动滤除过时的 replicas。并且 Master 通常会及时地对过时的 replicas 进行 garbage collected（垃圾回收）。
    
    Q：
        缓存未过期时 replica 出现过时的问题：因为在客户机存在缓存 cache 的缘故，在缓存被刷新之前，客户机还是有机会从stale replica上
        读取文件数据。这个时间穿窗口取决于缓存的超时时间设置以及下一次打开同一文件的限制。另一方面，GFS 中的大多数文件都是 append-only，
        因此 stale replica 上的读仅仅是返回一个 premature end of chunk，也就是说仅仅没有包含最新的追加内容的 chunk，而不是被覆写
        了的数据（因为无法覆写），这样造成的影响也不会很大。

        组件故障问题：Master 节点进行通过与所有 chunkserver 进行 regular handshake（定期握手）来检测出现故障的 chunkserver，
        通过 checksumming（校验和）来检测数据是否损坏。一旦出现问题，GFS 会尽快地从有效的 replicas 上进行数据恢复，除非在 Master
        节点检测到故障之前，存储相同内容的 3 块 replica 都出现故障时才会导致不可逆的数据丢失。不过即使是这样，GFS 系统也不会不可用，
        而是会及实地给客户端回应数据出错的响应，而不是返回出错的数据。

11、Implications for Applications

    1、尽量选择append追加，而不是overwrite覆写，这是因为GFS仅仅保证append操作的一致性，但是覆写操作没有一致性保证
    2、写入数据时使用额外的校验信息，比如校验和（或者其他hash技术）
    3、为了避免write at least once而造成的重复数据，可以通过加入额外的唯一标识符来达到效果

---
##  系统内部的交互

    尽可能减少Master节点的负载

1、Leases and Mutation Order-租赁和修改顺序

    修改操作包括追加和覆写，这两种写操作都会最终作用于所有的 replica（通常意义上说就是一个写指令对应于 3 个 chunk+replica 的 I/O 写）。
    对于 chunk 的写操作涉及两种修改：

    在相关 chunkserver 上进行 I/O 写操作；
    在 Master 节点上修改 metadata；

    Leases的含义是租赁，其用于Master确保多个节点之间写操作顺序的一致性，具体操作如下：
    1、Client 向Master请求本次写操作涉及到的chunk name，通过chunk header 拿到chunk location返回给Client；
    2、Client找到其中一个距离最近的chunk作为primary节点，其他的同种数据的replicas则是secondaries（从属节点），Master 在此消息中还会
        告知被选中的 primary 节点来自客户端的多个写操作请求；
    3、chunkserver收到lease以及写操作请求后，其才认为自己有权限进行写操作，并决定多个写操作的执行顺序，顺序被称为 serial order（串行执行顺序）
        更改本地数据后，同步传递给secondaries节点；
    4、当 primary 决定好顺序后，会将带有执行顺序的 lease 返回给 master 节点，master 节点随后会负责将顺序分发给其他两个 replica。其他两个
        replicas 并没有选择权，只能按照 primary 决定的顺序进行执行。

    这种 Lease 机制减少了 Master 的管理开销，同时也确保了线程安全性，因为执行顺序不再由 Master 去决定，而是由拥有具体 lease 租赁的
    chunkserver 节点决定。租赁的默认占用超时时间为 60s，但是如果 Master 又接收到对同一个 chunk 的写操作，那么可以延长当前 primary
    节点的剩余租赁时间。

    流程如下：

![](https://spongecaptain.cool/images/img_paper/image-20200720205146260.png)

![](../photo/070301.png)


2、Data Flow-数据流

    Data Flow传输模型的两个关键字：linearly、pipeline。
    关于数据流：GFS的目标是将数据流和控制流进行解藕，解藕的方式就是以线性的方式进行传输数据。具体来说就是Client给primary传递数据，而primary给
    replica传递数据，而下一个 replica 又负责给下下一个 replica 传输数据。这种方式能充分利用每台机器的网络带宽，避免网络瓶颈和高延迟链接，并最
    小化通过所有数据的延迟。

    具体来说，避免 network bottlenecks 与 high-latency links 的线型链路采用了如下做法：每一个节点都会将数据转发给最近且还没有收到数据的节
    点，节点通过 IP 地址来估算两个节点之间的链路距离。

    Data Flow 为了最小化延迟，还采用管道传输模型管道传输模型可以用生活中自来水管道来解释，水经过一个节点之后就会去下一个节点，在计算机中水就是字
    节。当一个 chunkserver 接收到一些数据后（不是一个字节就会触发，而是要有一定阈值）就会立即转发给下一个节点，而不是等到此次写操作的所有数据接
    收完毕，才开始向下一个节点转发。值得一提的是，GFS 系统使用的是全双工网络，即发送数据时不会降低数据接收速率。

3、Atomic Record Appends-原子的记录追加

    GFS 提供了原子的 Record Append（文件追加）操作。 在传统的写操纵中，客户端需要指定写入数据的地址偏移量，但是对同一地址的并发写操作并不能保证
    是序列化发生的，因此某一个文件区域可能包含来自多个客户端写操作的碎片。但是在 GFS 中的 Record Append 中，客户端仅仅负责指定要写的数据，GFS
    以 at least once 的原子操作进行写，写操作的相关数据一定是作为连续的字节序列存放在 GFS 选择的偏移量处。


    具体来说：record append 是指向已经存储的文件的末尾追加数据，因此客户端并不需要像读操作那样提供一个数据范围，因为 record append 操作总是
    在文件末尾追加数据，这个地址偏移量应当交给 chunksever 来确定。

    3.1、`GFS record append 操作的内部执行逻辑如下：`
    
        1、Client 确定 file name 以及要写入的 byte data（形式上可以选择一个 buffer 来存储要写入的字节数据）；
        2、Client 向 Master 发出期望进行 record 操作的请求，并附带上 file name，但是不用携带字节数据；
        3、Master 接收到请求之后得知是一个 append record 请求，并且得到 file name。Master 节点通过其内存中的 metadata 得到当前 file 分块存储
            的最后一个 chunk 的 chunk handle 以及 chunk 所在的所有 chunkserver；
        4、Master 之后将这些信息转发给 Client；
        …下面的操作就类似于 1 小节中的过程。

    
    3.2、`Record append操作涉及primary的选择步骤：`

        1、Master 节点在接受到修改请求时，会找此 file 文件最后一个 chunk 的 up-to-date 版本（最新版本），最新版本号应当等于 Master 节点的版本号；
            什么叫最新版本。chunk 使用一个 chunk version 进行版本管理（分布式环境下普遍使用版本号进行管理，比如 Lamport 逻辑时钟）。
            一个修改涉及 3 个 chunk 的修改，如果某一个 chunk 因为网络原因没能够修改成功，那么其 chunk version 就会落后于其他两个 chunk，
            此 chunk 会被认为是过时的。
        2、Master 在选择好 primary 节点后递增当前 chunk 的 chunk version，并通过 Master 的持久化机制持久化；
        3、通过 Primary 与其他 chunkserver，发送修改此 chunk 的版本号的通知，而节点接收到次通知后会修改版本号，然后持久化；
        4、Primary 然后开始选择 file 最后一个文件的 chunk 的末尾 offset 开始写入数据，写入后将此消息转发给其他 chunkserver，它们也对相同
            的 chunk 在 offset 处写入数据；

    3.3、`问题：`
        1、如果向file追加的数据超过了chunk剩余的容量怎么办？
            首先，这是一个经常发生的问题，因为 record append 操作实际上能一次添加的数据大小是被限制的，大小为 chunksize（64 MB）的 1/4，
            因此在大多数常见下，向 chunk append 数据并不会超出 64 MB 大小的限制；
            其次，如果真的发生了这个问题，那么 Primary 节点还是会向该 chunk append 数据，直到达到 64MB 大小上限，然后通知其他两个replicas
            执行相同的操作。最后响应客户端，告知客户端创建新 chunk 再继续填充，因为数据实际上没有完全消耗掉；

        2、Record Append 作为 GFS 中写操作的一种类型，自然准许遵循 3.1 小节中的数据流机制；

        3、At least once 机制引发的问题：

![](../photo/070401.png)

4、Read操作

    1、Client —> Master 请求：file name + range(或者说是 offset，总之是客户端打算读取的字节范围)；
    2、Master —-> Client 响应：chunk handle + a list of server
    3、这一步 Client 在本地存在缓存，它在缓存失效之前读取同一个 chunk 的数据并不需要向 Master 重新发出索要对应 metadata 数据的请求，而是直
        接使用缓存，即直接向 chunkserver 进行交互数据；
    4、client 在 a list of server 中跳出一个 chunkserver，GFS 论文指出，通过 chunkserver 的 IP 地址，能够 guess 出距离当前 client
        最近的 chunksver，然后 Client 会优先向最近的 chunkserver 请求数据读取；
    5、chunkserver 收到数据读取请求后，根据 clinet 发来的 chunk hanle 进行磁盘 I/O 最终返回数据给 client；

    问题：

![](../photo/070402.png)

5、Snapshot快照

![](../photo/070403.png)

---
## MASTER OPERATION-主节点操作

    Master 节点负责的工作有：
        所有 namespace 的管理工作；
        管理整个系统中的所有 chunk replicas：
        做出将 chunk 实际存储在哪里的决定；
        创建新的 chunk 和 replica；
        协调系统的各种操作（比如读、写、快照等），用于保证 chunk 正确且有效地进行备份；
        管理 chunkserver 之间的负载均衡；
        回收没有被收用的存储空间；
1、namespace management and locking-命名空间管理和锁机制
    
    Master 节点有很多操作都需要执行很长时间，比如：snapshot 操作必须向 chunkserver 撤回 snapshot 涉及的所有 chunk 的 lease。我们并不希望
    这些耗时的操作会影响 Master 节点的其他操作。出于这个目的，我们给 namespace 上锁来实现可以同时进行多个操作以及确保操作的正确串行执行顺序。

    不同于其他传统的文件系统，GFS 并没有为每一个目录创建一个用于记录当前目录拥有哪些文件的数据结构，也不支持文件和目录的别名。GFS 逻辑上将其
    namesapace 当做一个查询表，用于将 full pathname（要么是一个 absolute file name，要么是一个 absolute directory name） 映射为
    metadata。如果使用 prefix compression（前缀压缩），那么这个表可以在内存中被高效地表示。在 namespace 树中的每一个节点对应一个 full
    pathname，都拥有一个与之相关联的 read-write lock(读写锁)。

![](../photo/070404.png)

2、Replica Placement - 确定 Replica 的存放位置

    GFS 系统是一个分布式系统，一个 GFS 集群通常会有成百上千个 chunkserver，它们分布在很多 machine rack（机架上）。每一个 chunkserver 
    同一可以被成百上千个其他位于同一（或不同）的机架上的 chunkserver 访问。在两个不同机架上进行通信的 chunkserver 需要通过交换机。此外，机架
    的输入输出带宽可能小于机架内所有 chunkserver 的总带宽。这些多级分布对数据的可伸缩性、可靠性和可用性提出了挑战。

    chunk replica placement policy 有两个目的：
    
    最大化数据 reliability（可靠性）和 availability（可用性）；
    最大化 network bandwidth utilization（网络带宽利用率）；
    为了达到上述目的，在 chunkserver 和 chunkserver 之间传播 replicas 是不够的，这仅仅能够在机器或磁盘故障时保障可靠性和可用性，且会给单个
    rack 机架带来带宽使用上的压力，因为读取数据时，无论客户端最终选择哪一个 replicas，最终都是消耗着同一个 replicas 的带宽。我们还需要在不同的
    racks(机架) 上传输 chunk replicas，这能够在一整个机架故障时都能够确保可靠性和可用性。另一方面，由于 chunk replicas 分布存储在不同机架
    上的 chunkserver 上，因此降低了每一个 rack 提供 replicas 读写时的带宽压力，因为相同于多个机架平均了带宽压力。

3、Creation, Re-replication, Rebalancing-创建 重放置 均衡再建

    chunk replicas 出于三个原因被创建：

    chunk reation；
    Re-replication；
    reblancing；
    当 Master 节点创建了一个 chunk，它负责确定将这个 initially empty replicas 放置到哪里，它鉴于以下几个因素进行判断：
    
    默认情况下一个 Master 创建一个 chunk 对应 chunkserver 上创建 3 个 replica。
    
    选择将 replica 放置于磁盘空间利用率低于平均水平的 chunkserver，这样一来能够保持所有 chunkserver 在磁盘利用率上保持一致性；
    限制每一个 chunkserver 上最近创建的 chunk 的个数，虽然仅仅创建一个 chunk 代价不高，但是它通常是即将出现大量写操作的前兆，因为 chunk 通
    常是在写操作时被创建。
    正如 4.2 节所谈到的，我们期望将 replicas of chunk 分散放置在不同的 rack 上。
    另一方面，一旦可用的 replicas 数量下降到用户预设值（默认为 3），那么 Master 就会开始 re-replicate chunk 操作。这可能由于如下的原因造成：
    
    chunkserver unavailable（不可用），比如它给 Master 发送如下状态信息：它的 replica 崩溃了、某一个磁盘不可读。
    程序员修改了配置，动态增加了 replication 的个数的要求；
    当 chunk 需要被 re-replicated 时，Master 通过以下因素来确定执行优先级：
    
    根据距离 replication goal 的配置的距离来确定优先级。比如默认三个 replicas，有一组 replicas 中有两个出现了不可用，而另一组仅仅只有一个
    出现了不可用，因此前者比后有优先级高；
    最近活动的文件（被读、被写）比最近删除的文件的 chunk 有更高的优先级；
    如果 chunk 的读写可能阻塞客户端，那么该 chunk 将有较高的优先级，这能够减少 chunk 故障时对使用 GFS 的应用程序的影响；

    Master 节点选择优先级最高的 chunk 先进行 clone（克隆），方式是通知相关 chunkserver 直接从存在可用 replica 的 chunkserver 上进行
    chunk 数据的拷贝。

    注意：为了防止因为过多的 clone 操作占用过多的系统带宽，Master 节点既会限制当 chunkserver 进行的 clone 数量，又会限制整个 GFS 集群同时
    进行的 clone 操作数量。而且 chunkserver 自身也会对消耗在 clone 操作上的带宽占比，其方式是限制对复制源的请求数量。

    另一方面，Master 的 rebalancing 机制是指：Master 节点会负责检查当前 replica 的分布，然后会将相关 replicas 移动到更好的磁盘位置。因此，
    通过这个机制 Master 节点能将一个新加入的 chunkserver 自动地逐渐填充，而不是立即在其刚加入时用大量的写操作来填充它。新的 replicas 的放置
    原理和上面两种类似。当然 Master 还要选择从哪一个 chunkserver 上删除相关 replicas，以便将其移动到磁盘空闲的 chunksever 上，这里就挑选磁
    盘空闲水平低于平均值的 chunkserver。

4、Garbage Collection

![](../photo/070405.png)

5、Stale Replica Detection-过期 replica 检测

    当 chunkserver 故障了或者因为宕机没能够正确地实施写操作，那么 Chunk replicas 的状态就变为 stale。Master 节点为每一个 chunk 维护一个
    chunk verison nember（chunk 版本号）来辨别哪些 chunk 是 up-to-date，哪些 chunk 是 stale。

    当 Master 赋予一个 chunk 新的租赁时，其就会使 chunk version 自增，并将此版本号告知其他 replicas。
    
    这里也就是说 Master 节点、primary 节点的 chunk version 和其他 chunkserver 节点的 chunk 的 chunk version 会保持一致。
    
    版本号的通知优先于写操作执行前。如果其他 replica 此时不可用，那么这个 chunk version 的通知就无法到，因此其 chunk version 就不会增加。
    那么当此 Chunkserver 重启后的心跳消息中就会包含此 chunk version 信息，Master 就能借此发现 ChunkServer 拥有 stale 的 replica。
    如果 Master 发现高于自己记录的 chunk version number，那么 Master 会认为自己在授予 lease 时发生了错误，然后将高版本的 chunk version
    number 视为最新版本。
    
    Master 节点在定时的垃圾收集任务中删除 stale replicas，当客户端对 stale chunk 请求时，Master 节点会发出此 chunk 不存在的响应。

---
## FAULT TOLERANCE AND DIAGNOSIS-容错和诊断

    对于 GFS 系统设计的最大挑战是需要处理频繁的组件故障，组件的质量一般以及数量众多一起提高了挑战难度。我们总是不能完全信任主机与硬盘，组件的故障
    可以引发很严重的问题，比如 GFS 系统不可用，甚至是错误的数据。下面，我们将讨论我们如何来解决这些挑战，如何在错误不可避免地发生时进行问题诊断。

1、High Availability-高可用性
    
![](../photo/070406.png)

2、Data Integrity-数据完整性

![](../photo/070407.png)

3、 Diagnostic Tools-诊断工具

![](../photo/070408.png)
