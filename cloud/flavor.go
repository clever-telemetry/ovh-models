package cloud

// GENERATED SDK for cloud API

// Cloud Database flavor definition
type Flavor struct {
	Core    int64  `json:"core,omitempty"`
	Memory  int64  `json:"memory,omitempty"`
	Name    string `json:"name,omitempty"`
	Storage int64  `json:"storage,omitempty"`
}
