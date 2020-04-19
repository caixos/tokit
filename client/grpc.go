package client

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"caixin.app/tokit/contract"
	"caixin.app/tokit/server/transports/protobuf"
	"caixin.app/tokit/tools/idwork"
	"google.golang.org/grpc"
)

func NewGrpcClient(serviceAddress string, service string, params map[string]interface{}) (*protobuf.Response, error) {
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	jsonParam, _ := json.Marshal(params)
	in := &protobuf.Request{
		Id:    idwork.ID(),
		Param: string(jsonParam),
	}

	out := new(protobuf.Response)

	method := "/protobuf." + service + "/Handle"
	err = conn.Invoke(context.Background(), method, in, out)
	return out, err
}

func NewGrpcCall(host, service string, params map[string]interface{}) (ret contract.Response) {
	resp, err := NewGrpcClient(host, service, params)
	if err != nil {
		ret = contract.ResponseFailed(errors.New("没有响应的服务:" + service))
	} else {
		m := make(map[string]interface{})
		m["call_method"] = "grpc"
		err := json.Unmarshal([]byte(resp.GetData()), &m)
		if err != nil {
			ret = contract.ResponseFailed(err)
		} else {
			ret.Code = resp.Code
			ret.Ret = 200
			ret.Message = resp.Message
			ret.Data = m
		}
	}
	return
}
