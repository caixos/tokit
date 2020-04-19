package transports

import (
	"github.com/caixos/tokit/servers/commons"
	"github.com/caixos/tokit/servers/transports/codecs"
	"github.com/go-kit/kit/endpoint"
)

func NewQueue(endpoint endpoint.Endpoint) *commons.Server {
	return commons.NewServer(
		endpoint,
		codecs.QueueServerDecodeRequest,
		codecs.QueueServerEncodeResponse,
	)
}

