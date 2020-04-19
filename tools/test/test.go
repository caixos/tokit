package test

import (
	"caixin.app/tokit/contract"
	"caixin.app/tokit/filters"
	"caixin.app/tokit/tools/idwork"
	"errors"
	"context"
)

type TestStruct struct {
	controller contract.IController
	request    map[string]interface{}
}

func NewTest() *TestStruct {
	return &TestStruct{
		request: make(map[string]interface{}),
	}
}
func (s *TestStruct) Controller(controller contract.IController) *TestStruct {
	s.controller = controller
	return s
}
func (s *TestStruct) Request(m map[string]interface{}) *TestStruct {
	if m != nil {
		s.request = m
	}
	return s
}
func (s *TestStruct) Run() (contract.Response, error) {
	e := filters.Chain(
		&filters.ResponseEndpoint{},
		&filters.CommEndpoint{Controller: s.controller},
	)
	request := contract.Request{
		Id:   idwork.ID(),
		Data: s.request,
	}
	response, err := e(context.Background(), request)
	resp := response.(contract.Response)
	if err != nil {
		return resp, err
	}
	if resp.Code != "0000" {
		return resp, errors.New(resp.Message)
	}
	return resp, nil
}
