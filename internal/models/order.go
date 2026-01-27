package models

import "time"

// Order representa una orden en el sistema
type Order struct {
	ID        string    `json:"id"`
	Customer  string    `json:"customer"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
