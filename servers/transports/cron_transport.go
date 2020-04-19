package transports

import (
	"caixin.app/caixos/tokit/servers/commons"
	"caixin.app/caixos/tokit/servers/transports/codecs"
	"github.com/go-kit/kit/endpoint"
)

func NewCronJob(endpoint endpoint.Endpoint) *commons.Server {
	return commons.NewServer(
		endpoint,
		codecs.CronDecodeRequest,
		codecs.CronEncodeResponse,
	)
}

