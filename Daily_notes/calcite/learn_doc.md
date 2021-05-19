### 2020-12-15
-   sqlline在ignite中是使用：
    >[sqlline使用实例](https://apacheignite-sql.readme.io/docs/sqlline)

-   sqlline的具体信息
    >[sqlline官网介绍](http://sqlline.sourceforge.net/)

-   Calcite官网使用信息
    >[example_csv](../calcite/calcite_csv.md)
    
    >补充：官网的demo并不支持相关的insert、update、delete操作，因为没有实现相关的接口

### 2020-12-17
-   Calcite parser讲解
    >[web](https://github.com/quxiucheng/apache-calcite-tutorial/tree/master/calcite-tutorial-2-parser)

-   freeMarker 介绍
    >[web](http://fmpp.sourceforge.net/)
    
    
    >freeMarker 是一款 模板引擎： 即一种基于模板和要改变的数据， 并用来生成输出文本(HTML网页，电子邮件，配置文件，源代码等)的通用工具。 它不是面向最终用户的，而是一个Java类库，是一款程序员可以嵌入他们所开发产品的组件。

-  使用: fmpp作用是将模板和数据模型结合在一起并进行输出

    >xxx.tfl是我们所获取的模板文件，可以写任何形式的文件，比如java、html等
    [参考](http://freemarker.foofun.cn/pgui_quickstart_gettemplate.html)

    >config.fmpp 是freeMarker的配置所需要的数据模型data model，或者直接在java运行代码中直接定义也可以[web](https://github.com/quxiucheng/apache-calcite-tutorial/blob/6cd6c03cce7408ac2ba59ec75b58fb8133458c62/calcite-tutorial-2-parser/parser-1-fmpp-tutorial/src/main/codegen/templates/Main.ftl#L1-L4)

    >需要的配置文件目录、模板文件的加载目录、生成的输出文件目录在maven中可以提前定义[web](https://github.com/quxiucheng/apache-calcite-tutorial/blob/master/calcite-tutorial-2-parser/parser-1-fmpp-tutorial/README.md#1%E6%B7%BB%E5%8A%A0maven%E4%BE%9D%E8%B5%96)


-   JavaCC
    >JavaCC is a **parser generator** for the Java programming language.

    >JavaCC本身并不是词法分析器和语法分析器，它是一个生成器，通过读取一个后缀为.jj的描述文件中的特定描述，来生成词法分析器和语法分析器

    >[实例参考非常清晰](https://www.cnblogs.com/suhaha/p/11733487.html)
    
    >[doc](https://github.com/javacc/javacc/blob/master/docs/documentation/index.md)

    >文件描述，三个部分，1、Option块和class声明块以parser_begin......parser_end结束；
    2、词法描述器，主要声明TOKEN，提取对应的词传送给语法分析器；
    3、语法分析器，有BNF生产式构成，主要检测TOKEN的输入序列是否有错误。

    >JavaCC 语法文件格式如下：

    ```javaCC
        javacc_input ::= javacc_options
                    "PARSER_BEGIN" "(" <IDENTIFIER> ")"
                    java_compilation_unit
                    "PARSER_END" "(" <IDENTIFIER> ")"
                    ( production )*
                    <EOF>
    ```

    




-   todo:
-   javacc and fmpp parser.jj