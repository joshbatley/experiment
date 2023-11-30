package store

import (
	"errors"
	"math/rand"
	utils "shared"
	"time"
)

type InMemoryEventStore struct {
	store   map[string]*Entry
	keys    []string
	maxSize int
}

func NewInMemory() *InMemoryEventStore {
	return &InMemoryEventStore{
		store:   make(map[string]*Entry),
		keys:    []string{},
		maxSize: 10,
	}
}

func (i *InMemoryEventStore) Insert(e *Entry) error {
	if len(i.keys) >= i.maxSize {
		return errors.New("max capacity")
	}
	i.keys = append(i.keys, e.id)
	i.store[e.id] = e
	return nil
}

func (i *InMemoryEventStore) Update(e *Entry) error {
	i.store[e.id] = e
	return nil
}

func (i *InMemoryEventStore) GetRandom() (*Entry, error) {
	if len(i.keys) <= 0 {
		return nil, errors.New("empty")
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	rKey := i.keys[r.Intn(len(i.keys))]
	return i.store[rKey], nil
}

func (i *InMemoryEventStore) Delete(id string) error {
	if len(i.keys) <= 0 {
		return errors.New("empty")
	}
	idx := utils.FindIndex(i.keys, id)
	i.keys = append(i.keys[:idx], i.keys[idx+1:]...)
	delete(i.store, id)
	return nil
}