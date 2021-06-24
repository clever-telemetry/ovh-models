package me

// GENERATED SDK for me API

// List of documents added on your account
type Document struct {
	CreationDate   string        `json:"creationDate,omitempty"`
	ExpirationDate *string       `json:"expirationDate,omitempty"`
	GetUrl         string        `json:"getUrl,omitempty"`
	Id             string        `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	PutUrl         string        `json:"putUrl,omitempty"`
	Size           int64         `json:"size,omitempty"`
	Tags           []interface{} `json:"tags,omitempty"`
	ValidationDate *string       `json:"validationDate,omitempty"`
}
