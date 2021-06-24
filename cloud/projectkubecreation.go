package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectKubeCreation struct {
	Name             *string                      `json:"name,omitempty"`
	Nodepool         *ProjectKubeCreationNodePool `json:"nodepool,omitempty"`
	PrivateNetworkId *string                      `json:"privateNetworkId,omitempty"`
	Region           string                       `json:"region"`
	Version          *VersionEnum                 `json:"version,omitempty"`
}
