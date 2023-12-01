package store

import (
	"shared/event"
)

const ErrNoEvents = Error("invalid input")

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

type Error string

func (e Error) Error() string {
	return string(e)
}

type Store interface {
	UpdateOrInsert(e *Entry) error
	GetRandom() (*Entry, error)
	Delete(id string) error
}
