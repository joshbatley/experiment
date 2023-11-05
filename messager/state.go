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

type state struct {
	key             models.Action
	priority        int
	nextStates      []models.Action
	trigger         func(*models.Event) bool
	progressPayment func(*models.Event) *models.Event
}

func getCurrentState(action models.Action) *state {
	state, _ := utils.Find(States, func(s *state) bool {
		return s.key == action
	})

	return state
}

// TODO: Think about this
func (s *state) getNextState(p *payment) state {
	var possibleState []state
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

var States = []*state{Requested, Authorized, Captured, Refunded, Void, Expiry}

var Requested = &state{
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

var Authorized = &state{
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

var Captured = &state{
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

var Refunded = &state{
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

var Void = &state{
	key: models.ActionVoid,
	trigger: func(event *models.Event) bool {
		randomNum := rand.Intn(10)
		return randomNum == 0
	},
	nextStates:      []models.Action{},
	progressPayment: progressVoid,
	priority:        5,
}

var Expiry = &state{
	key: models.ActionExpiry,
	trigger: func(event *models.Event) bool {
		return event.Timestamp.After(time.Now().Add(ExpireAfter))
	},
	nextStates:      []models.Action{},
	progressPayment: progressExpiry,
	priority:        10,
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
