package cloud

// GENERATED SDK for cloud API

// List of required agreements to activate the lab
type LabAgreements struct {
	Accepted []int64 `json:"accepted,omitempty"`
	ToAccept []int64 `json:"toAccept,omitempty"`
}
