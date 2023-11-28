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

func (e *Event) updateIds() {
	id := utils.NewEventId()
	e.ID = id
	e.ActionID = utils.NewActionId(id)
}

func (e *Event) clearEventSpecific() {

}

func (e *Event) AsRequested() *Event {
	e.updateIds()
	e.Action = ActionRequested
	e.AuthorizedAmount = utils.GenerateRandomNumber()
	e.Status = StatusPending
	return e
}

func (e *Event) AsAuthorized() *Event {
	e.updateIds()
	e.ResponseCode = ResponseCodeSuccess
	e.Action = ActionAuthorize
	e.Status = StatusAuthorized
	return e
}

func (e *Event) AsCapture() *Event {
	e.updateIds()
	e.ResponseCode = ResponseCodeSuccess
	e.Action = ActionCapture

	if randomChance() {
		e.CapturedAmount = utils.GenerateRandomNumberBetween(e.AuthorizedAmount - e.CapturedAmount)
	} else {
		e.CapturedAmount = e.AuthorizedAmount
	}
	if e.CapturedAmount == e.AuthorizedAmount {
		e.Status = StatusCaptured
	} else {
		e.Status = StatusPartiallyCaptured
	}
	return e
}

func (e *Event) AsRefund() *Event {
	e.updateIds()
	e.ResponseCode = ResponseCodeSuccess
	e.Action = ActionRefund

	if randomChance() {
		e.RefundedAmount = utils.GenerateRandomNumberBetween(e.CapturedAmount - e.RefundedAmount)
	} else {
		e.RefundedAmount = e.CapturedAmount
	}
	if e.RefundedAmount == e.CapturedAmount {
		e.Status = StatusRefunded
	} else {
		e.Status = StatusPartiallyRefunded
	}
	return e
}

func (e *Event) AsVoid() *Event {
	e.updateIds()
	e.Action = ActionVoid
	e.ResponseCode = ResponseCodeSuccess
	e.Status = StatusCancelled
	return e
}

func (e *Event) AsExpiry() *Event {
	e.updateIds()
	e.Action = ActionExpiry
	e.ResponseCode = ResponseCodeSuccess
	return e
}

func (e *Event) withCustomer(customer Customer) *Event {
	e.Customer = customer
	return e
}

func (e *Event) withRecipient(recipient Recipient) *Event {
	e.Recipient = recipient
	return e
}

func (e *Event) withShipping(address Address) *Event {
	e.ShippingAddress = address
	return e
}

func (e *Event) withBilling(address Address) *Event {
	e.BillingAddress = address
	return e
}

func (e *Event) withCardDetails(details CardDetails) *Event {
	e.CardDetails = details
	return e
}

func (e *Event) withItems(items ...Item) *Event {
	e.Items = items
	return e
}

func (e *Event) withPayment(currency CurrencyCode, paymentMethod PaymentMethod) *Event {
	e.Currency = currency
	e.PaymentMethod = paymentMethod
	return e
}

func (e *Event) Log() {
	log.Info().
		Str("Id", e.ID).
		Str("Timestamp", e.Timestamp.String()).
		Str("PaymentId", e.PaymentID).
		Str("ActionId", e.ActionID).
		Str("ClientId", e.ClientId).
		Str("Action", string(e.Action)).
		Str("Status", string(e.Status)).
		Str("ResponseCode", string(e.ResponseCode)).
		Str("Reference", e.Reference).
		Str("Currency", string(e.Currency)).
		Str("PaymentMethod", string(e.PaymentMethod)).
		Float64("AuthorizedAmount", e.AuthorizedAmount).
		Float64("CapturedAmount", e.CapturedAmount).
		Float64("RefundedAmount", e.RefundedAmount).
		Interface("Metadata", e.Metadata).
		Interface("Items", e.Items).
		Interface("Customer", e.Customer).
		Interface("Recipient", e.Recipient).
		Interface("BillingAddress", e.BillingAddress).
		Interface("ShippingAddress", e.ShippingAddress).
		Interface("CardDetails", e.CardDetails).
		Send()
}
