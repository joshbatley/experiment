package event

type CardDetails struct {
	CardNumber     string `json:"card_number,omitempty"`
	ExpiryMonth    int    `json:"expiry_month,omitempty"`
	ExpiryYear     int    `json:"expiry_year,omitempty"`
	CVV            string `json:"cvv,omitempty"`
	CardholderName string `json:"cardholder_name,omitempty"`
	IssuingBank    string `json:"issuing_bank,omitempty"`
}
