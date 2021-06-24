package me

// GENERATED SDK for me API

// Describe U2F SignRequest
type U2FSignRequest struct {
	Challenge string `json:"challenge,omitempty"`
	KeyHandle string `json:"keyHandle,omitempty"`
	Version   string `json:"version,omitempty"`
}
