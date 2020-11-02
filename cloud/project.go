package cloud

// GENERATED SDK for cloud API

// Project
type Project struct {
	Access       AccessTypeEnum    `json:"access,omitempty"`
	CreationDate string            `json:"creationDate,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Expiration   *string           `json:"expiration,omitempty"`
	ManualQuota  bool              `json:"manualQuota,omitempty"`
	OrderId      *int64            `json:"orderId,omitempty"`
	PlanCode     string            `json:"planCode,omitempty"`
	ProjectName  *string           `json:"projectName,omitempty"`
	Project_id   string            `json:"project_id,omitempty"`
	Status       ProjectStatusEnum `json:"status,omitempty"`
	Unleash      bool              `json:"unleash,omitempty"`
}
