package cloud

// GENERATED SDK for cloud API

// MonthlyInstance
type MonthlyInstance struct {
	Details    []MonthlyInstanceDetail `json:"details,omitempty"`
	Reference  string                  `json:"reference,omitempty"`
	Region     string                  `json:"region,omitempty"`
	TotalPrice float64                 `json:"totalPrice,omitempty"`
}
