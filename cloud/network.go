package cloud

// GENERATED SDK for cloud API

// Network
type Network struct {
	Id      string            `json:"id,omitempty"`
	Name    string            `json:"name,omitempty"`
	Regions []NetworkRegion   `json:"regions,omitempty"`
	Status  NetworkStatusEnum `json:"status,omitempty"`
	Type    *NetworkTypeEnum  `json:"type,omitempty"`
	VlanId  int64             `json:"vlanId,omitempty"`
}
