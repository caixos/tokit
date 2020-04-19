package transports

import (
	"caixin.app/tokit/server/common"
	"github.com/go-kit/kit/endpoint"
	"caixin.app/tokit/server/transports/codecs"
)

func NewCommand(endpoint endpoint.Endpoint) *common.Server {
	return common.NewServer(
		endpoint,
		codecs.CommandDecodeRequest,
		codecs.CommandEncodeResponse,
	)
}
