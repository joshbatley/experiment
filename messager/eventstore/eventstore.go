package eventstore

import (
	"shared/event"
)

type Record struct {
	ID         string
	PastEvents []*event.Event
}

type EventStore interface {
	AddUnfinishedEvent(e Record) error
	UpdateEvent(e Record) error
	GetRandomEvent() (Record, error)
	RemoveEvent(id string) error
}
