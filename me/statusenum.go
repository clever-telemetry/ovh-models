package me

// GENERATED SDK for me API

// Tag status
type StatusEnum string

var (
	StatusEnumWAIT_REVOCATION StatusEnum = "WAIT_REVOCATION"
	StatusEnumWAIT_MODERATION StatusEnum = "WAIT_MODERATION"
	StatusEnumREVOCATING      StatusEnum = "REVOCATING"
	StatusEnumREFUSING        StatusEnum = "REFUSING"
	StatusEnumREFUSED         StatusEnum = "REFUSED"
	StatusEnumMODERATING      StatusEnum = "MODERATING"
	StatusEnumERROR           StatusEnum = "ERROR"
	StatusEnumDELETING        StatusEnum = "DELETING"
	StatusEnumDELETED         StatusEnum = "DELETED"
	StatusEnumCREATING        StatusEnum = "CREATING"
	StatusEnumCREATED         StatusEnum = "CREATED"
)
