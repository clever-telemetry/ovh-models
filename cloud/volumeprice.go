package cloud

// GENERATED SDK for cloud API

// Details about volume pricing
type VolumePrice struct {
	MonthlyPrice Price  `json:"monthlyPrice,omitempty"`
	Price        Price  `json:"price,omitempty"`
	Region       string `json:"region,omitempty"`
	VolumeName   string `json:"volumeName,omitempty"`
}
