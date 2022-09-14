package zrouter

import (
	"github.com/dxk1/edge-server/log"
	"github.com/dxk1/edge-server/net"
	"github.com/dxk1/edge-server/service"
)

//ping test 自定义路由
type PingRouter struct {
	net.BaseRouter
}

//Ping Handle
func (this *PingRouter) Handle(request service.IRequest) {

	log.Debug("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	log.Debug("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping[FromServer]"))
	if err != nil {
		log.Error(err)
	}
}
