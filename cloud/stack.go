package cloud

// GENERATED SDK for cloud API

// Stack
type Stack struct {
	Commit        string             `json:"commit,omitempty"`
	Description   string             `json:"description,omitempty"`
	GitRepository string             `json:"gitRepository,omitempty"`
	Instructions  []InstructionGuide `json:"instructions,omitempty"`
	Name          string             `json:"name,omitempty"`
	Release       string             `json:"release,omitempty"`
	Uuid          string             `json:"uuid,omitempty"`
}
