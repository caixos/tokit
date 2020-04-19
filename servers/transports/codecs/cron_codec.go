package codecs

import (
	"caixin.app/caixos/tokit/contracts"
	"caixin.app/caixos/tokit/tools/idwork"
	"context"
)

func CronDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(map[string]interface{})
	request["request_id"] = idwork.ID()
	return contracts.Request{
		Id:   request["request_id"].(string),
		Data: request,
	}, nil
}

func CronEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
