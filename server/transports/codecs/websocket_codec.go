package codecs

import (
	"caixin.app/caixos/tokit/contract"
	"context"
	"caixin.app/caixos/tokit/tools/idwork"
)

func WebSocketDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	id := idwork.ID()
	return contract.Request{
		Id:   id,
		Data: req.(map[string]interface{}),
	}, nil
}

func WebSocketEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
