package cloud

// GENERATED SDK for cloud API

// Node installed on your cluster
type Node struct {
	CreatedAt  string         `json:"createdAt,omitempty"`
	DeployedAt *string        `json:"deployedAt,omitempty"`
	Flavor     string         `json:"flavor,omitempty"`
	Id         string         `json:"id,omitempty"`
	InstanceId *string        `json:"instanceId,omitempty"`
	IsUpToDate bool           `json:"isUpToDate,omitempty"`
	Name       *string        `json:"name,omitempty"`
	NodePoolId string         `json:"nodePoolId,omitempty"`
	ProjectId  string         `json:"projectId,omitempty"`
	Status     NodeStatusEnum `json:"status,omitempty"`
	UpdatedAt  string         `json:"updatedAt,omitempty"`
	Version    string         `json:"version,omitempty"`
}
