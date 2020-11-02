package cloud

// GENERATED SDK for cloud API

// ContainerDetail
type ContainerDetail struct {
	Archive       bool              `json:"archive,omitempty"`
	ContainerType TypeEnum          `json:"containerType,omitempty"`
	Cors          []string          `json:"cors,omitempty"`
	Name          string            `json:"name,omitempty"`
	Objects       []ContainerObject `json:"objects,omitempty"`
	Public        bool              `json:"public,omitempty"`
	Region        string            `json:"region,omitempty"`
	StaticUrl     string            `json:"staticUrl,omitempty"`
	StoredBytes   int64             `json:"storedBytes,omitempty"`
	StoredObjects int64             `json:"storedObjects,omitempty"`
}
