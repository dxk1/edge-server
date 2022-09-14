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

//Test PreHandle
func (this *PingRouter) iceeHandle(request service.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test Handle
func (this *PingRouter) Handle(request service.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test PostHandle
func (this *PingRouter) PostHandle(request service.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	//创建一个server句柄
	// s := net2.NewServer("[zinx V0.3]")
	s := net2.NewServer()

	// s.AddRouter(&PingRouter{})
	s.AddRouter(3, &PingRouter{})
	//2 开启服务
	s.Serve()
}
