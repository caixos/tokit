package transports

import (
	"caixin.app/caixos/tokit/servers/commons"
	"caixin.app/caixos/tokit/servers/transports/codecs"
	"github.com/go-kit/kit/endpoint"
)

func NewTimer(endpoint endpoint.Endpoint) *commons.Server {
	return commons.NewServer(
		endpoint,
		codecs.TimerDecodeRequest,
		codecs.TimerEncodeResponse,
	)
}
