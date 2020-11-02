package cloud

// GENERATED SDK for cloud API

// Interface
type Interface struct {
	FixedIps   []FixedIp `json:"fixedIps,omitempty"`
	Id         string    `json:"id,omitempty"`
	MacAddress string    `json:"macAddress,omitempty"`
	NetworkId  string    `json:"networkId,omitempty"`
	State      string    `json:"state,omitempty"`
	Type       string    `json:"type,omitempty"`
}
