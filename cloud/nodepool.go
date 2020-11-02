package cloud

// GENERATED SDK for cloud API

// NodePool created on your cluster to manage your nodes
type NodePool struct {
	AntiAffinity   bool                   `json:"antiAffinity,omitempty"`
	AvailableNodes int64                  `json:"availableNodes,omitempty"`
	CreatedAt      string                 `json:"createdAt,omitempty"`
	CurrentNodes   int64                  `json:"currentNodes,omitempty"`
	DesiredNodes   int64                  `json:"desiredNodes,omitempty"`
	Flavor         string                 `json:"flavor,omitempty"`
	Id             string                 `json:"id,omitempty"`
	MaxNodes       int64                  `json:"maxNodes,omitempty"`
	MinNodes       int64                  `json:"minNodes,omitempty"`
	MonthlyBilled  bool                   `json:"monthlyBilled,omitempty"`
	Name           string                 `json:"name,omitempty"`
	ProjectId      string                 `json:"projectId,omitempty"`
	SizeStatus     NodePoolSizeStatusEnum `json:"sizeStatus,omitempty"`
	Status         NodePoolStatusEnum     `json:"status,omitempty"`
	UpToDateNodes  int64                  `json:"upToDateNodes,omitempty"`
	UpdatedAt      string                 `json:"updatedAt,omitempty"`
}
