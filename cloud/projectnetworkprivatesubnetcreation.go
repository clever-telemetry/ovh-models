package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectNetworkPrivateSubnetCreation struct {
	Dhcp      bool   `json:"dhcp"`
	End       string `json:"end"`
	Network   string `json:"network"`
	NoGateway bool   `json:"noGateway"`
	Region    string `json:"region"`
	Start     string `json:"start"`
}
