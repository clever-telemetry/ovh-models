package cloud

// GENERATED SDK for cloud API

// Public Cloud products availability
type Availability struct {
	Plans    []AvailabilityPlan    `json:"plans,omitempty"`
	Products []AvailabilityProduct `json:"products,omitempty"`
}
