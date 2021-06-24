package me

// GENERATED SDK for me API

// API Application
type Application struct {
	ApplicationId  int64                 `json:"applicationId,omitempty"`
	ApplicationKey string                `json:"applicationKey,omitempty"`
	Description    string                `json:"description,omitempty"`
	Name           string                `json:"name,omitempty"`
	Status         ApplicationStatusEnum `json:"status,omitempty"`
}
