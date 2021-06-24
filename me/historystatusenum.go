package me

// GENERATED SDK for me API

// History label of order follow-up
type HistoryStatusEnum string

var (
	HistoryStatusEnumREGISTERED_PAYMENT_INITIATED HistoryStatusEnum = "REGISTERED_PAYMENT_INITIATED"
	HistoryStatusEnumPAYMENT_RECEIVED             HistoryStatusEnum = "PAYMENT_RECEIVED"
	HistoryStatusEnumPAYMENT_INITIATED            HistoryStatusEnum = "PAYMENT_INITIATED"
	HistoryStatusEnumPAYMENT_CONFIRMED            HistoryStatusEnum = "PAYMENT_CONFIRMED"
	HistoryStatusEnumORDER_STARTED                HistoryStatusEnum = "ORDER_STARTED"
	HistoryStatusEnumORDER_ACCEPTED               HistoryStatusEnum = "ORDER_ACCEPTED"
	HistoryStatusEnumINVOICE_SENT                 HistoryStatusEnum = "INVOICE_SENT"
	HistoryStatusEnumINVOICE_IN_PROGRESS          HistoryStatusEnum = "INVOICE_IN_PROGRESS"
	HistoryStatusEnumFRAUD_REFUSED                HistoryStatusEnum = "FRAUD_REFUSED"
	HistoryStatusEnumFRAUD_MANUAL_REVIEW          HistoryStatusEnum = "FRAUD_MANUAL_REVIEW"
	HistoryStatusEnumFRAUD_DOCS_REQUESTED         HistoryStatusEnum = "FRAUD_DOCS_REQUESTED"
	HistoryStatusEnumFRAUD_CHECK                  HistoryStatusEnum = "FRAUD_CHECK"
	HistoryStatusEnumDELIVERY                     HistoryStatusEnum = "DELIVERY"
)
