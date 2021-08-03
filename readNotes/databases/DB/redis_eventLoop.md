# Redis-EventLoop

---

关于Redis的主流程相关的内容：

1、启动流程：

    redis-server的入口函数main在server.c中，主要做了下面工作(具体看源码):

    初始化配置为默认值: initServerConfig()
    解析命令行参数
    加载配置文件: loadServerConfig(configfile, options)
    初始化server: initServer()
    启动事件循环: aeMain(server.el)

2、server初始化：
    
    在server.c中定义了一个全局变量struct redisServer server保存了Redis运行时的状态和相关配置。
    server初始化流程如下:

    initServerConfig(): 使用常量默认值初始化server
    解析option: 解析命令行参数中格式--[name] value为name value\n，生成options
    loadServerConfig(configfile, options): 加载配置文件，并appendoptions生成config，解析
    config设置server
    
    initServer(): 初始化需要动态设置的部分:
        设置信号处理函数
        创建eventloop: aeCreateEventLoop(server.maxclients + CONFIG_FDSET_INCR)
        初始化db
        添加事件

3、EventLoop
    
    Redis使用I/O多路复用来处理两类事件：
        TimeEvent
        FileEvent
    在redisServer结构体中包含aeEventLoop *el这一成员，结构如下:

```
/* State of an event based program */
typedef struct aeEventLoop {
    int maxfd;   /* highest file descriptor currently registered */
    int setsize; /* max number of file descriptors tracked */
    long long timeEventNextId;
    time_t lastTime;     /* Used to detect system clock skew */
    aeFileEvent *events; /* Registered events */
    aeFiredEvent *fired; /* Fired events */
    aeTimeEvent *timeEventHead;
    int stop;
    void *apidata; /* This is used for polling API specific data */
    aeBeforeSleepProc *beforesleep;
    aeBeforeSleepProc *aftersleep;
} aeEventLoop;
```

    server.el使用AeEventLoop *aeCreateEventLoop(int setsize)创建，需要关注的是
    aeApiCreate函数，Redis为了实现跨平台， 封装了多个平台下相应的系统调用供EventLoop
    使用，见ae.c和config.h。生产环境中最常用的还是Linux，所以主要关注epoll。 

    ae_epoll.c中有如下实现：
    
```
typedef struct aeApiState {
    int epfd;
    struct epoll_event *events;
} aeApiState;
```

    aeApiState中封装了epoll所需要的结构，aeApiPoll()会调用对应的底层事件循环

4、TimeEvent

    时间事件在aeEventLoop中以链表保存，aeCreateTimeEvent()会将新创建的时间事件添加在链表头：

```
/* Time event structure */
typedef struct aeTimeEvent {
    long long id; /* time event identifier. */
    long when_sec; /* seconds */
    long when_ms; /* milliseconds */
    aeTimeProc *timeProc;
    aeEventFinalizerProc *finalizerProc;
    void *clientData;
    struct aeTimeEvent *next;
} aeTimeEvent;
```
    在initServer()中添加了几个事件，其中包含一个时间事件——serverCron。 serverCron以每秒
    server.hz次数执行，主要做了以下工作：

    信息统计
    客户端连接的处理：clientsCron()
    数据库的调整：databasesCron()
    持久化相关工作
    还有一些更复杂的工作，如复制和集群相关工作


5、FileEvent
    
    文件时间处理与套接字相关的工作，在aeEventLoop中以数组的形式保存，redis.conf中可以设置
    maxClients，根据这个配置分配数组的大小，对应fd的时间就保存 在events[fd]中，因为fd会按
    照从小到大分配。
    结构如下:

```
/* File event structure */
typedef struct aeFileEvent {
    int mask; /* one of AE_(READABLE|WRITABLE) */
    aeFileProc *rfileProc;
    aeFileProc *wfileProc;
    void *clientData;
} aeFileEvent;

typedef struct aeFiredEvent {
    int fd;
    int mask;
} aeFiredEvent;
```

    eFileEvent、aeFiredEvent和aeApiState三个结构实现了通用的事件循环机制：

    创建事件会添加到server.el.events和server.el.apidata中
    底层调用依赖对应的aeApiState，并将触发事件的套接字还有相应事件保存在server.el.fired
    处理事件时通过server.el.fired，在server.el.events中找到对应的handler执行
    在initServer()中创建了多个监听套接字，并创建了FileEvent用于接收客户端连接。

6、事件处理

    在初始化完成后，Redis就一直在aeMain()中处理事件循环：

```
void aeMain(aeEventLoop *eventLoop) {
    eventLoop->stop = 0;
    while (!eventLoop->stop) {
        if (eventLoop->beforesleep != NULL)
            eventLoop->beforesleep(eventLoop);
        aeProcessEvents(eventLoop, AE_ALL_EVENTS|AE_CALL_AFTER_SLEEP);
    }
}
```

    aeProcessEvents()为处理事件的函数，流程如下：
    
    查找最近的一个TimeEvent，因为以链表形式保存，耗时O(n)
    以最近的TimeEvent时间间隔为参数调用aeApiPoll()，保证能及时处理
    处理FileEvents
    处理TimeEvents: processTimeEvents()会处理所有到时的事件，并重新添加到链表头