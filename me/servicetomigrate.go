package me

// GENERATED SDK for me API

// Original service that can be migrated
type ServiceToMigrate struct {
	Description string        `json:"description,omitempty"`
	Metadata    []interface{} `json:"metadata,omitempty"`
	Route       *string       `json:"route,omitempty"`
	ServiceId   int64         `json:"serviceId,omitempty"`
	ServiceName string        `json:"serviceName,omitempty"`
}
