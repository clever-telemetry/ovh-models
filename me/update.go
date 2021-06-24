package me

// GENERATED SDK for me API

// Purchase order update payload
type Update struct {
	BillingGroupId *int64  `json:"billingGroupId,omitempty"`
	Description    *string `json:"description,omitempty"`
	EndDate        *string `json:"endDate,omitempty"`
	Reference      *string `json:"reference,omitempty"`
	StartDate      *string `json:"startDate,omitempty"`
}
