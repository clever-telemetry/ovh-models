package me

// GENERATED SDK for me API

// Credit card informations
type BankAccount struct {
	Bic                    string               `json:"bic,omitempty"`
	CreationDate           string               `json:"creationDate,omitempty"`
	DefaultPaymentMean     bool                 `json:"defaultPaymentMean,omitempty"`
	Description            *string              `json:"description,omitempty"`
	Iban                   string               `json:"iban,omitempty"`
	Icon                   *IconData            `json:"icon,omitempty"`
	Id                     int64                `json:"id,omitempty"`
	MandateSignatureDate   *string              `json:"mandateSignatureDate,omitempty"`
	OwnerAddress           string               `json:"ownerAddress,omitempty"`
	OwnerName              string               `json:"ownerName,omitempty"`
	State                  BankAccountStateEnum `json:"state,omitempty"`
	UniqueReference        string               `json:"uniqueReference,omitempty"`
	ValidationDocumentLink *string              `json:"validationDocumentLink,omitempty"`
}
