# 链路层

-   MAC地址和IP地址分别对应的是物理地址和网络中的逻辑地址

-   Q：为什么有了MAC地址还需要IP地址？

    A：MAC地址是48位的，需要我们记住每个子网的地址，内存消耗太大，对于路由器是不现实的，和 MAC 地址不同，IP 地址是和地域相关的，在一个子网中的设备，我们给其分配的 IP 地址前缀都是一样的，这样路由器就能根据 IP 地址的前缀知道这个设备属于哪个子网，剩下的寻址就交给子网内部实现，从而大大减少了路由器所需要的内存。

-   Q：路由器和交换机的区别？

    A：交换机主要工作在数据链路层，对象是MAC地址，目的是组建局域网

      路由器主要工作在网络层，对象是IP地址，目的是将局域网接入到网络中
    
      DNS主要是将域名和IP地址进行互换，比如：www.google.com-->116.213.120.232
    
-   为什么有了 IP 地址还需要 MAC 地址？

    A：同一个局域网的IP前缀是一样的，我们需要MAC地址来区分不同的设备


-   Q：私网地址和公网地址之间进行转换：同一个局域网内的两个私网地址，经过转换之后外面看到的一样吗？

    A：如果采用的是静态和动态转换，外面看到的是不一样的，如果采用多路复用，外面看到的公网地址是一样的，但是端口号是不一样的

-   Q：URL学习？其实就是网页的网址

    A：[参考解析](https://wangdoc.com/html/url.html)

