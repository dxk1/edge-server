package zrouter

import (
	"github.com/dxk1/edge-server/log"
	"github.com/dxk1/edge-server/net"
	"github.com/dxk1/edge-server/service"
)

type HelloZinxRouter struct {
	net.BaseRouter
}

//HelloZinxRouter Handle
func (this *HelloZinxRouter) Handle(request service.IRequest) {
	log.Debug("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	log.Debug("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(1, []byte("Hello Zinx Router V0.10"))
	if err != nil {
		log.Error(err)
	}
}
