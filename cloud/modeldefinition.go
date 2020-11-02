package cloud

// GENERATED SDK for cloud API

// Missing description
type ModelDefinition struct {
	AutoscalingSpec  *AutoscalingSpec      `json:"autoscalingSpec,omitempty"`
	Backend          *BackendIdEnum        `json:"backend,omitempty"`
	Flavor           string                `json:"flavor"`
	Framework        *FrameworkIdEnum      `json:"framework,omitempty"`
	Id               string                `json:"id"`
	ImageId          *string               `json:"imageId,omitempty"`
	StoragePath      *string               `json:"storagePath,omitempty"`
	WorkflowTemplate *WorkflowTemplateEnum `json:"workflowTemplate,omitempty"`
}
