package me

// GENERATED SDK for me API

// Describe all services impacted by SLA
type SlaOperationService struct {
	Description    string `json:"description,omitempty"`
	ServiceName    string `json:"serviceName,omitempty"`
	SlaApplication string `json:"slaApplication,omitempty"`
	SlaPlan        string `json:"slaPlan,omitempty"`
}
