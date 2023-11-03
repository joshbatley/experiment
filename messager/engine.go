package main

import (
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

// TODO: Add some sort of limiter for the new event chance

type Runner struct {
	Store          EventStore
	ticker         *time.Ticker
	newEventChance int
}

func NewRunner(store EventStore, tps int) *Runner {
	return &Runner{
		Store:          store,
		ticker:         time.NewTicker(time.Second / time.Duration(tps)),
		newEventChance: 2,
	}
}

func (e *Runner) StartUp() {
	go func() {
		for {
			select {
			case <-e.ticker.C:
				ev, err := e.Generate()
				if err != nil {
					log.Warn().Err(err).Send()
					e.newEventChance++
					continue
				}
				log.Print("%v", ev.currEvent.Action)
			}
		}
	}()
	select {}
	//time.Sleep(time.Second)
}

func (e *Runner) Generate() (*Payment, error) {
	if e.rndCreateNewPayment() {
		return createNewEvent(e.Store)
	}
	record, err := e.Store.GetRandomEvent()
	if err != nil {
		return createNewEvent(e.Store)
	}

	payment := PaymentFromRecord(record)
	payment.Progress()
	if payment.IsCompletedPayment {
		if err := e.Store.RemoveEvent(payment.currEvent.ID); err != nil {
			return nil, err
		}
		return payment, nil
	}
	if err := e.Store.UpdateEvent(payment.ToRecord()); err != nil {
		return nil, err
	}
	return payment, nil
}

func createNewEvent(store EventStore) (*Payment, error) {
	payment := NewPayment()
	if err := store.AddUnfinishedEvent(payment.ToRecord()); err != nil {
		return nil, err
	}
	return payment, nil
}

func (e *Runner) rndCreateNewPayment() bool {
	randomNum := rand.Intn(e.newEventChance)
	return randomNum == 0
}
