# Gradle
---
-   [官方网站](https://discuss.gradle.org/)
-   [参考网站](https://www.w3cschool.cn/gradle/6qo51htq.html)
---
## 介绍
-   一个基于 JVM 的富有突破性构建工具。
    
    它为您提供了: 
    
    一个像 ant 一样，通用的灵活的构建工具；
    一种可切换的，像 maven 一样的基于约定约定优于配置的构建框架；
    强大的多工程构建支持；
    强大的依赖管理(基于 ApacheIvy)；
    对已有的 maven 和 ivy 仓库的全面支持；
    支持传递性依赖管理，而不需要远程仓库或者 pom.xml 或者 ivy 配置文件；
    ant 式的任务和构建是 gradle 的第一公民；
    基于 groovy，其 build 脚本使用 groovy dsl 编写；
    具有广泛的领域模型支持你的构建；
---
## 安装
-   [环境配置参考](https://www.cnblogs.com/imyalost/p/8746527.html)
-   [安装参考链接](https://www.w3cschool.cn/gradle/ctgm1htw.html)
---

## 构建基础
-   Projects 和 tasks
    >projects 和 tasks是 Gradle 中最重要的两个概念。
    
    >任何一个 Gradle 构建都是由一个或多个 projects 组成。每个 project 包括许多可构建组成部分。 这完全取决于你要构建些什么。举个例子，每个 project 或许是一个 jar 包或者一个 web 应用，它也可以是一个由许多其他项目中产生的 jar 构成的 zip 压缩包。一个 project 不必描述它只能进行构建操作。它也可以部署你的应用或搭建你的环境。不要担心它像听上去的那样庞大。 Gradle 的 build-by-convention 可以让您来具体定义一个 project 到底该做什么。
    
    >每个 project 都由多个 tasks 组成。每个 task 都代表了构建执行过程中的一个原子性操作。如编译，打包，生成 javadoc，发布到某个仓库等操作。
---
### Gradle使用总结
-   问题——初始化模板问题：`gradle init` 这里可以直接在初始化的时候定义初步的模板如：eg: `gradle init --type java-library`

-   gradle提供了基础的四种模板
                                                                                               
    >1:basic        基础模板
                                                                                               
    >2:application      应用模板
                                                                                                                                                                                                                                              
    >3:library          依赖模板
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          
    >4:Gradle plugin        插件模板    

-   提供了两种：build script DSL
    >Groovy语言
    
    >Kotlin语言 
-   我们若想要创建基础的src目录，可以在基础模板上添加task进行创建
    >[网站参考](https://blog.csdn.net/biyiy929/article/details/89476711?utm_medium=distribute.pc_relevant.none-task-blog-title-2&spm=1001.2101.3001.4242#5.x)

```gradle task
    // 对java语言的应用
    apply plugin:'java'

    // 源兼容性
    sourceCompatibility =1.8

    // 仓库
    repositories {
        mavenCentral()
    }

    // 依赖关系
    dependencies{
        ...
    }

    // 创建目录的tasks，5.X版本之后不能用 >>> 改用 doLast
    task 'create-dirs' {
        doLast {
            sourceSets*.java.srcDirs*.each {
                it.mkdirs()
            }
            sourceSets*.resources.srcDirs*.each {
                it.mkdirs()
            }
        }
    }
```
-   结果-我们使用的标准目录
![](https://img-blog.csdn.net/20170322184320490?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveF9peWE=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

-   目录解析
    >src/main/java目录包含了项目的源代码。
    
    >src/main/resources目录包含了项目的资源（如属性文件）。

    >src/test/java目录包含了测试类。

    >src/test/resources目录包含了测试资源。所有我们构建生成的文件都会在build目录下被创建，这个目录涵盖了以下的子目录

-   进行`gradle build`之后会产生`build`目录，目录结构解析如下：
    
    >classes目录包含编译过的.class文件。
 
    >libs目录包含构建生成的jar或war文件。

-   可以直接在`build.gradle`中添加我们要构建的函数
```gradle
    apply plugin: 'java'
     
    jar {
        manifest {
            attributes 'Main-Class': 'xx.xx.xx.HelloWorld'
        }
```


---
 ### Gradle常见命令
 - compileJava      
 - processResources     
 - classes      
 - jar
 - assemble     编译并打包jar文件，但不会执行单元测试，一些其他插件可能会增强这个任务的功能，如果采用了war插件，这个任务变回为你的项目打出war包
 - compileTestJava
 - processTestResources
 - testClasses
 - test
 - check        编译并测试代码
 - build
 - clean        删除build目录以及所有构建完成的文件
 
 
 -  添加Maven仓库
```gradle
    repositories{
        mavenCentral()
    } 
```
-   添加依赖，这里我们添加了commons-collections和junit
```gradle
    dependencies {
        compile group: 'commons-collections', name: 'commons-collections', version: '3.2'
        testCompile group: 'junit', name: 'junit', version: '4.+'
    }
```
 
-   自定义项目，要指定版本号和所用的JDK版本，并添加一些属性到manifest中
```gradle
    apply plugin: 'java'
    sourceCompatibility = 1.5
    version = '1.0'
    jar {
        manifest {
            attributes 'Implementation-Title': 'Gradle Quickstart', 'Implementation-Version': version
        }
    }
```

- 发布jar包，如何发布 jar 包?你需要告诉 Gradle 发布到到哪。在 Gradle 中 jar 包通常被发布到某个仓库中
```gradle
   uploadArchives {
       repositories {
          flatDir {
              dirs 'repos'
          }
       }
   }
```
-   执行gradle uploadArchives发布

-   多工程构建-TODO

--- 
## gradle依赖管理基础

-   声明依赖
```gradle
   dependencies {
       compile group: 'org.hibernate', name: 'hibernate-core', version: '3.6.7.Final'
       testCompile group: 'junit', name: 'junit', version: '4.+'
   }
```
---
###  依赖配置
-   compile         编译范围依赖在所有的classpath中可用，同时他们也会被打包
-   runtime         依赖在运行和测试时需要，但在编译时不需要
-   testCompile         测试期编译需要的附加依赖
-   testRuntime         测试运行期需要

-
- 外部依赖

   >依赖的类型有很多种，其中有一种类型称之为外部依赖，这种依赖由外部的构建或者不同的仓库中
    ，需要像下面这样进行依赖配置
    
   >外部依赖包含 group，name 和 version 几个属性。根据选取仓库的不同，group 和 version 也可能是可选的。
   
   >当然，也有一种更加简洁的方式来声明外部依赖。采用：将三个属性拼接在一起即可。`"group:name:version"`

```gradle
    dependencies {
        // 属于编译时运行该依赖
        compile group: 'org.hibernate', name: 'hibernate-core', version: '3.6.7.Final'
    }
```
---
### 仓库
-   Gradle 是在一个被称之为仓库的地方找寻所需的外部依赖。仓库即是一个按 group，name 和 version 规则进行存储的一些文件。Gradle 可以支持不同的仓库存储格式，如 Maven 和 Ivy，并且还提供多种与仓库进行通信的方式，如通过本地文件系统或 HTTP。

-   默认情况下，Gradle 没有定义任何仓库，你需要在使用外部依赖之前至少定义一个仓库，例如 Maven 中央仓库。


-   使用Maven中央仓库
```gradle
   repositories {
       mavenCentral()
   }
```
-   使用Maven远程仓库
```gradle
   repositories {
       maven {
           url "http://repo.mycompany.com/maven2"
       }
   }
```
-   采用lvy远程仓库
 
```gradle
    repositories {
        ivy {
            url "http://repo.mycompany.com/repo"
        }
    }
```

-   采用本地lvy仓库
```gradle
    repositories {
        ivy {
           // URL can refer to a local directory
           url "../local-repo"
        }
    }
```
- 一个项目可以采用多个库。Gradle 会按照顺序从各个库里寻找所需的依赖文件，并且一旦找到第一个便停止搜索。
---
###　打包发布
- 插件对于打包提供了完美的支持，所以通常而言无需特别告诉 Gradle 需要做什么。但是你需要告诉 Gradle 发布到哪里。
- 这就需要在 uploadArchives 任务中添加一个仓库。
- 执行 gradle uploadArchives，Gradle 便会构建并上传你的 jar 包，同时会生成一个 ivy.xml 一起上传到目标仓库。
-   发布到lvy仓库
```gradle
    uploadArchives {
        repositories {
            ivy {
                credentials {
                    username "username"
                    password "pw"
                }
                url "http://repo.mycompany.com"
            }
        }
    }
```
-   发布到Maven仓库
```gradle
    apply plugin: 'maven'
    uploadArchives {
        repositories {
            mavenDeployer {
                repository(url: "file://localhost/tmp/myRepo/")
            }
        }
    }
```
---
### Gradle命令行的基本用法
-   [主要查看各种过程细节和信息](https://www.w3cschool.cn/gradle/ebfc1hto.html)
 
---
### Gradle-关于创建

-   创建目录
```gradle
    File classesdir = new File("build/classes")
    task resources { 
    	doLast{
    	    classesDirs.mkdirs()
    	    // do something
    	}
    }
    task compile(dependsOn: 'resources') {
    	doLast{
    	    if (classesDirs.isDirectory()) {
    		println 'The class directory exists. I can operate'
    	    }
    	    // do something
    	}  
    }
```
-   关于Gradle的属性设置
-   [网站链接](https://www.cnblogs.com/im-victor/p/10845245.html)

-   关于tasks的一些使用案例
```git
    git clone https://github.com/davenkin/gradle-learning.git
```
-   [实例代码](./gradle-learning)

---
###  关于gradle和ant的联合使用问题

- 问题描述：
  >在smart_sql中需要将calcite中的gradle task抽取出来，本身使用的是gradle kotlin文本，需要对其进行修改才能使用，从calcite-core中找到fmpp和javacc的task进行copy

-  问题描述：
![]()