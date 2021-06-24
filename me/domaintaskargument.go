package me

// GENERATED SDK for me API

// Domain operation argument
type DomainTaskArgument struct {
	AcceptedFormats *[]DocumentFormatsEnum `json:"acceptedFormats,omitempty"`
	AcceptedValues  *[]string              `json:"acceptedValues,omitempty"`
	Description     *string                `json:"description,omitempty"`
	Fields          *[]ContactFieldEnum    `json:"fields,omitempty"`
	Key             string                 `json:"key,omitempty"`
	MaximumSize     *int64                 `json:"maximumSize,omitempty"`
	MinimumSize     *int64                 `json:"minimumSize,omitempty"`
	ReadOnly        bool                   `json:"readOnly,omitempty"`
	Template        *string                `json:"template,omitempty"`
	Type            string                 `json:"type,omitempty"`
	Value           *string                `json:"value,omitempty"`
}
