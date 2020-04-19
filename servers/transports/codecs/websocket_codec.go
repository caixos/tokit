package codecs

import (
	"github.com/caixos/tokit/contracts"
	"context"
	"github.com/caixos/tokit/tools/idwork"
)

func WebSocketDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	id := idwork.ID()
	return contracts.Request{
		Id:   id,
		Data: req.(map[string]interface{}),
	}, nil
}

func WebSocketEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
