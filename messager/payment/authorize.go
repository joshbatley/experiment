package payment

import "shared/event"

var authorize = &state{
	trigger:         triggerAuthorized,
	progressPayment: progressAuthorization,
	nextStates: []event.Action{
		event.ActionCapture,
		event.ActionRefund,
		event.ActionVoid,
		event.ActionExpiry,
	},
	priority: 1,
}

func triggerAuthorized(*Payment) bool {
	return true
}

func progressAuthorization(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsAuthorized()
	return event, false
}
