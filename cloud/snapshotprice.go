package cloud

// GENERATED SDK for cloud API

// Details about snapshot pricing
type SnapshotPrice struct {
	MonthlyPrice Price  `json:"monthlyPrice,omitempty"`
	Price        Price  `json:"price,omitempty"`
	Region       string `json:"region,omitempty"`
}
