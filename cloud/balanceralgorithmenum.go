package cloud

// GENERATED SDK for cloud API

// Available load balancer backend balancer algorithm
type BalancerAlgorithmEnum string

var (
	BalancerAlgorithmEnumROUNDROBIN BalancerAlgorithmEnum = "roundrobin"
	BalancerAlgorithmEnumSTATIC_RR  BalancerAlgorithmEnum = "static-rr"
	BalancerAlgorithmEnumLEASTCONN  BalancerAlgorithmEnum = "leastconn"
	BalancerAlgorithmEnumFIRST      BalancerAlgorithmEnum = "first"
	BalancerAlgorithmEnumSOURCE     BalancerAlgorithmEnum = "source"
)
