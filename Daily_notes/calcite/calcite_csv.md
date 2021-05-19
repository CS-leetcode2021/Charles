# Calcite_Example_CSV
---
[官方Doc连接](https://calcite.apache.org/docs/tutorial.html)

### CSV主要包含几个重要的概念：

-   通过Schema Factory和Schema接口实现自定义的schema
    
-   使用json文件描述schema
    
-   使用json文件描述视图view

-   通过接口定义表

-   确定表的记录类型

-   创建一个表的简单方法是，使用scannableTable接口直接列举所有行
    
-   创建一个表的更高级的方法是，使用FilterTable接口，并根据简单的规则进行过滤

-   创建一个表的最高级的方法是，使用TranslateTable接口，这个类能使规划规则转换成关系操作符
---
### 安装&连接

-   我们需要使用sqline来连接calcite,这个工程包括SQL shell脚本。
```java
    ./sqlline
    sqlline> !connect jdbc:calcite:model=src/test/resources/model.json admin admin
```
```java
    >!tables
+-----------+-------------+------------+--------------+---------+----------+------------+-----------+---------------------------+----------------+
| TABLE_CAT | TABLE_SCHEM | TABLE_NAME |  TABLE_TYPE  | REMARKS | TYPE_CAT | TYPE_SCHEM | TYPE_NAME | SELF_REFERENCING_COL_NAME | REF_GENERATION |
+-----------+-------------+------------+--------------+---------+----------+------------+-----------+---------------------------+----------------+
|           | SALES       | DEPTS      | TABLE        |         |          |            |           |                           |                |
|           | SALES       | EMPS       | TABLE        |         |          |            |           |                           |                |
|           | SALES       | SDEPTS     | TABLE        |         |          |            |           |                           |                |
|           | metadata    | COLUMNS    | SYSTEM TABLE |         |          |            |           |                           |                |
|           | metadata    | TABLES     | SYSTEM TABLE |         |          |            |           |                           |                |
+-----------+-------------+------------+--------------+---------+----------+------------+-----------+---------------------------+----------------+

```
-   >在执行!table其实是在后台执行[DatabaseMetaData,getTables](https://docs.oracle.com/javase/7/docs/api/java/sql/DatabaseMetaData.html#getTables(java.lang.String,%20java.lang.String,%20java.lang.String,%20java.lang.String[]))

    >现在可以看到系统有5张表，分别是SDEPTS，EMPS，DEPTS在当前的SALES的schema中，并且COLUMNS和TABLES是系统表，系统表始终存在在calcite中，但是其他的是由特定的schema实现的，例如EMPS,HEDEPTS,是基于`src/test/resources/sales`下的 EMP.CSv和DEPTS.csv来决定的
---
### 模式发现
-   calcite是如何找到这些表的，这和CSV的内容是无关的，可以理解为calcite是一个不包含存储层的一个数据库，calcite之所以能够读到信息，是因为我们执行了calcite-example-csv项目中的代码。

-   整个执行链中有很多步骤，首先，我们基于model文件中定义的schema工厂类定义了一个schema,然后，schema工厂创建了一个schema,并且这个schema创建了多张表，每张表都清楚的扫描csv文件来获取数据，最后，calcite解析了查询语句并且创建了执行计划来使用这些表来执行查询时，calcite利用表来读取数据
-   下面举出更详细的过程:在JDBC连接字符串上，我们给出了JSON格式的模型路径，模型model.js如下：
```json
    {
        "version": "1.0",
        "defaultSchema": "SALES",
        "schemas": [
            {
            "name": "SALES",
            "type": "custom",
            "factory": "org.apache.calcite.adapter.csv.CsvSchemaFactory",
            "operand": {
                "directory": "sales"
            }
            }
        ]
    }
```
-   该模型定义了一个称为“ SALES”的模式。该模式由插件类[org.apache.calcite.adapter.csv.CsvSchemaFactory](https://github.com/apache/calcite/blob/master/example/csv/src/main/java/org/apache/calcite/adapter/csv/CsvSchemaFactory.java)驱动，该类 是calcite-example-csv项目的一部分，并实现Calcite接口 [SchemaFactory](https://calcite.apache.org/javadocAggregate/org/apache/calcite/schema/SchemaFactory.html)。它的create方法实例化一个模式，directory从模型文件中传入参数：

```java
    public class CsvSchemaFactory implements SchemaFactory {
        ...
    }

    public Schema create(SchemaPlus parentSchema, String name,
      Map<String, Object> operand) {
        final String directory = (String) operand.get("directory");
        final File base =
            (File) operand.get(ModelHandler.ExtraOperand.BASE_DIRECTORY.camelName);
        File directoryFile = new File(directory);
        if (base != null && !directoryFile.isAbsolute()) {
            directoryFile = new File(base, directory);
        }
        String flavorName = (String) operand.get("flavor");
        CsvTable.Flavor flavor;
        if (flavorName == null) {
            flavor = CsvTable.Flavor.SCANNABLE;
        } else {
            flavor = CsvTable.Flavor.valueOf(flavorName.toUpperCase(Locale.ROOT));
        }
        return new CsvSchema(directoryFile, flavor);
  }
```
-   在模型的驱动下，架构工厂实例化一个称为“ SALES”的架构。该架构是[org.apache.calcite.adapter.csv.CsvSchema](https://github.com/apache/calcite/blob/master/example/csv/src/main/java/org/apache/calcite/adapter/csv/CsvSchema.java)的实例， 并实现Calcite接口[Schema](https://calcite.apache.org/javadocAggregate/org/apache/calcite/schema/Schema.html)。

```java
    public class AbstractSchema extends Schema {
        ...
    }
    
    public class CsvSchema extends AbstractSchema {
        ...
    }

```
-   模式的工作是产生表列表（它也可以列出子模式和表功能，但是它们是高级功能，而calcite-example-csv不支持它们），这些表实现了calcite的table接口，csvSchema生成CsvTable及其子类的实例的表。

```java

    private Table createTable(File file) {
        switch (flavor) {
            case TRANSLATABLE:
                return new CsvTranslatableTable(file, null);
            case SCANNABLE:
                return new CsvScannableTable(file, null);
            case FILTERABLE:
                return new CsvFilterableTable(file, null);
            default:
                throw new AssertionError("Unknown flavor " + flavor);
        }
    }
```
-   该模式将扫描目录，并找到名称以“ .csv”结尾的所有文件，并为其创建表。在这种情况下，目录为sales并且包含文件EMPS.csv和DEPTS.csv，这些文件 将成为表EMPS和DEPTS。
---
### schemas中的Tables和views

-   应该关注到，我们怎样做到不需要在模型中定义任何表，但是schema本身自动创建了表，可以使用schema的table属性来定义除了自定创建的其他表。当你编写一个查询语句的时候，视图看上去就像一张表
-   编写查询时，视图看起来像表，但是它不存储数据。它通过执行查询得出其结果。在计划查询时会扩展视图，因此查询计划者通常可以执行优化操作，例如从SELECT子句中删除最终结果中未使用的表达式。

#### 这是定义视图的模式：VIEW
```json
    {
        version: '1.0',
        defaultSchema: 'SALES',
        schemas: [
            {
            name: 'SALES',
            type: 'custom',
            factory: 'org.apache.calcite.adapter.csv.CsvSchemaFactory',
            operand: {
                directory: 'sales'
            },
            tables: [
                {
                name: 'FEMALE_EMPS',
                type: 'view',
                sql: 'SELECT * FROM emps WHERE gender = \'F\''
                }
            ]
            }
        ]
    }

```
-   相当于在json文件中，直接通过自定义操作创建了一个view文件
```java
    // 连接该json文件，通过扫描可以看到多了一个view表
    > !tables
+-----------+-------------+-------------+--------------+---------+----------+------------+-----------+---------------------------+----------------+
| TABLE_CAT | TABLE_SCHEM | TABLE_NAME  |  TABLE_TYPE  | REMARKS | TYPE_CAT | TYPE_SCHEM | TYPE_NAME | SELF_REFERENCING_COL_NAME | REF_GENERATION |
+-----------+-------------+-------------+--------------+---------+----------+------------+-----------+---------------------------+----------------+
|           | SALES       | DEPTS       | TABLE        |         |          |            |           |                           |                |
|           | SALES       | EMPS        | TABLE        |         |          |            |           |                           |                |
|           | SALES       | SDEPTS      | TABLE        |         |          |            |           |                           |                |
|           | SALES       | FEMALE_EMPS | VIEW         |         |          |            |           |                           |                |
|           | metadata    | COLUMNS     | SYSTEM TABLE |         |          |            |           |                           |                |
|           | metadata    | TABLES      | SYSTEM TABLE |         |          |            |           |                           |                |
+-----------+-------------+-------------+--------------+---------+----------+------------+-----------+---------------------------+----------------+

```
-   需要注意的是view是通过我们给定的sql语言定义出来的，实际在底层的存储中它并不存在，删除view表并无实际意义，我们在查询中可以使用view，它看起来像一个表。

#### 自定义表模式：TABLE

-   链接model-with-custom-table.json
```json
    {   
        version: '1.0',
        defaultSchema: 'CUSTOM_TABLE',
        schemas: [
            {
            name: 'CUSTOM_TABLE',
            tables: [
                    {
                        name: 'EMPS',
                        type: 'custom',
                        factory: 'org.apache.calcite.adapter.csv.CsvTableFactory',
                        operand: {
                        file: 'sales/EMPS.csv.gz',
                        flavor: "scannable"
                        }
                    }
                ]
            }
        ]
    }
```

```java
    sqlline>!connect jdbc:calcite:model=src/test/resources/model-with-custom-table.json admin admin
    sqlline>!table
    sqlline>select * from custom_table.emps;
```
```java
    +-------+-------+--------+--------+---------------+-------+------+---------+---------+------------+
    | EMPNO | NAME  | DEPTNO | GENDER |     CITY      | EMPID | AGE  | SLACKER | MANAGER |  JOINEDAT  |
    +-------+-------+--------+--------+---------------+-------+------+---------+---------+------------+
    | 100   | Fred  | 10     |        |               | 30    | 25   | true    | false   | 1996-08-03 |
    | 110   | Eric  | 20     | M      | San Francisco | 3     | 80   |         | false   | 2001-01-01 |
    | 110   | John  | 40     | M      | Vancouver     | 2     | null | false   | true    | 2002-05-03 |
    | 120   | Wilma | 20     | F      |               | 1     | 5    |         | true    | 2005-09-07 |
    | 130   | Alice | 40     | F      | Vancouver     | 2     | null | false   | true    | 2007-01-01 |
    +-------+-------+--------+--------+---------------+-------+------+---------+---------+------------+

```
-   该模式是常规模式，包含一个由org.apache.calcite.adapter.csv.CsvTableFactory支持的自定义表，该表 实现了Calcite接口 TableFactory。它的create方法实例化a CsvScannableTable，file从模型文件传入参数：
```java
    public CsvTable create(SchemaPlus schema, String name,
        Map<String, Object> map, RelDataType rowType) {
        String fileName = (String) map.get("file");
        final File file = new File(fileName);
        final RelProtoDataType protoRowType =
            rowType != null ? RelDataTypeImpl.proto(rowType) : null;
        return new CsvScannableTable(file, protoRowType);
    }
```
-   实现定制表通常是实现定制方案的一种更简单的选择。两种方法最终都可能创建类似的Table接口实现，但是对于自定义表，**不需要实现元数据**。（CsvTableFactory 创建一个CsvScannableTable，就像创建一个CsvSchema一样，但是表实现不会在文件系统中扫描.csv文件。）

-   自定义表需要为模型的创建者做更多的工作（创建者需要显式指定每个表及其文件），但也要给作者更多的控制权（例如，为每个表提供不同的参数）


### 优化查询

-   当数据源信息特别大的时候，calcite可以通过支持添加查询计划规则来进行优化查询,规划器规则的运行方式是在查询分析树中查找模式（例如，某种表顶部的项目），然后用实现优化的一组新节点替换树中匹配的节点。
-   规划器规则也可以扩展，例如架构和表。因此，如果有要通过SQL访问的数据存储，则首先定义一个自定义表或架构，然后定义一些规则以提高访问效率

```java
    sqlline> !connect jdbc:calcite:model=src/test/resources/model.json admin admin
    sqlline> explain plan for select name from emps;
    +-----------------------------------------------------+
    | PLAN                                                |
    +-----------------------------------------------------+
    | EnumerableCalcRel(expr#0..9=[{inputs}], NAME=[$t1]) |
    |   EnumerableTableScan(table=[[SALES, EMPS]])        |
    +-----------------------------------------------------+
    sqlline> !connect jdbc:calcite:model=src/test/resources/smart.json admin admin
    sqlline> explain plan for select name from emps;
    +-----------------------------------------------------+
    | PLAN                                                |
    +-----------------------------------------------------+
    | EnumerableCalcRel(expr#0..9=[{inputs}], NAME=[$t1]) |
    |   CsvTableScan(table=[[SALES, EMPS]])               |
    +-----------------------------------------------------+
```

-   差异的原因主要是在smart.json中多了一行数据
 
    `flavor: "translatable"`

-   这将导致CsvSchema使用创建 flavor = TRANSLATABLE，并且其createTable方法将创建CsvTranslatableTable的实例， 而不是CsvScannableTable。scannableTable实现的是EnumerableTableScan操作，表扫描是查询运算符树的叶子，将数据所有数据一行行读出
-   CsvTranslatableTable实现TranslatableTable.toRel() 创建CsvTableScan的 方法 。但是我们创建了一个独特的子类型，该子类型将导致规则触发
```java
    //完成的规则，通过预定义的配置文件(config)加载扫描规则
    public class CsvProjectTableScanRule
        extends RelRule<CsvProjectTableScanRule.Config> {
    /** Creates a CsvProjectTableScanRule. */
    protected CsvProjectTableScanRule(Config config) {
        super(config);
    }

    @Override public void onMatch(RelOptRuleCall call) {
        final LogicalProject project = call.rel(0);
        final CsvTableScan scan = call.rel(1);
        int[] fields = getProjectFields(project.getProjects());
        if (fields == null) {
        // Project contains expressions more complex than just field references.
        return;
        }
        call.transformTo(
            new CsvTableScan(
                scan.getCluster(),
                scan.getTable(),
                scan.csvTable,
                fields));
    }

    private int[] getProjectFields(List<RexNode> exps) {
        final int[] fields = new int[exps.size()];
        for (int i = 0; i < exps.size(); i++) {
        final RexNode exp = exps.get(i);
        if (exp instanceof RexInputRef) {
            fields[i] = ((RexInputRef) exp).getIndex();
        } else {
            return null; // not a simple projection
        }
        }
        return fields;
    }

    /** Rule configuration. */
    public interface Config extends RelRule.Config {
        Config DEFAULT = EMPTY
            .withOperandSupplier(b0 ->
                b0.operand(LogicalProject.class).oneInput(b1 ->
                    b1.operand(CsvTableScan.class).noInputs()))
            .as(Config.class);

        @Override default CsvProjectTableScanRule toRule() {
        return new CsvProjectTableScanRule(this);
        }
    }
```

-   JDBC适配器当前仅下推表扫描操作。所有其他处理（过滤，联接，聚合等）都在方解石内部进行。我们的目标是尽可能减少对源系统的处理，尽可能地翻译语法，数据类型和内置函数。如果Calcite查询基于单个JDBC数据库中的表，则原则上整个查询应转到该数据库。如果表来自多个JDBC来源，或JDBC和非JDBC的混合，则Calcite将使用它可以使用的最高效的分布式查询方法。