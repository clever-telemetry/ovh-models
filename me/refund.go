package me

// GENERATED SDK for me API

// Details about a Refund
type Refund struct {
	Date            string  `json:"date,omitempty"`
	OrderId         int64   `json:"orderId,omitempty"`
	OriginalBillId  *string `json:"originalBillId,omitempty"`
	Password        string  `json:"password,omitempty"`
	PdfUrl          string  `json:"pdfUrl,omitempty"`
	PriceWithTax    Price   `json:"priceWithTax,omitempty"`
	PriceWithoutTax Price   `json:"priceWithoutTax,omitempty"`
	RefundId        string  `json:"refundId,omitempty"`
	Tax             Price   `json:"tax,omitempty"`
	Url             string  `json:"url,omitempty"`
}
