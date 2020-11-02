package cloud

// GENERATED SDK for cloud API

// Usage information for current month on your project
type CurrentUsage struct {
	Instances       InstancesUsage `json:"instances,omitempty"`
	Snapshots       SnapshotsUsage `json:"snapshots,omitempty"`
	Storage         StorageUsage   `json:"storage,omitempty"`
	Total           Price          `json:"total,omitempty"`
	VolumeSnapshots SnapshotsUsage `json:"volumeSnapshots,omitempty"`
	Volumes         VolumesUsage   `json:"volumes,omitempty"`
}
