package filter

import (
	"caixin.app/tokit/contract"
	"github.com/go-kit/kit/endpoint"
)

func Chain(endpoints ...contract.IFilter) endpoint.Endpoint {
	len := len(endpoints) - 1
	for i := 0; i < len; i++ {
		endpoints[i].Next(endpoints[i+1].Make())
	}
	return endpoints[0].Make()
}
