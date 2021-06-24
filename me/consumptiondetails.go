package me

// GENERATED SDK for me API

// Detailed consumption's retrieval information
type ConsumptionDetails struct {
	FileFormat *ConsumptionExportFormatsEnum   `json:"fileFormat,omitempty"`
	FileURL    *string                         `json:"fileURL,omitempty"`
	Message    *string                         `json:"message,omitempty"`
	TaskId     int64                           `json:"taskId,omitempty"`
	TaskStatus ConsumptionExportTaskStatusEnum `json:"taskStatus,omitempty"`
}
