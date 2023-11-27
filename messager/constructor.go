package main

import (
	"math/rand"
	"messager/eventstore"
	"shared/models"
	"sort"
)

type Record struct {
	events             []*models.Event
	currEvent          *models.Event
	isCompletedPayment bool
}

func newRecord() *Record {
	newEvent := constructRequested()
	return &Record{
		events:             []*models.Event{newEvent},
		currEvent:          newEvent,
		isCompletedPayment: false,
	}
}

func fromEventstoreRecord(record eventstore.Record) *Record {
	sort.Slice(record.PastEvents, func(i, j int) bool {
		return record.PastEvents[i].Timestamp.Before(record.PastEvents[j].Timestamp)
	})
	return &Record{
		currEvent:          record.PastEvents[len(record.PastEvents)-1],
		events:             record.PastEvents,
		isCompletedPayment: false,
	}
}

func (p *Record) toEventstoreRecord() eventstore.Record {
	return eventstore.Record{
		ID:         p.currEvent.PaymentID,
		PastEvents: p.events,
	}
}

func (p *Record) progress() {
	if p.isCompletedPayment {
		return
	}
	next := getCurrentState(p.currEvent.Action).getNextState(p)
	if len(next) == 0 {
		p.isCompletedPayment = true
		return
	}
	newEvent := next[0].ProgressPayment(p.currEvent)
	if len(next[0].nextStates) == 0 || endEventEarly() {
		p.isCompletedPayment = true
	}
	p.currEvent = newEvent
	p.events = append(p.events, newEvent)
}

func endEventEarly() bool {
	randomNum := rand.Intn(30)
	return randomNum == 0
}
