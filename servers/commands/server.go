package commands

import (
	"caixin.app/tokit/servers/commons"
	"caixin.app/tokit/args"
	"caixin.app/tokit/constant"
	"caixin.app/tokit/contract"
	"context"
	"errors"

)

type Server struct {
	handlers map[string]*commons.CommHandler
	Logger   contract.ILogger
}

func NewServer() *Server {
	//初始化,logger,redis池
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
			return errors.New(constant.ErrRoute)
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
