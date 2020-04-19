package timers

import (
	"github.com/caixos/tokit/contracts"
	"github.com/caixos/tokit/servers/commons"
	"github.com/caixos/tokit/tools/idwork"
	"time"
	"context"
)

type Server struct {
	handlers map[string]*service
	Logger   contracts.ILogger
}
type service struct {
	freq    int
	handler *commons.CommHandler
	params  map[string]interface{}
}

func NewServer() *Server {
	ss := &Server{
		handlers: make(map[string]*service),
	}

	return ss
}
func (s *Server) Register(name string, freq int, handler *commons.CommHandler, params map[string]interface{}) {

	s.handlers[name] = &service{
		freq:    freq,
		handler: handler,
		params:  params,
	}
}

func (s *Server) Serve() error {
	errChans := make(map[string]chan error)
	for name, svr := range s.handlers {
		errChans[name] = make(chan error)
		ticker := time.NewTicker(time.Duration(svr.freq) * time.Second)
		go func(name string, svr *service, t *time.Ticker, errChan chan error) {
			for {
				select {
				case <-t.C:
					id := idwork.ID()
					ctx := context.Background()
					params := svr.params
					params["request_id"] = id
					resp, err := svr.handler.Handle(ctx, params)
					if err != nil {
						s.Logger.Info(err.Error())
					} else {
						s.Logger.Info("定时任务:", resp)
					}
				}
			}
		}(name, svr, ticker, errChans[name])
	}
	for _, errChan := range errChans {
		e := <-errChan
		if e != nil {
			return e
		}
	}
	return nil
}

func (s *Server) Close() {

}
