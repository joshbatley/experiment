package main

import (
	"errors"
	"math/rand"
	utils "shared"
	"shared/models"
	"time"
)

type Event struct {
	ID         string
	PastEvents []models.Payment
}

type EventStore interface {
	AddUnfinishedEvent(e Event) error
	GetRandomEvent() (Event, error)
	RemoveEvent(id string) error
}

type InMemoryEventStore struct {
	store map[string]Event
	keys  []string
}

func NewInMemory() *InMemoryEventStore {
	return &InMemoryEventStore{
		store: make(map[string]Event),
		keys:  []string{},
	}
}

func (i *InMemoryEventStore) AddUnfinishedEvent(ev Event) error {
	i.keys = append(i.keys, ev.ID)
	i.store[ev.ID] = ev
	return nil
}

func (i *InMemoryEventStore) GetRandomEvent() (Event, error) {
	if len(i.keys) <= 0 {
		return Event{}, errors.New("empty")
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	rKey := i.keys[r.Intn(len(i.keys))]
	return i.store[rKey], nil
}

func (i *InMemoryEventStore) RemoveEvent(id string) error {
	if len(i.keys) <= 0 {
		return errors.New("empty")
	}
	idx := utils.FindIndex(i.keys, id)
	i.keys = append(i.keys[:idx], i.keys[idx+1:]...)
	delete(i.store, id)
	return nil
}
