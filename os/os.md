# OS学习记录

---
## Linux IO 模式及select、poll、epoll详解

----
[链接](https://segmentfault.com/a/1190000003063859)

1、概念说明

    1.1 用户空间和内核空间
    现在操作系统都是采用虚拟存储器，那么对32位操作系统而言，它的寻址空间（虚拟存储空间）为4G（2的32次方）。
    操作系统的核心是内核，独立于普通的应用程序，可以访问受保护的内存空间，也有访问底层硬件设备的所有权限。为
    了保证用户进程不能直接操作内核（kernel），保证内核的安全，操心系统将虚拟空间划分为两部分，一部分为内核空
    间，一部分为用户空间。针对linux操作系统而言，将最高的1G字节（从虚拟地址0xC0000000到0xFFFFFFFF），供
    内核使用，称为内核空间，而将较低的3G字节（从虚拟地址0x00000000到0xBFFFFFFF），供各个进程使用，称为用户空间。

    1.2 进程切换
    为了控制进程的执行，内核必须有能力挂起正在CPU上运行的进程，并恢复以前挂起的某个进程执行，这种行为被称为进程切换。
    因此可以说，任何进程都是在操作系统内核的支持下运行的，是与内核密切相关的。

    从一个进程的运行转到另一个进程上运行，这个过程经过下面的这些变化：
        1、保存处理机的上下文，包括程序计数器和其他寄存器
        2、更新PCB信息
        3、把进程的PCB移入相应的队列，如就绪、阻塞等队列
        4、选择另一个就绪的进程执行，并更新其PCB信息
        5、更新内存管理的数据结构
        6、恢复处理机的上下文
    总而言之，就是很耗费资源

    1.3 进程阻塞
    正在执行的进程，由于期待的某些时间未发生，如请求系统资源失败、等待某种操作的完成、新数据尚未到达或无新工作做等，则
    由系统自动执行阻塞原语，是自己由运行态变为阻塞状态，是一种主动行为，阻塞状态不占用CPU资源。

    1.4 文件描述符fd--是一个用于表述指向文件的引用的抽象化概念
    文件描述符在形式上是一个非负整数。实际上，它是一个索引值，指向内核为每一个进程所维护的该进程打开文件的记录表。
    当程序打开一个现有文件或者创建一个新文件时，内核向进程返回一个文件描述符。在程序设计中，一些涉及底层的程序编
    写往往会围绕着文件描述符展开。但是文件描述符这一概念往往只适用于UNIX、Linux这样的操作系统。

    1.5 缓存IO
    缓存 I/O 又被称作标准 I/O，大多数文件系统的默认 I/O 操作都是缓存 I/O。在 Linux 的缓存 I/O 机制中，操作
    系统会将 I/O 的数据缓存在文件系统的页缓存（ page cache ）中，也就是说，数据会先被拷贝到操作系统内核的缓冲
    区中，然后才会从操作系统内核的缓冲区拷贝到应用程序的地址空间。

    缓存 I/O 的缺点：
    数据在传输过程中需要在应用程序地址空间和内核进行多次数据拷贝操作，这些数据拷贝操作所带来的 CPU 以及内存开销是
    非常大的。

2、IO模式

    对于一次IO访问，数据会先被拷贝到操作系统内核的缓存区中，然后才会从操作系统内核的缓冲区拷贝到应用程序的地址空间，
    所以说，当一个read操作发生时，会经历两个阶段：
        1、等待数据准备
        2、将数据从内核拷贝到进程中

    正式因为这两个阶段，linux系统产生了下面五种网络模式的方案：
    - 阻塞
    - 非阻塞
    - IO多路复用
    - 信号驱动IO
    - 异步IO

    阻塞IO（blocking IO）
        当用户进程调用了recvfrom这个系统调用，kernel就开始了IO的第一个阶段：准备数据，这个过程需要等待
        也就是说，数据被拷贝到操作系统内核的缓冲区中是需要一个过程的，而用户晋城这边，整个进程会被阻塞，等
        kernel一支等待的数据准备好了，他舅会将数据从kernel中拷贝到用户内存，然后kernel返回结果，用户解
        除block状态，重新运行起来

    非阻塞IO
        当用户进程发出read操作时，如果kernel中的数据还没有准备好，那么它并不会block用户进程，而是立刻返
        回一个error。从用户进程角度讲 ，它发起一个read操作后，并不需要等待，而是马上就得到了一个结果。用
        户进程判断结果是一个error时，它就知道数据还没有准备好，于是它可以再次发送read操作。一旦kernel中
        的数据准备好了，并且又再次收到了用户进程的system call，那么它马上就将数据拷贝到了用户内存，然后返回。

        所以，nonblocking IO的特点是用户进程需要不断的主动询问kernel数据好了没有。

    IO多路复用（IO multiplexing）
        IO multiplexing就是我们说的select，poll，epoll，有些地方也称这种IO方式为event driven IO。
        select/epoll的好处就在于单个process就可以同时处理多个网络连接的IO。它的基本原理就是select，
        poll，epoll这个function会不断的轮询所负责的所有socket，当某个socket有数据到达了，就通知用户进程。

        当用户进程调用了select，那么整个进程会被block，而同时，kernel会“监视”所有select负责的socket，当
        任何一个socket中的数据准备好了，select就会返回。这个时候用户进程再调用read操作，将数据从kernel拷贝
        到用户进程。
        所以，I/O 多路复用的特点是通过一种机制一个进程能同时等待多个文件描述符，而这些文件描述符（套接字描述符）
        其中的任意一个进入读就绪状态，select()函数就可以返回。
        这个图和blocking IO的图其实并没有太大的不同，事实上，还更差一些。因为这里需要使用两个system call 
        (select 和 recvfrom)，`而blocking IO只调用了一个system call (recvfrom)。但是，用select的优势
        在于它可以同时处理多个connection。`
        所以，如果处理的连接数不是很高的话，使用select/epoll的web server不一定比使用multi-threading + 
        blocking IO的web server性能更好，可能延迟还更大。select/epoll的优势并不是对于单个连接能处理得更快，
        而是在于能处理更多的连接。）
        在IO multiplexing Model中，实际中，对于每一个socket，一般都设置成为non-blocking，但是，如上图所示，
        整个用户的process其实是一直被block的。只不过process是被select这个函数block，而不是被socket IO给block。

    异步IO（asynchronous IO）
        用户进程发起read操作之后，立刻就可以开始去做其它的事。而另一方面，从kernel的角度，当它受到一个asynchronous
        read之后，首先它会立刻返回，所以不会对用户进程产生任何block。然后，kernel会等待数据准备完成，然后将数据拷贝到
        用户内存，当这一切都完成之后，kernel会给用户进程发送一个signal，告诉它read操作完成了。

    总结：
        blocking 和 non-blocking的区别：
        前者会一支block住对应的进程直到操作完成，而non-blocking IO 在kernel还准备数据的情况下会立刻返回
        
        non-blocking IO和asynchronous IO的区别还是很明显的。在non-blocking IO中，虽然进程大部分时间
        都不会被block，但是它仍然要求进程去主动的check，并且当数据准备完成以后，也需要进程主动的再次调用recvfrom
        来将数据拷贝到用户内存。而asynchronous IO则完全不同。它就像是用户进程将整个IO操作交给了他人（kernel）完
        成，然后他人做完后发信号通知。在此期间，用户进程不需要去检查IO操作的状态，也不需要主动的去拷贝数据。

3、IO 多路复用之select、poll、epoll详解

    select、poll、epoll都是多路复用的机制，IO多路复用就是通过一种机制，一个进程可以监视多个描述符，一旦某热描述符就绪，
    能够通知程序进行相应的读写操作。但select，poll，epoll本质上都是同步I/O，因为他们都需要在读写事件就绪后自己负责进行
    读写，也就是说这个读写过程是阻塞的，而异步I/O则无需自己负责进行读写，异步I/O的实现会负责把数据从内核拷贝到用户空间。

    Select：
        select函数监视文件描述符有3类：分别是writefds、readfds、和exceptfds
        调用后select函数会阻塞，直到有描述副就绪（有数据 可读、可写、或者有except），或者超时（timeout指定等待时间，如果立即
        返回设为null即可），函数返回。当select函数返回后，可以 通过遍历fdset，来找到就绪的描述符。

    select目前几乎在所有的平台上支持，其良好跨平台支持也是它的一个优点。select的一个缺点在于单个进程能够监视的文件描述符
    的数量存在最大限制，在Linux上一般为1024，可以通过修改宏定义甚至重新编译内核的方式提升这一限制，但是这样也会造成效率的降低。

    Poll:
        不同与select使用三个位图来表示三个fdset的方式，poll使用一个 pollfd的指针实现。

```c
    struct pollfd{
        int fd;
        short events;
        short revents;
    }
```
        pollfd结构包含了要监视的event和发生的event，不再使用select“参数-值”传递的方式。
        同时，pollfd并没有最大数量限制（但是数量过大后性能也是会下降）。 和select函数一样，
        poll返回后，需要轮询pollfd来获取就绪的描述符。

        从上面看，select和poll都需要在返回后，通过遍历文件描述符来获取已经就绪的socket。
        事实上，同时连接的大量客户端在一时刻可能只有很少的处于就绪状态，因此随着监视的描述符
        数量的增长，其效率也会线性下降。

    Epoll：是之前的select和poll的增强版本
        epoll更加灵活，没有描述符限制。epoll使用一个文件描述符管理多个描述符，将用户关系的文件描述符的事件存放到内核的一个事件表中，
        这样在用户空间和内核空间的copy只需一次。

        一 epoll操作过程
        epoll操作过程需要三个接口，分别如下：
        
        int epoll_create(int size)；//创建一个epoll的句柄，size用来告诉内核这个监听的数目一共有多大
        int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)；
        int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
        1. int epoll_create(int size);
            创建一个epoll的句柄，size用来告诉内核这个监听的数目一共有多大，这个参数不同于select()中的第一个参数，
            给出最大监听的fd+1的值，参数size并不是限制了epoll所能监听的描述符最大个数，只是对内核初始分配内部数据结构的一个建议。
            当创建好epoll句柄后，它就会占用一个fd值，在linux下如果查看/proc/进程id/fd/，是能够看到这个fd的，所以在使用完epoll后，
            必须调用close()关闭，否则可能导致fd被耗尽。
        
        2. int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)；
           函数是对指定描述符fd执行op操作。
        - epfd：是epoll_create()的返回值。
        - op：表示op操作，用三个宏来表示：添加EPOLL_CTL_ADD，删除EPOLL_CTL_DEL，修改EPOLL_CTL_MOD。分别添加、删除和修改对fd的监听事件。
        - fd：是需要监听的fd（文件描述符）
        - epoll_event：是告诉内核需要监听什么事，struct epoll_event结构如下：
```c
    struct epoll_event {
      __uint32_t events;  /* Epoll events */
      epoll_data_t data;  /* User data variable */
    };

    //events可以是以下几个宏的集合：
    EPOLLIN ：表示对应的文件描述符可以读（包括对端SOCKET正常关闭）；
    EPOLLOUT：表示对应的文件描述符可以写；
    EPOLLPRI：表示对应的文件描述符有紧急的数据可读（这里应该表示有带外数据到来）；
    EPOLLERR：表示对应的文件描述符发生错误；
    EPOLLHUP：表示对应的文件描述符被挂断；
    EPOLLET： 将EPOLL设为边缘触发(Edge Triggered)模式，这是相对于水平触发(Level Triggered)来说的。
    EPOLLONESHOT：只监听一次事件，当监听完这次事件之后，如果还需要继续监听这个socket的话，需要再次把这个socket加入到EPOLL队列里

```
        3. int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
        等待epfd上的io事件，最多返回maxevents个事件。
        参数events用来从内核得到事件的集合，maxevents告之内核这个events有多大，这个maxevents的值不能大于创建
        epoll_create()时的size，参数timeout是超时时间（毫秒，0会立即返回，-1将不确定，也有说法说是永久阻塞）。
        该函数返回需要处理的事件数目，如返回0表示已超时
    工作模式有两种：
        LT模式（应用程序收到就绪的通知可以不立即处理该事件）
        ET模式（应用程序收到就绪的通知立即处理该事件）

    EPoll总结：
        在 select/poll中，进程只有在调用一定的方法后，内核才对所有监视的文件描述符进行扫描，而epoll事先通过epoll_ctl()来注册一个文件描述符，
        一旦基于某个文件描述符就绪时，内核会采用类似callback的回调机制，迅速激活这个文件描述符，当进程调用epoll_wait() 时便得到通知。
        (此处去掉了遍历文件描述符，而是通过监听回调的的机制。这正是epoll的魅力所在。)

    epoll的优点主要是一下几个方面：
        1. 监视的描述符数量不受限制，它所支持的FD上限是最大可以打开文件的数目，这个数字一般远大于2048,举个例子,在1GB内存的机器上大约是10万左右，
        具体数目可以cat /proc/sys/fs/file-max察看,一般来说这个数目和系统内存关系很大。select的最大缺点就是进程打开的fd是有数量限制的。
        这对于连接数量比较大的服务器来说根本不能满足。虽然也可以选择多进程的解决方案( Apache就是这样实现的)，不过虽然linux上面创建进程的代价
        比较小，但仍旧是不可忽视的，加上进程间数据同步远比不上线程间同步的高效，所以也不是一种完美的方案。

        IO的效率不会随着监视fd的数量的增长而下降。epoll不同于select和poll轮询的方式，而是通过每个fd定义的回调函数来实现的。
        只有就绪的fd才会执行回调函数。
        如果没有大量的idle -connection或者dead-connection，epoll的效率并不会比select/poll高很多，但是当遇到大量的idle-connection，
        就会发现epoll的效率大大高于select/poll。