package servers

import (
	"github.com/go-kit/kit/endpoint"
	"caixin.app/caixos/tokit/servers/gateways"
	"caixin.app/caixos/tokit/filters"
	"caixin.app/caixos/tokit/loggers"
)

type GateWayCommServer struct {
	*gateways.Server
}

func NewGateWayCommServer() *GateWayCommServer {
	ss := &GateWayCommServer{
		Server: gateways.NewServer(),
	}
	ss.Logger = loggers.GetLog()
	return ss
}

func (s *GateWayCommServer) Route(method, path string, endpoint endpoint.Endpoint) {
	//如果有本地注册的路由,则跑本地,2种情况组合endpoint filter
	//1 跑本地服务
	//2 只跑本地endpoint filter
	s.Register(method, path, endpoint)
}

func (s *GateWayCommServer) Load() {
	//注册通用路由,consul 心跳检测
	s.Route("GET", "/health", (&filters.HealthEndpoint{}).Make())

}

func (s *GateWayCommServer) Start() error {
	return s.Serve()
}
func (s *GateWayCommServer) Close() {
	s.Server.Close()
}
