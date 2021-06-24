package cloud

// GENERATED SDK for cloud API

// AI Solutions Progress Object
type Progress struct {
	Completed        int64                     `json:"completed,omitempty"`
	CreatedAt        string                    `json:"createdAt,omitempty"`
	Direction        DataSyncEnum              `json:"direction,omitempty"`
	Eta              int64                     `json:"eta,omitempty"`
	Failed           int64                     `json:"failed,omitempty"`
	Id               string                    `json:"id,omitempty"`
	Info             string                    `json:"info,omitempty"`
	Processed        int64                     `json:"processed,omitempty"`
	Skipped          int64                     `json:"skipped,omitempty"`
	State            DataSyncProgressStateEnum `json:"state,omitempty"`
	Total            int64                     `json:"total,omitempty"`
	TransferredBytes int64                     `json:"transferredBytes,omitempty"`
	UpdatedAt        string                    `json:"updatedAt,omitempty"`
	VolumeIndex      int64                     `json:"volumeIndex,omitempty"`
}
