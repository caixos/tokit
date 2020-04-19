package commands

import (
	"github.com/caixos/tokit/servers/commons"
	"github.com/caixos/tokit/args"
	"github.com/caixos/tokit/constants"
	"github.com/caixos/tokit/contracts"
	"context"
	"errors"

)

type Server struct {
	handlers map[string]*commons.CommHandler
	Logger   contracts.ILogger
}

func NewServer() *Server {
	//初始化,loggers,redis池
	s := &Server{
		handlers: make(map[string]*commons.CommHandler),
	}
	return s
}

func (s *Server) Register(name string, handler *commons.CommHandler) {
	s.handlers[name] = handler

}

func (s *Server) Serve() error {
	if args.Cmd != "" {
		//调用服务
		handler, isExist := s.handlers[args.Cmd]
		if isExist == false {
			return errors.New(constants.ErrRoute)
		}
		ctx := context.Background()
		response, err := handler.Handle(ctx, args.Args)
		if err != nil {
			return err
		}
		s.Logger.Info(response)
	}
	return nil
}
func (s *Server) Close() {

}
