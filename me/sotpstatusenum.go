package me

// GENERATED SDK for me API

// Status of SOTP account
type SOTPStatusEnum string

var (
	SOTPStatusEnumNEEDEMAILVALIDATION SOTPStatusEnum = "needEmailValidation"
	SOTPStatusEnumNEEDCODEVALIDATION  SOTPStatusEnum = "needCodeValidation"
	SOTPStatusEnumENABLED             SOTPStatusEnum = "enabled"
	SOTPStatusEnumDISABLED            SOTPStatusEnum = "disabled"
)
