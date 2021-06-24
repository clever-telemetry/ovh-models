package me

// GENERATED SDK for me API

// Transaction object
type Transaction struct {
	Amount          Price      `json:"amount,omitempty"`
	CreationDate    string     `json:"creationDate,omitempty"`
	LastUpdate      string     `json:"lastUpdate,omitempty"`
	PaymentMethodId int64      `json:"paymentMethodId,omitempty"`
	Status          StatusEnum `json:"status,omitempty"`
	TransactionId   int64      `json:"transactionId,omitempty"`
	Type            TypeEnum   `json:"type,omitempty"`
}
