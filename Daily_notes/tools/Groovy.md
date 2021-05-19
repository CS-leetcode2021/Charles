# Groovy
---
-   是一门编程语言
-   [学习网站](https://www.w3cschool.cn/groovy/groovy_overview.html)
-   [在Idea中使用](https://blog.csdn.net/w_y_x_y/article/details/102930248)
---
### 特点
```markdown
    Groovy中有以下特点:
    同时支持静态和动态类型。
    支持运算符重载。
    本地语法列表和关联数组。
    对正则表达式的本地支持。
    各种标记语言，如XML和HTML原生支持。
    Groovy对于Java开发人员来说很简单，因为Java和Groovy的语法非常相似。
    您可以使用现有的Java库。
    Groovy扩展了java.lang.Object。
```
---
### 使用
-   身份标识 `def`
```groovy
    def employeename 
    def student1 
    def student_name
```
   > `DEF` 是在 Groovy 用来定义标识符的关键字。
     
   > 下面是一个如何在我们的 Hello World 程序中使用标识符的代码示例。
```groovy
    class Example {
       static void main(String[] args) {
          // One can see the use of a semi-colon after each statement
          def x = 5;
          println('Hello World'); 
       }
    }
```
-   方法定义：
    >Groovy 中的方法是使用返回类型或使用 def 关键字定义的。方法可以接收任意数量的参数。定义参数时，不必显式定义类型。可以添加修饰符，如 public，private 和 protected。
     默认情况下，如果未提供可见性修饰符，则该方法为 public。
     
    >最简单的方法是没有参数的方法，如下所示：
```groovy
    def methodName() { 
       //Method code 
    }
```
```groovy
    class Example {
       static def DisplayName() {
          println("This is how methods work in groovy");
          println("This is an example of a simple method");
       } 
    	
       static void main(String[] args) {
          DisplayName();
       } 
    }
```
...

### 支持
-   方法
-   文件I/O
-   可选
-   字符串方法
-   范围判定方法
-   `list` 列表
-   映射
-   面向对象：getter、setter、实例方法、创建多个对象、继承、扩展（extends）、内部类（Quter）、抽象类（不能被实例化）、接口

-   泛型
-   特征
-   闭包
-   注释
-   XML
-   JMX
-   JSON
-   DSLS
-   数据库
-   构建器
-   命令行
-   单元测试、模板引擎、元对象编程
