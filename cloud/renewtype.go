package cloud

// GENERATED SDK for cloud API

// Map a possible renew for a specific service
type RenewType struct {
	Automatic          bool   `json:"automatic,omitempty"`
	DeleteAtExpiration bool   `json:"deleteAtExpiration,omitempty"`
	Forced             bool   `json:"forced,omitempty"`
	ManualPayment      *bool  `json:"manualPayment,omitempty"`
	Period             *int64 `json:"period,omitempty"`
}
