package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectKubeCreationNodePool struct {
	AntiAffinity  *bool   `json:"antiAffinity,omitempty"`
	Autoscale     *bool   `json:"autoscale,omitempty"`
	DesiredNodes  *int64  `json:"desiredNodes,omitempty"`
	FlavorName    *string `json:"flavorName,omitempty"`
	MaxNodes      *int64  `json:"maxNodes,omitempty"`
	MinNodes      *int64  `json:"minNodes,omitempty"`
	MonthlyBilled *bool   `json:"monthlyBilled,omitempty"`
	Name          *string `json:"name,omitempty"`
}
