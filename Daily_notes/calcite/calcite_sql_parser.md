# SQL-Parser

## 代码结构：
```html
    --- src
        ---main
            --- codegen [这里存放FreeMarker模板目录]
                ---includes [自定义解析代码目录]
                    --- compoundIdentifier.ftl  []
                    --- parserImpls.ftl
                ---templates [SQL解析核心文件，javaCC的解析文件]
            --- config.fmpp [配置FreeMarker和JavaCC相关配置]
         // 以上部分已经实现
            --- java
            --- resources
        --- test
    --- target
        --- classes
        --- generated-sources
            --- fmpp  [生产环境FMPP生成的文件]
                --- javacc
                    --- Parser.jj
            --- javacc [生产环境JavaCC编译后的目录]
                --- package [会产生一系列的java文件]
    --- generated-test-sources
        --- fmpp  [测试环境FMPP生成后的文件]
        --- javacc  [测试环境JavaCC编译后生成的目录]

```

-   解析
    >config.fmpp FM的配置模板
    
    >Parser.jj  JavaCC解析器

    >parserImpls.ftl/compoundldentifier.ftl 自定义JavaCC语法格式的解析SQL代码

-   自定义SQl语句实现的流程：
    >1、需要在 `parserImpls.ftl` 和 `compoundldentifier.ftl`两个文件中添加自定义解析器语句、文本或数据类型

    >2、结合templates/Parser.jj 和 comfig.fmpp文件进行fmpp配置文件，该文件将data model 和data temples 进行整合为Parser.jj

    >3、JavaCC本身并不是词法分析器和语法生成器，是一个生成器，主要的作用是将 *.jj文件生成对应的词法语法描述器，进行sql语言的检测和分析

    >4、生成的词法语法描述器在target/generated-sources/javacc目录下

-   **我们当下目标:** 是要在`parserImpls.ftl` 和 `compoundldentifier.ftl`文件中添加自定义sql语句

## 解析示例代码

```java
public class SqlParserSample {
    public static void main(String[] args) throws SqlParseException {
        // Sql语句
        String sql = "select * from emps where id = 1";
        // 解析配置
        SqlParser.Config mysqlConfig = SqlParser.configBuilder().setLex(Lex.MYSQL).build();
        // 创建解析器
        SqlParser parser = SqlParser.create(sql, mysqlConfig);
        // 解析sql
        SqlNode sqlNode = parser.parseQuery();
        // 还原某个方言的SQL
        System.out.println(sqlNode.toSqlString(OracleSqlDialect.DEFAULT));
    }
}
```

## 解析流程

1. 首先生成SQL解析器`SqlParser.Config`,`SqlParser.Config`中存在获取解析工厂类`SqlParser.Config#parserFactory()`方法,可以在`SqlParser.configBuilder()`配置类中设置解析工厂
2. `SqlParserImplFactory`解析工厂中调用`getParser`方法获取解析器
3. `SqlAbstractParserImpl`抽象解析器,JavaCC中生成的解析器的父类,Calcite中默认的解析类名为`SqlParserImpl`
4. `SqlParserImpl`中,有静态字段`FACTORY`,主要是实现`SqlParserImplFactory`,并创建解析器
5. `SqlParser`调用`create`方法,从`SqlParser.Config`中获取工厂`SqlParserImplFactory`,并创建解析器
6. 调用`SqlParser#parseQuery`方法,解析SQL,最终调用`SqlAbstractParserImpl`(默认实现类`SqlParserImpl`)的`parseSqlStmtEof`或者`parseSqlExpressionEof`方法,获取解析后的抽象语法树`SqlNode`


Parser.jj 解析简单介绍
1. 调用`SqlParserImplFactory`的`SqlAbstractParserImpl getParser(Reader stream);`方法,解析获取解器,
   或者,直接调用`SqlParser#parseQuery`传入sql语句,解析器重新传入sql`parser.ReInit(new StringReader(sql));`
2. 解析器入口类`SqlAbstractParserImpl#parseSqlExpressionEof`或者`SqlAbstractParserImpl#parseSqlStmtEof`
3. Parser.jj解析SQL语句入口`SqlStmtEof()` 解析SQL语句，直到文件结束符,`SqlStmtEof()`调用`SqlStmt()`
4. `SqlStmt()`中定义各个类型的解析,例如 `SqlExplain()`(explain语句),`OrderedQueryOrExpr()`(select语句),之后解析各个关键字
