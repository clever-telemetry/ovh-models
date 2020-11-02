package cloud

// GENERATED SDK for cloud API

// Create a consumer on a stream
type Subscription struct {
	Id   string               `json:"id,omitempty"`
	Kind SubscriptionKindEnum `json:"kind,omitempty"`
	Name string               `json:"name,omitempty"`
}
