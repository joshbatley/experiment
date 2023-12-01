package payment

import (
	utils "shared"
	"shared/event"
)

var refund = &state{
	trigger:         triggerRefund,
	progressPayment: progressRefund,
	nextStates: []event.Action{
		event.ActionCapture,
		event.ActionRefund,
		event.ActionVoid,
		event.ActionExpiry,
	},
	priority: 3,
}

func triggerRefund(p *Payment) bool {
	if p.CanRefund() {
		return utils.RandomChance(3)
	}
	return false
}

func progressRefund(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsRefund(p.MaxRefundable(), p.IsCaptured())
	return event, false
}
