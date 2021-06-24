package cloud

// GENERATED SDK for cloud API

// AI Solutions Notebook Spec Object to create a notebook
type NotebookSpec struct {
	Env          NotebookEnv           `json:"env,omitempty"`
	Flavor       *string               `json:"flavor,omitempty"`
	Labels       *interface{}          `json:"labels,omitempty"`
	Name         string                `json:"name,omitempty"`
	Region       string                `json:"region,omitempty"`
	Resources    Resources             `json:"resources,omitempty"`
	Shutdown     *ShutdownStrategyEnum `json:"shutdown,omitempty"`
	UnsecureHttp *bool                 `json:"unsecureHttp,omitempty"`
	Volumes      *[]Volume             `json:"volumes,omitempty"`
}
