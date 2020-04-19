package codecs

import (
	"github.com/caixos/tokit/constants"
	"github.com/caixos/tokit/contracts"
	"github.com/caixos/tokit/tools/idwork"
	"encoding/json"
	"errors"
	"context"
)

func CommandDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(req.(string)), &mapResult)
	if err != nil {
		return nil, errors.New(constants.ErrJson)
	}
	return contracts.Request{
		Id:   idwork.ID(),
		Data: mapResult,
	}, nil
}

func CommandEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

