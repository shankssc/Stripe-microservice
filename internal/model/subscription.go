package model

import "time"

// Subscription represents a subscription in the Stripe system.
type Subscription struct {
	ID                 string     `json:"id"`
	CustomerID         string     `json:"customer_id"`
	Status             string     `json:"status"`  // e.g., "active", "canceled", "past_due"
	PlanID             string     `json:"plan_id"` // Reference to the plan
	CurrentPeriodStart *time.Time `json:"current_period_start"`
	CurrentPeriodEnd   *time.Time `json:"current_period_end"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
}
