package cloud

// GENERATED SDK for cloud API

// Possible eligibility actions
type EligibilityAction string

var (
	EligibilityActionVERIFYPAYPAL             EligibilityAction = "verifyPaypal"
	EligibilityActionCHALLENGEPAYMENTMETHOD   EligibilityAction = "challengePaymentMethod"
	EligibilityActionASKINCREASEPROJECTSQUOTA EligibilityAction = "askIncreaseProjectsQuota"
	EligibilityActionADDPAYMENTMETHOD         EligibilityAction = "addPaymentMethod"
)
