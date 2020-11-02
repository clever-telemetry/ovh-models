package cloud

// GENERATED SDK for cloud API

// A serving engine namespace
type Namespace struct {
	ClusterId   string `json:"clusterId,omitempty"`
	Container   string `json:"container,omitempty"`
	ContainerId string `json:"containerId,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	Description string `json:"description,omitempty"`
	HubUrl      string `json:"hubUrl,omitempty"`
	Id          string `json:"id,omitempty"`
	Region      string `json:"region,omitempty"`
	Url         string `json:"url,omitempty"`
}
