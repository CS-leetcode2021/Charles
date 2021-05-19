# Transport

---
### tcp/ip

-   Q:三次握手和四次挥手？

    A:**[重点内容](https://leetcode-cn.com/leetbook/read/networks-interview-highlights/esegch/)**

    *三次握手*：是tcp建立连接的过程
    ![](https://pic.leetcode-cn.com/1614160878-FiFlkq-image.png)

    *四次挥手*：释放tcp连接的过程（客户端主动释放）
    ![](https://pic.leetcode-cn.com/1612459478-ajInIu-%E5%9B%9B%E6%AC%A1%E6%8C%A5%E6%89%8B.png)
    
-   Q：如果三次握手的时候每次握手信息对方没有收到会怎么样？

    A：等待阈值+重传

-   Q：为什么要进行三次握手？两次握手可以吗？

    A：为了确认彼此都是的发送和接受都是正常的，如果只是进行了两次握手，服务器并不知道客户端是否收到自己发送的确认消息，会一直等待并监听，长此以往，这样的端口会越来越多

-   Q： 第 2 次握手传回了 ACK，为什么还要传回 SYN？

    A：ACK 是为了告诉客户端发来的数据已经接收无误，而传回 SYN 是为了告诉客户端，服务端收到的消息确实是客户端发送的消息

-   Q：为什么要四次挥手？

    A：A 和 B 打电话，通话即将结束后，A 说“我没啥要说的了”，B回答“我知道了”，但是 B 可能还会有要说的话，A 不能要求 B 跟着自己的节奏结束通话，于是 B 可能又巴拉巴拉说了一通，最后 B 说“我说完了”，A 回答“知道了”，这样通话才算结束。

-   Q：CLOSE-WAIT 和 TIME-WAIT 的状态和意义？

    A：close-wiat，可能是服务器还有剩余的数据没有传送结束；time-wait是要等待响应的服务器窗口并没欧关闭，需要等待服务器先关闭
    
-   Q：TIME-WAIT 为什么是 2MSL？

    A：因为客户端并不清楚服务器是否收到了自己的确认信息，如果没有收到，需要等待2个tcp最长的周期，万一服务器重新发送一次呢
    
-   Q：TCP 和 UDP 的区别？

    A：[答案解析](https://leetcode-cn.com/leetbook/read/networks-interview-highlights/eswcie/)
    
-   Q：TCP 是如何保证可靠性的？

    A：将数据进行分块编号打包发送，每次接收和确认都要核查信息是否丢失，如果丢失就会触发超时重发机制，并且在网络链路层中还有拥塞控制机制

-   Q：UDP 为什么是不可靠的？bind 和 connect 对于 UDP 的作用是什么

    A：UDP 只有一个 socket 接收缓冲区，没有 socket 发送缓冲区，即只要有数据就发，不管对方是否可以正确接收。而在对方的 socket 接收缓冲区满了之后，新来的数据报无法进入到 socket 接受缓冲区，此数据报就会被丢弃，因此 UDP 不能保证数据能够到达目的地
    ，此外，UDP 也没有流量控制和重传机制。
    
-   Q：TCP 拥塞控制采用的四种算法？

    A：慢开始、拥塞避免、快重传、快恢复

-   Q：TCP 拥塞控制采用的四种算法？

    A：[答案解析](https://leetcode-cn.com/leetbook/read/networks-interview-highlights/esqf51/)

-   Q：SYN FLOOD是什么？

    A：SYN Flood 是种典型的 DoS（拒绝服务）攻击，其目的是通过消耗服务器所有可用资源使服务器无法用于处理合法请求。通过重复发送初始连接请求（SYN）数据包，攻击者能够压倒目标服务器上的所有可用端口，导致目标设备根本不响应合法请求。

-   Q：为什么服务端易受到 SYN 攻击？

    A：三次握手中的第二次握手，服务器发送给客户端的确认链接消息，如果丢失，客户端并没有回复，服务器是要维持数据进行超时等待的，如果大量的虚假IP发送大量的SYN请求，却不回复第三次握手的确认信息，便是服务器资源大量被占有
    
    解决方案：减少超时等待的时间、或者部署能识别恶意攻击的IP地址进行过滤
    
-   Q：高并发服务器客户端主动关闭连接和服务端主动关闭连接的区别？

    A：[答案解析](https://leetcode-cn.com/leetbook/read/networks-interview-highlights/esgkj6/)
    