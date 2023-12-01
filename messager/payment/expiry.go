package payment

import "shared/event"

var expiry = &state{
	trigger:         triggerExpiry,
	progressPayment: progressExpiry,
	nextStates:      []event.Action{},
	priority:        10,
}

func triggerExpiry(p *Payment) bool {
	return p.CanExpire()
}

func progressExpiry(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsExpiry()
	return event, true
}
