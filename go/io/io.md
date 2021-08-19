# GO IO

- 阻塞IO
- 非阻塞IO
- 信号驱动
- 异步IO
- IO多路复用
    - select 最多同时监听1024个fd，内存拷贝开销大（需要维护一个较大的数据结构，该结构拷贝到内核中），时间复杂度高，需要每次遍历
    - poll 用链表的形式摆脱了数量上限

### 实现以下5个接口：

- func netpollinit()
    - 初始化网络轮询器，通过 sync.Once 和 netpollInited 变量保证函数只会调用一次；

- func netpollopen(fd uintptr, pd *pollDesc) int32
    - 监听文件描述符上的边缘触发事件，创建事件并加入监听；

- func netpoll(delta int64) gList
    - 轮询网络并返回一组已经准备就绪的 Goroutine，传入的参数会决定它的行为
        - 如果参数小于 0，无限期等待文件描述符就绪
        - 如果参数等于 0，非阻塞地轮询网络
        - 如果参数大于 0，阻塞特定时间轮询网络

- func netpollBreak()
    - 唤醒网络轮询器，例如：计时器向前修改时间时会通过该函数中断网络轮询器

- func netpollIsPollDescriptor(fd uintptr) bool
    - 判断文件描述符是否被轮询器使用

- 数据结构：
    - pollDesc/pollCache

### 多路复用

- 网络轮询器实际上是对 I/O 多路复用技术的封装，本节将通过以下的三个过程分析网络轮询器的实现原理
    - 网络轮询器的初始化
    - 如何向网络轮询器加入待监控的任务
    - 如何从网络轮询器获取触发的事件

- 因为文件 I/O、网络 I/O 以及计时器都依赖网络轮询器，所以 Go 语言会通过以下两条不同路径初始化网络轮询器：
    - internal/poll.pollDesc.init — 通过 net.netFD.init 和 os.newFile 初始化网络 I/O 和文件 I/O 的轮询信息时；
    - runtime.doaddtimer — 向处理器中增加新的计时器时；
    
- runtime.netpollGenericInit 会调用平台上特定实现的 runtime.netpollinit，即 Linux 上的 epoll，它主要做了以下几件事情：
    - 是调用 epollcreate1 创建一个新的 epoll 文件描述符，这个文件描述符会在整个程序的生命周期中使用
    - 通过 runtime.nonblockingPipe 创建一个用于通信的管道
    - 使用 epollctl 将用于读取数据的文件描述符打包成 epollevent 事件加入监听
    
- 轮询事件
    - 调用 internal/poll.pollDesc.init 初始化文件描述符时不止会初始化网络轮询器，还会通过 runtime.poll_runtime_pollOpen 重置轮询信息 runtime.pollDesc 并调用 runtime.netpollopen 初始化轮询事件：
    - runtime.netpollopen 的实现非常简单，它会调用 epollctl 向全局的轮询文件描述符 epfd 中加入新的轮询事件监听文件描述符的可读和可写状态
    - 从全局的 epfd 中删除待监听的文件描述符可以使用 runtime.netpollclose，因为该函数的实现与 runtime.netpollopen 比较相似，所以这里不展开分析了
    
