## GO
---
### 解决国内 go get 无法下载的问题（已解决）
-   [参考文章](https://www.sunzhongwei.com/problem-of-domestic-go-get-unable-to-download?from=sidebar_new)

---
### VScode安装go所需插件
```go
    go env -w GO111MODULE=on
    go get -v github.com/mdempsky/gocode 
    go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs 
    go get -v github.com/ramya-rao-a/go-outline 
    go get -v github.com/acroca/go-symbols 
    go get -v golang.org/x/tools/cmd/guru 
    go get -v golang.org/x/tools/cmd/gorename 
    go get -v github.com/cweill/gotests/... 
    go get -v github.com/fatih/gomodifytags 
    go get -v github.com/josharian/impl FAILED
    go get -v ithub.com/davidrjenni/reftools/cmd/fillstruct 
    go get -v github.com/haya14busa/goplay/cmd/goplay 
    go get -v github.com/godoctor/godoctor 
    go get -v github.com/go-delve/delve/cmd/dlv 
    go get -v github.com/stamblerre/gocode 
    go get -v github.com/rogpeppe/godef 
    go get -v github.com/sqs/goreturns 
    go get -v golang.org/x/lint/golint 
    go get -v golang.org/x/tools/gopls
    go get -v golang.org/x/tools/cmd/godoc
```

---
### 设置go环境的代理（已解决）
```go
    export GOPROXY=https://mirrors.aliyun.com/goproxy/
```

---
### go语言命令
-   [参考网站](https://www.cntofu.com/book/19/readme.html)
-   [go命令教程](https://github.com/hyper0x/go_command_tutorial.git)
-   [教程参考代码](https://github.com/hyper0x/goc2p.git)

-   >Go语言的源码文件有三大类，即：命令源码文件、库源码文件和测试源码文件。他们的功用各不相同，而写法也各有各的特点。
	
    >命令源码文件总是作为可执行的程序的入口。
	
    >库源码文件一般用于集中放置各种待被使用的程序实体（全局常量、全局变量、接口、结构体、函数等等）。
	
    >测试源码文件主要用于对前两种源码文件中的程序实体的功能和性能进行测试。另外，后者也可以用于展现前两者中程序的使用方法。


#### go build命令：命令用于编译我们指定的源码文件或代码包以及它们的依赖包
-   >1、执行时如果不跟任何项，将试图编译当前目录所对应的代码包。
	
    >2、如果当前文件已经设置在GOPATH或GOROOT下，便可以直接go build XXX 进行编译。
	
    >3、go build可以一次性编译多个文件，但是必须是同一个目录下的多个文件，同时编译两个包是不可以的。
	
    >4、go build -v XXX.go  直接将编译目标代码的名字命名给生成的可执行程序。
		  
    >5、go build -o test XXX.go 可以自定义任意的名称给生成的可执行程序，可以指定输出文件（在这个示例中指的是可执行文件）的名称，当使用标记-o的时候，不能同时对多个代码包进行编译。
		  
    >6、-i会使go build命令安装那些编译目标依赖的且还未被安装的代码包。
	
    >7、-a 强行对所有涉及到的代码包（包含标准库中的代码包）进行重新构建，即使它们已经是最新的了。
	
    >8、-n 打印编译期间所用到的其它命令，但是并不真正执行它们。
		  
    >9、-p n 指定编译过程中执行各任务的并行数量（确切地说应该是并发数量）。在默认情况下，该数量等于CPU的逻辑核数。但是在darwin/arm平台（即iPhone和iPad所用的平台）下，该数量默认是1。
	
    >10、-race 	开启竞态条件的检测。不过此标记目前仅在linux/amd64、freebsd/amd64、darwin/amd64和windows/amd64平台下受到支持。
	
    >11、-work 	打印出编译时生成的临时工作目录的路径，并在编译结束时保留它。在默认情况下，编译结束时会删除该目录。
		  
    >12、-x 	打印编译期间所用到的其它命令。注意它与-n标记的区别。

#### go install：用于编译并安装指定的代码包及它们的依赖包 
-   >1、只比go build命令多做了一件事，即：安装编译后的结果文件到指定目录。
	
    >2、如果go install 命令后面跟的代码包中仅包含库源码文件，那么go install命令会把编译后的结果文件保存在源码文件所在工作区的pkg目录下，也叫静态链接库文件。
	
    >3、如果go install 命令后面跟的代码包是命令源文件，那么go install命令会把生成的可执行程序保存在源码文件所在工作区的bin目录下。
	
    >4、安装VScode go插件可以先从github.com下载需要的源码，然后用go install 命令直接进行编译安装在GOPATH目录下的bin目录，为可执行程序。

#### go get ：可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装

-   >1、-d 	让命令程序只执行下载动作，而不执行安装动作。
	
    >2、-f 	仅在使用-u标记时才有效。该标记会让命令程序忽略掉对已下载代码包的导入路径的检查。如果下载并安装的代码包所属的项目是你从别人那里Fork过来的，那么这样做就尤为重要了。
	
    >3、-fix 	让命令程序在下载代码包后先执行修正动作，而后再进行编译和安装。
	
    >4、-insecure 	允许命令程序使用非安全的scheme（如HTTP）去下载指定的代码包。如果你用的代码仓库（如公司内部的Gitlab）没有HTTPS支持，可以添加此标记。请在确定安全的情况下使用它。
	
    >5、-t 	让命令程序同时下载并安装指定的代码包中的测试源码文件中依赖的代码包。
	
    >6、-u 	让命令利用网络来更新已有代码包及其依赖包。默认情况下，该命令只会从网络上下载本地不存在的代码包，而不会更新已有的代码包。
	
    >7、-v   标记意味着会打印出被构建的代码包的名字。
	
    >8、-x   可以看到go get命令执行过程中所使用的所有命令。

#### go clean：会删除掉执行其它命令时产生的一些文件和目录
-   >1、带有标记-i，则会同时删除安装当前代码包时所产生的结果文件。
	
    >2、带有标记-r，则还包括当前代码包的所有依赖包的上述目录和文件。
	
    >3、-n会让命令在执行过程中打印用到的系统命令，但不会真正执行它们。
	
    >4、既打印命令又执行命令则需使用标记-x。

#### go doc和godoc：查看程序或者函数的实体

-   >可以打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。

#### go run ：可以编译并运行命令源码文件

-   >由于它其中包含了编译动作，因此它也可以接受所有可用于go build命令的标记。除了标记之外，go run命令只接受Go源码文件作为参数，而不接受代码包。
	
    >与go build命令和go install命令一样，go run命令也不允许多个命令源码文件作为参数，即使它们在同一个代码包中也是如此。而原因也是一致的，多个命令源码文件会都有main函数声明

#### go test：用于对Go语言编写的程序进行测试
-   >这种测试是以代码包为单位的。当然，这还需要测试源码文件的帮助


#### go list：列出指定的代码包的信息
-   >与其他命令相同，我们需要以代码包导入路径的方式给定代码包。被给定的代码包可以有多个

#### go mod 使用规则（已解决）
[参考文章](https://blog.csdn.net/Q_QuanTing/article/details/102958977)

-   go mod init [module_name]	// 创建模块
-   go list -m all 		// 可以列出当前module以及它的所有依赖包
-   go sum 			// 该文件主要包含特定module版本内容的Hash

-   升级依赖包 ！！！！！
    >在Go Module中，版本通过语义版本标签引用，一个语义版本有三个部分：
	major（主版本号）、minor（次版本号）和patch（修订号）。
	例如，对v0.1.2来说：
	其major（主版本号）就是0，其minor（次版本号）就是1，其patch（修订号）就是2。
    
    >补充：
	主版本号：当你做了不兼容的 API 修改，
	次版本号：当你做了向下兼容的功能性新增，
	修订号：当你做了向下兼容的问题修正。

    >升级module的最新版本，直接go get [module_name]
其中indirect表示这个依赖包不是被我们直接使用，而是其他依赖包用到。

-   >如果升级之后的版本不兼容程序，就要查看该module所有的版本，看那个可以用。
	
    **go list -m -versions [module_name]**
    >更改相应的版本号：`go get [module_name]@[版本号]`

	>eg:
```go
    go list -m -versions rsc.io/sampler
	rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99
    go get rsc.io/sampler@v1.3.1
```

-   添加另一个主板本依赖包
    >在文件中加入新的函数，并import新的依赖包名。go run 会直接下载的,因为主版本进行了更新和修正，大多数都是向后兼容的，少部分bug不可避免。
```
   F:\Users\QQT\Documents\Go Projects\NoGOPATH\hello>go list -m rsc.io/q...
	rsc.io/quote v1.5.2
	rsc.io/quote/v3 v3.1.0
```
	
-   查看两个主版本依赖有什么不同，查看module不同版本的方法体差别
	>go doc [module_name]，返回该module的主方法替

-   删除不用的依赖包，虽然进行过引用的更新，但是在go.mod中，引用的module还是在的，尽管不使用，因此要删除
    >go mod tidy
	
-   go mod 总结：
```go mod
	go mod download	        //下载依赖包
	go mod edit		//工具或脚本编辑go.mod
	go mod graph		//打印模块依赖图
	go mod init		//在当前目录初始化mod
	go mod tidy		//拉取缺少的模块，移除不用的模块。
	go mod vendor 		//将依赖复制到vendor下
	go mod verify	  	//验证依赖是否正确
	go mod why 		//解释为什么需要依赖
	go list -m  all	        //依赖详情
```




