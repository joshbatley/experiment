package payment

import (
	utils "shared"
	"shared/event"
)

func constructRequested() *event.Event {
	return event.New("cli_123123", "").AsRequested()
}

func progressAuthorization(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsAuthorized()
	return event, false
}

func progressCapture(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsCapture(p.MaxCapturable())
	return event, false
}

func progressRefund(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsRefund(p.MaxRefundable(), p.IsCaptured())
	return event, false
}

func progressVoid(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsVoid()
	return event, true
}

func progressExpiry(p *Payment) (*event.Event, bool) {
	event := p.latestEvent.AsExpiry()
	return event, true
}

func alwaysTrigger(*Payment) bool {
	return true
}

func triggerCapture(p *Payment) bool {
	if p.CanCapture() {
		return true
	}
	return false
}

func triggerRefund(p *Payment) bool {
	if p.CanRefund() {
		return utils.RandomChance(3)
	}
	return false
}

func triggerVoid(*Payment) bool {
	return utils.RandomChance(1000)
}

func triggerExpiry(p *Payment) bool {
	return p.CanExpire()
}
