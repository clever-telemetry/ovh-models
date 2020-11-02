package cloud

// GENERATED SDK for cloud API

// Details about flavor pricing
type FlavorPrice struct {
	FlavorId     string `json:"flavorId,omitempty"`
	FlavorName   string `json:"flavorName,omitempty"`
	MonthlyPrice *Price `json:"monthlyPrice,omitempty"`
	Price        Price  `json:"price,omitempty"`
	Region       string `json:"region,omitempty"`
}
