# DESC

---
处理问题：
1、数据库性能对比：
    
    1、boltDB
    2、levelDB
    3、rocksDB
    4、redisDB

2、索引探究：

    1、Adaptive Radix Tree
    2、Masstree
    3、Bw-Tree

3、并发跳表

    主要讨论跳表的多读多写特性：高并发跳表

---

## 数据库性能对比

1、boltDB

    boltdb 有如下性质：
        K/V 型存储，使用B+树索引（改进）。
        支持namespace，每对K/V存放在一个Bucket下，不同Bucket可以有相同的key，支持嵌套的Bucket。
        支持事务(ACID)，使用MVCC和COW，允许多个读事务和一个写事务并发执行，但是读事务有可能会阻塞写事务，适合读多写少的场景

    特点：


