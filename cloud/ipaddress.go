package cloud

// GENERATED SDK for cloud API

// IpAddress
type IpAddress struct {
	GatewayIp *string `json:"gatewayIp,omitempty"`
	Ip        string  `json:"ip,omitempty"`
	NetworkId string  `json:"networkId,omitempty"`
	Type      string  `json:"type,omitempty"`
	Version   int64   `json:"version,omitempty"`
}
