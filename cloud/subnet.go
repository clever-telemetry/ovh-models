package cloud

// GENERATED SDK for cloud API

// Subnet
type Subnet struct {
	Cidr      string   `json:"cidr,omitempty"`
	GatewayIp *string  `json:"gatewayIp,omitempty"`
	Id        string   `json:"id,omitempty"`
	IpPools   []IPPool `json:"ipPools,omitempty"`
}
