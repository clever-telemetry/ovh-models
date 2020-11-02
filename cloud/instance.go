package cloud

// GENERATED SDK for cloud API

// Instance
type Instance struct {
	Created        string             `json:"created,omitempty"`
	FlavorId       string             `json:"flavorId,omitempty"`
	Id             string             `json:"id,omitempty"`
	ImageId        string             `json:"imageId,omitempty"`
	IpAddresses    []IpAddress        `json:"ipAddresses,omitempty"`
	MonthlyBilling *MonthlyBilling    `json:"monthlyBilling,omitempty"`
	Name           string             `json:"name,omitempty"`
	OperationIds   []string           `json:"operationIds,omitempty"`
	PlanCode       *string            `json:"planCode,omitempty"`
	Region         string             `json:"region,omitempty"`
	SshKeyId       *string            `json:"sshKeyId,omitempty"`
	Status         InstanceStatusEnum `json:"status,omitempty"`
}
