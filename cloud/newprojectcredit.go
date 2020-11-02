package cloud

// GENERATED SDK for cloud API

// Credit details
type NewProjectCredit struct {
	Description  *string          `json:"description,omitempty"`
	Id           int64            `json:"id,omitempty"`
	Products     *[]string        `json:"products,omitempty"`
	Total_credit Price            `json:"total_credit,omitempty"`
	Validity     *VoucherValidity `json:"validity,omitempty"`
}
