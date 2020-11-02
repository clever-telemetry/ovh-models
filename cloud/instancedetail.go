package cloud

// GENERATED SDK for cloud API

// InstanceDetail
type InstanceDetail struct {
	Created        string             `json:"created,omitempty"`
	Flavor         Flavor             `json:"flavor,omitempty"`
	Id             string             `json:"id,omitempty"`
	Image          Image              `json:"image,omitempty"`
	IpAddresses    []IpAddress        `json:"ipAddresses,omitempty"`
	MonthlyBilling *MonthlyBilling    `json:"monthlyBilling,omitempty"`
	Name           string             `json:"name,omitempty"`
	OperationIds   []string           `json:"operationIds,omitempty"`
	PlanCode       *string            `json:"planCode,omitempty"`
	Region         string             `json:"region,omitempty"`
	SshKey         *SshKeyDetail      `json:"sshKey,omitempty"`
	Status         InstanceStatusEnum `json:"status,omitempty"`
}
