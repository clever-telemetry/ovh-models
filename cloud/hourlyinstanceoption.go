package cloud

// GENERATED SDK for cloud API

// HourlyInstanceOption
type HourlyInstanceOption struct {
	Details    []HourlyInstanceOptionDetail `json:"details,omitempty"`
	Quantity   Quantity                     `json:"quantity,omitempty"`
	Reference  string                       `json:"reference,omitempty"`
	Region     string                       `json:"region,omitempty"`
	TotalPrice float64                      `json:"totalPrice,omitempty"`
}
