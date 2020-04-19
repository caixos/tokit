package filter

import (
	"caixin.app/caixos/tokit/contract"
	"github.com/go-kit/kit/endpoint"
	"context"
)

type HealthEndpoint struct {
}

func (s *HealthEndpoint) Next(next endpoint.Endpoint) contract.IFilter {
	return s
}

func (s *HealthEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return contract.ResponseSucess("SERVING"), nil
	}
}
