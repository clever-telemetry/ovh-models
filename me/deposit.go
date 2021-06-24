package me

// GENERATED SDK for me API

// Details about a deposit
type Deposit struct {
	Amount      Price        `json:"amount,omitempty"`
	Date        string       `json:"date,omitempty"`
	DepositId   string       `json:"depositId,omitempty"`
	OrderId     int64        `json:"orderId,omitempty"`
	Password    string       `json:"password,omitempty"`
	PaymentInfo *PaymentInfo `json:"paymentInfo,omitempty"`
	PdfUrl      string       `json:"pdfUrl,omitempty"`
	Url         string       `json:"url,omitempty"`
}
