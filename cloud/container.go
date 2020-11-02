package cloud

// GENERATED SDK for cloud API

// Container
type Container struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Region        string `json:"region,omitempty"`
	StoredBytes   int64  `json:"storedBytes,omitempty"`
	StoredObjects int64  `json:"storedObjects,omitempty"`
}
