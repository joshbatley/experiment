package event

type Recipient struct {
	Name          string `json:"name,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	BankName      string `json:"bank_name,omitempty"`
}
