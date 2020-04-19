package events

import "caixin.app/tokit/contract"

func addEvent(event *contract.Payload) {
	go func() {
		eventChan <- event
	}()
}

func newEvent(event *contract.Payload) *contract.Payload {
	e := eventPool.Get()
	if e == nil {
		return event
	} else {
		ret := e.(*contract.Payload)
		(*ret).Route = event.Route
		(*ret).Params = event.Params
		return ret
	}
}
