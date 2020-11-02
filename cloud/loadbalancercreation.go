package cloud

// GENERATED SDK for cloud API

// A load balancer to handle workload
type LoadBalancerCreation struct {
	Description *string `json:"description,omitempty"`
	Id          string  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Region      string  `json:"region"`
}
