package event

type Customer struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
}
