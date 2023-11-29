package main

import (
	"errors"
	"math/rand"
	"shared/event"
	"sort"
)

type state struct {
	key             event.Action
	priority        int
	nextStates      []event.Action
	trigger         func(*Record) bool
	ProgressPayment func(*Record) (*event.Event, bool)
}

var states = map[event.Action]*state{
	event.ActionRequested: {
		key: event.ActionRequested,
		trigger: func(r *Record) bool {
			return true
		},
		nextStates: []event.Action{
			event.ActionAuthorize,
			event.ActionVoid,
			event.ActionExpiry,
		},
		priority: 1,
	},
	event.ActionAuthorize: {
		key: event.ActionAuthorize,
		trigger: func(r *Record) bool {
			return true
		},
		nextStates: []event.Action{
			event.ActionCapture,
			event.ActionRefund,
			event.ActionVoid,
			event.ActionExpiry,
		},
		ProgressPayment: progressAuthorization,
		priority:        1,
	},
	event.ActionCapture: {
		key: event.ActionCapture,
		trigger: func(r *Record) bool {
			if r.CanCapture() {
				return true
			}
			return false
		},
		nextStates: []event.Action{
			event.ActionCapture,
			event.ActionRefund,
			event.ActionVoid,
			event.ActionExpiry,
		},
		ProgressPayment: progressCapture,
		priority:        1,
	},
	event.ActionRefund: {
		trigger: func(r *Record) bool {
			if r.CanRefund() {
				randomNum := rand.Intn(3)
				return randomNum == 0
			}
			return false
		},
		nextStates: []event.Action{
			event.ActionCapture,
			event.ActionRefund,
			event.ActionVoid,
			event.ActionExpiry,
		},
		ProgressPayment: progressRefund,
		priority:        3,
	},
	event.ActionVoid: {
		key: event.ActionVoid,
		trigger: func(r *Record) bool {
			return false
			randomNum := rand.Intn(1000)
			return randomNum == 0
		},
		nextStates:      []event.Action{},
		ProgressPayment: progressVoid,
		priority:        5,
	},
	event.ActionExpiry: {
		key: event.ActionExpiry,
		trigger: func(r *Record) bool {
			return r.CanExpire()
		},
		nextStates:      []event.Action{},
		ProgressPayment: progressExpiry,
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

func getNewState(s *state, r *Record) (*state, error) {
	var possibleState []*state
	for _, c := range s.nextStates {
		currState := getCurrentState(c)
		if currState.trigger(r) {
			possibleState = append(possibleState, currState)
		}
	}

	if len(possibleState) == 0 {
		return nil, errors.New("no more states")
	}

	return sortStates(possibleState), nil
}

func progressPayment(r *Record) (*event.Event, bool) {
	s := getCurrentState(r.latestEvent.Action)
	next, err := getNewState(s, r)
	if err != nil {
		return nil, true
	}
	return next.ProgressPayment(r)
}
