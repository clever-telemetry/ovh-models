package cloud

// GENERATED SDK for cloud API

// Instance usage
type InstanceUsageDetail struct {
	Hourly         *Price                  `json:"hourly,omitempty"`
	InstanceId     string                  `json:"instanceId,omitempty"`
	Monthly        *InstanceMonthlyBilling `json:"monthly,omitempty"`
	MonthlyBilling bool                    `json:"monthlyBilling,omitempty"`
	Reference      string                  `json:"reference,omitempty"`
}
