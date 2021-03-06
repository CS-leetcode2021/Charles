# 四次握手

-------

![](https://user-gold-cdn.xitu.io/2019/11/29/16eb5e598f308f2d?imageView2/0/w/1280/h/960/format/webp/ignore-error/1)

    1、客户端Client进程发送连接释放报文，并且停止发送数据。其中FIN标志位为1,顺序号为Seq = m （等于前面已经传送过来的数据的最后一个字节的序号加1），
    此时，客户端Client进入FIN—WAIT-1（终止等待1）状态。 TCP规定，FIN报文段即使不携带数据，也要消耗一个序号。

    2、服务器Server收到连接释放报文，发出确认报文，ACK=1，ack=m+1，并且带上自己的顺序号seq=n，此时，服务器Server就进入了CLOSE-WAIT（关闭等待）状态。
    TCP服务器通知高层的应用进程，客户端Client向服务器的方向就释放了，这时候处于半关闭状态，即客户端Client已经没有数据要发送了，但是服务器Server若发送数据，
    客户端Client依然要接受。这个状态还要持续一段时间，也就是整个CLOSE-WAIT状态持续的时间。

    3、客户端Client收到服务器Server的确认信息后，此时，客户端Client就进入FIN-WAIT-2（终止等待2）状态，等待服务器Server发送连接释放报文
    （在这之前还需要接受服务器Server发送的最后的数据）。
    
    4、服务器Server将最后的数据发送完毕后，就向客户端发送连接释放报文，FIN=1，ack=m+1，由于在半关闭状态，服务器Server很可能又发送了一些数据，
    假定此时的顺序号为seq=p，此时，服务器Server就进入了LAST-ACK（最后确认）状态，等待客户端Client的确认。

    5、客户端Client收到服务器Server的连接释放报文后，必须发出确认，ACK=1，ack=p+1，而自己的顺序号是seq=m+1，此时，客户端Client就进入了TIME-WAIT
    （时间等待）状态。注意此时TCP连接还没有释放，必须经过2*MSL（最长报文段寿命）的时间后，当客户端Client撤销相应的TCB（保护程序）后，才进入CLOSED状态。
    服务器Server只要收到了客户端Client发出的确认，立即进入CLOSED状态。同样，撤销TCB后，就结束了这次的TCP连接。
    
    可以看到，服务器Server结束TCP连接的时间要比客户端Client早一些。


-----
# 常见的面试题
    
    1.为什么连接的时候是三次握手，关闭的时候却是四次握手？
    答：因为当客户端发起关闭连接的请求时，发出的FIN，仅代表客户端没有需要发送给服务器端的数据了。而如果服务器端如果仍有数据需要发送给客户端的话，
    响应报文ACK和结束报文FIN则就不能同时发送给客户端了。此时，服务器端会先返回一个响应报文，代表接收到了客户端发出的FIN请求，而后在数据传输完了
    之后，再发出FIN请求，表示服务器端已经准备好断开连接了。所以关闭连接的时候是四次握手。

    2.为什么TIME_WAIT状态需要经过2MSL(最大报文段生存时间)才能返回到CLOSE状态？
    答：按照前面所说，当四个报文全部发送完毕后，理论上就算是结束了。但是实际情况往往不会那么可靠，比如最后一条报文发出后丢失了，
    那么服务器端就不会接收到这一报文，每隔一段时间，服务器端会再次发出FIN报文，此时如果客户端已经断开了，那么就无法响应服务器的二次请求，
    这样服务器会继续发出FIN报文，从而变成了死循环。所以需要设置一个时间段，如果在这个时间段内接收到了服务器端的再次请求，则代表客户端发出的ACK报文
    没有接收成功。反之，则代表服务器端成功接收响应报文，客户端进入CLOSED状态，此次连接成功关闭。而这个时间，就规定为了2MSL，即客户端发出ACK报文到
    服务器端的最大时间 + 服务器没有接收到ACK报文再次发出FIN的最大时间 = 2MSL

    3.为什么不能用两次握手进行连接？
    答：三次握手有两个重要的功能，一是要双方做好发送数据的准备工作且双方都知道彼此已准备好，二要允许双方就初始顺序号进行协商，这个顺序号在握手过程中
    被发送和确认。如果改为了两次握手，是有可能发生死锁的。在两次握手的设定下，服务器端在成功接受客户端的连接请求SYN后，向客户端发出ACK确定报文时，
    如果因为网络原因客户端没有接收到，则会一直等待服务器端的ACK报文，而服务器端则认为连接成功建立了，便开始向客户端发送数据。
    但是客户端因为没有收到服务器端的ACK报文，且不知道服务器的顺序号seq，则会认为连接未成功建立，忽略服务器发出的任何数据。
    如此客户端一直等待服务器端的ACK报文，而服务器端因为客户端一直没有接收数据，而不断地重复发送数据，从而造成死锁。
    
    4.如果已经建立了连接，但是客户端突然出现故障了怎么办？
    答：TCP还设有一个保活计时器，显然，客户端如果出现故障，服务器不能一直等下去，白白浪费资源。服务器每收到一次客户端的请求后都会重新复位这个计时器，
    时间通常是设置为2小时，若两小时还没有收到客户端的任何数据，服务器就会发送一个探测报文段，以后每隔75秒钟发送一次。
    若一连发送10个探测报文仍然没反应，服务器就认为客户端出了故障，接着就关闭连接。
