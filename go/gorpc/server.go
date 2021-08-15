package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
 *  @ClassName:main
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/14 下午6:37
 */



func main() {
	// rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数
	// 所有注册的方法会放在“HelloService”服务空间之下。
	rpc.RegisterName("HelloService", new(HelloService))
	//err := RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("Listen error:", err)
	}

	num := 0
	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// 然后我们建立一个唯一的TCP链接，并且通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
		//rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		fmt.Println(num)
		num++

	}
}
