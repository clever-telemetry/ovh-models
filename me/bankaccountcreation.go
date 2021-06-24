package me

// GENERATED SDK for me API

// Missing description
type BankAccountCreation struct {
	Bic          string  `json:"bic"`
	Description  *string `json:"description,omitempty"`
	Iban         string  `json:"iban"`
	OwnerAddress string  `json:"ownerAddress"`
	OwnerName    string  `json:"ownerName"`
	SetDefault   *bool   `json:"setDefault,omitempty"`
}
