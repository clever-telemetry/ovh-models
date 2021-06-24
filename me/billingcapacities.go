package me

// GENERATED SDK for me API

// Internal customer billing capacities for customer control panel
type BillingCapacities struct {
	CanUseDebtSystem            bool                      `json:"canUseDebtSystem,omitempty"`
	CanUsePostalMailForInvoices bool                      `json:"canUsePostalMailForInvoices,omitempty"`
	RequiredPaymentMethod       RequiredPaymentMethodEnum `json:"requiredPaymentMethod,omitempty"`
}
