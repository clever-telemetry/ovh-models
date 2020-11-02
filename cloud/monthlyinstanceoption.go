package cloud

// GENERATED SDK for cloud API

// MonthlyInstanceOption
type MonthlyInstanceOption struct {
	Details    []MonthlyInstanceOptionDetail `json:"details,omitempty"`
	Reference  string                        `json:"reference,omitempty"`
	Region     string                        `json:"region,omitempty"`
	TotalPrice float64                       `json:"totalPrice,omitempty"`
}
