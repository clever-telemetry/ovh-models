package cloud

// GENERATED SDK for cloud API

// Cloud credit
type Credit struct {
	Available_credit Price            `json:"available_credit,omitempty"`
	Bill             *string          `json:"bill,omitempty"`
	Description      *string          `json:"description,omitempty"`
	Id               int64            `json:"id,omitempty"`
	Products         *[]string        `json:"products,omitempty"`
	Total_credit     Price            `json:"total_credit,omitempty"`
	Used_credit      Price            `json:"used_credit,omitempty"`
	Validity         *VoucherValidity `json:"validity,omitempty"`
	Voucher          *string          `json:"voucher,omitempty"`
}
