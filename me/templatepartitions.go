package me

// GENERATED SDK for me API

//  Partitions defined in this partitioning scheme
type templatePartitions struct {
	Filesystem TemplateOsFileSystemEnum  `json:"filesystem,omitempty"`
	Mountpoint string                    `json:"mountpoint,omitempty"`
	Order      int64                     `json:"order,omitempty"`
	Raid       *PartitionRaidEnum        `json:"raid,omitempty"`
	Size       interface{}               `json:"size,omitempty"`
	Type       TemplatePartitionTypeEnum `json:"type,omitempty"`
	VolumeName *string                   `json:"volumeName,omitempty"`
}
