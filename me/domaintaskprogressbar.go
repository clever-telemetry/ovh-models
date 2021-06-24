package me

// GENERATED SDK for me API

// Domain operation progress
type DomainTaskProgressBar struct {
	CurrentStep      OperationStep         `json:"currentStep,omitempty"`
	ExpectedDoneDate *string               `json:"expectedDoneDate,omitempty"`
	FollowUpSteps    *[]OperationStep      `json:"followUpSteps,omitempty"`
	LastUpdateDate   *string               `json:"lastUpdateDate,omitempty"`
	Progress         int64                 `json:"progress,omitempty"`
	TaskActions      []OperationActionEnum `json:"taskActions,omitempty"`
	TaskStatus       OperationStatusEnum   `json:"taskStatus,omitempty"`
}
