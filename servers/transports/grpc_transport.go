package transports

import (
	GrpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/endpoint"
	"caixin.app/tokit/servers/transports/codecs"
)

func NewGRPC(endpoint endpoint.Endpoint) *GrpcTransport.Server {
	return GrpcTransport.NewServer(
		endpoint,
		codecs.GprcDecodeRequest,
		codecs.GprcEncodeResponse,
	)
}

