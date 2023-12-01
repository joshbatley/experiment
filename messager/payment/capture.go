package payment

import "shared/event"

var capture = &state{
	trigger:         triggerCapture,
	progressPayment: progressCapture,
	nextStates: []event.Action{
		event.ActionCapture,
		event.ActionRefund,
		event.ActionVoid,
		event.ActionExpiry,
	},
	priority: 1,
}

func triggerCapture(p *Payment) bool {
	if p.CanCapture() {
		return true
	}
	return false
}

func progressCapture(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsCapture(p.MaxCapturable())
	return event, false
}
