package codecs

import (
	"github.com/caixos/tokit/contracts"
	"github.com/caixos/tokit/tools/idwork"
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
