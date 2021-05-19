# Mysql
---


---
### explain

-   使用：`explain + SQLStmt`，返回执行计划包含的信息


-   `ID`
    
    1、id如果相同，可以认为是一组，从上到下顺序执行

    2、在所有组中，id值越大，优先级越高，越先执行

-   `select_type`

    1、Simple：简单的查询不包括子查询或者UNION

    2、Primary:查询中如果包含任何复杂的子部分，最外层查询被标记为primary

    3、SubQuery：在select和where列表中包含的子查询

    4、Derived：查询过程中的中间数据临时表

    5、Union：

    6、Union Result:从union表获取的结果集

-   `type`

    system>const>eq_ref>ref>range>index>all

    system:表只有一行记录（等于系统表）。这是const类型的特例

    const:表示通过索引一次就找到，const用于比较primary key 或者unique索引，因为只匹配一行数据，
    所以很快将主键置于where列表中，mysql就能将该查询转换为一个常量
    
    eq_ref:唯一性索引扫描，对于每个索引键，表中只有一条记录与之匹配。
    常见于主键或唯一索引扫描
    
    ref:非唯一性索引扫描，返回匹配某个单位值得所有行，所以结果可能是多个符合条件的行，属于查找和扫描的混合体
    
    range:值检索给定范围的行

    index:full index scan，只遍历索引树

    all:full table scan

-   `possible_keys`: 显示可能应用在这张表中的索引，一个或者多个
    查询涉及到的字段上若存在索引，则该索引将被列出，但不一定被查询实际使用
    
-   `key` :实际中使用的索引

-   `key_len`:表示索引的最大可用长度，并非实际中的使用长度,同样的查询条件下，使用的精度越小越好

-   `ref`:显示索引的那一列被使用了，如果可能的话，是一个常数，那些列或常量被用于查找索引列上的值

-   `rows`:根据表统计信息及索引选用情况，大致估算出找到所需的记录所需要读取的行数

-   `extra`:
    
    1、Using FileSort:说明mysql会对数据使用一个外部的索引排序，而不是按照表内的索引顺序进行读取，是十分危险的

    2、Using Temporary:使用了临时表保存中间结果，mysql在对查询结果排序时使用临时表，很增加负担

    3、Using Index: 表示相应的select操作中使用了覆盖索引，避免访问了表的数据行，效率不错
    
    3.1、如果同时出现using where，表明索引被用来执行索引键值的查找

    3.2、如果没有出现，表明索引用来读取数据而非执行查找动作

    4、Using Where:表明使用了where过滤

    5、Using join buffer:使用了连接缓存

    6、impossible where:where的结果是false

    7、select tables optimized away：

    8、distinct:


