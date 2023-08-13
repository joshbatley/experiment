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

func (s *InMemoryEventStore) AddUnfinishedEvent(e Event) error {
	s.keys = append(s.keys, e.ID)
	s.store[e.ID] = e
	return nil
}

func (s *InMemoryEventStore) GetRandomEvent() (Event, error) {
	if len(s.keys) <= 0 {
		return Event{}, errors.New("empty")
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	rKey := s.keys[r.Intn(len(s.keys))]
	return s.store[rKey], nil
}

func (s *InMemoryEventStore) RemoveEvent(id string) error {
	if len(s.keys) <= 0 {
		return errors.New("empty")
	}
	idx := utils.FindIndex(s.keys, id)
	s.keys = append(s.keys[:idx], s.keys[idx+1:]...)
	delete(s.store, id)
	return nil
}
