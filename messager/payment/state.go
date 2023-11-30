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
	progressPayment func(*Payment) (*event.Event, bool)
}

var states = map[event.Action]*state{
	event.ActionRequested: {
		trigger: alwaysTrigger,
		nextStates: []event.Action{
			event.ActionAuthorize,
			event.ActionVoid,
			event.ActionExpiry,
		},
		priority: 1,
	},
	event.ActionAuthorize: {
		trigger:         alwaysTrigger,
		progressPayment: progressAuthorization,
		nextStates: []event.Action{
			event.ActionCapture,
			event.ActionRefund,
			event.ActionVoid,
			event.ActionExpiry,
		},
		priority: 1,
	},
	event.ActionCapture: {
		trigger:         triggerCapture,
		progressPayment: progressCapture,
		nextStates: []event.Action{
			event.ActionCapture,
			event.ActionRefund,
			event.ActionVoid,
			event.ActionExpiry,
		},
		priority: 1,
	},
	event.ActionRefund: {
		trigger:         triggerRefund,
		progressPayment: progressRefund,
		nextStates: []event.Action{
			event.ActionCapture,
			event.ActionRefund,
			event.ActionVoid,
			event.ActionExpiry,
		},
		priority: 3,
	},
	event.ActionVoid: {
		trigger:         triggerVoid,
		progressPayment: progressVoid,
		nextStates:      []event.Action{},
		priority:        5,
	},
	event.ActionExpiry: {
		trigger:         triggerExpiry,
		progressPayment: progressExpiry,
		nextStates:      []event.Action{},
		priority:        10,
	},
}

func getCurrentState(action event.Action) *state {
	return states[action]
}

func sortStates(states []*state) *state {
	sort.Slice(states, func(i, j int) bool {
		return states[i].priority > states[j].priority
	})
	return states[0]
}

func getNextState(s *state, p *Payment) (*state, error) {
	var possibleState []*state
	for _, c := range s.nextStates {
		currState := getCurrentState(c)
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
	s := getCurrentState(p.latestEvent.Action)
	next, err := getNextState(s, p)
	if err != nil {
		return nil, true
	}
	return next.progressPayment(p)
}
