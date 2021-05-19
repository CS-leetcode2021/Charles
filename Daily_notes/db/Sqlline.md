# SQLline  
---
- [资料链接](http://sqlline.sourceforge.net/#sect_command_all)
- [参考链接](https://apacheignite-sql.readme.io/docs/sqlline)
- [github源代码链接](https://github.com/julianhyde/sqlline)
-   用于进行数据库的链接：
    >先通过jdbc加载数据库驱动，进行数据库链接，输入sql语言进行数据库操作
---
### 介绍
-   >SQLLine是一个基于纯java控制台的实用程序，用于连接到关系数据库和执行SQL命令。它类似于其他命令行数据库访问程序，如sqlplus用于Oracle, mysql用于mysql, isql用于Sybase/SQL Server。因为它是纯Java的，所以它是平台独立的，并且可以运行在任何可以运行Java 1.3或更高版本的平台上
-   相关的命令介绍可以直接看官方文档
  
-   Q:为什么不直接在使用的数据库上用对应官方提供的平台直接操作数据库，例如mysql—client操作MySql？

    >对应官方提供的操作平台只能操作一个对应的数据库，并不能灵活的切换到多个数据库链接

-   Q:为什么不能单单使用JDBC进行数据库的驱动、链接、操作
    
    >JDBC有对应的client端，直接对底层进行操作，为什么不可？
    
    >个人理解：可能因为JDBC只能仅仅支持内部写好的数据库驱动？而sqlline可以更好的支持自定义模式，通过JDBC进行数据库链接并操作，这样可以更加灵活...