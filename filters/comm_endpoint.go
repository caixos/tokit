package filters

import (
	"caixin.app/tokit/args"
	"caixin.app/tokit/constant"
	"caixin.app/tokit/contract"
	"caixin.app/tokit/logger"
	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
	"context"
)

type CommEndpoint struct {
	Controller contract.IController
	next       endpoint.Endpoint
}

func (s *CommEndpoint) Next(next endpoint.Endpoint) contract.IFilter {
	s.next = next
	return s
}

func (s *CommEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//生成请求参数
		req := request.(contract.Request)
		_, has := (req.Data)["mock"]
		if has && args.Mode == "dev" {
			if v, ok := s.Controller.(contract.IMock); ok {
				return v.Mock(), nil
			}
		}
		//生成context上下文
		cc := s.makeContext(ctx, req)
		//成成线程log,统一处理ip,request_id等
		cc.Log = s.makeLog(req)
		//参数验证
		err := s.valid(cc, req)
		if err != nil {
			cc.Log.Info(err.Error())
			return nil, err
		}
		//逻辑处理
		ret, err := s.Controller.Handle(cc)
		if err != nil {
			cc.Log.Info(err.Error())
		}
		return ret, err
	}
}
func (s *CommEndpoint) valid(ctx contract.Context, request contract.Request) error {
	var obj interface{}
	if v, ok := s.Controller.(contract.IValid); ok {
		obj = v.GetRules()
	}
	if obj != nil {
		// Map2Struct 时自动验证
		err := convert.Map2Struct(request.Data, obj)
		if err != nil {
			return err
		}
		ctx.Set(constant.RequestDto, obj)
	}
	return nil
}

func (s *CommEndpoint) makeLog(req contract.Request) *logrus.Entry {
	//初始化日志字段,放到context中
	ip, has := (req.Data)["client_ip"]
	if !has || ip == nil {
		ip = "LAN"
	}
	entity := logger.GetLog().WithFields(logrus.Fields{
		"request_id": req.Id,
		"client_ip":  ip,
	})
	return entity
}

func (s *CommEndpoint) makeContext(ctx context.Context, req contract.Request) contract.Context {
	cc := contract.Context{
		Context: ctx,
		Keys:    make(map[string]interface{}),
	}
	cc.Set("request", req.Data)
	cc.Set("request.id", req.Id)

	return cc
}
