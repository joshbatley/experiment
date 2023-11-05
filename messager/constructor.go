package main

import (
	"shared"
	"shared/models"
	"sort"
)

type payment struct {
	events             []*models.Event
	currEvent          *models.Event
	isCompletedPayment bool
}

func NewPayment() *payment {
	newEvent := constructRequested()
	return &payment{
		events:             []*models.Event{newEvent},
		currEvent:          newEvent,
		isCompletedPayment: false,
	}
}

func PaymentFromRecord(record Record) *payment {
	sort.Slice(record.PastEvents, func(i, j int) bool {
		return record.PastEvents[i].Timestamp.Before(record.PastEvents[j].Timestamp)
	})
	return &payment{
		currEvent:          record.PastEvents[len(record.PastEvents)-1],
		events:             record.PastEvents,
		isCompletedPayment: false,
	}
}

func (p *payment) ToRecord() Record {
	return Record{
		ID:         p.currEvent.ID,
		PastEvents: p.events,
	}
}

func (p *payment) Progress() {
	if p.isCompletedPayment {
		return
	}
	nextState := getCurrentState(p.currEvent.Action).getNextState(p)
	newEvent := nextState.progressPayment(p.currEvent)
	if len(nextState.nextStates) == 0 {
		p.isCompletedPayment = true
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
