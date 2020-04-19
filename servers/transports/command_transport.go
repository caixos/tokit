package transports

import (
	"caixin.app/caixos/tokit/servers/commons"
	"github.com/go-kit/kit/endpoint"
	"caixin.app/caixos/tokit/servers/transports/codecs"
)

func NewCommand(endpoint endpoint.Endpoint) *commons.Server {
	return commons.NewServer(
		endpoint,
		codecs.CommandDecodeRequest,
		codecs.CommandEncodeResponse,
	)
}
