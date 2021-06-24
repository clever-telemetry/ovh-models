package cloud

// GENERATED SDK for cloud API

// A load balancer backend
type Backend struct {
	Balancer      *BalancerAlgorithmEnum `json:"balancer,omitempty"`
	Name          string                 `json:"name"`
	ProxyProtocol *ProxyProtocolEnum     `json:"proxyProtocol,omitempty"`
	Servers       []Server               `json:"servers"`
	Sticky        *bool                  `json:"sticky,omitempty"`
}
