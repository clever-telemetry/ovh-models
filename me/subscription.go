package me

// GENERATED SDK for me API

// List of all OVH things you can subscribe to
type Subscription struct {
	Registered *bool  `json:"registered,omitempty"`
	Type       string `json:"type,omitempty"`
}
