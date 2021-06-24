package me

// GENERATED SDK for me API

// SMS notifications
type NichandleSmsNotification struct {
	Abuse        bool                   `json:"abuse,omitempty"`
	CreationDate string                 `json:"creationDate,omitempty"`
	PhoneNumber  string                 `json:"phoneNumber,omitempty"`
	Status       NotificationStatusEnum `json:"status,omitempty"`
	UpdateDate   *string                `json:"updateDate,omitempty"`
}
