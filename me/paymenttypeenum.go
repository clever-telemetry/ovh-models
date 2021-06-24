package me

// GENERATED SDK for me API

// List of payment type enum
type PaymentTypeEnum string

var (
	PaymentTypeEnumPAYPAL                   PaymentTypeEnum = "PAYPAL"
	PaymentTypeEnumINTERNAL_TRUSTED_ACCOUNT PaymentTypeEnum = "INTERNAL_TRUSTED_ACCOUNT"
	PaymentTypeEnumENTERPRISE               PaymentTypeEnum = "ENTERPRISE"
	PaymentTypeEnumDEFERRED_PAYMENT_ACCOUNT PaymentTypeEnum = "DEFERRED_PAYMENT_ACCOUNT"
	PaymentTypeEnumCURRENT_ACCOUNT          PaymentTypeEnum = "CURRENT_ACCOUNT"
	PaymentTypeEnumCREDIT_CARD              PaymentTypeEnum = "CREDIT_CARD"
	PaymentTypeEnumBANK_ACCOUNT             PaymentTypeEnum = "BANK_ACCOUNT"
)
