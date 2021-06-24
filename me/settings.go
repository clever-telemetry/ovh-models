package me

// GENERATED SDK for me API

// Telephony settings linked to the customer account
type Settings struct {
	BillingPolicies         BillingSettings         `json:"billingPolicies,omitempty"`
	LineDescriptionPolicies LineDescriptionSettings `json:"lineDescriptionPolicies,omitempty"`
}
