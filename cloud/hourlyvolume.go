package cloud

// GENERATED SDK for cloud API

// HourlyVolume
type HourlyVolume struct {
	Details    []HourlyVolumeDetail `json:"details,omitempty"`
	Quantity   Quantity             `json:"quantity,omitempty"`
	Region     string               `json:"region,omitempty"`
	TotalPrice float64              `json:"totalPrice,omitempty"`
	Type       string               `json:"type,omitempty"`
}
