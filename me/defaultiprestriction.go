package me

// GENERATED SDK for me API

// Default IP restriction of a VoIP line
type DefaultIpRestriction struct {
	Id     int64        `json:"id,omitempty"`
	Subnet string       `json:"subnet,omitempty"`
	Type   ProtocolEnum `json:"type,omitempty"`
}
