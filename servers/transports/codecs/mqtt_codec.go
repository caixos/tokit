package codecs

import (
	"caixin.app/caixos/tokit/constants"
	"caixin.app/caixos/tokit/contracts"
	"caixin.app/caixos/tokit/tools/idwork"
	"encoding/json"
	"errors"
	"context"
)

func MqttSubscribeDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal(req.([]byte), &mapResult)
	if err != nil {
		return nil, errors.New(constants.ErrJson)
	}
	requestId, ok := mapResult["request_id"].(string)
	if ok == false {
		requestId = idwork.ID()
	}
	return contracts.Request{
		Id:   requestId,
		Data: mapResult,
	}, nil
}

func MqttSubscribeEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
