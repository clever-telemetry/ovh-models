package cloud

// GENERATED SDK for cloud API

// Training Platform Job Spec Object to create a job
type JobSpec struct {
	Command         *[]string    `json:"command,omitempty"`
	DefaultHttpPort *int64       `json:"defaultHttpPort,omitempty"`
	Env             *[]JobEnv    `json:"env,omitempty"`
	Image           string       `json:"image"`
	Name            string       `json:"name,omitempty"`
	Region          string       `json:"region"`
	Resources       JobResource  `json:"resources"`
	Shutdown        *string      `json:"shutdown,omitempty"`
	Timeout         *int64       `json:"timeout,omitempty"`
	Volumes         *[]JobVolume `json:"volumes,omitempty"`
}
