# 第二次学习tidb

---

## 三大技术-存储

    - 保存数据-如果仅仅是一个简单的数组，那么直接在内存中创建数据，然后进行性能上的读写效果肯定是不错的 ，但是一旦停机或者服务器重启，所有的数据就会永久丢失（redis就是采用的这种的方式来进行数据的读写）

    - 需要进行多个数据副本的创建和复制，过程中是否能保证副本之间的一致性？

      能否支持跨数据中心的容灾？

      写入速度是否够快？

      数据保存下来后，是否方便读取？

      保存的数据如何修改？如何支持并发的修改？

      如何原子地修改多条记录

- Key-Value

    - 作为保存数据的系统，首先要决定的是数据的存储模型，也就是数据以什么样的形式存储下来， TiKV 的选择是 Key-Value 模型，并且提供有序遍历方法。 其中 Key 和 Value 都是原始的 Byte 数组，在这个 Map
      中，Key 按照 Byte 数组总的原始二进制比特位比较顺序排列。
      
      注意！！！这里进行排序的是Key不是全局的value，因为是hash映射，怎么能保证映射后的数据有序的，同理key的全局排序就轻松很多了
    
      这里的全局映射是按照行数据进行创建键值对信息，然后进行相关k-v的映射
      
- RocksDB
    - iKV 没有选择直接向磁盘上写数据，而是把数据保存在 RocksDB 中，具体的数据落地由 RocksDB 负责。这个选择的原因是开发一个单机存储引擎工作量很大，
      特别是要做一个高性能的单机引擎，需要做各种细致的优化，而 RocksDB 是一个非常优秀的开源的单机存储引擎，可以满足我们对单机引擎的各种要求，
      而且还有 Facebook 的团队在做持续的优化，这样我们只投入很少的精力，就能享受到一个十分强大且在不断进步的单机引擎。
      
- Raft

    - 主要的作用是负责数据同步的，进行多个存储副本的负载均衡，这里所有的数据副本都是按照Regin进行分区存储的，
      通过单机的 RocksDB，我们可以将数据快速地存储在磁盘上；通过 Raft，我们可以将数据复制到多台机器上，以防单机失效。数据的写入是通过 Raft 这一层的接口写入，
      而不是直接写 RocksDB。通过实现 Raft，我们拥有了一个分布式的 KV，现在再也不用担心某台机器挂掉了。
    
- Regin
    - 分 Range，某一段连续的 Key 都保存在一个存储节点上。TiKV 选择了第二种方式，将整个 Key-Value 空间分成很多段，每一段是一系列连续的 Key，
      我们将每一段叫做一个 Region，并且我们会尽量保持每个 Region 中保存的数据不超过一定的大小(这个大小可以配置，目前默认是 96mb)。
      每一个 Region 都可以用 **StartKey** 到 **EndKey** 这样一个左闭右开区间来描述。
    
        ![](https://download.pingcap.com/images/blog-cn/tidb-internal-1/3.png)
      
    - **以 Region 为单位，将数据分散在集群中所有的节点上，并且尽量保证每个节点上服务的 Region 数量差不多**
      
      **以 Region 为单位做 Raft 的复制和成员管理**
    
    - 按照Key切分成很多的region，通过raft进行均匀的分布，实现了存储的水平扩展，实现了负载平衡，
      同时会有相关的组件记录region在节点上面的分布情况
- MVCC
    - 多版本控制，通过在key后面添加version来实现，对于同一个 Key 的多个版本，我们把版本号较大的放在前面，版本号小的放在后面
    
- 事务
    - TiKV 的事务采用乐观锁，事务的执行过程中，不会检测写写冲突，只有在提交过程中，才会做冲突检测，冲突的双方中比较早完成提交的会写入成功，
      另一方会尝试重新执行整个事务。当业务的写入冲突不严重的情况下，这种模型性能会很好，比如随机更新表中某一行的数据，并且表很大。
      
---
## 说计算

-   关系模型到key-value模型的映射

    ```mysql
    CREATE TABLE User {
	ID int,
	Name varchar(20),
	Role varchar(20),
	Age int,
	PRIMARY KEY (ID),
	Key idxAge (age)    
    };
    ```
    
    - **对于任何一个table来说，需要存储的数据包括以下三部分：**
    
        1、表的元信息：记录表的列、元素各种属性信息
    
        2、Table中的Row
    
        3、索引数据

    - 表的信息映射方案：tableID_rowID

        ```
        Key: tablePrefix{tableID}_recordPrefixSep{rowID}
        Value: [col1, col2, col3, col4]
        ```
    - 索引信息的映射方案：tableID_indexID_columnsValue
    
        ```
        Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue
        Value: rowID
        ```
    - Unique Index：
    
         ```
        Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue_rowID
        Value: null
        ```
      
    - 
      ```
        var(
            tablePrefix     = []byte{'t'}
            recordPrefixSep = []byte("_r")
            indexPrefixSep  = []byte("_i")
        )
       ```
      
-   元信息管理：
    Database/Table 都有元信息，也就是其定义以及各项属性，这些信息也需要持久化，我们也将这些信息存储在 TiKV 中。
    每个 Database/Table 都被分配了一个唯一的 ID，这个 ID 作为唯一标识，并且在编码为 Key-Value 时，这个 ID 都会编码到 Key 中，再加上 m_ 前缀。
    
-   分布式SQL运算

    避免大量的RPC调用，将Filter下推到存储节点进行计算，只返回有效行，聚合函数和GroupBy也下推到存储节点

-   SQL层架构

    用户的 SQL 请求会直接或者通过 Load Balancer 发送到 tidb-server，tidb-server 会解析 MySQL Protocol Packet，
    获取请求内容，然后做语法解析、查询计划制定和优化、执行查询计划获取和处理数据。数据全部存储在 TiKV 集群中，
    所以在这个过程中 tidb-server 需要和 tikv-server 交互，获取数据。最后 tidb-server 需要将查询结果返回给用户。

    
    
