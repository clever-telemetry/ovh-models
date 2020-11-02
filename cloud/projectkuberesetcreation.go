package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectKubeResetCreation struct {
	Version           *VersionEnum                `json:"version,omitempty"`
	WorkerNodesPolicy *ResetWorkerNodesPolicyEnum `json:"workerNodesPolicy,omitempty"`
}
