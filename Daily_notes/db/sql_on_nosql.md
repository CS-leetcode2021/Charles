# 项目当下的设计

---

## 原因背景

1、 nosql数据基本是无预订模式的，对结构化的数据存取并不友好
2、 SQL语句的优势和强大的查找功能
3、 K-V的特点：操作方式类似于map，具有查找速度快，模型简单，但不支持SQL语言，B+tree的mysql（不支持高并发），以及以LSM-tree的levelDB、rocksdb（是基于leveldb开发的），CockroachDB底层基于rocksdb实现的，redis是基于缓存查询实现的
4、 国内外现状：google基于Spanner基础实现了F1数据库
    pingcap基于rocksDB构建了处理引擎TiKV，兼容了mysql协议以及在线处理TiDB
    Cockroach Labs基于RocksDB构建了使用PostgreSQL 协议的CockroachDB
    Phoenix也是基于HBase构建了一套SQL引擎
    Facebook在mysql的基础上基于rocksDB实现了MyRocks
5、 Gson：google发布的用于java序列化的java库
6、目前键值对存储系统的底层实现主要有三种算法：B+树、LSM（读性能低，避免频繁的io）（Big Table、LevelDB、RocksDB）、log & index（内存中进行索引，牺牲了范围查找的读性能）
7、查询编译器首先对查询进行语法分析，亦即将查询语句转换成按某种方式表示查询语句结构的语法树，之后根据逻辑优化将语法分析树转换成关系代数表达式树（逻辑查询计划），最后利用元数据和数据统计信息来进行物理优化来将逻辑查询计划转换成物理查询计划。物理查询计划不仅指明了要执行的操作，而且也找出了这些操作执行的顺序、执行每步所用的算法、获得所存储数据的方式以及数据从一个操作传递给另一个操作的方式。

8、查询编译器包括三个主要的部分，查询分析器、查询预处理器、查询优化器（初始计划转化为逻辑计划和屋里计划）
9、执行引擎：负责执行上一步中最终形成的物理查询计划的每一步。
    需要与其它各部分进行交互，与调度器进行交互，避免访问加锁的数据；与日志管理器交互，确保所有数据操作的一致性
10、 存储的信息：数据、元数据、日志记录、统计信息、索引
11、 RocksDB通过JNI实现RocksDB的Java API
12、 TiDB的逻辑框架：TiDB :MySQL Protocol Server---》SQL Layer
     TiKV的逻辑框架：Transaction、MVCC、Raft、Local KV Storage（RocksDB）
13、 元信息：GloablID_DBID_TableID
     行信息：key:tablePrefix_TableID_recordFix_RowID; value:  (c1,c2,c3)
     index信息：分 主键和非主键 key：tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue value：rowID
14、 Calcite:可以以使用传统SQL数据库的方式完成对不同底层的数据操作
15、 Calcite使用JavaCC生成词法解析器和语法解析器，SQL语句经过转化后生成SQL Node———》RelNode（用于表达逻辑表达式）———》Expression(Node)———》通过Jinano Compiler转化为字节码执行，这就是Calcite的三次解析一次编译执行过程