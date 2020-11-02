package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectVolumeCreation struct {
	Description *string        `json:"description,omitempty"`
	ImageId     *string        `json:"imageId,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Region      string         `json:"region"`
	Size        int64          `json:"size"`
	SnapshotId  *string        `json:"snapshotId,omitempty"`
	Type        VolumeTypeEnum `json:"type"`
}
