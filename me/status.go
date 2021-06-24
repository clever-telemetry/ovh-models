package me

// GENERATED SDK for me API

// Payment transaction status enum
type Status string

var (
	StatusSUCCESS    Status = "SUCCESS"
	StatusREADY      Status = "READY"
	StatusFAILED     Status = "FAILED"
	StatusERROR      Status = "ERROR"
	StatusCREATED    Status = "CREATED"
	StatusCONFIRMING Status = "CONFIRMING"
	StatusCANCELING  Status = "CANCELING"
	StatusCANCELED   Status = "CANCELED"
)
