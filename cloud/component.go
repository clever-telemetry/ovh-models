package cloud

// GENERATED SDK for cloud API

// Information about the different components available in the region
type Component struct {
	Endpoint string            `json:"endpoint,omitempty"`
	Name     string            `json:"name,omitempty"`
	Status   ServiceStatusEnum `json:"status,omitempty"`
}
