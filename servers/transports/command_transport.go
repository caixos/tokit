package transports

import (
	"caixin.app/tokit/servers/commons"
	"github.com/go-kit/kit/endpoint"
	"caixin.app/tokit/servers/transports/codecs"
)

func NewCommand(endpoint endpoint.Endpoint) *commons.Server {
	return commons.NewServer(
		endpoint,
		codecs.CommandDecodeRequest,
		codecs.CommandEncodeResponse,
	)
}
