package main

import (
	"math/rand"
	"shared/models"
	"sort"
	"time"
)

const (
	ExpireAfter = time.Second * 3000
)

var states = map[models.Action]*state{
	models.ActionRequested: {
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
	},
	models.ActionAuthorize: {
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
		ProgressPayment: progressAuthorization,
		priority:        1,
	},
	models.ActionCapture: {
		key: models.ActionCapture,
		trigger: func(e *models.Event) bool {
			if e.CapturedAmount <= e.AuthorizedAmount {
				return true
			}
			return false
		},
		nextStates: []models.Action{
			models.ActionCapture,
			models.ActionRefund,
			models.ActionVoid,
			models.ActionExpiry,
		},
		ProgressPayment: progressCapture,
		priority:        1,
	},
	models.ActionRefund: {
		trigger: func(e *models.Event) bool {
			if e.RefundedAmount <= e.CapturedAmount {
				randomNum := rand.Intn(3)
				return randomNum == 0
			}
			return false
		},
		nextStates: []models.Action{
			models.ActionCapture,
			models.ActionRefund,
			models.ActionVoid,
			models.ActionExpiry,
		},
		ProgressPayment: progressRefund,
		priority:        3,
	},
	models.ActionVoid: {
		key: models.ActionVoid,
		trigger: func(event *models.Event) bool {
			randomNum := rand.Intn(10)
			return randomNum == 0
		},
		nextStates:      []models.Action{},
		ProgressPayment: progressVoid,
		priority:        5,
	},
	models.ActionExpiry: {
		key: models.ActionExpiry,
		trigger: func(event *models.Event) bool {
			return event.Timestamp.After(time.Now().Add(ExpireAfter))
		},
		nextStates:      []models.Action{},
		ProgressPayment: progressExpiry,
		priority:        10,
	},
}

type state struct {
	key             models.Action
	priority        int
	nextStates      []models.Action
	trigger         func(*models.Event) bool
	ProgressPayment func(*models.Event) (*models.Event, bool)
}

func getCurrentState(action models.Action) *state {
	return states[action]
}

func sortStates(states []*state) *state {
	sort.Slice(states, func(i, j int) bool {
		return states[i].priority > states[j].priority
	})
	return states[0]
}

func getNewState(s *state, ev *models.Event) *state {
	var possibleState []*state
	for _, c := range s.nextStates {
		currState := getCurrentState(c)
		if currState.trigger(ev) {
			possibleState = append(possibleState, currState)
		}
	}
	return sortStates(possibleState)
}

func progressPayment(ev *models.Event) (*models.Event, bool) {
	current := getCurrentState(ev.Action)
	next := getNewState(current, ev)
	return next.ProgressPayment(ev)
}
