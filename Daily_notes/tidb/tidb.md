### tidb

-----------------------------------------------------------------------
[pd-server命令行参数解析](https://docs.pingcap.com/zh/tidb/v3.0/command-line-flags-for-pd-configuration)

[tikv-server命令行参数解析](https://docs.pingcap.com/zh/tidb/v3.0/command-line-flags-for-tikv-configuration)

[tidb-server命令行参数解析](https://docs.pingcap.com/zh/tidb/v3.0/command-line-flags-for-tidb-configuration)


#### Binary的单节点部署（已成功）

直接下载官网的安装包，一键式安装启动

-   下载安装包：
```
	wget http://download.pingcap.org/tidb-latest-linux-amd64.tar.gz
	wget http://download.pingcap.org/tidb-latest-linux-amd64.sha256
```
-   检查文件完整性，返回ok则正确：
```
	sha256sum -c tidb-latest-linux-amd64.sha256
```
-   解开压缩包：
```
	tar -xzf tidb-latest-linux-amd64.tar.gz
	cd tidb-latest-linux-amd64
```
-   tidb分三部分启动，PD、Tikv、TiDB
   
-   PD启动：
```
	./bin/pd-server --data-dir=pd --log-file=pd.log &
```
-   Tikv启动：
```
	./bin/tikv-server --pd="127.0.0.1:2379" --data-dir=tikv --log-file=tikv.log &
```
-   TiDB启动：
```
	./bin/tidb-server --store=tikv --path="127.0.0.1:2379" --log-file=tidb.log &
```
-   使用mysql连接Tidb（已成功）
```
	mysql -h 127.0.0.1 -P 4000 -u root -D test 
```

#### tidb 源码手动编译启动（已解决）部署三个KV节点
[参考文章](https://blog.csdn.net/damanchen/article/details/108223095)

-   TiDB
```
	git clone tidb...（这里不是官方的源码，已添加rocksdb） & cd tidb
	make     // 直接进行编译，显示Build TiDB Server successfully!
```
-   PD
``` 
	git clone https://github.com/tikv/pd.git
	make
```
- 进行pd编译的时候，go语言自带的io.fs package出现问题，导致无法编译，更新go -version到1.16.x可以解决
-   TiKV
```
	git clone https://github.com/tikv/tikv.git
```
[需要先安装rust然后再编译](../linux/linux.md)tikv因为是rust语言写的,然后 make，时间比较久


-   deploy（部署）参考Tiup的启动方式
```
	// PD启动：
	mkdir pd-1
	./bin/pd-server --name=pd-1 
		--data-dir=./pd-1/data 
		--peer-urls=http://127.0.0.1:2380 
		--advertise-peer-urls=http://127.0.0.1:2380 
		--client-urls=http://127.0.0.1:2379 
		--advertise-client-urls=http://127.0.0.1:2379 
		--log-file=./pd-1/pd.log 
		--initial-cluster=pd-1=http://127.0.0.1:2380 &
```
	
```	//Tikv启动：
	mkdir tikv-1
	./target/release/tikv-server 
		--addr=127.0.0.1:20160 
		--advertise-addr=127.0.0.1:20160 
		--status-addr=127.0.0.1:20180 
		--pd=http://127.0.0.1:2379 
		--config=./tikv-1/tikv.toml 
		--data-dir=./tikv-1/data 
		--log-file=./tikv-1/tikv.log &
```	
        *****
        文件: tikv.toml
            [rocksdb]
            max-open-files = 256
            [raftdb]
            max-open-files = 256
            [storage]
            reserve-space = 0
        *****
```
	//TiDB启动：
	mkdir tidb-1
	./bin/tidb-server -P 4000 
		--store=tikv 
		--host=127.0.0.1 
		--status=10080 
		--path=127.0.0.1:2379 
		--log-file=./tidb-1/tidb.log &
```
-   连接mysql测试（已成功）
```
    mysql -h127.0.0.1 -P4000 -uroot
```

#### tidb 源码切换rocksdb存储引擎（已解决）

-   pd启动
```
    ./bin/pd-server --name=pd-2 --data-dir=./pd-2/data --log-file=./pd-2/pd.log &
```	
-   rocksdb启动 ([已经配置完全](../db/db.md))纯内存引擎的TiDB不需要配置PATH,默认数据存放路径在 /tmp/tidb
-   Tidb启动
```
    ./bin/pd-server -P 4000 
			--store=rocksstore
			--host=127.0.0.1
			--status=10080
			--log-file=./tidb-2/tidb.log &
```
-   连接mysql(已成功)
```
    mysql -h127.0.0.1 -P4000 -uroot
```
-----------------------------------------------------------------------
