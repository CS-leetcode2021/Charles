# Maven
--- 
-   [官方文档](http://maven.apache.org/index.html)
-   [参考网站](https://www.w3cschool.cn/maven/)
-   [参考文档](https://www.cnblogs.com/hihtml5/p/6248327.html)
---
### Maven仓库地址
-   `https://mvnrepository.com/`
---

### Maven 是什么？
   >Maven 是一个项目管理和整合工具。Maven 为开发者提供了一套完整的构建生命周期框架。
   
   >开发团队几乎不用花多少时间就能够自动完成工程的基础构建配置，因为 Maven 使用了一个标准的目录结构和一个默认的构建生命周期。
   在有多个开发团队环境的情况下，Maven 能够在很短的时间内使得每项工作都按照标准进行。因为大部分的工程配置操作都非常简单并且可复用，在创建报告、检查、构建和测试自动配置时，Maven 可以让开发者的工作变得更简单。
   
   >Maven 能够帮助开发者完成以下工作：
   - 构建
   - 文档生成
   - 报告
   - 依赖
   - SCMs
   - 发布
   - 分发
   - 邮件列表
   >总的来说，Maven 简化了工程的构建过程，并对其标准化。它无缝衔接了编译、发布、文档生成、团队合作和其他任务。Maven 提高了重用性，负责了大部分构建相关的任务。
---
### Maven 的目标
   >在开发过程中，为了保证编译通过，我们会到处去找包，当编译通过了，却发现还是有部分class没有找到
   >这个时候我们还会想到？缺包？
   >每个Java项目都没有一个统一的标准，配置文件到处都是，单元测试文件也不知道到底放在哪里

   >Maven 的主要目的是为开发者提供:
    
   -    一个可复用、可维护、更易理解的工程综合模型
   -    与这个模型交互的插件或者工具
   >Maven 工程结构和内容被定义在一个 xml 文件中 － pom.xml，是 Project Object Model (POM) 的简称，此文件是整个 Maven 系统的基础组件
    
---
### 示例：

```
    配置项	            默认值
    source code	        ${basedir}/src/main/java
    resources	        ${basedir}/src/main/resources
    Tests	        ${basedir}/src/test
    Complied byte code	${basedir}/target
    distributable JAR	${basedir}/target/classes
```
---
### 安装
-   [参考网址链接](https://www.w3cschool.cn/maven/j3x41ht2.html)
-   [安装参考链接](https://blog.csdn.net/qq_38270106/article/details/97764483)
---
### Maven-Project Object Model (POM)
   >POM 代表工程对象模型。它是使用 Maven 工作时的基本组建，是一个 xml 文件。它被放在工程根目录下，文件命名为 pom.xml。

   >POM 包含了关于工程和各种配置细节的信息，Maven 使用这些信息构建工程。

   >POM 也包含了目标和插件。当执行一个任务或者目标时，Maven 会查找当前目录下的 POM，从其中读取所需要的配置信息，然后执行目标。能够在 POM 中设置的一些配置如下：
   - project dependencies
   - plugins
   - goals
   - build profiles
   - project version
   - developers
   - mailing list
   >在创建 POM 之前，我们首先确定工程组（groupId），及其名称（artifactId）和版本，在仓库中这些属性是工程的唯一标识。

```maven.eg
    <project xmlns="http://maven.apache.org/POM/4.0.0"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
         
      <modelVersion>4.0.0</modelVersion>
      <groupId>com.xrq.withmaven</groupId>
      <artifactId>withmaven</artifactId>
      <version>0.0.1-SNAPSHOT</version>
      <build/>
    </project>
```
- **解析**

-   modelVersion
    >指定了当前Maven模型的版本号，对于Maven2和Maven3来说，它只能是4.0.0
-   groupId
    >顾名思义，这个应该是公司名或是组织名。一般来说groupId是由三个部分组成，每个部分之间以"."分隔，第一部分是项目用途，比如用于商业的就是"com"，用于非营利性组织的就　　是"org"；第二部分是公司名，比如"tengxun"、"baidu"、"alibaba"；第三部分是你的项目名
-   artifactId
    >可以认为是Maven构建的项目名，比如你的项目中有子项目，就可以使用"项目名-子项目名"的命名方式
-   version
    >版本号，SNAPSHOT意为快照，说明该项目还在开发中，是不稳定的版本。在Maven中很重要的一点是，groupId、artifactId、version三个元素生成了一个Maven项目的基本坐标，这非常重要，我在使用和研究Maven的时候多次感受到了这点。

-   **在上面的这些元素之外，还有一些元素，同样罗列一下：**
-   packing
    >项目打包的类型，可以使jar、war、rar、ear、pom，默认是jar
-   dependencies和dependency
    >前者包含后者。前面说了，Maven的一个重要作用就是统一管理jar包，为了一个项目可以build或运行，项目中不可避免的，会依赖很多其他的jar包，在Maven中，这些依赖就被称为dependency。
    
    >说到这里，就有一个本地仓库和远程仓库的概念了。官方下载的本地仓库的配置在"%MAVEN_HOME%\conf\settings.xml"里面，找一下"localRepository"就可以了；MyEclipse默认的本地仓库的地址在"{user.home}/.m2/repository"路径下，同样找一下"localRepository"就可以找到MyEclipse默认的本地仓库了。
 
    >本地仓库和远程仓库是这样的，Maven工程首先会从本地仓库中获取jar包，当无法获取指定jar包时，本地仓库会从远程仓库（中央仓库）中下载jar包，并放入本地仓库以备将来使用。
 
    >举个例子，比方说项目中用到了`MyBatis`，那么可以这么配置：
```
    <dependencies>
        <dependency>
            <groupId>org.mybatis</groupId>
            <artifactId>mybatis</artifactId>
            <version>3.2.5</version>
        </dependency>
    </dependencies>
```                                                                                                                                                                                                                                                                                                                                                                                                                                                        
   >之前有说过groupId、artifactId、version唯一标识一个Maven项目，有了这三个元素，我们就可以去远程仓库下载MyBatis3.2.5.jar到本地仓库了。回想我们之前的做法，如果要MyBatis的jar包，发现没有，然后去网上下载一个，需要另外的jar包，然后去网上下载一个，但是有了Maven，就方便多了，只需要配置jar包对应的dependency依赖，Maven会自动帮助我们去远程仓库中下载jar包到本地仓库中。                                                                                                                                                                                                                                                                                                                                                                                                                                                    
 
-   properties
    >properties是用来定义一些配置属性的，例如project.build.sourceEncoding（项目构建源码编码方式），可以设置为UTF-8，防止中文乱码，也可定义相关构建版本号，便于日后统一升级。
-   build
    >build表示与构建相关的配置，比如build下有finalName，表示的就是最终构建之后的名称。

---
### 了解Maven仓库
-   使用 Maven 给我们带来的最直接的帮助，就是 jar 包得到了统一管理，那么这些 jar 包存放在哪里呢？它们就在您的 本地仓库 中，位于 C:\Users\用户名\.m2 目录下（当然也可以修改这个默认地址）。
    
-   实际上可将本地仓库理解“缓存”，因为项目首先会从本地仓库中获取 jar 包，当无法获取指定 jar 包的时候，本地仓库会从远程仓库（或 中央仓库） 中下载 jar 包，并放入本地仓库中以备将来使用。这个远程仓库是 Maven 官方提供的，可通过 http://search.maven.org/ 来访问。这样一来，本地仓库会随着项目的积累越来越大。通过下面这张图可以清晰地表达项目、本地仓库、远程仓库之间的关系。
![](http://static.oschina.net/uploads/space/2014/0120/222601_cKl0_223750.png)
---
### 创建Maven项目
-   终端命令：`mvn archetype:generate -DinteractiveMode=false -DarchetypeArtifactId=maven-archetype-webapp -DgroupId=com.smart -DartifactId=smart-demo -Dversion=1.0`

-   结果：

![](http://static.oschina.net/uploads/space/2014/0118/224720_DNpK_223750.png)
-   改造结果：当然也可以在IDEA中直接配置Maven项目

![](http://static.oschina.net/uploads/space/2014/0118/225115_hglq_223750.png)

---
### 理解pom.xml
```
    <project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
      <modelVersion>4.0.0</modelVersion>
      <groupId>com.smart</groupId>
      <artifactId>smart-demo</artifactId>
      <packaging>war</packaging>
      <version>1.0</version>
      <name>smart-demo Maven Webapp</name>
      <url>http://maven.apache.org</url>
      <dependencies>
        <dependency>
          <groupId>junit</groupId>
          <artifactId>junit</artifactId>
          <version>3.8.1</version>
          <scope>test</scope>
        </dependency>
      </dependencies>
      <build>
        <finalName>smart-demo</finalName>
      </build>
    </project>
```
-   从上往下简要说明一下：

>modelVersion：这个是 POM 的版本号，现在都是 4.0.0 的，必须得有，但不需要修改。

>groupId、artifactId、version：分别表示 Maven 项目的组织名、构件名、版本号，它们三个合起来就是 Maven 坐标，根据这个坐标可以在 Maven 仓库中对应唯一的 Maven 构件。

>packaging：表示该项目的打包方式，war 表示打包为 war 文件，默认为 jar，表示打包为 jar 文件。

>name、url：表示该项目的名称与 URL 地址，意义不大，可以省略。

>dependencies：定义该项目的依赖关系，其中每一个 dependency 对应一个 Maven 项目，可见 Maven 坐标再次出现，还多了一个 scope，表示作用域（下面会描述）。

>build：表示与构建相关的配置，这里的 finalName 表示最终构建后的名称 smart-demo.war，这里的 finalName 还可以使用另一种方式来定义（下面会描述）。



---
### Maven 构建生命周期
-   构建生命周期是一组阶段的序列（sequence of phases），每个阶段定义了目标被执行的顺序。这里的阶段是生命周期的一部分。
    
    >举例说明，一个典型的 Maven 构建生命周期是由以下几个阶段的序列组成的：
``` 
    阶段	                处理	                    描述
    prepare-resources	资源拷贝	                本阶段可以自定义需要拷贝的资源
    compile	        编译	                本阶段完成源代码编译
    package	        打包	                本阶段根据 pom.xml 中描述的打包配置创建 JAR / WAR 包
    install	        安装	                本阶段在本地 / 远程仓库中安装工程包
```
---
### 使用Maven命令：
-   前面我们已经使用了几个 Maven 命令，例如：mvn archetype:generate，mvn tomcat7:run-war 等。其实，可使用两种不同的方式来执行 Maven 命令：
    
    >方式一：mvn <插件>:<目标> [参数]
    
    >方式二：mvn <阶段>
    
    >现在我们接触到的都是第一种方式，而第二种方式才是我们日常中使用最频繁的，例如：
    
    >mvn clean：清空输出目录（即 target 目录）
    
    >mvn compile：编译源代码
    
    >mvn package：生成构件包（一般为 jar 包或 war 包）
    
    >mvn install：将构件包安装到本地仓库
    
    >mvn deploy：将构件包部署到远程仓库
    
    >执行 Maven 命令需要注意的是：必须在 Maven 项目的根目录处执行，也就是当前目录下一定存在一个名为 pom.xml 的文件

---
### 问题——运行打包后的jar文件出现没有主清单
-   因为你在打包的时候并没有告诉它主函数在哪里，要在pom.xml中告知主函数的入口
-   [问题解决参考](https://www.jianshu.com/p/72f914a617ec)

### 问题——进行mvn site时候出现错误
-   因为缺少对应的插件，需要在pom.xml中添加插件
-   [问题解决参考](https://blog.csdn.net/weixin_43936652/article/details/90067814)

---
### 后续更新--2020-12-10
