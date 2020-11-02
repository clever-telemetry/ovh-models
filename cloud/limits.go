package cloud

// GENERATED SDK for cloud API

// Limitation of a docker registry
type Limits struct {
	ImageStorage    int64 `json:"imageStorage,omitempty"`
	ParallelRequest int64 `json:"parallelRequest,omitempty"`
}
