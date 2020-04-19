package filters

import (
	"github.com/caixos/tokit/contracts"
	"github.com/go-kit/kit/endpoint"
	"context"
)

type HealthEndpoint struct {
}

func (s *HealthEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	return s
}

func (s *HealthEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return contracts.ResponseSucess("SERVING"), nil
	}
}
