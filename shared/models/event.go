package models

import (
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
