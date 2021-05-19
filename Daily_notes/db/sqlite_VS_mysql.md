# SQLite与Mysql的区别
---
-   >数据库是任何网站或开发设计的核心部分， SQLite和MySQL都是主流的开源数据库。数据库管理系统也称为DBMS，通常称为RDBMS。RDBMS代表关系数据库管理系统，在最基本的层次上，所有数据库都用于管理、维护和操作数据。
---
### 数据模型可以是以下两种之一：

-   >NoSQL: 一个非结构化的，仍在发展中的模型

-   >Relational DBMS:  一种结构化的、更常用的模型

### 什么是数据库管理系统？
-   DBMS是与数据库交互的软件，它有助于对其管理的数据库进行多项操作，其中一些操作是：
    
    >管理数据库的访问权限:运行SQL查询;将数据写入数据库;更新数据库;从数据库中删除数据
    
    >在大多数情况下，数据库和DBMS可以互换使用。但是，数据库是存储数据的集合，而DBMS是用于访问数据库的软件。每个DBMS都有一个底层模型，它决定数据库的结构以及如何检索数据。Relational DBMS使用Relational数据模型，这里的数据以表格的形式组织。每个表都有一组属性或列，每一行也称为元组有一个关系。由于这个原因，结构化被称为RDBMS。要使用RDBMS，您必须使用SQL或结构化查询语言，每个RDBMS都有不同的语法

### SQLite与MySQL的区别
-   SQLite和MySQL都是开源的RDBMS
-   架构差异 – SQLite与MySQL
    >SQLite是一个在公共领域中可用的开源项目;MySQL是一个由Oracle拥有的开源项目
    
    >SQLite是一个无服务器的数据库，是自包含的。这也称为嵌入式数据库，这意味着数据库引擎作为应用程序的一部分运行。另一方面，MySQL需要运行服务器，MySQL将需要客户端和服务器架构通过网络进行交互。

-   数据类型支持 – SQLite与MySQL
    >SQLite支持以下数据类型：Blob，Integer，Null，Text，Real。

    >MySQL支持下面提到的数据类型：Tinyint, Smallint, Mediumint, Int, Bigint, Double, Float, Real, Decimal, Double precision, Numeric, Timestamp, Date, Datetime, Char, Varchar, Year, Tinytext, Tinyblob, Blob, Text, MediumBlob, MediumText, Enum, Set, Longblob, Longtext.
MySQL在数据类型方面更加灵活

-   存储和可移植性 – SQLite与MySQL

    >SQLite库大小约为250 KB，而MySQL服务器大约为600 MB。SQLite直接将信息存储在单个文件中，使其易于复制。不需要任何配置，并且可以使用最少的支持来完成该过程

    >在复制或导出MySQL之前，您需要将其压缩为单个文件。对于较大的数据库，这将是一项耗时的活动

-   多种访问和可伸缩性 – SQLite与MySQL

    >SQLite没有任何特定的用户管理功能，因此不适合多用户访问。MySQL有一个构造良好的用户管理系统，可以处理多个用户并授予不同级别的权限

    >SQLite适用于较小的数据库，随着数据库的增长，使用SQLite时内存需求也会变大。使用SQLite时，性能优化更加困难。相反，MySQL易于扩展，可以轻松处理更大的数据库

-   安全性和易于安装 – SQLite与MySQL
    >SQLite没有内置的身份验证机制，任何人都可以访问数据库文件。但是，MySQL带有许多内置的安全功能。这包括使用用户名，密码和SSH进行身份验证

    >SQLite不需要太多配置，易于设置。与SQLite相比，MySQL需要更多配置

### 优点和缺点 – SQLite与MySQL

-   SQLite的优点：

    >基于文件，易于设置和使用;适合基础开发和测试;轻松携带;使用标准SQL语法进行微小更改;使用方便

-   SQLite的缺点：

    >缺乏用户管理和安全功能;不容易扩展;不适合大数据;无法定制

-   MySQL的优点：

    >使用方便;提供了许多与数据库相关的功能;良好的安全功能;易于扩展，适用于大型数据库;提供良好的速度和性能;提供良好的用户管理和多种访问控制

-   MySQL的缺点：

    >需要一些技术专业知识来设置;与传统SQL相比，语法略有不同
---                 