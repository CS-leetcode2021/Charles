# 详解Java 中try，catch，finally的用法及解析
---
[参考](https://cloud.tencent.com/developer/article/1088141)
---
```java
    try
        {
            // 需要检测的异常代码
        }
    catch
        {
            // 异常处理，即处理异常代码
        }
    finally
        {
            // 一定会执行的代码
        }
```

```java
    try
        {
        int i = 1/0;
        }
    catch(Exception e)
        {
        ........
        }
```
---
```java
public class CatchTest {
    public static void main(String[] args) {
        CatchTest catchTest = new CatchTest();
        try {
            catchTest.testEx();
        }
        catch (Exception e){
            e.printStackTrace();
        }
    }
    public CatchTest(){

    }

    boolean testEx() throws Exception{
        boolean ret = true;
        try{
            ret = testEx1();
        }
        catch (Exception e){
            System.out.println("testEx, catch exception");
            ret = false;
            throw e;
        }
        finally {
            System.out.println("testEx,finally.return value="+ret);
            return  ret;
        }
    }

    boolean testEx1() throws Exception{
        boolean ret = true;
        try {
            ret = testEx2();
            if (!ret){
                return false;
            }
            System.out.println("testEx1,at the end of try");
            return true;
        }
        catch (Exception e){
            System.out.println("testEx1,catch exception");
            ret = false;
            throw e;
        }
        finally {
            System.out.println("testEx1,finally; return value="+ ret);
            return ret;
        }
    }

    boolean testEx2() throws Exception{
        boolean ret = true;
        try {
            int b = 12;
            int c;
            for (int i=2;i>= -2;i--){
                c = b/i;
                System.out.println("i="+i);
            }
            return true;
        }
        catch (Exception e){
            System.out.println("testEx2,catch exception");
            ret = false;
            throw e;
        }
        finally {
            System.out.println("testEx2,finally; return value="+ ret);
            return ret;
        }
    }
}

```
result:
```
i=2
i=1
testEx2,catch exception
testEx2,finally; return value=false
testEx1,finally; return value=false
testEx,finally.return value=false
```
---
## 相关概念

1、 例外是在程序运行过程中发生的异常，比如除0溢出，数组越界、文件找不到等，会阻止程序的正常运行

2、 Java用过面向对象的方法来处理例外，在一个方法运行过程中，如果发生了例外，则这个方法生成代表了该例外的一个对象
    并把它交给运行时系统，运行时系统寻找相应的代码来处理这一例外。

3、 我们把生成例外对象并把它提交给运行时系统的过程称之为抛弃（throw）一个例外，运行时系统在方法的调用栈中查找，
    从生成例外的方法开始进行回溯，直到找到包含响应例外的处理方法为止，这一个过程称为捕获（catch）一个例外

## 关键字

try、catch、throw、throws、finally

