### 关于CGO的使用细节问题
-----------------------------------------------------------------------

[使用参考网址](https://www.cntofu.com/book/73/index.html)


#### 2020-12-02---issues:CGO的静态链接（已解决）

-   >在当前目录下创建C函数目录 eg：number


	>编写c函数：eg：number.c number.h

    >gcc编译：
```
        gcc -c -o number.o number.c	        // 该过程主要是编译源码
        ar rcs libnumber.a number.o	        // 进行静态库的创建
```
-   >回到go语言主函数：引用CGO静态库

```
        //#cgo CFLAGS: -I./number
        //#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
        //#include "number.h"
        import"C"
```
-   >参数解析：
	
    >两个#cgo命令，分别是编译和链接参数。

    >CFLAGS通过-I./number将number库对应头文件所在的目录加入头文件检索路径。

	>LDFLAGS通过-L${SRCDIR}/number将编译后number静态库所在目录加为链接库检索路径
	 
    >-lnumber表示链接libnumber.a静态库。
	
    >需要注意的是，在链接部分的检索路径不能使用相对路径（C/C++代码的链接程序所限制）。
	
    >我们必须通过cgo特有的${SRCDIR}变量将源文件对应的当前目录路径展开为绝对路径（因此在windows平台中绝对路径不能有空白符号）。
	
    >因为我们有number库的全部代码，所以我们可以用go generate工具来生成静态库，或者是通过Makefile来构建静态库。因此发布CGO源码包时，我们并不需要提前构建C静态库。

    >因为多了一个静态库的构建步骤，这种使用了自定义静态库并已经包含了静态库全部代码的Go包无法直接用go get安装。
	
    >不过我们依然可以通过go get下载，然后用go generate触发静态库构建，最后才是go install来完成安装。
	
    >为了支持go get命令直接下载并安装，我们C语言的#include语法可以将number库的源文件链接到当前的包。

	>创建z_link_number_c.c文件如下：
	>#include "./number/number.c".
    >然后在执行go get或go build之类命令的时候，CGO就是自动构建number库对应的代码。这种技术是在不改变静态库源代码组织结构的前提下，将静态库转化为了源代码方式引用。这种CGO包是最完美的。
	
	
#### CGO的动态库链接（已解决）
	
-   >动态库出现的初衷是对于相同的库，多个进程可以共享同一个，以节省内存和磁盘资源。但是在磁盘和内存已经白菜价的今天，这两个作用已经显得微不足道了，那么除此之外动态库还有哪些存在的价值呢？从库开发角度来说，动态库可以隔离不同动态库之间的关系，减少链接时出现符号冲突的风险。而且对于windows等平台，动态库是跨越VC和GCC不同编译器平台的唯一的可行方式。
-	>对于CGO来说，使用动态库和静态库是一样的，因为动态库也必须要有一个小的静态导出库用于链接动态库（Linux下可以直接链接so文件，但是在Windows下必须为dll创建一个.a文件用于链接）。我们还是以前面的number库为例来说明如何以动态库方式使用。
	
    
-   >对于在macOS和Linux系统下的gcc环境，我们可以用以下命令创建number库的的动态库
```
        cd number
        gcc -shared -o libnumber.so number.c
```
-   >因为动态库和静态库的基础名称都是libnumber，只是后缀名不同而已。因此Go语言部分的代码和静态库版本完全一样：
```
        package main

        //#cgo CFLAGS: -I./number
        //#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
        //
        //#include "number.h"
        import "C"
        import "fmt"

        func main() {
            fmt.Println(C.number_add_mod(10, 5, 12))
        }
```
-   >编译时GCC会自动找到libnumber.a或libnumber.so进行链接。
	
-----------------------------------------------------------------------

#### 静态库和动态库区别（已解决）简洁版理解

-   >静态库：即静态链接库。以.a 为文件后缀名。在程序编译时会被链接到目标代码中，程序运行时将不再需要该静态库。在编译的时候就已经将所需要的依赖包含进去，运行所需内存小，运行时间比较快。
  
-	>动态库：即动态链接库。以.tbd(之前叫.dylib) 为文件后缀名。在程序编译时并不会被链接到目标代码中，而是在程序运行是才被载入，因此在程序运行时还需要动态库存在,在运行是才将所需要的依赖包含进去，运行时需要内存比较大。

-----------------------------------------------------------------------

#### 如何直接引入C语言代码（未解决）