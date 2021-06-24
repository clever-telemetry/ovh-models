package me

// GENERATED SDK for me API

// Available installation templates
type Templates struct {
	AvailableLanguages         []TemplateOsLanguageEnum   `json:"availableLanguages,omitempty"`
	BitFormat                  BitFormatEnum              `json:"bitFormat,omitempty"`
	Category                   TemplateOsUsageEnum        `json:"category,omitempty"`
	Customization              *TemplateOsProperties      `json:"customization,omitempty"`
	DefaultLanguage            TemplateOsLanguageEnum     `json:"defaultLanguage,omitempty"`
	Description                string                     `json:"description,omitempty"`
	Distribution               string                     `json:"distribution,omitempty"`
	Family                     TemplateOsTypeEnum         `json:"family,omitempty"`
	Filesystems                []TemplateOsFileSystemEnum `json:"filesystems,omitempty"`
	HardRaidConfiguration      *bool                      `json:"hardRaidConfiguration,omitempty"`
	LastModification           *string                    `json:"lastModification,omitempty"`
	License                    *TemplateOsInfoLicense     `json:"license,omitempty"`
	LvmReady                   *bool                      `json:"lvmReady,omitempty"`
	NoPartitioning             bool                       `json:"noPartitioning,omitempty"`
	Project                    *TemplateOsInfoProject     `json:"project,omitempty"`
	Subfamily                  TemplateOsSubfamilyEnum    `json:"subfamily,omitempty"`
	SupportsDistributionKernel *bool                      `json:"supportsDistributionKernel,omitempty"`
	SupportsGptLabel           *bool                      `json:"supportsGptLabel,omitempty"`
	SupportsRTM                bool                       `json:"supportsRTM,omitempty"`
	SupportsSqlServer          *bool                      `json:"supportsSqlServer,omitempty"`
	SupportsUEFI               *SupportsUEFIEnum          `json:"supportsUEFI,omitempty"`
	TemplateName               string                     `json:"templateName,omitempty"`
}
