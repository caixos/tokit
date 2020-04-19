package transports

import (
	"caixin.app/tokit/servers/commons"
	"caixin.app/tokit/servers/transports/codecs"
	"github.com/go-kit/kit/endpoint"
)

func NewMqttSubscribe(endpoint endpoint.Endpoint) *commons.Server {
	return commons.NewServer(
		endpoint,
		codecs.MqttSubscribeDecodeRequest,
		codecs.MqttSubscribeEncodeResponse,
	)
}
