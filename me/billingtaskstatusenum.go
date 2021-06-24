package me

// GENERATED SDK for me API

// billing task status
type BillingTaskStatusEnum string

var (
	BillingTaskStatusEnumTODO          BillingTaskStatusEnum = "todo"
	BillingTaskStatusEnumOVHERROR      BillingTaskStatusEnum = "ovhError"
	BillingTaskStatusEnumINIT          BillingTaskStatusEnum = "init"
	BillingTaskStatusEnumDONE          BillingTaskStatusEnum = "done"
	BillingTaskStatusEnumDOING         BillingTaskStatusEnum = "doing"
	BillingTaskStatusEnumCUSTOMERERROR BillingTaskStatusEnum = "customerError"
	BillingTaskStatusEnumCANCELLED     BillingTaskStatusEnum = "cancelled"
)
