package payment

import (
	"errors"
	"shared/event"
	"sort"
)

type state struct {
	priority        int
	nextStates      []event.Action
	trigger         func(*Payment) bool
	progressPayment func(*Payment) (e *event.Event, isComplete bool)
}

var states = map[event.Action]*state{
	event.ActionRequest:   request,
	event.ActionAuthorize: authorize,
	event.ActionCapture:   capture,
	event.ActionRefund:    refund,
	event.ActionVoid:      void,
	event.ActionExpiry:    expiry,
}

func sortStates(states []*state) *state {
	sort.Slice(states, func(i, j int) bool {
		return states[i].priority > states[j].priority
	})
	return states[0]
}

func getNextState(p *Payment) (*state, error) {
	var possibleState []*state
	// Can we refactor this out?
	if p.GetLatestEvent() == nil {
		return request, nil
	}
	for _, c := range states[p.latestEvent.Action].nextStates {
		currState := states[c]
		if currState.trigger(p) {
			possibleState = append(possibleState, currState)
		}
	}

	if len(possibleState) == 0 {
		return nil, errors.New("no more states")
	}

	return sortStates(possibleState), nil
}

func createNewEvent(p *Payment) (*event.Event, bool) {
	next, err := getNextState(p)
	if err != nil {
		return nil, true
	}
	return next.progressPayment(p)
}
