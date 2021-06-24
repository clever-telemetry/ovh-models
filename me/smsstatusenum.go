package me

// GENERATED SDK for me API

// Status of the Sms account
type SmsStatusEnum string

var (
	SmsStatusEnumNEEDEMAILVALIDATION SmsStatusEnum = "needEmailValidation"
	SmsStatusEnumNEEDCODEVALIDATION  SmsStatusEnum = "needCodeValidation"
	SmsStatusEnumENABLED             SmsStatusEnum = "enabled"
	SmsStatusEnumDISABLED            SmsStatusEnum = "disabled"
)
