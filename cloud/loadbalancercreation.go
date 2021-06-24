package cloud

// GENERATED SDK for cloud API

// A load balancer to handle workload
type LoadBalancerCreation struct {
	Description     *string             `json:"description,omitempty"`
	Id              string              `json:"id,omitempty"`
	Name            *string             `json:"name,omitempty"`
	Networking      *NetworkingCreation `json:"networking,omitempty"`
	OpenstackRegion string              `json:"openstackRegion,omitempty"`
	Region          string              `json:"region"`
	Size            *SizeEnum           `json:"size,omitempty"`
}
