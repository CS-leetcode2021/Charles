# Calcite——one size fits all

-   [参考网站](https://lixiyan4633.gitee.io/2020/03/18/calcite%E6%95%99%E7%A8%8B01-calcite%E7%9A%84%E5%9F%BA%E6%9C%AC%E6%A6%82%E5%BF%B5/
)
---
-   >上层尽量简单的封装请求，定义为标准的SQL，中间通过构建JDBC或ODBC来访问Calcite，下层对接数据库
-   >其实Calcite并没有数据库，需要我们在代码里告诉Calcite，虚拟出来的表是什么、字段是什么、字段类型是什么等，整体抽象为一个个Schema，对于我们来说就查Calcite虚拟出来东西，不用关心底层真正对接了哪些数据源，底层通过定义各种adapter，来对接不同的查询和存储引擎（比如es、habse、redis、mysql，甚至是CSV、HDFS等），这也就决定了它其实没有真正的“物理执行计划”
---

### 特点
-   Calcite支持异构数据源查询
-   可以不做SQL解析，不做优化，只要接进来就是可以工作的
-   独立的编程语言和数据源，可以支持不同的前端和后端
-   支持关系代数，可定制的逻辑规划规则和基于成本模型优化的查询引擎
-   支持视图
-   跨数据源查询
-   Calcite本身会缓存Schema、Function等信息（在内存里缓存）
-   支持复合指标计算(a+b as c)、常用聚合函数(sum、count、distinct)、sort、group by、join、limit等 
-   架构比较精简，利用Calcite写几百行代码就可以实现一个SQL查询方案
-   灵活绑定优化规则，对于一个条件，我们可以自定义多个优化规则，只要命中，可以根据不同的规则多次优化同一个查询条件 

---
### Calcite的架构
-   关于calcite的架构，相对于传统数据库，他将数据存储和数据处理算法和元信息存储忽略了，这样使得calcite更加兼容，适合统一查询服务。
![架构](https://tva1.sinaimg.cn/large/00831rSTly1gcyeie5z1zj30il0fqgom.jpg)

-   rules:也就是匹配规则，calcite目前内置上百种rule来优化relations expression,也支持自定义的rules。
-   Metadata providers:主要向优化器提供信息,有助于指导优化器向着目标（减少整体cost）进行优化，信息可以包括行数，table哪一列是唯一列，也包括计算relNode树，执行subexpression cost的函数。

-   Planner engines:触发rules来达到指定目标，比如cost-base expression optimizer(CBO)的目标是减少cost（减少数据行数，CPU cost,IO cost等）。
  
-   对SQL执行完整流程，分为4部分:

![](https://tva1.sinaimg.cn/large/00831rSTly1gcyeiiozxej30xc02qjsa.jpg)


-   >Parser: 解析，calcite通过javaCC将SQL解析成未经校验的AST(Abstract Syntax Tree,即抽象语法树)。
-   >Validate: 验证，主要负责校验上一步parser是否合法，如验证SQL的schema信息，字段，函数等是否存在，SQL是否合法，此步骤完成后生成RealNode。
-   >Optimize：优化，也是整个核心所在，将逻辑计划RelNode转化为物理计划，主要涉及SQL优化：基于规则优化（RBO），基于代价优化（CBO）,optimizer按道理是可以选的。此步骤得到物理计划。
-   >Execute：执行阶段，将物理计划转成程序，变成自己平台的可执行代码。

---
### catalog
-   Schema
    >主要定义schema与表的结合，有点像database，但是sechma不一定是强制的，比如同名表T1就需要加上A.T1,B.T1。
```java
    public interface Schema {

        Table getTable(String name);

        Set<String> getTableNames();

        RelProtoDataType getType(String name);

        Set<String> getTypeNames();

        Collection<Function> getFunctions(String name);

        Set<String> getFunctionNames();

        Schema getSubSchema(String name);

        Set<String> getSubSchemaNames();

        Expression getExpression(SchemaPlus parentSchema, String name);

        boolean isMutable();

        Schema snapshot(SchemaVersion version);

        /** Table type. */
        enum TableType {}
    }
```

-   Table
    >对应数据库的表。
```java
    public interface Table {
        RelDataType getRowType(RelDataTypeFactory typeFactory);

        Statistic getStatistic();

        Schema.TableType getJdbcTableType();

        boolean isRolledUp(String column);

        boolean rolledUpColumnValidInsideAgg(String column, SqlCall call,
            SqlNode parent, CalciteConnectionConfig config);
    }
```
-   RelDataType
    >代表表的数据定义，对应表的数据列名和类型，RelDataType代表数据表的行Row的数据类型，Statistic用于统计表的相关数据，特别用于CBO用于代价计算。

	>eg：`selcct id, name, cast(age as bigint) from A.INFO`
	
```java
    Data type field: id,name;	
    data type: bigint;	
    A为Schema;	
    INFO:table。
```
---
### Table
-   ScannableTable
    >a simple implementation of Table, using the ScannableTable interface, that enumerates all rows directly 

    >这种方式基本不会用，原因是查询数据库的时候没有任何条件限制，默认会先把全部数据拉到内存，然后再根据filter条件在内存中过滤。使用方式：实现Enumerable scan(DataContext root)，该函数返回Enumerable对象，通过该对象可以一行行的获取这个Table的全部数据。

-   FilterableTable
    >a more advanced implementation that implements FilterableTable, and can filter out rows according to simple predicates
	
    >初级用法，我们能拿到filter条件，即能再查询底层DB时进行一部分的数据过滤，一般开始介入calcite可以用这种方式（translatable方式学习成本较高）。
	使用方式：实现Enumerable scan(DataContext root, List filters )。
   
   
-   TranslatableTable
    >advanced implementation of Table, using TranslatableTable, that translates to relational operators using planner rules.
	
    >高阶用法，有些查询用上面的方式都支持不了或支持的不好，比如join、聚合、或对于select的字段筛选等，需要用这种方式来支持，好处是可以支持更全的功能，代价是所有的解析都要自己写，“承上启下”，上面解析sql的各个部件，下面要根据不同的DB（es\mysql\drudi..）来写不同的语法查询。
	
    >当使用ScannableTable的时候，我们只需要实现函数Enumerable scan(DataContext root)，该函数返回Enumerable对象，通过该对象可以一行行的获取这个Table的全部数据（也就意味着每次的查询都是扫描这个表的数据，我们干涉不了任何执行过程）
    
    >当使用FilterableTable的时候，我们需要实现函数Enumerable scan(DataContext root, Listfilters )，参数中多了filters数组，这个数据包含了针对这个表的过滤条件，这样我们根据过滤条件只返回过滤之后的行，减少上层进行其它运算的数据集
    
    >当使用TranslatableTable的时候，我们需要实现RelNode toRel( RelOptTable.ToRelContext context, RelOptTable relOptTable)，该函数可以让我们根据上下文自己定义表扫描的物理执行计划
    
    >至于为什么不在返回一个Enumerable对象了，因为上面两种其实使用的是默认的执行计划，转换成EnumerableTableAccessRel算子，通过TranslatableTable我们可以实现自定义的算子，以及执行一些其他的rule，Kylin就是使用这个类型的Table实现查询。

---

### SQL Parse(SQL -> SQL Node)
-   Calcite使用javaCC作为SQL解析，javaCC通过Calcite中定义的Parse.jj文件，生成一系列的java代码，生成的代码会把SQL转换成AST的数据结构（SQL Node类型）
	
    >补充：javaCC主要有两点功能：
	
    >1、设计词法和语义，定义SQL中的具体元素
	
    >2、实现词法分析器（Lexer）和语法分析器（Parse）,完成对SQL的解析，完成相应的转换
		
```java
        //org.apache.calcite.prepare.CalcitePrepareImpl
        //... 
        //解析sql->SqlNode
        SqlParser parser = createParser(query.sql,  parserConfig);
        SqlNode sqlNode;
        try {
            sqlNode = parser.parseStmt();
            statementType = getStatementType(sqlNode.getKind());
        } catch (SqlParseException e) {
            hrow new RuntimeException(
                "parse failed: " + e.getMessage(), e);
        }
        //...
```
-   >解析：query.sql主要经过语法分析器产生parse，然后经过parseStmt产生一个SqlNode，此结构便是SQL拆解后的Node
    
    >Debug:
![](https://tva1.sinaimg.cn/large/00831rSTly1gcyeioqmloj30ps07pad5.jpg)
---

### 验证和转换 （SQL Node -> RelNode）
-   经过上面的解析，会将SQL翻译成SQL 的抽象语法树 AST,SQL Node就是抽象语法树的节点Node,而Rel代表关系表达式（Relation Expression）,所以从sql node -> rel node 的过程，是一个转换，一个校验的过程。
	
-   >校验Validate：语法检查前我们要先知道元信息，这个检查会包括表名，字段名，字段类型，函数名，进行语法检查的实现如下：
```java
    //org.apache.calcite.sql.validate.SqlValidatorImpl
    public SqlNode validate(SqlNode topNode) {
        SqlValidatorScope scope = new EmptyScope(this);
        scope = new CatalogScope(scope, ImmutableList.of("CATALOG"));
        final SqlNode topNode2 = validateScopedExpression(topNode, scope);
        final RelDataType type = getValidatedNodeType(topNode2);
        Util.discard(type);
        return topNode2;
    }
```
-   >在对SQLnode的校验过程中，会对AST的语法进行补充，eg：
```java
    //校验前
    select province,avg(cast(age as double)) from t_user group by province
    //校验后
    SELECT `T_USER`.`PROVINCE`, AVG(CAST(`T_USER`.`AGE` AS DOUBLE))
    FROM `TEST_PRO`.`T_USER` AS `T_USER`
    GROUP BY `T_USER`.`PROVINCE`
```	
-   >SQL Node->Rel Node 在经过上面进行合法的补充了一定语义的AST树之后，而validate中已经准备好各种映射信息后，就开始具体的RelNode的转换了。
```java
    //org.apache.calcite.sql2rel.SqlToRelConverter
    protected RelRoot convertQueryRecursive(SqlNode query, boolean top,
        RelDataType targetRowType) {
    final SqlKind kind = query.getKind();
    switch (kind) {
        case SELECT:
            return RelRoot.of(convertSelect((SqlSelect) query, top), kind);
        case INSERT:
            return RelRoot.of(convertInsert((SqlInsert) query), kind);
        case DELETE:
            return RelRoot.of(convertDelete((SqlDelete) query), kind);
        case UPDATE:
            return RelRoot.of(convertUpdate((SqlUpdate) query), kind);
        case MERGE:
            return RelRoot.of(convertMerge((SqlMerge) query), kind);
        case UNION:
        case INTERSECT:
        case EXCEPT:
            return RelRoot.of(convertSetOp((SqlCall) query), kind);
        case WITH:
            return convertWith((SqlWith) query, top);
        case VALUES:
            return RelRoot.of(convertValues((SqlCall) query, targetRowType), kind);
        default:
            throw new AssertionError("not a query: " + query);
        }
    }
```
-   >测试:
	`select province,avg(cast(age as double)) from t_user group by province`
	   
    >结果
```java
    Plan after converting SqlNode to RelNode
		LogicalAggregate(group=[{0}], EXPR$1=[AVG($1)])
		  LogicalProject(PROVINCE=[$3], $f1=[CAST($2):DOUBLE])
		    LogicalTableScan(table=[[TEST_PRO, T_USER]])
```

---
### Optimizer
-   >在calcite中有两种优化器：
	>HepPlanner(RBO)：它是一个启发式的优化器，按照规则进行匹配，知道达到次数限制，或者遍历一遍，不在出现rule match为止。

	>VolcanoPlanner(RBO+CBO)：他会一直迭代rule,知道找到最小的cost的plan
-   >RBO：
	
    >基于规则的优化器（Rule-Based Optimizer，RBO）：根据优化规则对关系表达式进行转换，这里说的是，一个表达式A经过优化 得到另一个表达式B，得到B的同时A被裁掉，经过一系列转换得到最终的结果。
	RBO有着一套严格顺序的优化规则，同样一条SQL，无论读取表中数据是什么样的，最后生成的执行计划都是相同的，所以SQL的性能可能和SQL本身的书写有很大的关系。
	
-   >CBO：
	
    >基于代价的优化器(Cost-Based Optimizer，CBO)：也是根据优化规则对表达式进行转换，这里不同的是，表达式A在转换表达式B后，会被保存下来，根据不同的规则生成多个不同的表达式，CBO会根据统计模型和代价模型，在多个表达式选择一个cost最小的计划作为整个执行计划。
	
-   >由上可知，CBO有两个依赖，一个是统计模型，一个是代价模型，目前大多数的计算引擎都趋向CBO，但是CBO比较有难度，因为不能明确预知数据量大小，所以CBO主要用于离线的场景。

---