package cloud

// GENERATED SDK for cloud API

// Managed Kubernetes cluster description
type Cluster struct {
	ControlPlaneIsUpToDate bool                  `json:"controlPlaneIsUpToDate,omitempty"`
	CreatedAt              string                `json:"createdAt,omitempty"`
	Id                     string                `json:"id,omitempty"`
	IsUpToDate             bool                  `json:"isUpToDate,omitempty"`
	Name                   string                `json:"name,omitempty"`
	NextUpgradeVersions    *[]UpgradeVersionEnum `json:"nextUpgradeVersions,omitempty"`
	NodesUrl               string                `json:"nodesUrl,omitempty"`
	PrivateNetworkId       *string               `json:"privateNetworkId,omitempty"`
	Region                 RegionEnum            `json:"region,omitempty"`
	Status                 ClusterStatusEnum     `json:"status,omitempty"`
	UpdatePolicy           string                `json:"updatePolicy,omitempty"`
	UpdatedAt              string                `json:"updatedAt,omitempty"`
	Url                    string                `json:"url,omitempty"`
	Version                string                `json:"version,omitempty"`
}
