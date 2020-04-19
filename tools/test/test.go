package test

import (
	"caixin.app/caixos/tokit/contracts"
	"caixin.app/caixos/tokit/filters"
	"caixin.app/caixos/tokit/tools/idwork"
	"errors"
	"context"
)

type TestStruct struct {
	controller contracts.IController
	request    map[string]interface{}
}

func NewTest() *TestStruct {
	return &TestStruct{
		request: make(map[string]interface{}),
	}
}
func (s *TestStruct) Controller(controller contracts.IController) *TestStruct {
	s.controller = controller
	return s
}
func (s *TestStruct) Request(m map[string]interface{}) *TestStruct {
	if m != nil {
		s.request = m
	}
	return s
}
func (s *TestStruct) Run() (contracts.Response, error) {
	e := filters.Chain(
		&filters.ResponseEndpoint{},
		&filters.CommEndpoint{Controller: s.controller},
	)
	request := contracts.Request{
		Id:   idwork.ID(),
		Data: s.request,
	}
	response, err := e(context.Background(), request)
	resp := response.(contracts.Response)
	if err != nil {
		return resp, err
	}
	if resp.Code != "0000" {
		return resp, errors.New(resp.Message)
	}
	return resp, nil
}
