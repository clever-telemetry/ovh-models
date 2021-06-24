package me

// GENERATED SDK for me API

// Hardware RAID defined in this partitioning scheme
type hardwareRaid struct {
	Disks []string                   `json:"disks,omitempty"`
	Mode  TemplateOsHardwareRaidEnum `json:"mode,omitempty"`
	Name  string                     `json:"name,omitempty"`
	Step  int64                      `json:"step,omitempty"`
}
