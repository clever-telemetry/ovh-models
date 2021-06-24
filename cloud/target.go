package cloud

// GENERATED SDK for cloud API

// A load balancer target
type Target struct {
	Balancer      *BalancerAlgorithmEnum `json:"balancer,omitempty"`
	Name          string                 `json:"name"`
	ProxyProtocol *ProxyProtocolEnum     `json:"proxyProtocol,omitempty"`
	Servers       []Server               `json:"servers"`
	Sticky        *bool                  `json:"sticky,omitempty"`
}
