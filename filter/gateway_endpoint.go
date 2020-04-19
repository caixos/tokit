package filter

import (
	"caixin.app/caixos/tokit/contract"
	"github.com/go-kit/kit/endpoint"
	"context"
)

type GateWayEndpoint struct {
	next endpoint.Endpoint
}

func (s *GateWayEndpoint) Next(next endpoint.Endpoint) contract.IFilter {
	s.next = next
	return s
}

func (s *GateWayEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(contract.Request)
		req.Data["GATEWAY"] = "GATEWAY"
		if s.next == nil {
			response = contract.ResponseSucess(req.Data)
		} else {
			response, err = s.next(ctx, req)
			res := response.(contract.Response)
			if res.Code == "0000" {
				m, b := res.Data.(map[string]interface{})
				if b && m != nil {
					for k, v := range m {
						req.Data[k] = v
					}
					response = contract.ResponseSucess(req.Data)
				}
			}
		}
		return
	}
}
