package filters

import (
	"caixin.app/tokit/contract"
	"github.com/go-kit/kit/endpoint"
)

func New(controller contract.IController) endpoint.Endpoint {
	return Chain(
		&ResponseEndpoint{},
		&CommEndpoint{Controller: controller},
	)
}

func Auth(controller contract.IController) endpoint.Endpoint {
	return Chain(
		&ResponseEndpoint{},
		&JwtEndpoint{},
		&CommEndpoint{Controller: controller},
	)
}

func Limit(controller contract.IController) endpoint.Endpoint {
	return Chain(
		&ResponseEndpoint{},
		&LimitEndpoint{},
		&CommEndpoint{Controller: controller},
	)
}

