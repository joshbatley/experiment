package main

import (
	"shared/event"
)

func constructRequested() *event.Event {
	return event.New("cli_123123", "").AsRequested()
}

func progressAuthorization(r *Record) (*event.Event, bool) {
	event := r.latestEvent.AsAuthorized()
	return event, false
}

func progressCapture(r *Record) (*event.Event, bool) {
	event := r.latestEvent.AsCapture(r.MaxCapturableAmount())
	return event, false
}

func progressRefund(r *Record) (*event.Event, bool) {
	event := r.latestEvent.AsRefund(r.MaxRefundableAmount(), r.IsFullyCaptured())
	return event, false
}

func progressVoid(r *Record) (*event.Event, bool) {
	event := r.latestEvent.AsVoid()
	return event, true
}

func progressExpiry(r *Record) (*event.Event, bool) {
	event := r.latestEvent.AsExpiry()
	return event, true
}
