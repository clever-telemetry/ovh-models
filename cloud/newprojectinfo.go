package cloud

// GENERATED SDK for cloud API

// New cloud project informations
type NewProjectInfo struct {
	Agreements *[]int64               `json:"agreements,omitempty"`
	Error      *NewProjectInfoError   `json:"error,omitempty"`
	Order      *Price                 `json:"order,omitempty"`
	Voucher    *NewProjectInfoVoucher `json:"voucher,omitempty"`
}
