package me

// GENERATED SDK for me API

// Details about a voucher account
type VoucherMovement struct {
	Amount          Price         `json:"amount,omitempty"`
	Balance         Price         `json:"balance,omitempty"`
	Date            string        `json:"date,omitempty"`
	Description     string        `json:"description,omitempty"`
	MovementId      int64         `json:"movementId,omitempty"`
	Operation       OperationEnum `json:"operation,omitempty"`
	Order           int64         `json:"order,omitempty"`
	PreviousBalance Price         `json:"previousBalance,omitempty"`
}
