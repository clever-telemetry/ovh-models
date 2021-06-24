package me

// GENERATED SDK for me API

// Missing description
type Balance struct {
	Amount         Price              `json:"amount,omitempty"`
	BalanceDetails []BalanceDetails   `json:"balanceDetails,omitempty"`
	BalanceName    string             `json:"balanceName,omitempty"`
	Booked         []BookedMovement   `json:"booked,omitempty"`
	CreationDate   string             `json:"creationDate,omitempty"`
	Expiring       []ExpiringMovement `json:"expiring,omitempty"`
	LastUpdate     string             `json:"lastUpdate,omitempty"`
	Type           TypeEnum           `json:"type,omitempty"`
}
