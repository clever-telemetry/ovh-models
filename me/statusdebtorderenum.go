package me

// GENERATED SDK for me API

// All status a debt HistoryOrder entry can be in
type StatusDebtOrderEnum string

var (
	StatusDebtOrderEnumUNPAID     StatusDebtOrderEnum = "UNPAID"
	StatusDebtOrderEnumUNMATURED  StatusDebtOrderEnum = "UNMATURED"
	StatusDebtOrderEnumTO_BE_PAID StatusDebtOrderEnum = "TO_BE_PAID"
	StatusDebtOrderEnumREFUNDED   StatusDebtOrderEnum = "REFUNDED"
	StatusDebtOrderEnumPAID       StatusDebtOrderEnum = "PAID"
)
