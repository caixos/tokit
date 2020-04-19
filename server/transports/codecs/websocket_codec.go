package codecs

import (
	"caixin.app/tokit/contract"
	"context"
	"caixin.app/tokit/tools/idwork"
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
