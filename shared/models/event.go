package models

import "time"

type Event struct {
	ID               string                 `json:"id,omitempty"`
	PaymentID        string                 `json:"payment_id,omitempty"`
	ActionID         string                 `json:"action_id,omitempty"`
	ClientId         string                 `json:"client_id,omitempty"`
	Description      string                 `json:"description,omitempty"`
	CaptureAmount    float64                `json:"capture_amount,omitempty"`
	AuthorizedAmount float64                `json:"authorized_amount,omitempty"`
	RefundAmount     float64                `json:"refund_amount,omitempty"`
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

func NewEvent(
	ID string, paymentID string, actionID string, clientId string,
	description string, currency CurrencyCode, paymentMethod PaymentMethod,
) *Event {
	return &Event{
		ID:            ID,
		PaymentID:     paymentID,
		ActionID:      actionID,
		ClientId:      clientId,
		Description:   description,
		Currency:      currency,
		PaymentMethod: paymentMethod,
		Timestamp:     time.Now(),
	}
}

func (p *Event) AsRequested(amount float64, code ResponseCode) *Event {
	p.CaptureAmount = amount
	p.ResponseCode = code
	p.Action = ActionRequested
	return p
}

func (p *Event) AsAuthorized(amount float64, code ResponseCode) *Event {
	p.AuthorizedAmount = amount
	p.ResponseCode = code
	p.Action = ActionAuthorize
	return p
}

func (p *Event) AsCapture(amount float64, code ResponseCode) *Event {
	p.CaptureAmount = amount
	p.ResponseCode = code
	p.Action = ActionCapture
	return p
}

func (p *Event) AsRefund(amount float64, code ResponseCode) *Event {
	p.RefundAmount = amount
	p.ResponseCode = code
	p.Action = ActionRefund
	return p
}

func (p *Event) AsVoid(code ResponseCode) *Event {
	p.Action = ActionVoid
	p.ResponseCode = code
	return p
}

func (p *Event) AsExpiry(code ResponseCode) *Event {
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
