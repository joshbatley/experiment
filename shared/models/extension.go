package models

import (
	"github.com/rs/zerolog/log"
	"math/rand"
	utils "shared"
	"time"
)

func NewEvent(clientId string) *Event {
	id := utils.NewEventId()
	return &Event{
		ID:        id,
		PaymentID: utils.NewPaymentId(),
		ActionID:  utils.NewActionId(id),
		ClientId:  clientId,
		Reference: utils.GenerateRandomReference(8),
		Timestamp: time.Now(),
	}
}

func randomChance() bool {
	randomNum := rand.Intn(10)
	return randomNum == 0
}

func (p *Event) updateIds() {
	id := utils.NewEventId()
	p.ID = id
	p.ActionID = utils.NewActionId(id)
}

func (p *Event) AsRequested() *Event {
	p.updateIds()
	p.Action = ActionRequested
	p.AuthorizedAmount = utils.GenerateRandomNumber()
	p.Status = StatusPending
	return p
}

func (p *Event) AsAuthorized() *Event {
	p.updateIds()
	p.ResponseCode = ResponseCodeSuccess
	p.Action = ActionAuthorize
	p.Status = StatusAuthorized
	return p
}

func (p *Event) AsCapture() *Event {
	p.updateIds()
	p.ResponseCode = ResponseCodeSuccess
	p.Action = ActionCapture

	if randomChance() {
		p.CapturedAmount = utils.GenerateRandomNumberBetween(p.AuthorizedAmount - p.CapturedAmount)
	} else {
		p.CapturedAmount = p.AuthorizedAmount
	}
	if p.CapturedAmount == p.AuthorizedAmount {
		p.Status = StatusCaptured
	} else {
		p.Status = StatusPartiallyCaptured
	}
	return p
}

func (p *Event) AsRefund() *Event {
	p.updateIds()
	p.ResponseCode = ResponseCodeSuccess
	p.Action = ActionRefund

	if randomChance() {
		p.RefundedAmount = utils.GenerateRandomNumberBetween(p.CapturedAmount - p.RefundedAmount)
	} else {
		p.RefundedAmount = p.CapturedAmount
	}
	if p.RefundedAmount == p.CapturedAmount {
		p.Status = StatusRefunded
	} else {
		p.Status = StatusPartiallyRefunded
	}
	return p
}

func (p *Event) AsVoid() *Event {
	p.updateIds()
	p.Action = ActionVoid
	p.ResponseCode = ResponseCodeSuccess
	p.Status = StatusCancelled
	return p
}

func (p *Event) AsExpiry() *Event {
	p.updateIds()
	p.Action = ActionExpiry
	p.ResponseCode = ResponseCodeSuccess
	return p
}

func (p *Event) withCustomer(customer Customer) *Event {
	p.Customer = customer
	return p
}

func (p *Event) withRecipient(recipient Recipient) *Event {
	p.Recipient = recipient
	return p
}

func (p *Event) withShipping(address Address) *Event {
	p.ShippingAddress = address
	return p
}

func (p *Event) withBilling(address Address) *Event {
	p.BillingAddress = address
	return p
}

func (p *Event) withCardDetails(details CardDetails) *Event {
	p.CardDetails = details
	return p
}

func (p *Event) withItems(items ...Item) *Event {
	p.Items = items
	return p
}

func (p *Event) withPayment(currency CurrencyCode, paymentMethod PaymentMethod) *Event {
	p.Currency = currency
	p.PaymentMethod = paymentMethod
	return p
}

func (p *Event) Log() {
	log.Info().
		Str("Id", p.ID).
		Str("Timestamp", p.Timestamp.String()).
		Str("PaymentId", p.PaymentID).
		Str("ActionId", p.ActionID).
		Str("ClientId", p.ClientId).
		Str("Action", string(p.Action)).
		Str("Status", string(p.Status)).
		Str("ResponseCode", string(p.ResponseCode)).
		Str("Reference", p.Reference).
		Str("Currency", string(p.Currency)).
		Str("PaymentMethod", string(p.PaymentMethod)).
		Float64("AuthorizedAmount", p.AuthorizedAmount).
		Float64("CapturedAmount", p.CapturedAmount).
		Float64("RefundedAmount", p.RefundedAmount).
		Interface("Metadata", p.Metadata).
		Interface("Items", p.Items).
		Interface("Customer", p.Customer).
		Interface("Recipient", p.Recipient).
		Interface("BillingAddress", p.BillingAddress).
		Interface("ShippingAddress", p.ShippingAddress).
		Interface("CardDetails", p.CardDetails).
		Send()
}
