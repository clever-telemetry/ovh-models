package cloud

// GENERATED SDK for cloud API

// HTTP load balancer reject action
type ActionReject struct {
	Name       string               `json:"name"`
	StatusCode RejectStatusCodeEnum `json:"statusCode"`
}
