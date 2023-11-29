package event

import (
	"time"
)

type Event struct {
	ID               string                 `json:"id,omitempty" log:"omitempty" log:"omitempty"`
	PaymentID        string                 `json:"payment_id,omitempty" log:"omitempty"`
	ActionID         string                 `json:"action_id,omitempty" log:"omitempty"`
	ClientId         string                 `json:"client_id,omitempty" log:"omitempty"`
	Reference        string                 `json:"description,omitempty" log:"omitempty"`
	CapturedAmount   int                    `json:"captured_amount,omitempty" log:"omitempty"`
	AuthorizedAmount int                    `json:"authorized_amount,omitempty" log:"omitempty"`
	RefundedAmount   int                    `json:"refunded_amount,omitempty" log:"omitempty"`
	Timestamp        time.Time              `json:"timestamp" log:"omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty" log:"omitempty"`
	Currency         CurrencyCode           `json:"currency,omitempty" log:"omitempty"`
	Customer         Customer               `json:"customer" log:"omitempty"`
	Recipient        Recipient              `json:"recipient" log:"omitempty"`
	PaymentMethod    PaymentMethod          `json:"payment_method,omitempty" log:"omitempty"`
	Action           Action                 `json:"action,omitempty" log:"omitempty"`
	Status           Status                 `json:"status,omitempty" log:"omitempty"` // This is never set in events
	BillingAddress   Address                `json:"billing_address" log:"omitempty"`
	ShippingAddress  Address                `json:"shipping_address" log:"omitempty"`
	CardDetails      CardDetails            `json:"card_details" log:"omitempty"`
	Items            []Item                 `json:"items,omitempty" log:"omitempty"`
	ResponseCode     ResponseCode           `json:"response_code,omitempty" log:"omitempty"`
}
