package me

// GENERATED SDK for me API

// Operation that happend on a debt
type Operation struct {
	Amount         Price          `json:"amount,omitempty"`
	Date           string         `json:"date,omitempty"`
	DepositOrderId int64          `json:"depositOrderId,omitempty"`
	OperationId    int64          `json:"operationId,omitempty"`
	Status         StatusEnum     `json:"status,omitempty"`
	Type           *OperationEnum `json:"type,omitempty"`
}
