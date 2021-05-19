# SQL_Custom

---
## 使用参考
-   本次使用的是mvn构建脚本，后面会转向gradle.kt脚本构建
    >主要实现以下几个步骤：

    >1. dependencies：`calcite-core`

    >2. `copy-fmpp-resources`

    >3. `unpack-parser-templates`,output:`target/codegen/templates`

    >4. `generate-fmpp-sources`, output:`target/generated-sources`

    >5. `javacc`,output:`target/generated-sources`


-   实现过程细节
    >1. 可以自定义生成parser的包名：eg: package `org.apache.calcite.sql.parser.impl`,class `CustomSqlParserImpl`,
    在使用的时候直接引入对应的报名即可。

    >2. 需要导入需要处理的自定义语句，主要是JavaCode形式存在

    >3. 如果想添加新的TOKEN，可以直接在config.fmpp中data项中，直接添加，如果parser.jj中已经定义过的TOKEN，可以不需要再次定义，如果是没有的，直接在keywords项中直接添加，eg：`DATABASES` `TABLES`

    >4. 解析类型支持很多，joinTypes,statementParserMethods,literalParserMethods,alterStatementParserMethods,createStatementParserMethods,dropStatementParserMethods

    >5. 需要在`implementationFiles`中添加具有解析器方法的文件列表，主要是实现语法解析文件，前面的keywords已经实现了词法解析。在`freemarkerLinks`中导入语法解析文件所在路径

    >6. **后续添加自定义语句流程**：添加keywords，添加JavaCode文件，添加解析列表，添加语法解析器文件列表，进行fmpp和javacc的生成即可，使用的时候添加对应的parser解析器package和class即可