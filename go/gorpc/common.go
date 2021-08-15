package main

import "net/rpc"

/**
 *  @ClassName:common.go
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/14 下午7:10
 */

// 服务器的名字
const HelloServiceName = "path/to/pkg.HelloService"

// 服务器要实现的方法
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// 注册函数
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
