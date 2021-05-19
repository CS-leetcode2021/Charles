# SQLite
---
-   [参考链接](https://www.runoob.com/sqlite/sqlite-intro.html)
-   [与mysql的区别](./sqlite_VS_mysql.md)
---
## 使用
-   [使用](../db/db.md)
-   `.help`查看点命令所用清单
-   SQLite不区分大小写
-   [SQLite 语句](https://www.runoob.com/sqlite/sqlite-syntax.html)
## 基础语法：
-   创建数据库是在当前目录下直接创建：命令`sqlite3 databases_name.db`
```sqlite
   >sqlite3 testDB.db
   SQLite version SQLite version 3.34.0 2020-12-01 16:14:00
   Enter ".help" for instructions
   sqlite>               
```
-   `.dump`命令,实现数据库的备份和回复
```sqlite
    //将数据库数据备份到testDB.sql
   $sqlite3 testDB.db .dump > testDB.sql 
    //从备份数据恢复到数据库数据
   $sqlite3 testDB.db < testDB.sql
```
-   SQLite 附加数据库:`ATTACH DATABASE file_name AS database_name;`
-   SQLite 分离数据库：`DETACH DATABASE 'Alias-Name';`
-   创建表、删除表和MySQL用法基本一样，查看当前数据表使用：`.tables`
-   增删该差和mysql基本一样,where、order by、limit、like、and/or、group by、distinct、having基本一致
--- 

## SQL高级用法
-   join、index、alter、explain、view、transaction
---

## SQL对外接口

-   当下接口只要是对C/C++、Java、Python、PHP