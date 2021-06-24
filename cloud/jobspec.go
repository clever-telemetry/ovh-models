package cloud

// GENERATED SDK for cloud API

// AI Solutions Job Spec Object to create a job
type JobSpec struct {
	Command         *[]string             `json:"command,omitempty"`
	DefaultHttpPort *int64                `json:"defaultHttpPort,omitempty"`
	Env             *[]JobEnv             `json:"env,omitempty"`
	Image           string                `json:"image,omitempty"`
	Labels          *interface{}          `json:"labels,omitempty"`
	Name            string                `json:"name,omitempty"`
	ReadUser        *string               `json:"readUser,omitempty"`
	Region          string                `json:"region,omitempty"`
	Resources       Resources             `json:"resources,omitempty"`
	Shutdown        *ShutdownStrategyEnum `json:"shutdown,omitempty"`
	SshPublicKeys   *[]string             `json:"sshPublicKeys,omitempty"`
	Timeout         *int64                `json:"timeout,omitempty"`
	UnsecureHttp    *bool                 `json:"unsecureHttp,omitempty"`
	Volumes         *[]Volume             `json:"volumes,omitempty"`
}
