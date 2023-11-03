package main

import (
	utils "shared"
	"shared/models"
	"sort"
	"time"
)

type State struct {
	name     models.Action
	priority int
	states   []models.Action
	trigger  func(*models.Event) bool
	progress func(*models.Event) *models.Event
}

func getCurrentState(action models.Action) *State {
	return utils.Find(States, func(s *State) bool {
		return s.name == action
	})
}

// TODO: Think about this
func (s *State) getNextState(p *Payment) *State {
	var possibleState []*State
	for _, c := range s.states {
		state := getCurrentState(c)
		if state.trigger(p.currEvent) {
			possibleState = append(possibleState, state)
		}
	}
	sort.Slice(possibleState, func(i, j int) bool {
		return possibleState[i].priority > possibleState[j].priority
	})
	return possibleState[0]
}

var States = []*State{Requested, Authorized, Captured, Refunded, Void, Expiry}

var Requested = &State{
	name: models.ActionRequested,
	trigger: func(event *models.Event) bool {
		return true
	},
	states: []models.Action{
		models.ActionAuthorize,
		models.ActionVoid,
		models.ActionExpiry,
	},
	priority: 1,
}

var Authorized = &State{
	name: models.ActionAuthorize,
	trigger: func(event *models.Event) bool {
		return false
	},
	states: []models.Action{
		models.ActionCapture,
		models.ActionRefund,
		models.ActionVoid,
		models.ActionExpiry,
	},
	progress: progressAuthorization,
	priority: 1,
}

var Captured = &State{
	name: models.ActionCapture,
	trigger: func(event *models.Event) bool {
		return false
	},
	states: []models.Action{
		models.ActionCapture,
		models.ActionRefund,
		models.ActionVoid,
		models.ActionExpiry,
	},
	progress: progressCapture,
	priority: 1,
}

var Refunded = &State{
	name: models.ActionRefund,
	trigger: func(event *models.Event) bool {
		return false
	},
	states: []models.Action{
		models.ActionCapture,
		models.ActionRefund,
		models.ActionVoid,
		models.ActionExpiry,
	},
	progress: progressRefund,
	priority: 3,
}

var Void = &State{
	name: models.ActionVoid,
	trigger: func(event *models.Event) bool {
		return false
	},
	states:   []models.Action{},
	progress: progressVoid,
	priority: 5,
}

var Expiry = &State{
	name: models.ActionVoid,
	trigger: func(event *models.Event) bool {
		return event.Timestamp.After(time.Now().Add(ExpireAfter))
	},
	states:   []models.Action{},
	priority: 10,
}
