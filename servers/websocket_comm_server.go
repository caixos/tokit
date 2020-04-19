package servers

import (
	"github.com/caixos/tokit/loggers"
	"github.com/caixos/tokit/servers/commons"
	"github.com/caixos/tokit/servers/transports"
	"github.com/caixos/tokit/servers/websockets"
	"github.com/go-kit/kit/endpoint"
)

//websocket 服务器尽量采用 emqx mqtt broker
type WebSocketCommServer struct {
	*websockets.Server
}

func NewWebSocketCommServer() *WebSocketCommServer {

	ss := &WebSocketCommServer{
		Server: websockets.NewServer(),
	}
	ss.Logger = loggers.GetLog()
	return ss
}

func (s *WebSocketCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewWebSocket(endpoint),
	}
	s.Register(name, handler)
}

func (s *WebSocketCommServer) Load() {
	//注册通用路由
}

func (s *WebSocketCommServer) Start() error {
	return s.Serve()
}
func (s *WebSocketCommServer) Close() {
	s.Server.Close()
}

