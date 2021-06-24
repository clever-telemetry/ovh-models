package me

// GENERATED SDK for me API

// Task running an email change on an account
type Task struct {
	DateDone    *string       `json:"dateDone,omitempty"`
	DateRequest string        `json:"dateRequest,omitempty"`
	Id          int64         `json:"id,omitempty"`
	NewEmail    string        `json:"newEmail,omitempty"`
	State       TaskStateEnum `json:"state,omitempty"`
}
