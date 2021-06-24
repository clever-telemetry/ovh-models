package me

// GENERATED SDK for me API

// Details about a fidelity account
type FidelityMovement struct {
	Amount          int64         `json:"amount,omitempty"`
	Balance         int64         `json:"balance,omitempty"`
	Date            string        `json:"date,omitempty"`
	Description     string        `json:"description,omitempty"`
	MovementId      int64         `json:"movementId,omitempty"`
	Operation       OperationEnum `json:"operation,omitempty"`
	Order           int64         `json:"order,omitempty"`
	PreviousBalance int64         `json:"previousBalance,omitempty"`
}
