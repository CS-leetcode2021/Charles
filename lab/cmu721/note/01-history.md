# History
---

1、 忽略分布式数据库，只讨论基于单节点的内存数据库，讨论最先进的主题，不是一个关于DBMS经典的课程

    for this course we are gonna focus on getting the single node system working
    as fast as possible and it is correct and correctly and what we don't worry 
    about going to stripping it right.
2、 我们主要关注的问题：

    Concurrency Control     并发处理
    Indexing                索引
    Storage Models, Compression     存储模型、压缩
    Parallel Join Algorithms        并行连接算法
    Networking Protocols            网络协议
    Logging & Recovery Methods      日志、恢复
    Query Optimization, Execution, Compilation  查找优化、执行、汇编

3、 所要求的知识：

    two-phase locking
    B plus trees

4、 两个project:

    Project #1 will be completed individually.
    Project #2 will be done in a group of three.

5、 History of databases:

    The SQL vs. NoSQL debate is reminiscent of
    Relational vs. CODASYL debate from the 1970s.

    1960s- IDS:
    Integrated Data Store   集中数据库
    Network data model.Tuple-at-a-time queries.网络数据模型、元组查询

    1960s- CODASYL
    trying to define a standard API

    1960S- IBM IMS
    Hierarchical data model.        层次模型
    Programmer-defined physical storage format. 可定义的物理存储格式
    Tuple-at-a-time queries.        元组
    
    Quesition：Duplicate Data - No Indepandence 
    because if multiple suppliers provide the same 
    I have to have different instances of that you know that record in this case large batteries
    over and over again
    right for every single supplier that sells that part and that means now if the name of the part
    ever change I have to write extra code to ho find all instances of the batteries the large batteries
    and make sure that theit names all get changed together at the same time and that they're ther in
    sync the other big problem in addition to sort of the tuple at a time for queries is that there 
    was no independence between the physical data structure of the databases of the database and the 
    logical abstraction that programmers interacted with so

    1970s- Relational Model
    Database abstraction to avoid this maintenance:
    → Store database in simple data structures.
    → Access data through high-level language.
    → Physical storage left up to implementation.

[Derivability, redundancy and consistency of relations stored in large data banks](https://dl.acm.org/doi/10.1145/1558334.1558336)

[A relational model of data for large shared data banks](https://dl.acm.org/doi/10.1145/362384.362685)

    1980s- Relational Model
    The relational model wins.
    → IBM comes out with DB2 in 1983.
    → “SEQUEL” becomes the standard (SQL).
    
    Many new “enterprise” DBMSs but Oracle wins marketplace.
    Stonebraker creates Postgres.
    
    1980s- Object-Orienred Data Bases
    Few of these original DBMSs from the 1980s still exist today but many of the technologies exist in
    other forms (JSON, XML)
    类似于Java结构体，都是按照结构体封装存储的

    Questions:Complex Queries、No Standard API
    如果仅仅是single Object 很轻松，但是是多个Object查询会导致复杂程度
    No API people use to interact with these object on the databases MongoDB or Redis

    2000s- INTERNET BOOM
    All the big players were heavyweight and expensive. Open-source databases were missing important features.
    Many companies wrote their own custom middleware to scale out database across single-node DBMS instances.

    2000s- Data WareHouses（OLAP-数据仓库、数据分析）
    → Distributed / Shared-Nothing
    → Relational / SQL
    → Usually closed-source.

    2000s- NoSQL Systems（放弃ACID，事务）
    Focus on high-availability & high-scalability:
    Schemaless (i.e., “Schema Last”)
    Non-relational data models (document, key/value, etc)
    No ACID transactions
    Custom APIs instead of SQL
    Usually open-source

    2010s- NewSql
    Provide same performance for OLTP workloads as
    NoSQL DBMSs without giving up ACID:     (NewSQL并没有放弃事务和ACID)
    → Relational / SQL
    → Distributed
    → Usually closed-source

    2010s-HYBRID SYSTEMS
    Hybrid Transactional-Analytical Processing.
    Execute fast OLTP like a NewSQL system while
    also executing complex OLAP queries like a data
    warehouse system.
    → Distributed / Shared-Nothing
    → Relational / SQL
    → Mixed open/closed-source.

    2010s-CLOUD SYSTEMS
    First database-as-a-service (DBaaS) offerings were
    "containerized" versions of existing DBMSs.
    There are new DBMSs that are designed from
    scratch explicitly for running in a cloud
    environment.

    2010s-SHARED-DISK ENGINES
    Instead of writing a custom storage manager, the
    DBMS leverages distributed storage.
    → Scale execution layer independently of storage.
    → Favors log-structured approaches.
    This is what most people think of when they talk
    about a data lake.

    2010s-GRAPH SYSTEMS
    Systems for storing and querying graph data.
    Their main advantage over other data models is to
    provide a graph-centric query API
    → Recent research demonstrated that is unclear whether
    there is any benefit to using a graph-centric execution
    engine and storage manager.

    2010s-SPECIALIZED SYSTEMS
    Embedded DBMSs
    Multi-Model DBMSs
    Blockchain DBMSs
    Hardware Acceleration

    The demarcation lines of DBMS categories will
    continue to blur over time as specialized systems
    expand the scope of their domains.
    I believe that the relational model and declarative
    query languages promote better data engineering.