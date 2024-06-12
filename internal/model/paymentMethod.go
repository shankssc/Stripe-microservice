package model

import "time"

// PaymentMethod represents a payment method in the Stripe system.
type PaymentMethod struct {
	ID         string     `json:"id"`
	CustomerID string     `json:"customer_id"`
	Type       string     `json:"type"`       // e.g., "card", "bank_account"
	CardBrand  string     `json:"card_brand"` // e.g., "Visa", "Mastercard" (if type is card)
	CardLast4  string     `json:"card_last4"` // Last 4 digits of card (if type is card)
	ExpMonth   int        `json:"exp_month"`  // Expiry month (if type is card)
	ExpYear    int        `json:"exp_year"`   // Expiry year (if type is card)
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
