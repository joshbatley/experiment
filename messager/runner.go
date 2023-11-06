package main

import (
	"github.com/rs/zerolog/log"
	"math/rand"
	"shared/models"
	"time"
)

// TODO: Add some sort of limiter for the new event chance

type Runner struct {
	store          EventStore
	ticker         *time.Ticker
	newEventChance int
}

func NewRunner(store EventStore, tps int) *Runner {
	return &Runner{
		store:          store,
		ticker:         time.NewTicker(time.Second / time.Duration(tps)),
		newEventChance: 2,
	}
}

func (e *Runner) StartUp() {
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
				logEvent(ev.currEvent)
			}
		}
	}()
	select {}
}

func (e *Runner) generate() (*Payment, error) {
	if e.rndCreateNewPayment() {
		return createNewEvent(e.store)
	}
	record, err := e.store.GetRandomEvent()
	if err != nil {
		return createNewEvent(e.store)
	}

	payment := PaymentFromRecord(record)
	payment.Progress()
	if payment.isCompletedPayment {
		if err := e.store.RemoveEvent(payment.currEvent.ID); err != nil {
			return nil, err
		}
		return payment, nil
	}
	if err := e.store.UpdateEvent(payment.ToRecord()); err != nil {
		return nil, err
	}
	return payment, nil
}

func (e *Runner) rndCreateNewPayment() bool {
	randomNum := rand.Intn(e.newEventChance)
	return randomNum == 0
}

func createNewEvent(store EventStore) (*Payment, error) {
	payment := NewPayment()
	if err := store.AddUnfinishedEvent(payment.ToRecord()); err != nil {
		return nil, err
	}
	return payment, nil
}

func logEvent(ev *models.Event) {
	log.Info().
		Str("Id", ev.ID).
		Str("Timestamp", ev.Timestamp.String()).
		Str("PaymentId", ev.PaymentID).
		Str("ActionId", ev.ActionID).
		Str("ClientId", ev.ClientId).
		Str("Action", string(ev.Action)).
		Str("Status", string(ev.Status)).
		Str("ResponseCode", string(ev.ResponseCode)).
		Str("Description", ev.Description).
		Str("Currency", string(ev.Currency)).
		Str("PaymentMethod", string(ev.PaymentMethod)).
		Float64("AuthorizedAmount", ev.AuthorizedAmount).
		Float64("CapturedAmount", ev.CapturedAmount).
		Float64("RefundedAmount", ev.RefundedAmount).
		Interface("Metadata", ev.Metadata).
		Interface("Items", ev.Items).
		Interface("Customer", ev.Customer).
		Interface("Recipient", ev.Recipient).
		Interface("BillingAddress", ev.BillingAddress).
		Interface("ShippingAddress", ev.ShippingAddress).
		Interface("CardDetails", ev.CardDetails).
		Send()
}
