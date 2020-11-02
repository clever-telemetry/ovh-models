package cloud

// GENERATED SDK for cloud API

// Framework of the model
type Framework struct {
	Backends           []BackendIdEnum `json:"backends,omitempty"`
	DocPage            string          `json:"docPage,omitempty"`
	Id                 FrameworkIdEnum `json:"id,omitempty"`
	Link               string          `json:"link,omitempty"`
	Logo               string          `json:"logo,omitempty"`
	Name               string          `json:"name,omitempty"`
	RecommendedBackend BackendIdEnum   `json:"recommendedBackend,omitempty"`
	Version            string          `json:"version,omitempty"`
}
