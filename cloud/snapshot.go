package cloud

// GENERATED SDK for cloud API

// Snapshot
type Snapshot struct {
	CreationDate string             `json:"creationDate,omitempty"`
	Description  string             `json:"description,omitempty"`
	Id           string             `json:"id,omitempty"`
	Name         string             `json:"name,omitempty"`
	PlanCode     *string            `json:"planCode,omitempty"`
	Region       string             `json:"region,omitempty"`
	Size         int64              `json:"size,omitempty"`
	Status       SnapshotStatusEnum `json:"status,omitempty"`
	VolumeId     string             `json:"volumeId,omitempty"`
}
