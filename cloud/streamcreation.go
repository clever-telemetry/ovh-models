package cloud

// GENERATED SDK for cloud API

// Create a stream of data
type StreamCreation struct {
	Description string         `json:"description"`
	Kind        StreamKindEnum `json:"kind"`
	Name        string         `json:"name"`
	Region      string         `json:"region"`
}
