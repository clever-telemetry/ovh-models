package me

// GENERATED SDK for me API

// List of all IP Restrictions
type IpRestriction struct {
	Id      int64                 `json:"id,omitempty"`
	Ip      string                `json:"ip,omitempty"`
	Rule    IpRestrictionRuleEnum `json:"rule,omitempty"`
	Warning bool                  `json:"warning,omitempty"`
}
