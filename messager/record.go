package main

import (
	"messager/eventstore"
	"shared/event"
	"sort"
	"time"
)

const (
	ExpireAfter = time.Second * 3000
)

type Record struct {
	events      []*event.Event
	latestEvent *event.Event
	isCompleted bool
}

func NewRecord() *Record {
	newEvent := constructRequested()
	return &Record{
		events:      []*event.Event{newEvent},
		latestEvent: newEvent,
		isCompleted: false,
	}
}

func FromEventstoreRecord(record eventstore.Record) *Record {
	sort.Slice(record.PastEvents, func(i, j int) bool {
		return record.PastEvents[i].Timestamp.Before(record.PastEvents[j].Timestamp)
	})
	return &Record{
		latestEvent: record.PastEvents[len(record.PastEvents)-1],
		events:      record.PastEvents,
		isCompleted: false,
	}
}

func (r *Record) ToEventstoreRecord() eventstore.Record {
	return eventstore.Record{
		ID:         r.latestEvent.PaymentID,
		PastEvents: r.events,
	}
}

func (r *Record) updateEvent(newEvent *event.Event) {
	r.latestEvent = newEvent
	r.events = append(r.events, newEvent)
}

func (r *Record) Progress() {
	newEvent, isComplete := progressPayment(r)
	if newEvent != nil {
		r.updateEvent(newEvent)
	}
	r.isCompleted = isComplete
}

func (r *Record) getTotalAuthorisedAmount() (total int) {
	return r.events[0].AuthorizedAmount
}

func (r *Record) getTotalCapturedAmount() (total int) {
	for _, ev := range r.events {
		total += ev.CapturedAmount
	}
	return total
}

func (r *Record) getTotalRefundedAmount() (total int) {
	for _, ev := range r.events {
		total += ev.RefundedAmount
	}
	return total
}

func (r *Record) MaxCapturableAmount() int {
	if r.IsFullyCaptured() || r.IsFullyRefunded() {
		return 0
	}
	return r.getTotalAuthorisedAmount() - (r.getTotalCapturedAmount() - r.getTotalRefundedAmount())
}

func (r *Record) MaxRefundableAmount() int {
	if !r.hasCapturedEvent() || r.IsFullyRefunded() {
		return 0
	}
	return r.getTotalCapturedAmount() - r.getTotalRefundedAmount()
}

func (r *Record) hasCapturedEvent() bool {
	for _, ev := range r.events {
		if ev.Status == event.StatusCaptured || ev.Status == event.StatusPartiallyCaptured {
			return true
		}
	}
	return false
}

func (r *Record) IsFullyCaptured() bool {
	for _, ev := range r.events {
		if ev.Status == event.StatusCaptured {
			return true
		}
	}
	return false
}

func (r *Record) IsFullyRefunded() bool {
	for _, ev := range r.events {
		if ev.Status == event.StatusRefunded {
			return true
		}
	}
	return false
}

func (r *Record) CanCapture() bool {
	return !r.isCompleted && !r.IsFullyCaptured() && r.MaxCapturableAmount() != 0
}

func (r *Record) CanRefund() bool {
	return !r.isCompleted && r.hasCapturedEvent() && !r.IsFullyRefunded() && r.MaxRefundableAmount() != 0
}

func (r *Record) CanExpire() bool {
	return (r.latestEvent.Action == event.ActionAuthorize || r.latestEvent.Action == event.ActionRequested) &&
		r.latestEvent.Timestamp.After(time.Now().Add(ExpireAfter))
}
