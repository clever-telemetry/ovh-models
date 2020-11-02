package cloud

// GENERATED SDK for cloud API

// List of accepted payment methods
type PaymentMethodAuthorized string

var (
	PaymentMethodAuthorizedBANKACCOUNT PaymentMethodAuthorized = "bankAccount"
	PaymentMethodAuthorizedCREDIT      PaymentMethodAuthorized = "credit"
	PaymentMethodAuthorizedCREDITCARD  PaymentMethodAuthorized = "creditCard"
	PaymentMethodAuthorizedPAYPAL      PaymentMethodAuthorized = "paypal"
)
