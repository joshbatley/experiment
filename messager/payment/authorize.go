package payment

import "shared/event"

var authorize = &state{
	trigger:         triggerAuthorized,
	progressPayment: progressAuthorization,
	nextStates: []event.Action{
		event.ActionCapture,
		event.ActionVoid,
		event.ActionExpiry,
	},
	priority: 1,
}

func triggerAuthorized(p *Payment) bool {
	if p.GetLatestEvent() == nil {
		return false
	}
	var requested *event.Event
	for _, e := range p.events {
		if e.Action == event.ActionRequest {
			requested = e
		}
	}
	if requested.Status == event.StatusCancelled || requested.Status == event.StatusFailed {
		return false
	}
	return true
}

func progressAuthorization(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsAuthorized()
	return event, event.IsFailureResponseCode()
}
