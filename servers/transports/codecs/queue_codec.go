package codecs

import (
	"caixin.app/caixos/tokit/contracts"
	"caixin.app/caixos/tokit/tools/idwork"
	"context"
)

func QueueServerDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	id := idwork.ID()
	return contracts.Request{
		Id:   id,
		Data: req.(map[string]interface{}),
	}, nil
}

func QueueServerEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
