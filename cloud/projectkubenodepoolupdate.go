package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectKubeNodePoolUpdate struct {
	Autoscale     *bool     `json:"autoscale,omitempty"`
	DesiredNodes  *int64    `json:"desiredNodes,omitempty"`
	MaxNodes      *int64    `json:"maxNodes,omitempty"`
	MinNodes      *int64    `json:"minNodes,omitempty"`
	NodesToRemove *[]string `json:"nodesToRemove,omitempty"`
}
