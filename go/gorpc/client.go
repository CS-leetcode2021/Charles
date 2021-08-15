package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
 *  @ClassName:client
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/14 下午6:44
 */

//type HelloServiceClient struct {
//	*rpc.Client
//}
////
////var _ HelloServiceInterface = (*HelloServiceClient)(nil)
//
//func DialHelloService(network, address string) (*HelloServiceClient, error) {
//	c, err := rpc.Dial(network, address)
//	if err != nil {
//		return nil, err
//	}
//	return &HelloServiceClient{Client: c}, nil
//}

//func (p *HelloServiceClient) Hello(request string, reply *string) error {
//	return p.Client.Call(HelloServiceName+".Hello", request, reply)
//}

func main() {

	// client, err := DialHelloService("tcp", "localhost:1234")
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

}
