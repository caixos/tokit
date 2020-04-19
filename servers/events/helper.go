package events

import "caixin.app/caixos/tokit/contracts"

func Fire(payload *contracts.Payload) {
	//发送事件需要判断是否有处理器,否则不处理
	_, isExist := Handlers[payload.Route]
	if isExist {
		event := newEvent(payload)
		addEvent(event)
	}
}
