package me

// GENERATED SDK for me API

// Indicates the mandatory nature of having a valid payment method
type RequiredPaymentMethodEnum string

var (
	RequiredPaymentMethodEnumNOTMANDATORY          RequiredPaymentMethodEnum = "notMandatory"
	RequiredPaymentMethodEnumMANDATORYFORPOSTPAID  RequiredPaymentMethodEnum = "mandatoryForPostpaid"
	RequiredPaymentMethodEnumMANDATORYFORAUTORENEW RequiredPaymentMethodEnum = "mandatoryForAutorenew"
)
