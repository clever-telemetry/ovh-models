package me

// GENERATED SDK for me API

// Available automatic payment means
type AutomaticPaymentMean struct {
	BankAccount            bool `json:"bankAccount,omitempty"`
	CreditCard             bool `json:"creditCard,omitempty"`
	DeferredPaymentAccount bool `json:"deferredPaymentAccount,omitempty"`
	Paypal                 bool `json:"paypal,omitempty"`
}
