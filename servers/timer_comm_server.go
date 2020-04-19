package servers

import (
	"github.com/go-kit/kit/endpoint"
	"caixin.app/caixos/tokit/loggers"
	"caixin.app/caixos/tokit/servers/commons"
	"caixin.app/caixos/tokit/servers/timers"
	"caixin.app/caixos/tokit/servers/transports"
)

/**
定时器
*/
type TimerCommServer struct {
	*timers.Server
}

func NewTimerCommServer() *TimerCommServer {
	ss := &TimerCommServer{
		Server: timers.NewServer(),
	}
	ss.Logger = loggers.GetLog()
	return ss
}

func (s *TimerCommServer) Load() {

	//注册通用路由
}

func (s *TimerCommServer) Route(name string, freq int, endpoint endpoint.Endpoint, params map[string]interface{}) {

	handler := &commons.CommHandler{
		Handler: transports.NewTimer(endpoint),
	}
	s.Register(name, freq, handler, params)
}

func (s *TimerCommServer) Start() error {
	return s.Serve()
}
func (s *TimerCommServer) Close() {
	s.Server.Close()
}
