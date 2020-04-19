package commands

import (
	"caixin.app/caixos/tokit/server/common"
	"caixin.app/caixos/tokit/args"
	"caixin.app/caixos/tokit/constant"
	"caixin.app/caixos/tokit/contract"
	"context"
	"errors"

)

type Server struct {
	handlers map[string]*common.CommHandler
	Logger   contract.ILogger
}

func NewServer() *Server {
	//初始化,logger,redis池
	s := &Server{
		handlers: make(map[string]*common.CommHandler),
	}
	return s
}

func (s *Server) Register(name string, handler *common.CommHandler) {
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
