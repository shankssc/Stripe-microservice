package model

import "time"

// Charge represents a charge in the Stripe system.
type Charge struct {
	ID              string     `json:"id"`
	CustomerID      string     `json:"customer_id"`
	PaymentMethodID string     `json:"payment_method_id"`
	Amount          int64      `json:"amount"`      // Amount in smallest currency unit (e.g., cents)
	Currency        string     `json:"currency"`    // Currency code (e.g., "USD", "EUR")
	Status          string     `json:"status"`      // e.g., "succeeded", "pending", "failed"
	Description     string     `json:"description"` // Optional description
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
