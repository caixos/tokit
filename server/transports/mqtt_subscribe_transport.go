package transports

import (
	"caixin.app/caixos/tokit/server/common"
	"caixin.app/caixos/tokit/server/transports/codecs"
	"github.com/go-kit/kit/endpoint"
)

func NewMqttSubscribe(endpoint endpoint.Endpoint) *common.Server {
	return common.NewServer(
		endpoint,
		codecs.MqttSubscribeDecodeRequest,
		codecs.MqttSubscribeEncodeResponse,
	)
}
