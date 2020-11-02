package cloud

// GENERATED SDK for cloud API

// Address to reach the load balancer
type Address struct {
	Ipv4 string  `json:"ipv4,omitempty"`
	Ipv6 *string `json:"ipv6,omitempty"`
}
