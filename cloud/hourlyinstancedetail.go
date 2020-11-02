package cloud

// GENERATED SDK for cloud API

// HourlyInstanceDetail
type HourlyInstanceDetail struct {
	InstanceId string   `json:"instanceId,omitempty"`
	Quantity   Quantity `json:"quantity,omitempty"`
	TotalPrice float64  `json:"totalPrice,omitempty"`
}
