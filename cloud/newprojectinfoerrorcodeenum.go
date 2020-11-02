package cloud

// GENERATED SDK for cloud API

// Possible values for error code on project creation
type NewProjectInfoErrorCodeEnum string

var (
	NewProjectInfoErrorCodeEnumACCOUNTNOTELIGIBLE              NewProjectInfoErrorCodeEnum = "accountNotEligible"
	NewProjectInfoErrorCodeEnumCHALLENGEPAYMENTMETHODREQUESTED NewProjectInfoErrorCodeEnum = "challengePaymentMethodRequested"
	NewProjectInfoErrorCodeEnumINVALIDPAYMENTMEAN              NewProjectInfoErrorCodeEnum = "invalidPaymentMean"
	NewProjectInfoErrorCodeEnumMAXPROJECTSLIMITREACHED         NewProjectInfoErrorCodeEnum = "maxProjectsLimitReached"
	NewProjectInfoErrorCodeEnumPAYPALACCOUNTNOTVERIFIED        NewProjectInfoErrorCodeEnum = "paypalAccountNotVerified"
	NewProjectInfoErrorCodeEnumUNPAIDDEBTS                     NewProjectInfoErrorCodeEnum = "unpaidDebts"
)
