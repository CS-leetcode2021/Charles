# Mysql

## -2020-12-07-记录
### 一条sql的执行过程
-   主要讲解了mysql进行sql语言操作时的流程
-   [参考文章](https://zhuanlan.zhihu.com/p/50597960)
#### server层主要包含以下部分：
-   连接器——缓存（缓存功能可以进行自定义设置）——sql语言分析器——sql语言优化器——执行器
-   解析：
    >sql语言分析器主要功能是进行输入的sql语言进行关键字提取

    >sql语言优化器主要功能是对分析器提取的相关表达式进行等价转换，如是否存在匹配的索引和执行代价相对较小的计划

    >执行器主要功能是对优化器形成的计划进行执行，调用存储引擎进行sql语句操作


#### 存储引擎层：Mysql使用的是InnoDB
-   undo log 与MVCC
    >undo log是记录每行数据事务执行前的数据


    >MVCC是保证多版本并发控制实现，保证事务隔离级别的读已提交和读未提交级别

-   redo log 与 Buffer Pool
    >redo log 是InnoDB内部的一个缓冲池，大小固定采用循环写

    >Buffer Pool 
    [参考文章](https://www.cnblogs.com/mengxinJ/p/14071262.html)

-   bin log(Server 层)
-   >redo log 因为大小固定，所以不能存储过多的数据，它只能用于未更新的数据落盘，而数据操作的备份恢复、以及主从复制是靠 bin log
  
    >[参考文章](https://zhuanlan.zhihu.com/p/50597960)

    >主要分为三种状态：

    >基于SQL语句的复制（statement-based replication，SBR）

    >基于行的复制（row-based replication，RBR)

    >混合模式复制（mixed-based replication,MBR)

#### 三种log的比较（undo、redo、bin）
1. undo log是用于事务的回滚、保证事务隔离级别读已提交、可重复读实现的。redo log是用于对暂不更新到磁盘上的操作进行记录，使得其可以延迟落盘，保证程序的效率。bin log是对数据操作进行备份恢复（并不能依靠 bin log 直接完成数据恢复）。

2. undo log 与 redo log 是存储引擎层的日志，只能在 InnoDB 下使用；而bin log 是 Server 层的日志，可以在任何引擎下使用。

3. redo log 大小有限，超过后会循环写；另外两个大小不会。

4. undo log 记录的是行记录变化前的数据；redo log 记录的是 sql 或者是数据页修改逻辑或 sql（个人理解）；bin log记录的是修改后的行记录（5.7默认）或者sql语句。

### 执行过程
#### where语句的执行
##### 会将where语句拆解成Index Key,Index Filter,Table Filter
-   Index Key 主要确定SQL查询在索引中的连续范围的查询条件，分为First Key和Last Key

    >Index First Key：提取规则：从索引的第一个键值开始，检查其在 where 条件中是否存在，若存在并且条件是 =、>=，则将对应的条件加入Index First Key之中，继续读取索引的下一个键值，使用同样的提取规则；若存在并且条件是 >，则将对应的条件加入 Index First Key 中，同时终止 Index First Key 的提取；若不存在，同样终止 Index First Key 的提取

    >Index Last Key：与 Index First Key 正好相反；提取规则：从索引的第一个键值开始，检查其在 where 条件中是否存在，若存在并且条件是 =、<=，则将对应条件加入到 Index Last Key 中，继续提取索引的下一个键值，使用同样的提取规则；若存在并且条件是 < ，则将条件加入到 Index Last Key 中，同时终止提取；若不存在，同样终止Index Last Key的提取

-   Index Filter 用于索引范围确定后，确定 SQL 中还有哪些条件可以使用索引来过滤
    
    >提取规则：从索引列的第一列开始，检查其在 where 条件中是否存在，若存在并且 where 条件仅为 =，则跳过第一列继续检查索引下一列，下一索引列采取与索引第一列同样的提取规则；若 where 条件为 >=、>、<、<= 其中的几种，则跳过索引第一列，将其余 where 条件中索引相关列全部加入到 Index Filter 之中；若索引第一列的 where 条件包含 =、>=、>、<、<= 之外的条件，则将此条件以及其余 where 条件中索引相关列全部加入到 Index Filter 之中；若第一列不包含查询条件，则将所有索引相关条件均加入到 Index Filter之中

-   Table Filter where 中不能被索引过滤的条件都归为此中；提取规则：所有不属于索引列的查询条件，均归为 Table Filter 之中


-   在5.6 之前，是不分 Table Filter 与 Index Filter 的，这两个条件都直接分配到 Server 层进行筛选。筛选过程是先根据 Index Key 的条件先在引擎层进行初步筛选，然后得到对应的主键值进行回表查询得到初筛的行记录，传入 Server 层进行后续的筛选，在 Server 层的筛选因为没有用到索引所以会进行全表扫描。而索引下推的优化就是将 Index Filter 的条件下推到引擎层，在使用  Index First Key 与 Index Last Key 进行筛选时，就带上 Index Filter 的条件再次筛选，以此来过滤掉不符合条件的记录对应的主键值，减少回表的次数，同时发给 Server 层的记录也会更少，全表扫描筛选的效率也会变高。

-   未使用索引下推

![](https://mmbiz.qpic.cn/mmbiz_png/gjnldtnoHOpMWPWMOySlK82E56xRLtzU45B21uPHAZCGqDUB9ovuOo0Lh9P4q9sbxnliaicPw8wnOgn1fSibzF0oQ/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)


-   使用索引下推

![](https://mmbiz.qpic.cn/mmbiz_png/gjnldtnoHOpMWPWMOySlK82E56xRLtzUiaeXk4HvmA2atic6QNzcL2uxaUBVTswrkEfyPOjwoQlbYhjSsicqDh1UA/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)