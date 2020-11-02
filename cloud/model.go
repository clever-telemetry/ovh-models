package cloud

// GENERATED SDK for cloud API

// A deployed machine learning model
type Model struct {
	ApiStatus                  APIStatusEnum                  `json:"apiStatus,omitempty"`
	CreatedAt                  string                         `json:"createdAt,omitempty"`
	Id                         string                         `json:"id,omitempty"`
	Replicas                   *int64                         `json:"replicas,omitempty"`
	Url                        *string                        `json:"url,omitempty"`
	Version                    *int64                         `json:"version,omitempty"`
	VersionStatus              VersionStatusEnum              `json:"versionStatus,omitempty"`
	WorkflowTemplate           *WorkflowTemplateEnum          `json:"workflowTemplate,omitempty"`
	WorkflowTemplateParameters ModelWorkflowTemplateParameter `json:"workflowTemplateParameters,omitempty"`
}
