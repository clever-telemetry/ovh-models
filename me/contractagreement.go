package me

// GENERATED SDK for me API

// Contract Agreement
type ContractAgreement struct {
	Agreed     AgreementStatusEnum `json:"agreed,omitempty"`
	ContractId int64               `json:"contractId,omitempty"`
	Date       string              `json:"date,omitempty"`
	Id         int64               `json:"id,omitempty"`
}
