package me

// GENERATED SDK for me API

// Status of TOTP account
type TOTPStatusEnum string

var (
	TOTPStatusEnumNEEDEMAILVALIDATION TOTPStatusEnum = "needEmailValidation"
	TOTPStatusEnumNEEDCODEVALIDATION  TOTPStatusEnum = "needCodeValidation"
	TOTPStatusEnumENABLED             TOTPStatusEnum = "enabled"
	TOTPStatusEnumDISABLED            TOTPStatusEnum = "disabled"
)
