package cloud

// GENERATED SDK for cloud API

// AI Solutions Notebook Spec Object to create a notebook
type NotebookSpecInput struct {
	Env          NotebookEnv           `json:"env"`
	Labels       *interface{}          `json:"labels,omitempty"`
	Name         string                `json:"name,omitempty"`
	Region       string                `json:"region"`
	Resources    ResourcesInput        `json:"resources"`
	Shutdown     *ShutdownStrategyEnum `json:"shutdown,omitempty"`
	UnsecureHttp *bool                 `json:"unsecureHttp,omitempty"`
	Volumes      *[]Volume             `json:"volumes,omitempty"`
}
