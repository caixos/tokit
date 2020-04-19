package filters

import (
	"caixin.app/tokit"
	"caixin.app/tokit/constant"
	"caixin.app/tokit/contract"
	"caixin.app/tokit/logger"
	"github.com/go-kit/kit/endpoint"
	"errors"
	"context"
)

type ResponseEndpoint struct {
	next endpoint.Endpoint
}

func (s *ResponseEndpoint) Next(next endpoint.Endpoint) contract.IFilter {
	s.next = next
	return s
}

func (s *ResponseEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//全局扑捉错误
		defer func() {
			if err := recover(); err != nil {
				logger.GetLog().Info(err)
				response = contract.MakeResponse(nil, err.(error))
			}
		}()
		if tokit.App.Status == false {
			err := errors.New(constant.ErrStop)
			return contract.MakeResponse(nil, err), nil
		}
		response, err = s.next(ctx, request)
		return contract.MakeResponse(response, err), nil
	}
}
