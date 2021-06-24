package me

// GENERATED SDK for me API

// API Log
type Log struct {
	Account string     `json:"account,omitempty"`
	Date    string     `json:"date,omitempty"`
	Ip      *string    `json:"ip,omitempty"`
	LogId   int64      `json:"logId,omitempty"`
	Method  MethodEnum `json:"method,omitempty"`
	Path    string     `json:"path,omitempty"`
	Route   string     `json:"route,omitempty"`
}
