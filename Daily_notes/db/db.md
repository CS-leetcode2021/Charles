# DB

-----------------------------------------------------------------------
### mysql
Ubuntu20.04   &  mysql8.0

[Linux安装mysql](https://blog.csdn.net/Thanlon/article/details/100436317)
```
    user     = root
    password = 123456
```
-----------------------------------------------------------------------

### redis

[Linux安装redis](https://blog.csdn.net/hzlarm/article/details/99432240)
```
    redis-server:/etc/init.d/redis-server restart
    redis-cli: redis-cli 
```
服务器后台常开,全局无密码启动

-----------------------------------------------------------------------

### gorocksdb

[CentOS安装（已解决）](https://studygolang.com/articles/28236?fr=sidebar)

[Ubuntu安装](https://blog.csdn.net/taroyoven/article/details/88813386)

#### Issues：部分教程是有错误的。
1. rocksdb最新要求版本是5.16+，这里使用的是最新版的rocksdb master分支
2. 可以使用动态库编译也可以使用静态库编译，编译之前使用make clean清理一下残缺内容，再进行编译
3. 进行 portable=1 make shared_lib 或者portable=1 make static_lib, 
   然后 INSTALL_PATH=/usr/local make install_shared / install_static
	注意这里安装的路径是在/usr/local的include和lib路径下
    
    这里不需要进行全局变量设定
	一个是include导入包的库，一个是lib的依赖库
	// 这里不需要进行全局变量设定

[关于CGO的使用细节问题](../go/CGO_tool.md)
```
	CGO_CFLAGS="-I/usr/local/include"
	CGO_LDFLAGS="-L/usr/local/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" 
	go get github.com/tecbot/gorocksdb
```
4. 全局变量一定要配置rocksdb的路径,和CGO的路径一样
   因为rocksdb的安装路径在/usr/local/include和/usr/local/lib下
   
   配置环境
```
	export CPLUS_INCLUDE_PATH=${CPLUS_INCLUDE_PATH}:/usr/local/include
	export LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib
	export LIBRARY_PATH=${LIBRARY_PATH}:/usr/local/lib
```

5. 这里配置的全是全局环境PATH,在~/.bashrc文件中直接修改，然后source ~/.bashrc
	/etc/profile 文件中也可以修改，此次没有修改
-----------------------------------------------------------------------

### rocksdb

关于rocksdb中文文档：官方主要提供的API是基于C++和java两种，go语言可以通过CGO的方式实现调用rocksdb的API

后期设想是否可以使用calcite直接调用rocksdb的java_API实现sql的转化

[rocksdb官方网站](https://rocksdb.org.cn/doc/Direct-IO.html)

[rocksdb的基本使用方法](https://www.cnblogs.com/wanshuafe/p/11564148.html)

-----------------------------------------------------------------------

### myrocksdb

[CentOS安装myrocksdb（已解决）](https://blog.csdn.net/tdhxhzhz/article/details/107408416)
-----------------------------------------------------------------------
### sqlite

-   全局可以使用
-   启动命令：`sqlite3`；`.exit`退出。
-----------------------------------------------------------------------
