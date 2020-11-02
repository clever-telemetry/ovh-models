package cloud

// GENERATED SDK for cloud API

// Order
type Order struct {
	Date        string     `json:"date,omitempty"`
	OrderId     int64      `json:"orderId,omitempty"`
	PlanCode    string     `json:"planCode,omitempty"`
	ServiceName *string    `json:"serviceName,omitempty"`
	Status      StatusEnum `json:"status,omitempty"`
}
