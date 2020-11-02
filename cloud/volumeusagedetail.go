package cloud

// GENERATED SDK for cloud API

// Volume usage
type VolumeUsageDetail struct {
	Price          Price       `json:"price,omitempty"`
	VolumeCapacity interface{} `json:"volumeCapacity,omitempty"`
	VolumeId       string      `json:"volumeId,omitempty"`
	VolumeType     VolumeType  `json:"volumeType,omitempty"`
}
