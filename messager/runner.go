package main

import (
	"github.com/rs/zerolog/log"
	"math/rand"
	"messager/eventstore"
	"time"
)

// TODO: Add some sort of limiter for the new event chance

type runner struct {
	store          eventstore.EventStore
	ticker         *time.Ticker
	newEventChance int
}

func newRunner(store eventstore.EventStore, tps int) *runner {
	return &runner{
		store:          store,
		ticker:         time.NewTicker(time.Second / time.Duration(tps)),
		newEventChance: 2,
	}
}

func (e *runner) StartUp() {
	go func() {
		for {
			select {
			case <-e.ticker.C:
				ev, err := e.generate()
				if err != nil {
					log.Warn().Err(err).Send()
					e.newEventChance++
					continue
				}
				ev.currEvent.Log()
			}
		}
	}()
	select {}
}

func (e *runner) generate() (*Record, error) {
	if e.shouldCreateNewRecord() {
		return createNewEvent(e.store)
	}
	record, err := e.store.GetRandomEvent()
	if err != nil {
		return createNewEvent(e.store)
	}

	payment := fromEventstoreRecord(record)
	payment.progress()
	if payment.isCompletedPayment {
		if err := e.store.RemoveEvent(payment.currEvent.ID); err != nil {
			return nil, err
		}
		return payment, nil
	}
	if err := e.store.UpdateEvent(payment.toEventstoreRecord()); err != nil {
		return nil, err
	}
	return payment, nil
}

func (e *runner) shouldCreateNewRecord() bool {
	randomNum := rand.Intn(e.newEventChance)
	return randomNum == 0
}

func createNewEvent(store eventstore.EventStore) (*Record, error) {
	payment := newRecord()
	if err := store.AddUnfinishedEvent(payment.toEventstoreRecord()); err != nil {
		return nil, err
	}
	return payment, nil
}
