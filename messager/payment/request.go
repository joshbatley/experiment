package payment

import "shared/event"

var request = &state{
	trigger: triggerRequest,
	nextStates: []event.Action{
		event.ActionAuthorize,
		event.ActionVoid,
		event.ActionExpiry,
	},
	priority: 1,
}

func triggerRequest(*Payment) bool {
	return true
}
