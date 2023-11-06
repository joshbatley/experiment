package eventstore

import (
	"shared/models"
)

type Record struct {
	ID         string
	PastEvents []*models.Event
}

type EventStore interface {
	AddUnfinishedEvent(e Record) error
	UpdateEvent(e Record) error
	GetRandomEvent() (Record, error)
	RemoveEvent(id string) error
}
