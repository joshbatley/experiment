package models

import "time"

type Payment struct {
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

func NewPayment(
	ID string, paymentID string, actionID string, clientId string,
	description string, currency CurrencyCode, paymentMethod PaymentMethod, responseCode ResponseCode,
) *Payment {
	return &Payment{
		ID:            ID,
		PaymentID:     paymentID,
		ActionID:      actionID,
		ClientId:      clientId,
		Description:   description,
		Currency:      currency,
		PaymentMethod: paymentMethod,
		ResponseCode:  responseCode,
		Timestamp:     time.Now(),
	}
}

func (p *Payment) AsCapture(amount float64, code ResponseCode) *Payment {
	p.CaptureAmount = amount
	p.ResponseCode = code
	p.Action = ActionCapture
	return p
}

func (p *Payment) AsAuthorized(amount float64, code ResponseCode) *Payment {
	p.AuthorizedAmount = amount
	p.ResponseCode = code
	p.Action = ActionAuthorize
	return p
}

func (p *Payment) AsRefund(amount float64, code ResponseCode) *Payment {
	p.RefundAmount = amount
	p.ResponseCode = code
	p.Action = ActionRefund
	return p
}

func (p *Payment) AsVoid(code ResponseCode) *Payment {
	p.Action = ActionVoid
	p.ResponseCode = code
	return p
}

func (p *Payment) WithCustomer(customer Customer) *Payment {
	p.Customer = customer
	return p
}

func (p *Payment) WithRecipient(recipient Recipient) *Payment {
	p.Recipient = recipient
	return p
}

func (p *Payment) WithShipping(address Address) *Payment {
	p.ShippingAddress = address
	return p
}

func (p *Payment) WithBilling(address Address) *Payment {
	p.BillingAddress = address
	return p
}

func (p *Payment) WithCardDetails(details CardDetails) *Payment {
	p.CardDetails = details
	return p
}

func (p *Payment) WithItems(items ...Item) *Payment {
	p.Items = items
	return p
}
