# 数据库发展历程

1、层次数据库/网状数据库---关系型数据库---关系型数据库数据仓库---开源数据库/Nosql数据库---云数据库

2、NewSQL：Spanner、F1、CockroachDB、TiDB

3、NoSQL放弃强事务性和一致性来保证高可用性和可扩展性

    许多应用不能放弃强事务性和一致性，比如财务系统、订单系统、人力资源系统

    如何保证高可用性和可扩展性的同时，又不损失强事务性和一致性？ NewSQL

4、NewSQL 论文基石
    
    Spanner：Google's Globally-Distributed Database

    F1: A Distributed SQL Database That Scales

    Spanner目标：跨数据中心的管理及复制数据、数据的重分片及均衡能力、主机间数据的自动迁移、
    
    F1设计的宗旨：SQL的全支持，同时支持索引，对事务支持ACID特性、无需改变应用程序就具备数据分片及均衡能力、
    系统可以添加资源进行扩展

5、TiDB的整体介绍：
    
    分布式、水平扩展能力、支持HTAP以及ACID事务的关系型数据库
    
    物理上是分布式的、逻辑上却是集中的

    高可用性：需要做到冗余、自动故障转移（人为介入恢复势必会增加系统服务的不可用时间，需要系统进行自动检测和故障转移）

    ACID：通过Primary Key所在的Region的原子性来保证分布式事务的原子性

    一致性：写入数据之前，会校验数据的一致性，校验通过才会写入内存并返回成功

    只支持一种隔离级别：重复读

    数据一旦提交成功，所有的数据全部持久化存储到TIKV上
    
