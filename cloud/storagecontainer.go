package cloud

// GENERATED SDK for cloud API

// Container
type StorageContainer struct {
	Name         string          `json:"name,omitempty"`
	Objects      []StorageObject `json:"objects,omitempty"`
	ObjectsCount int64           `json:"objectsCount,omitempty"`
	ObjectsSize  int64           `json:"objectsSize,omitempty"`
}
