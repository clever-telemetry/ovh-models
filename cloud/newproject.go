package cloud

// GENERATED SDK for cloud API

// New cloud project
type NewProject struct {
	Agreements  *[]int64             `json:"agreements,omitempty"`
	Credit      *NewProjectCredit    `json:"credit,omitempty"`
	Description *string              `json:"description,omitempty"`
	OrderId     *int64               `json:"orderId,omitempty"`
	Project     *string              `json:"project,omitempty"`
	Status      NewProjectStatusEnum `json:"status,omitempty"`
}
