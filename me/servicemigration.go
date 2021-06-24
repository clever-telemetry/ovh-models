package me

// GENERATED SDK for me API

// Service you may migrate to a given offer
type ServiceMigration struct {
	Addons           []ServiceMigration `json:"addons,omitempty"`
	OrderId          *int64             `json:"orderId,omitempty"`
	ProposedOffer    ProposedOffer      `json:"proposedOffer,omitempty"`
	ServiceToMigrate *ServiceToMigrate  `json:"serviceToMigrate,omitempty"`
}
