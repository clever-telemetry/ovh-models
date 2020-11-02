package cloud

// GENERATED SDK for cloud API

// HourlyInstance
type HourlyInstance struct {
	Details    []HourlyInstanceDetail `json:"details,omitempty"`
	Quantity   Quantity               `json:"quantity,omitempty"`
	Reference  string                 `json:"reference,omitempty"`
	Region     string                 `json:"region,omitempty"`
	TotalPrice float64                `json:"totalPrice,omitempty"`
}
