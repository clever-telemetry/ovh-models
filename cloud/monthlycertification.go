package cloud

// GENERATED SDK for cloud API

// MonthlyCertification
type MonthlyCertification struct {
	Details    []MonthlyCertificationDetail `json:"details,omitempty"`
	Reference  string                       `json:"reference,omitempty"`
	TotalPrice float64                      `json:"totalPrice,omitempty"`
}
