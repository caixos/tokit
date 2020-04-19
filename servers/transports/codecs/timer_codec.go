package codecs

import (
	"caixin.app/tokit/contract"
	"context"
)

func TimerDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(map[string]interface{})
	return contract.Request{
		Id:   request["request_id"].(string),
		Data: request,
	}, nil
}

func TimerEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
