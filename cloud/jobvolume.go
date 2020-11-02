package cloud

// GENERATED SDK for cloud API

// Training Platform Job Volume Object
type JobVolume struct {
	Container  string                  `json:"container"`
	MountPath  string                  `json:"mountPath"`
	Permission JobVolumePermissionEnum `json:"permission"`
	Region     string                  `json:"region"`
}
