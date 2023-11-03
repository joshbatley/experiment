package main

import (
	"shared"
	"shared/models"
)

type Payment struct {
	events             []*models.Event
	currEvent          *models.Event
	IsCompletedPayment bool
}

func NewPayment() *Payment {
	newEvent := constructRequested()
	return &Payment{
		events:             []*models.Event{newEvent},
		currEvent:          newEvent,
		IsCompletedPayment: false,
	}
}

func (p *Payment) Progress() {
	nextState := getCurrentState(p.currEvent.Action).getNextState(p)
	newEvent := nextState.progress(p.currEvent)
	p.currEvent = newEvent
	p.events = append(p.events, newEvent)
}

func constructRequested() *models.Event {
	return models.NewEvent(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsRequested(10, models.ResponseCodeSuccess)
}

func progressAuthorization(ev *models.Event) *models.Event {
	return models.NewEvent(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsAuthorized(10, models.ResponseCodeSuccess)
}

func progressCapture(ev *models.Event) *models.Event {
	return models.NewEvent(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
}

func progressRefund(ev *models.Event) *models.Event {
	return models.NewEvent(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
}

func progressVoid(ev *models.Event) *models.Event {
	return models.NewEvent(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
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
