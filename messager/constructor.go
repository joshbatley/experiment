package main

import (
	"shared"
	"shared/models"
	"sort"
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

func PaymentFromRecord(record Record) *Payment {
	sort.Slice(record.PastEvents, func(i, j int) bool {
		return record.PastEvents[i].Timestamp.Before(record.PastEvents[j].Timestamp)
	})
	return &Payment{
		currEvent:          record.PastEvents[len(record.PastEvents)-1],
		events:             record.PastEvents,
		IsCompletedPayment: false,
	}
}

func (p *Payment) ToRecord() Record {
	return Record{
		ID:         p.currEvent.ID,
		PastEvents: p.events,
	}
}

func (p *Payment) Progress() {
	if p.IsCompletedPayment {
		return
	}
	nextState := getCurrentState(p.currEvent.Action).getNextState(p)
	newEvent := nextState.progressPayment(p.currEvent)
	if len(nextState.nextStates) == 0 {
		p.IsCompletedPayment = true
	}
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
