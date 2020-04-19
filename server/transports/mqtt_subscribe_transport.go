package transports

import (
	"caixin.app/tokit/server/common"
	"caixin.app/tokit/server/transports/codecs"
	"github.com/go-kit/kit/endpoint"
)

func NewMqttSubscribe(endpoint endpoint.Endpoint) *common.Server {
	return common.NewServer(
		endpoint,
		codecs.MqttSubscribeDecodeRequest,
		codecs.MqttSubscribeEncodeResponse,
	)
}