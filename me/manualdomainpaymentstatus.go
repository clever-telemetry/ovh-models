package me

// GENERATED SDK for me API

// Status of your manual domain payment migration
type ManualDomainPaymentStatus struct {
	DomainsToMigrate *int64                `json:"domainsToMigrate,omitempty"`
	MigratedDomains  *int64                `json:"migratedDomains,omitempty"`
	Status           BillingTaskStatusEnum `json:"status,omitempty"`
}
