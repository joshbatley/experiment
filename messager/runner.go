package main

import (
	"errors"
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

func (r *runner) updateStore(p *payment.Payment) error {
	if p.HasFinalEvent() {
		return r.store.Delete(p.GetLatestEvent().PaymentID)
	}
	return r.store.UpdateOrInsert(p.ToEntry())
}

func (r *runner) getOrCreatePayment() (*payment.Payment, error) {
	if utils.RandomChance(r.newEventChance) {
		return payment.New("cli_123").CreateNewEvent(), nil
	}
	ev, err := r.store.GetRandom()
	if errors.Is(err, store.ErrNoEvents) {
		return payment.New("cli_123").CreateNewEvent(), nil
	}
	if err != nil {
		return nil, err
	}

	return payment.NewFromStore(ev).CreateNewEvent(), nil
}

func (r *runner) generate() (*event.Event, error) {
	p, err := r.getOrCreatePayment()
	if err != nil {
		return nil, err
	}
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
