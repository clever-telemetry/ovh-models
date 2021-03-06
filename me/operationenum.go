package me

// GENERATED SDK for me API

// All operations a debt entry can represent
type OperationEnum string

var (
	OperationEnumWITHDRAW_AUTOMATIC          OperationEnum = "WITHDRAW_AUTOMATIC"
	OperationEnumWARRANT_MANUAL              OperationEnum = "WARRANT_MANUAL"
	OperationEnumUNPAID_WITHDRAW             OperationEnum = "UNPAID_WITHDRAW"
	OperationEnumUNPAID_SEPA                 OperationEnum = "UNPAID_SEPA"
	OperationEnumUNPAID_PAYU                 OperationEnum = "UNPAID_PAYU"
	OperationEnumUNPAID_PAYPAL               OperationEnum = "UNPAID_PAYPAL"
	OperationEnumUNPAID_MULTIBANCO           OperationEnum = "UNPAID_MULTIBANCO"
	OperationEnumUNPAID_IDEAL                OperationEnum = "UNPAID_IDEAL"
	OperationEnumUNPAID_CREDITCARD           OperationEnum = "UNPAID_CREDITCARD"
	OperationEnumUNPAID_CREDIT_ACCOUNT       OperationEnum = "UNPAID_CREDIT_ACCOUNT"
	OperationEnumUNPAID_CHECK                OperationEnum = "UNPAID_CHECK"
	OperationEnumTRANSFER_MANUAL             OperationEnum = "TRANSFER_MANUAL"
	OperationEnumSEPA_AUTOMATIC              OperationEnum = "SEPA_AUTOMATIC"
	OperationEnumREFUND_UNKNOWN              OperationEnum = "REFUND_UNKNOWN"
	OperationEnumREFUND_TRANSFER             OperationEnum = "REFUND_TRANSFER"
	OperationEnumREFUND_SEPA                 OperationEnum = "REFUND_SEPA"
	OperationEnumREFUND_PAYU                 OperationEnum = "REFUND_PAYU"
	OperationEnumREFUND_PAYPAL               OperationEnum = "REFUND_PAYPAL"
	OperationEnumREFUND_MULTIBANCO           OperationEnum = "REFUND_MULTIBANCO"
	OperationEnumREFUND_LOSS                 OperationEnum = "REFUND_LOSS"
	OperationEnumREFUND_IDEAL                OperationEnum = "REFUND_IDEAL"
	OperationEnumREFUND_CREDITCARD           OperationEnum = "REFUND_CREDITCARD"
	OperationEnumREFUND_CREDIT_ACCOUNT       OperationEnum = "REFUND_CREDIT_ACCOUNT"
	OperationEnumREFUND_CHECK                OperationEnum = "REFUND_CHECK"
	OperationEnumREFUND                      OperationEnum = "REFUND"
	OperationEnumRECOVERY_TRANSFER_AUTOMATIC OperationEnum = "RECOVERY_TRANSFER_AUTOMATIC"
	OperationEnumPAYU_MANUAL                 OperationEnum = "PAYU_MANUAL"
	OperationEnumPAYU_AUTOMATIC              OperationEnum = "PAYU_AUTOMATIC"
	OperationEnumPAYPAL_MANUAL               OperationEnum = "PAYPAL_MANUAL"
	OperationEnumPAYPAL_AUTOMATIC            OperationEnum = "PAYPAL_AUTOMATIC"
	OperationEnumORDER                       OperationEnum = "ORDER"
	OperationEnumMULTIBANCO_MANUAL           OperationEnum = "MULTIBANCO_MANUAL"
	OperationEnumMULTIBANCO_AUTOMATIC        OperationEnum = "MULTIBANCO_AUTOMATIC"
	OperationEnumIDEAL_MANUAL                OperationEnum = "IDEAL_MANUAL"
	OperationEnumIDEAL_AUTOMATIC             OperationEnum = "IDEAL_AUTOMATIC"
	OperationEnumEDINAR_MANUAL               OperationEnum = "EDINAR_MANUAL"
	OperationEnumCREDITCARD_MANUAL           OperationEnum = "CREDITCARD_MANUAL"
	OperationEnumCREDITCARD_AUTOMATIC        OperationEnum = "CREDITCARD_AUTOMATIC"
	OperationEnumCREDITCARD                  OperationEnum = "CREDITCARD"
	OperationEnumCREDIT_ACCOUNT_AUTOMATIC    OperationEnum = "CREDIT_ACCOUNT_AUTOMATIC"
	OperationEnumCHECK_MANUAL                OperationEnum = "CHECK_MANUAL"
	OperationEnumCASH_MANUAL                 OperationEnum = "CASH_MANUAL"
	OperationEnumCANCEL                      OperationEnum = "CANCEL"
)
