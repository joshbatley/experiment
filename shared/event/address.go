package event

type Address struct {
	Street1    string `json:"street_1,omitempty"`
	Street2    string `json:"street_2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"payment,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
}
