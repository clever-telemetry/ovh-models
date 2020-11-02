package cloud

// GENERATED SDK for cloud API

// InstructionGuide
type InstructionGuide struct {
	Content  []Content `json:"content,omitempty"`
	Language string    `json:"language,omitempty"`
	Sections []Section `json:"sections,omitempty"`
	Title    string    `json:"title,omitempty"`
}
