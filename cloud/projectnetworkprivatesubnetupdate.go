package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectNetworkPrivateSubnetUpdate struct {
	Dhcp           bool    `json:"dhcp"`
	DisableGateway bool    `json:"disableGateway"`
	GatewayIp      *string `json:"gatewayIp,omitempty"`
}
