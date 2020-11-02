package cloud

// GENERATED SDK for cloud API

// A load balancer to handle workload
type LoadBalancer struct {
	Address       Address              `json:"address,omitempty"`
	Configuration ConfigurationVersion `json:"configuration,omitempty"`
	CreatedAt     string               `json:"createdAt,omitempty"`
	Description   *string              `json:"description,omitempty"`
	EgressAddress Addresses            `json:"egressAddress,omitempty"`
	Id            string               `json:"id,omitempty"`
	Name          *string              `json:"name,omitempty"`
	Region        string               `json:"region,omitempty"`
	Status        StatusEnum           `json:"status,omitempty"`
}
