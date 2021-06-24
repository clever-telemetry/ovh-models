package me

// GENERATED SDK for me API

// Purchase Order
type PurchaseOrder struct {
	BillingGroupId *int64     `json:"billingGroupId,omitempty"`
	CreationDate   string     `json:"creationDate,omitempty"`
	Description    *string    `json:"description,omitempty"`
	EndDate        *string    `json:"endDate,omitempty"`
	Id             int64      `json:"id,omitempty"`
	LastUpdate     string     `json:"lastUpdate,omitempty"`
	Reference      string     `json:"reference,omitempty"`
	StartDate      string     `json:"startDate,omitempty"`
	Status         StatusEnum `json:"status,omitempty"`
}
