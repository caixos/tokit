package servers

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/caixos/tokit/loggers"
	"github.com/caixos/tokit/servers/commons"
	"github.com/caixos/tokit/servers/timers"
	"github.com/caixos/tokit/servers/transports"
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
