package me

// GENERATED SDK for me API

// All possible order status
type OrderStatusEnum string

var (
	OrderStatusEnumUNKNOWN            OrderStatusEnum = "unknown"
	OrderStatusEnumNOTPAID            OrderStatusEnum = "notPaid"
	OrderStatusEnumDOCUMENTSREQUESTED OrderStatusEnum = "documentsRequested"
	OrderStatusEnumDELIVERING         OrderStatusEnum = "delivering"
	OrderStatusEnumDELIVERED          OrderStatusEnum = "delivered"
	OrderStatusEnumCHECKING           OrderStatusEnum = "checking"
	OrderStatusEnumCANCELLING         OrderStatusEnum = "cancelling"
	OrderStatusEnumCANCELLED          OrderStatusEnum = "cancelled"
)
