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

var states = map[models.Action]*State{
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
		trigger: func(event *models.Event) bool {
			return true
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

type State struct {
	key             models.Action
	priority        int
	nextStates      []models.Action
	trigger         func(*models.Event) bool
	ProgressPayment func(*models.Event) *models.Event
}

func getCurrentState(action models.Action) *State {
	return states[action]
}

// TODO: Think about this
func (s *State) getNextState(p *Record) State {
	var possibleState []State
	for _, c := range s.nextStates {
		currState := getCurrentState(c)
		if currState.trigger(p.currEvent) {
			possibleState = append(possibleState, *currState)
		}
	}
	sort.Slice(possibleState, func(i, j int) bool {
		return possibleState[i].priority > possibleState[j].priority
	})
	return possibleState[0]
}

func progressAuthorization(ev *models.Event) *models.Event {
	return models.NewEvent(
		ev.ID, utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsAuthorized(10, models.ResponseCodeSuccess)
}

func progressCapture(ev *models.Event) *models.Event {
	return models.NewEvent(
		ev.ID, utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsCapture(0.3, "")
}

func progressRefund(ev *models.Event) *models.Event {
	return models.NewEvent(
		ev.ID, utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsRefund(0.1, "")
}

func progressVoid(ev *models.Event) *models.Event {
	return models.NewEvent(
		ev.ID, utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsVoid("2000")
}

func progressExpiry(ev *models.Event) *models.Event {
	return models.NewEvent(
		ev.ID, utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsExpiry("2000")
}

//func progressSuccessfulResponse(ev *models.Event) models.ResponseCode {
//	return models.SuccessfulResponseCodes[0]
//}
//
//func progressFailureResponse(ev *models.Event) models.ResponseCode {
//	return models.FailureResponseCodes[0]
//}
//
//func progressInfoResponse(ev *models.Event) models.ResponseCode {
//	return models.InformationResponseCode[0]
//}
//
//func progressFraudResponse(ev *models.Event) models.ResponseCode {
//	return models.FraudResponseCode[0]
//}
