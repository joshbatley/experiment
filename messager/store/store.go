package store

import (
	"shared/event"
)

type Entry struct {
	id     string
	Events []*event.Event
}

func NewEntry(id string, e []*event.Event) *Entry {
	return &Entry{
		id:     id,
		Events: e,
	}
}

type Store interface {
	Insert(e *Entry) error
	Update(e *Entry) error
	GetRandom() (*Entry, error)
	Delete(id string) error
}
