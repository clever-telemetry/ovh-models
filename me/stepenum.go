package me

// GENERATED SDK for me API

// Status of order follow-up
type StepEnum string

var (
	StepEnumVALIDATING StepEnum = "VALIDATING"
	StepEnumVALIDATED  StepEnum = "VALIDATED"
	StepEnumDELIVERING StepEnum = "DELIVERING"
	StepEnumAVAILABLE  StepEnum = "AVAILABLE"
)
