package cloud

// GENERATED SDK for cloud API

// HourlyStorage
type HourlyStorage struct {
	IncomingBandwidth *BandwidthStorage `json:"incomingBandwidth,omitempty"`
	OutgoingBandwidth *BandwidthStorage `json:"outgoingBandwidth,omitempty"`
	Region            string            `json:"region,omitempty"`
	Stored            *StoredStorage    `json:"stored,omitempty"`
	TotalPrice        float64           `json:"totalPrice,omitempty"`
	Type              StorageTypeEnum   `json:"type,omitempty"`
}
