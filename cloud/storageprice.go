package cloud

// GENERATED SDK for cloud API

// Details about storage pricing
type StoragePrice struct {
	MonthlyPrice Price  `json:"monthlyPrice,omitempty"`
	Price        Price  `json:"price,omitempty"`
	Region       string `json:"region,omitempty"`
}
