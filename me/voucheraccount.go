package me

// GENERATED SDK for me API

// Details about a Voucher account
type VoucherAccount struct {
	Balance          Price  `json:"balance,omitempty"`
	LastUpdate       string `json:"lastUpdate,omitempty"`
	OpenDate         string `json:"openDate,omitempty"`
	VoucherAccountId string `json:"voucherAccountId,omitempty"`
}
