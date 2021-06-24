package cloud

// GENERATED SDK for cloud API

// A load balancer to handle application workload
type ApplicationLoadBalancer struct {
	Address         Address              `json:"address,omitempty"`
	Configuration   ConfigurationVersion `json:"configuration,omitempty"`
	CreatedAt       string               `json:"createdAt,omitempty"`
	Description     *string              `json:"description,omitempty"`
	EgressAddress   Addresses            `json:"egressAddress,omitempty"`
	Id              string               `json:"id,omitempty"`
	Name            *string              `json:"name,omitempty"`
	OpenstackRegion string               `json:"openstackRegion,omitempty"`
	Region          string               `json:"region,omitempty"`
	Size            SizeEnum             `json:"size,omitempty"`
	Status          StatusEnum           `json:"status,omitempty"`
}
