package me

// GENERATED SDK for me API

// State of your bank account
type BankAccountStateEnum string

var (
	BankAccountStateEnumVALID               BankAccountStateEnum = "valid"
	BankAccountStateEnumPENDINGVALIDATION   BankAccountStateEnum = "pendingValidation"
	BankAccountStateEnumBLOCKEDFORINCIDENTS BankAccountStateEnum = "blockedForIncidents"
)
