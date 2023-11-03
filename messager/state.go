package main

import (
	"math/rand"
	utils "shared"
	"shared/models"
	"sort"
	"time"
)

const (
	ExpireAfter = time.Second * 3000
)

type State struct {
	key             models.Action
	priority        int
	nextStates      []models.Action
	trigger         func(*models.Event) bool
	progressPayment func(*models.Event) *models.Event
}

func getCurrentState(action models.Action) *State {
	state, _ := utils.Find(States, func(s *State) bool {
		return s.key == action
	})

	return state
}

// TODO: Think about this
func (s *State) getNextState(p *Payment) State {
	var possibleState []State
	for _, c := range s.nextStates {
		state := getCurrentState(c)
		if state.trigger(p.currEvent) {
			possibleState = append(possibleState, *state)
		}
	}
	sort.Slice(possibleState, func(i, j int) bool {
		return possibleState[i].priority > possibleState[j].priority
	})
	return possibleState[0]
}

var States = []*State{Requested, Authorized, Captured, Refunded, Void, Expiry}

var Requested = &State{
	key: models.ActionRequested,
	trigger: func(event *models.Event) bool {
		return true
	},
	nextStates: []models.Action{
		models.ActionAuthorize,
		models.ActionVoid,
		models.ActionExpiry,
	},
	priority: 1,
}

var Authorized = &State{
	key: models.ActionAuthorize,
	trigger: func(event *models.Event) bool {
		return true
	},
	nextStates: []models.Action{
		models.ActionCapture,
		models.ActionRefund,
		models.ActionVoid,
		models.ActionExpiry,
	},
	progressPayment: progressAuthorization,
	priority:        1,
}

var Captured = &State{
	key: models.ActionCapture,
	trigger: func(event *models.Event) bool {
		return true
	},
	nextStates: []models.Action{
		models.ActionCapture,
		models.ActionRefund,
		models.ActionVoid,
		models.ActionExpiry,
	},
	progressPayment: progressCapture,
	priority:        1,
}

var Refunded = &State{
	key: models.ActionRefund,
	trigger: func(event *models.Event) bool {
		randomNum := rand.Intn(3)
		return randomNum == 0
	},
	nextStates: []models.Action{
		models.ActionCapture,
		models.ActionRefund,
		models.ActionVoid,
		models.ActionExpiry,
	},
	progressPayment: progressRefund,
	priority:        3,
}

var Void = &State{
	key: models.ActionVoid,
	trigger: func(event *models.Event) bool {
		randomNum := rand.Intn(10)
		return randomNum == 0
	},
	nextStates:      []models.Action{},
	progressPayment: progressVoid,
	priority:        5,
}

var Expiry = &State{
	key: models.ActionExpiry,
	trigger: func(event *models.Event) bool {
		return event.Timestamp.After(time.Now().Add(ExpireAfter))
	},
	nextStates:      []models.Action{},
	progressPayment: progressExpiry,
	priority:        10,
}
