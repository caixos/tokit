package service

import (
	"context"
	"caixin.app/caixos/tokit/contract"
)

func New(service contract.IService) *commonService {
	var s []contract.IService
	s = append(s, service)
	return &commonService{
		services: s,
	}
}

func Pipe() *commonService {
	var s []contract.IService
	return &commonService{
		services: s,
	}
}

type commonService struct {
	services []contract.IService
}

func (s *commonService) Middle(services ...contract.IService) *commonService {
	for _, service := range services {
		s.services = append(s.services, service)
	}
	return s
}

func (s *commonService) Call(ctx contract.Context) error {
	return s.services[0].Handle(ctx)
}

func (s *commonService) Line(ctx contract.Context) error {
	for _, service := range s.services {
		err := service.Handle(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *commonService) Parallel(ctx contract.Context) error {

	type st struct {
		contract.Context
		err error
	}
	ch := make([]chan st, len(s.services))
	for k, service := range s.services {
		ch[k] = make(chan st)
		cc := contract.Context{
			Context: context.Background(),
			Log:     ctx.Log,
			Keys:    ctx.Keys,
		}
		go func(cc contract.Context, s contract.IService, c chan st) {
			err := s.Handle(cc)
			ret := st{
				Context: cc,
				err:     err,
			}
			c <- ret
		}(cc, service, ch[k])
	}
	m := make(map[string]interface{})
	for _, c := range ch {
		res := <-c
		if res.err != nil {
			return res.err
		}
		for key, value := range res.Keys {
			m[key] = value
		}
	}
	ctx.Keys = m
	return nil
}
