package transports

import (
	"caixin.app/caixos/tokit/server/common"
	"github.com/go-kit/kit/endpoint"
	"caixin.app/caixos/tokit/server/transports/codecs"
)

func NewCommand(endpoint endpoint.Endpoint) *common.Server {
	return common.NewServer(
		endpoint,
		codecs.CommandDecodeRequest,
		codecs.CommandEncodeResponse,
	)
}
