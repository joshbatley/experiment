package main

import (
	"github.com/rs/zerolog/log"
	"messager/payment"
	"messager/store"
	utils "shared"
	"shared/event"
	"time"
)

// TODO: Insert some sort of limiter for the new event chance

type runner struct {
	store          store.Store
	ticker         *time.Ticker
	newEventChance int
}

func NewRunner(store store.Store, tps int) *runner {
	return &runner{
		store:          store,
		ticker:         time.NewTicker(time.Second / time.Duration(tps)),
		newEventChance: 2,
	}
}

func createNewPayment(store store.Store) (*event.Event, error) {
	p := payment.New()
	if err := store.Insert(p.ToEntry()); err != nil {
		return nil, err
	}
	return p.GetLatestEvent(), nil
}

func (r *runner) updateStore(p *payment.Payment) error {
	if p.HasFinalEvent() {
		return r.store.Delete(p.GetLatestEvent().PaymentID)
	} else {
		return r.store.Update(p.ToEntry())
	}
}

func (r *runner) generate() (*event.Event, error) {
	if utils.RandomChance(r.newEventChance) {
		return createNewPayment(r.store)
	}
	ev, err := r.store.GetRandom()
	if err != nil {
		return createNewPayment(r.store)
	}

	p := payment.NewFromEntry(ev).CreateNewEvent()
	if err := r.updateStore(p); err != nil {
		return nil, err
	}

	return p.GetLatestEvent(), nil
}

func (r *runner) StartUp() {
	go func() {
		for {
			select {
			case <-r.ticker.C:
				ev, err := r.generate()
				if err != nil {
					log.Warn().Err(err).Send()
					r.newEventChance++
					continue
				}
				utils.LogTags(ev)
			}
		}
	}()
	select {}
}
