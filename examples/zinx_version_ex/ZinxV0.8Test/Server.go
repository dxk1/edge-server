package main

import (
	"fmt"
	net2 "github.com/dxk1/edge-server/net"
	"github.com/dxk1/edge-server/service"
)

//ping test 自定义路由
type PingRouter struct {
	net2.BaseRouter
}

//Ping Handle
func (this *PingRouter) Handle(request service.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	net2.BaseRouter
}

//HelloZinxRouter Handle
func (this *HelloZinxRouter) Handle(request service.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.8"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//创建一个server句柄
	s := net2.NewServer()

	//配置路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//开启服务
	s.Serve()
}
