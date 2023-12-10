package payment

import "shared/event"

var request = &state{
	trigger:         triggerRequest,
	progressPayment: progressRequested,
	nextStates: []event.Action{
		event.ActionAuthorize,
		event.ActionVoid,
		event.ActionExpiry,
	},
	priority: 1,
}

func triggerRequest(p *Payment) bool {
	if p.latestEvent == nil {
		return true
	}
	return false
}

func progressRequested(p *Payment) (*event.Event, bool) {
	event := event.New(p.clientId, "").AsRequested()
	return event, event.IsFailureResponseCode()
}
