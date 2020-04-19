package events

import (
	"caixin.app/caixos/tokit/contract"
	"caixin.app/caixos/tokit/tools/idwork"
	"github.com/go-kit/kit/endpoint"
	"runtime"
	"sync"
	"time"
	"context"

)

/**
通过channel方式传递event,而不是通过共享内存传递
*/
var Handlers map[string]endpoint.Endpoint
var eventPool sync.Pool
var eventChan chan *contract.Payload

type Server struct {
	Concurrency int
	After       <-chan time.Time
	Logger      contract.ILogger
}

func NewServer() *Server {
	Handlers = make(map[string]endpoint.Endpoint)
	eventChan = make(chan *contract.Payload, runtime.NumCPU())
	ss := &Server{}
	return ss
}

func (s *Server) Serve() error {
	errChan := make(chan error)
	for i := 0; i < s.Concurrency; i++ {
		go s.handleEventReceive(errChan)
	}
	err := <-errChan
	if err != nil {
		s.Logger.Info(err)
	}
	return nil
}
func (s *Server) handleEventReceive(errChan chan error) {
	for {
		select {
		case event := <-eventChan:
			filter, ok := Handlers[event.Route]
			if ok {
				ctx := context.Background()
				id := idwork.ID()
				request := contract.Request{
					Id:   id,
					Data: event.Params,
				}
				resp, err := filter(ctx, request)
				if err != nil {
					eventPool.Put(event)
					s.Logger.Info("event error:", err)
					//errChan <- err // 退出协程了
				} else {
					s.Logger.Info("event response:", resp)
				}
			}
			eventPool.Put(event)
		case <-s.After:
			s.Logger.Info("event wait ......")
		}
	}
}
func (s *Server) Close() {

}

