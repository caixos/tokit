package codecs

import (
	"caixin.app/caixos/tokit/contract"
	"caixin.app/caixos/tokit/tools/idwork"
	"context"
)

func QueueServerDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	id := idwork.ID()
	return contract.Request{
		Id:   id,
		Data: req.(map[string]interface{}),
	}, nil
}

func QueueServerEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
