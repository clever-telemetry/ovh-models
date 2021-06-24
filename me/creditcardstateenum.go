package me

// GENERATED SDK for me API

// State of your credit card
type CreditCardStateEnum string

var (
	CreditCardStateEnumVALID           CreditCardStateEnum = "valid"
	CreditCardStateEnumTOOMANYFAILURES CreditCardStateEnum = "tooManyFailures"
	CreditCardStateEnumEXPIRED         CreditCardStateEnum = "expired"
)
