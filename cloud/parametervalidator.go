package cloud

// GENERATED SDK for cloud API

// Conditions to which the value of parameter must conform
type ParameterValidator struct {
	Max   *float64 `json:"max,omitempty"`
	Min   *float64 `json:"min,omitempty"`
	Regex *string  `json:"regex,omitempty"`
}
