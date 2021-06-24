package me

// GENERATED SDK for me API

// U2F Register Request
type U2FSignChallenge struct {
	ApplicationId string         `json:"applicationId,omitempty"`
	Request       U2FSignRequest `json:"request,omitempty"`
}
