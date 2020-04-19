package servers

import (
	"github.com/caixos/tokit/loggers"
	"github.com/caixos/tokit/servers/commons"
	"github.com/caixos/tokit/servers/mqtts"
	"github.com/caixos/tokit/servers/transports"
	"github.com/go-kit/kit/endpoint"
)

type MqttSubscribeCommCommServer struct {
	*mqtts.Server
}

func NewMqttSubscribeCommCommServer() *MqttSubscribeCommCommServer {
	ss := &MqttSubscribeCommCommServer{
		Server: mqtts.NewServer(),
	}
	ss.Logger = loggers.GetLog()
	return ss
}

func (s *MqttSubscribeCommCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewMqttSubscribe(endpoint),
	}
	s.Register(name, handler)
}

func (s *MqttSubscribeCommCommServer) Load() {

	//注册通用路由
}

func (s *MqttSubscribeCommCommServer) Start() error {
	return s.Serve()

}

func (s *MqttSubscribeCommCommServer) Close() {
	s.Server.Close()
}
