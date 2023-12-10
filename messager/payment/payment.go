package payment

import (
	"messager/store"
	"shared/event"
	"sort"
	"time"
)

const (
	ExpireAfter = time.Second * 3000
)

type Payment struct {
	events      []*event.Event
	latestEvent *event.Event
	isFinalised bool
	clientId    string
}

func New(clientId string) *Payment {
	return &Payment{
		events:      []*event.Event{},
		latestEvent: nil,
		isFinalised: false,
		clientId:    clientId,
	}
}

func NewFromStore(e *store.Entry) *Payment {
	sort.Slice(e.Events, func(i, j int) bool {
		return e.Events[i].Timestamp.Before(e.Events[j].Timestamp)
	})
	return &Payment{
		latestEvent: e.Events[len(e.Events)-1],
		events:      e.Events,
		isFinalised: false,
		clientId:    e.Events[0].ClientId,
	}
}

func (p *Payment) addNewEvent(e *event.Event) {
	p.latestEvent = e
	p.events = append(p.events, e)
}

func (p *Payment) getAuthorisedAmount() (total int) {
	if p.events[0].Status == event.StatusFailed {
		return 0
	}
	return p.events[0].AuthorizedAmount
}

func (p *Payment) getCapturedAmount() (total int) {
	for _, ev := range p.events {
		if ev.Status == event.StatusFailed {
			continue
		}
		total += ev.CapturedAmount
	}
	return total
}

func (p *Payment) getRefundedAmount() (total int) {
	for _, ev := range p.events {
		if ev.Status == event.StatusFailed {
			continue
		}
		total += ev.RefundedAmount
	}
	return total
}

func (p *Payment) hasCapturedEvent() bool {
	for _, ev := range p.events {
		if ev.Status == event.StatusCaptured || ev.Status == event.StatusPartiallyCaptured {
			return true
		}
	}
	return false
}

func (p *Payment) ToEntry() *store.Entry {
	return store.NewEntry(p.latestEvent.PaymentID, p.events)
}

func (p *Payment) MaxCapturable() int {
	if p.IsCaptured() || p.IsRefunded() {
		return 0
	}
	return p.getAuthorisedAmount() - (p.getCapturedAmount() - p.getRefundedAmount())
}

func (p *Payment) MaxRefundable() int {
	if !p.hasCapturedEvent() || p.IsRefunded() {
		return 0
	}
	return p.getCapturedAmount() - p.getRefundedAmount()
}

func (p *Payment) IsCaptured() bool {
	for _, ev := range p.events {
		if ev.Status == event.StatusCaptured {
			return true
		}
	}
	return false
}

func (p *Payment) IsRefunded() bool {
	for _, ev := range p.events {
		if ev.Status == event.StatusRefunded {
			return true
		}
	}
	return false
}

func (p *Payment) IsAuthorised() bool {
	for _, ev := range p.events {
		if ev.Status == event.StatusAuthorized {
			return true
		}
	}
	return false
}

func (p *Payment) CanCapture() bool {
	return !p.isFinalised && p.IsAuthorised() && !p.IsCaptured() && p.MaxCapturable() != 0
}

func (p *Payment) CanRefund() bool {
	return !p.isFinalised && p.IsAuthorised() && p.hasCapturedEvent() && !p.IsRefunded() && p.MaxRefundable() != 0
}

func (p *Payment) CanExpire() bool {
	return (p.latestEvent.Status == event.StatusAuthorized || p.latestEvent.Status == event.StatusPending) &&
		p.latestEvent.Timestamp.After(time.Now().Add(ExpireAfter))
}

func (p *Payment) HasFinalEvent() bool {
	return p.isFinalised
}

func (p *Payment) CreateNewEvent() *Payment {
	e, isComplete := createNewEvent(p)
	if e != nil {
		p.addNewEvent(e)
	}
	p.isFinalised = isComplete
	return p
}

func (p *Payment) GetLatestEvent() *event.Event {
	return p.latestEvent
}

func (p *Payment) NumOfEvent() int {
	return len(p.events)
}
