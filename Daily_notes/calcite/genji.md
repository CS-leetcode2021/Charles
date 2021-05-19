# genji
-   基于key-value存储的数据库，用纯go语言所写
-    Genji relies on [BoltDB](https://github.com/etcd-io/bbolt.git) and [Badger](https://github.com/dgraph-io/badger.git) to manage data.
-   *BoltDB* 是一个fork[bolt](https://github.com/boltdb/bolt.git)的一个改进过的k-v数据库
-   *Badger* 是一个纯go语言写的基于k-v操作的数据库
-   [官网文档](https://genji.dev/docs/genji-sql/)


# learning
- db_test.go
```go
    // 打开一个链接  
    // open做判断是否是默认的 boltengine.NewEngine，还是memoryengine.NewEngine
    db, err := genji.Open(":memory:") 
    


```