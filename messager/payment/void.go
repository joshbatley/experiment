package payment

import (
	utils "shared"
	"shared/event"
)

var void = &state{
	trigger:         triggerVoid,
	progressPayment: progressVoid,
	nextStates:      []event.Action{},
	priority:        5,
}

func triggerVoid(*Payment) bool {
	return utils.RandomChance(1000)
}

func progressVoid(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsVoid()
	return event, true
}
