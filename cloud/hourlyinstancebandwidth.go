package cloud

// GENERATED SDK for cloud API

// HourlyInstanceBandwidth
type HourlyInstanceBandwidth struct {
	IncomingBandwidth *BandwidthInstance `json:"incomingBandwidth,omitempty"`
	OutgoingBandwidth *BandwidthInstance `json:"outgoingBandwidth,omitempty"`
	Region            string             `json:"region,omitempty"`
	TotalPrice        float64            `json:"totalPrice,omitempty"`
}
