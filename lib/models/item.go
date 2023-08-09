package models

type Item struct {
	Name     string  `json:"name,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
	Price    float64 `json:"price,omitempty"`
}
