package cloud

// GENERATED SDK for cloud API

// Details about archive storage pricing
type ArchiveStoragePrice struct {
	MonthlyPrice Price  `json:"monthlyPrice,omitempty"`
	Region       string `json:"region,omitempty"`
}
