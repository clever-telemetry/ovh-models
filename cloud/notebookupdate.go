package cloud

// GENERATED SDK for cloud API

// AI Solutions Notebook Spec Object to create a notebook
type NotebookUpdate struct {
	Labels       *interface{}    `json:"labels,omitempty"`
	Resources    *ResourcesInput `json:"resources,omitempty"`
	UnsecureHttp *bool           `json:"unsecureHttp,omitempty"`
	Volumes      *[]Volume       `json:"volumes,omitempty"`
}
