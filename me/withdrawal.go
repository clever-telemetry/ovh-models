package me

// GENERATED SDK for me API

// Details about a withdrawal
type Withdrawal struct {
	Amount       Price  `json:"amount,omitempty"`
	Country      string `json:"country,omitempty"`
	Date         string `json:"date,omitempty"`
	OrderId      int64  `json:"orderId,omitempty"`
	Password     string `json:"password,omitempty"`
	PdfUrl       string `json:"pdfUrl,omitempty"`
	Url          string `json:"url,omitempty"`
	WithdrawalId string `json:"withdrawalId,omitempty"`
}
