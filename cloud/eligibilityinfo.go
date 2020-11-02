package cloud

// GENERATED SDK for cloud API

// Eligibility information
type EligibilityInfo struct {
	ActionsRequired          *[]EligibilityAction       `json:"actionsRequired,omitempty"`
	MinimumCredit            *Price                     `json:"minimumCredit,omitempty"`
	PaymentMethodsAuthorized *[]PaymentMethodAuthorized `json:"paymentMethodsAuthorized,omitempty"`
	Voucher                  *NewProjectInfoVoucher     `json:"voucher,omitempty"`
}
