package eventstore

import (
	"errors"
	"math/rand"
	utils "shared"
	"time"
)

type InMemoryEventStore struct {
	store   map[string]Record
	keys    []string
	maxSize int
}

func NewInMemory() *InMemoryEventStore {
	return &InMemoryEventStore{
		store:   make(map[string]Record),
		keys:    []string{},
		maxSize: 10,
	}
}

func (i *InMemoryEventStore) AddUnfinishedEvent(ev Record) error {
	if len(i.keys) >= i.maxSize {
		return errors.New("max capacity")
	}
	i.keys = append(i.keys, ev.ID)
	i.store[ev.ID] = ev
	return nil
}

func (i *InMemoryEventStore) UpdateEvent(ev Record) error {
	i.store[ev.ID] = ev
	return nil
}

func (i *InMemoryEventStore) GetRandomEvent() (Record, error) {
	if len(i.keys) <= 0 {
		return Record{}, errors.New("empty")
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
