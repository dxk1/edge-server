/**
* @Author: Aceld
* @Date: 2020/12/24 00:24
* @Mail: danbing.at@gmail.com
*    zinx server demo
 */
package main

import (
	"edge-server/examples/zinx_server/zrouter"
	"edge-server/log"
	"edge-server/net"
	"edge-server/service"
)

//创建连接的时候执行
func DoConnectionBegin(conn service.IConnection) {
	log.Debug("DoConnecionBegin is Called ... ")

	//设置两个链接属性，在连接创建之后
	log.Debug("Set conn Name, Home done!")
	conn.SetProperty("Name", "Aceld")
	conn.SetProperty("Home", "https://www.kancloud.cn/@aceld")

	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		log.Error(err)
	}
}

//连接断开的时候执行
func DoConnectionLost(conn service.IConnection) {
	//在连接销毁之前，查询conn的Name，Home属性
	if name, err := conn.GetProperty("Name"); err == nil {
		log.Error("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		log.Error("Conn Property Home = ", home)
	}

	log.Debug("DoConneciotnLost is Called ... ")
}

func main() {
	//创建一个server句柄
	s := net.NewServer()

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//配置路由
	s.AddRouter(0, &zrouter.PingRouter{})
	s.AddRouter(1, &zrouter.HelloZinxRouter{})

	//开启服务
	s.Serve()
}
