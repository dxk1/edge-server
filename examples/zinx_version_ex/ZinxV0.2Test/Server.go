package main

import (
	net2 "github.com/dxk1/edge-server/net"
)

//Server 模块的测试函数
func main() {

	/*
		服务端测试
	*/
	//1 创建一个server 句柄 s
	// s := znet.NewServer("[zinx V0.2]")

	s := net2.NewServer()

	//2 开启服务
	s.Serve()
}
