package cloud

// GENERATED SDK for cloud API

// UsageBill
type UsageBill struct {
	Bill_id      string          `json:"bill_id,omitempty"`
	Credit       float64         `json:"credit,omitempty"`
	Part         float64         `json:"part,omitempty"`
	Payment_type PaymentTypeEnum `json:"payment_type,omitempty"`
	Total        float64         `json:"total,omitempty"`
}
