package cloud

// GENERATED SDK for cloud API

// A entrypoint rule
type Rule struct {
	Action     string    `json:"action"`
	Conditions *[]string `json:"conditions,omitempty"`
}
