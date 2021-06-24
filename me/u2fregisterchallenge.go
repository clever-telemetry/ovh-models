package me

// GENERATED SDK for me API

// U2F Register Request
type U2FRegisterChallenge struct {
	ApplicationId string                 `json:"applicationId,omitempty"`
	Id            int64                  `json:"id,omitempty"`
	Request       U2FRegistrationRequest `json:"request,omitempty"`
}
