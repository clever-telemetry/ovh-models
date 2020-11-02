package cloud

// GENERATED SDK for cloud API

// Instance monthly billing details
type InstanceMonthlyBilling struct {
	ActivatedOn string `json:"activatedOn,omitempty"`
	Cost        Price  `json:"cost,omitempty"`
}
