package me

// GENERATED SDK for me API

// Missing description
type Movement struct {
	Amount            Price      `json:"amount,omitempty"`
	BalanceName       string     `json:"balanceName,omitempty"`
	CreationDate      string     `json:"creationDate,omitempty"`
	DestinationObject *SubObject `json:"destinationObject,omitempty"`
	ExpirationDate    *string    `json:"expirationDate,omitempty"`
	LastUpdate        string     `json:"lastUpdate,omitempty"`
	MovementId        int64      `json:"movementId,omitempty"`
	OrderId           *int64     `json:"orderId,omitempty"`
	SourceObject      SubObject  `json:"sourceObject,omitempty"`
	Type              TypeEnum   `json:"type,omitempty"`
}
