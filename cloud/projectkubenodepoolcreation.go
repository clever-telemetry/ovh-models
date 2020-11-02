package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectKubeNodePoolCreation struct {
	AntiAffinity  *bool   `json:"antiAffinity,omitempty"`
	DesiredNodes  *int64  `json:"desiredNodes,omitempty"`
	FlavorName    string  `json:"flavorName"`
	MaxNodes      *int64  `json:"maxNodes,omitempty"`
	MinNodes      *int64  `json:"minNodes,omitempty"`
	MonthlyBilled *bool   `json:"monthlyBilled,omitempty"`
	Name          *string `json:"name,omitempty"`
}
