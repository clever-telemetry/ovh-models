package me

// GENERATED SDK for me API

// Details about a Bill
type Bill struct {
	BillId          string       `json:"billId,omitempty"`
	Category        CategoryEnum `json:"category,omitempty"`
	Date            string       `json:"date,omitempty"`
	OrderId         int64        `json:"orderId,omitempty"`
	Password        string       `json:"password,omitempty"`
	PdfUrl          string       `json:"pdfUrl,omitempty"`
	PriceWithTax    Price        `json:"priceWithTax,omitempty"`
	PriceWithoutTax Price        `json:"priceWithoutTax,omitempty"`
	Tax             Price        `json:"tax,omitempty"`
	Url             string       `json:"url,omitempty"`
}
