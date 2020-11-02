package cloud

// GENERATED SDK for cloud API

// HourlySnapshot
type HourlySnapshot struct {
	Instance   *InstanceSnapshot `json:"instance,omitempty"`
	Region     string            `json:"region,omitempty"`
	TotalPrice float64           `json:"totalPrice,omitempty"`
	Volume     *VolumeSnapshot   `json:"volume,omitempty"`
}
