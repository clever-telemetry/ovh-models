package cloud

// GENERATED SDK for cloud API

// Parameters of the Workflow that build
type ModelWorkflowTemplateParameter struct {
	Backend     *BackendIdEnum   `json:"backend,omitempty"`
	Framework   *FrameworkIdEnum `json:"framework,omitempty"`
	ImageId     *string          `json:"imageId,omitempty"`
	StoragePath *string          `json:"storagePath,omitempty"`
}
