package cloud

// GENERATED SDK for cloud API

// NetworkRegion
type NetworkRegion struct {
	OpenstackId *string                 `json:"openstackId,omitempty"`
	Region      string                  `json:"region,omitempty"`
	Status      NetworkRegionStatusEnum `json:"status,omitempty"`
}
