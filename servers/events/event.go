package events

import "github.com/caixos/tokit/contracts"

func addEvent(event *contracts.Payload) {
	go func() {
		eventChan <- event
	}()
}

func newEvent(event *contracts.Payload) *contracts.Payload {
	e := eventPool.Get()
	if e == nil {
		return event
	} else {
		ret := e.(*contracts.Payload)
		(*ret).Route = event.Route
		(*ret).Params = event.Params
		return ret
	}
}
