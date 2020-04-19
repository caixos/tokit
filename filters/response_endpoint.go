package filters

import (
	"caixin.app/caixos/tokit"
	"caixin.app/caixos/tokit/constants"
	"caixin.app/caixos/tokit/contracts"
	"caixin.app/caixos/tokit/loggers"
	"github.com/go-kit/kit/endpoint"
	"errors"
	"context"
)

type ResponseEndpoint struct {
	next endpoint.Endpoint
}

func (s *ResponseEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	s.next = next
	return s
}

func (s *ResponseEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//全局扑捉错误
		defer func() {
			if err := recover(); err != nil {
				loggers.GetLog().Info(err)
				response = contracts.MakeResponse(nil, err.(error))
			}
		}()
		if tokit.App.Status == false {
			err := errors.New(constants.ErrStop)
			return contracts.MakeResponse(nil, err), nil
		}
		response, err = s.next(ctx, request)
		return contracts.MakeResponse(response, err), nil
	}
}
