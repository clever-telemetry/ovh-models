package me

// GENERATED SDK for me API

// Customer consent information for a campaign
type Consent struct {
	Campaign string           `json:"campaign,omitempty"`
	History  []Decision       `json:"history,omitempty"`
	Type     CampaignTypeEnum `json:"type,omitempty"`
	Value    bool             `json:"value,omitempty"`
}
