package me

// GENERATED SDK for me API

// All states a Credential can be in
type CredentialStateEnum string

var (
	CredentialStateEnumVALIDATED         CredentialStateEnum = "validated"
	CredentialStateEnumREFUSED           CredentialStateEnum = "refused"
	CredentialStateEnumPENDINGVALIDATION CredentialStateEnum = "pendingValidation"
	CredentialStateEnumEXPIRED           CredentialStateEnum = "expired"
)
