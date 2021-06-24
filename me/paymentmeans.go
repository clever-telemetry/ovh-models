package me

// GENERATED SDK for me API

// TODO
type PaymentMeans struct {
	CreditCard     *[]PaymentMean `json:"creditCard,omitempty"`
	Edinar         *[]PaymentMean `json:"edinar,omitempty"`
	FidelityPoints *[]PaymentMean `json:"fidelityPoints,omitempty"`
	Ideal          *[]PaymentMean `json:"ideal,omitempty"`
	Multibanco     *[]PaymentMean `json:"multibanco,omitempty"`
	OvhAccount     *[]PaymentMean `json:"ovhAccount,omitempty"`
	Paypal         *[]PaymentMean `json:"paypal,omitempty"`
	Promotion      *[]PaymentMean `json:"promotion,omitempty"`
}
