# 概述

转化目标：通过一个全局有序的分布式 Key-Value 引擎来实现SQL语句的DDL、DQL、DML操作。

对于一个 Table 来说，需要存储的数据包括三部分：
- 表的元信息
- Table 中的 Row
- 索引数据

下面详细介绍这三部分的设计与实现（以下设计**不考虑并发**）。

# 元信息

## 元信息储存方案
```
Meta structure:
	NextGlobalID -> int64
	DBs -> {
		DB:1 -> db meta data []byte
		DB:2 -> db meta data []byte
	}
	DB:1 -> {
		Table:1 -> table meta data []byte
		Table:2 -> table meta data []byte
		TID:1 -> int64
		TID:2 -> int64
	}
```

字段解释：

- NextGlobalID：用来记录为下一个Table生成的ID（给autoid Rebase用）。

- DBs：记录有哪些Database，储存Database的元信息，将[DBID]DBMetaData的Hash转化为key进行存储，value内容包括：id、db_name

- DB：记录Database中Tables的元信息，将```[TableID]base```的Hash转化为key进行存储，TID用来记录为column生成id的base，将```[TableID]TableMetaData```的Hash转化为key进行存储，value以json格式来记录数据，包括：id表的全局ID、name表名、clos列的元信息，其中clos列的元信息包括：id、列名、数据格式。

### 具体转化为键的过程

```go
var (
	mMetaPrefix       = []byte("m")
	mNextGlobalIDKey  = []byte("NextGlobalID")
	mDBs              = []byte("DBs")
	mDBPrefix         = "DB"
	mTablePrefix      = "Table"
	mTableIDPrefix    = "TID"
	mBootstrapKey     = []byte("BootstrapKey")
)
```
- `mMetaPrefix` 为所有元数据key的前缀
- `mDBs` 用来标记所有 Database 的元信息储存
- `mDBPrefix` 用来标记每个 Database 的信息储存
- `mTablePrefix` 用来标记Table的信息储存
- `mTableIDPrefix` 用来标记TID的信息储存
- `mBootstrapKey` 用来标记数据库是否进行过初始化

通过以上变量来组合形成key，DBs为例，形成的key为：
- `mDBsH` uint64 储存 Hash 的个数
- `mDBshDB:{db_id}` 储存Table的元信息

其中的 `'H'` 、与 `'h'` 是用来表示储存的是hash的元信息还是数据

```go
const (
	// HashMeta is the flag for hash meta.
	HashMeta TypeFlag = 'H'
	// HashData is the flag for hash data.
	HashData TypeFlag = 'h'
)
```

#### Memcomparable

所有 key 均要进行对应的 Encode 方法来进行 memcomparable 转化，memcomparable 转换意味着编码前的原始类型比较结果和编码成 byte 数组后的比较结果保持一致，这样能够缩小很多查询使用的key的范围从而提高效率。

### 系统表

information-schema等系统表本身实际是调用查询元信息的API，并不与普通表一样按照下面的转化进行储存。

# Row

## 行存
- 首先定义key各个部分的前缀
```go
var(
	tablePrefix     = []byte{'t'}
	recordPrefixSep = []byte("r")
	indexPrefixSep  = []byte("i")
)
```
- 每行数据按照如下规则进行编码成 Key-Value pair：
```
Key: tablePrefix{tableID}_recordPrefixSep{rowID}
Value: [col1, col2, col3, col4]
```

# 索引

- Unique Index 按照如下规则编码成 Key-Value pair：
```
Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue
Value: rowID
```

- 非 Unique Index 按照如下规则编码成 Key-Value pair：
```
Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue_rowID
Value: null
```

# 参考资料
- [三篇文章了解 TiDB 技术内幕 - 说计算](https://pingcap.com/blog-cn/tidb-internal-2/)
- [TiDB系统表 mysql Schema](https://docs.pingcap.com/zh/tidb/stable/mysql-schema)
- [TiDB系统表 Information Schema](https://docs.pingcap.com/zh/tidb/stable/information-schema)
