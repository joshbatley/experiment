package models

import (
	"github.com/rs/zerolog/log"
	utils "shared"
	"time"
)

type Event struct {
	ID               string                 `json:"id,omitempty"`
	PaymentID        string                 `json:"payment_id,omitempty"`
	ActionID         string                 `json:"action_id,omitempty"`
	ClientId         string                 `json:"client_id,omitempty"`
	Reference        string                 `json:"description,omitempty"`
	CapturedAmount   float64                `json:"captured_amount,omitempty"`
	AuthorizedAmount float64                `json:"authorized_amount,omitempty"`
	RefundedAmount   float64                `json:"refunded_amount,omitempty"`
	Timestamp        time.Time              `json:"timestamp"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	Currency         CurrencyCode           `json:"currency,omitempty"`
	Customer         Customer               `json:"customer"`
	Recipient        Recipient              `json:"recipient"`
	PaymentMethod    PaymentMethod          `json:"payment_method,omitempty"`
	Action           Action                 `json:"action,omitempty"`
	Status           Status                 `json:"status,omitempty"` // This is never set in events
	BillingAddress   Address                `json:"billing_address"`
	ShippingAddress  Address                `json:"shipping_address"`
	CardDetails      CardDetails            `json:"card_details"`
	Items            []Item                 `json:"items,omitempty"`
	ResponseCode     ResponseCode           `json:"response_code,omitempty"`
}

func NewEvent(clientId string, reference string) *Event {
	id := utils.NewEventId()
	return &Event{
		ID:        id,
		PaymentID: utils.NewPaymentId(),
		ActionID:  utils.NewActionId(id),
		ClientId:  clientId,
		Reference: reference,
		Timestamp: time.Now(),
	}
}

func (p *Event) updateIds() {
	id := utils.NewEventId()
	p.ID = id
	p.ActionID = utils.NewActionId(id)
}

func (p *Event) AsRequested() *Event {
	p.updateIds()
	p.Action = ActionRequested
	return p
}

func (p *Event) AsAuthorized(amount float64, code ResponseCode) *Event {
	p.updateIds()
	p.AuthorizedAmount = amount
	p.ResponseCode = code
	p.Action = ActionAuthorize
	return p
}

func (p *Event) AsCapture(amount float64, code ResponseCode) *Event {
	p.updateIds()
	p.CapturedAmount = amount
	p.ResponseCode = code
	p.Action = ActionCapture
	return p
}

func (p *Event) AsRefund(amount float64, code ResponseCode) *Event {
	p.updateIds()
	p.RefundedAmount = amount
	p.ResponseCode = code
	p.Action = ActionRefund
	return p
}

func (p *Event) AsVoid(code ResponseCode) *Event {
	p.updateIds()
	p.Action = ActionVoid
	p.ResponseCode = code
	return p
}

func (p *Event) AsExpiry(code ResponseCode) *Event {
	p.updateIds()
	p.Action = ActionExpiry
	p.ResponseCode = code
	return p
}

func (p *Event) WithCustomer(customer Customer) *Event {
	p.Customer = customer
	return p
}

func (p *Event) WithRecipient(recipient Recipient) *Event {
	p.Recipient = recipient
	return p
}

func (p *Event) WithShipping(address Address) *Event {
	p.ShippingAddress = address
	return p
}

func (p *Event) WithBilling(address Address) *Event {
	p.BillingAddress = address
	return p
}

func (p *Event) WithCardDetails(details CardDetails) *Event {
	p.CardDetails = details
	return p
}

func (p *Event) WithItems(items ...Item) *Event {
	p.Items = items
	return p
}

func (p *Event) WithPayment(currency CurrencyCode, paymentMethod PaymentMethod) *Event {
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
