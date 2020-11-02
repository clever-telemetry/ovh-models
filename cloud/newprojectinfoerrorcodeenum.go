package cloud

// GENERATED SDK for cloud API

// Possible values for error code on project creation
type NewProjectInfoErrorCodeEnum string

var (
	NewProjectInfoErrorCodeEnumUNPAIDDEBTS                     NewProjectInfoErrorCodeEnum = "unpaidDebts"
	NewProjectInfoErrorCodeEnumPAYPALACCOUNTNOTVERIFIED        NewProjectInfoErrorCodeEnum = "paypalAccountNotVerified"
	NewProjectInfoErrorCodeEnumMAXPROJECTSLIMITREACHED         NewProjectInfoErrorCodeEnum = "maxProjectsLimitReached"
	NewProjectInfoErrorCodeEnumINVALIDPAYMENTMEAN              NewProjectInfoErrorCodeEnum = "invalidPaymentMean"
	NewProjectInfoErrorCodeEnumCHALLENGEPAYMENTMETHODREQUESTED NewProjectInfoErrorCodeEnum = "challengePaymentMethodRequested"
	NewProjectInfoErrorCodeEnumACCOUNTNOTELIGIBLE              NewProjectInfoErrorCodeEnum = "accountNotEligible"
)
