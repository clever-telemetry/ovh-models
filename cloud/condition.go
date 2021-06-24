package cloud

// GENERATED SDK for cloud API

// A condition
type Condition struct {
	Key    *string   `json:"key,omitempty"`
	Match  MatchEnum `json:"match"`
	Name   string    `json:"name"`
	Negate *bool     `json:"negate,omitempty"`
	Type   TypeEnum  `json:"type"`
	Values []string  `json:"values"`
}
