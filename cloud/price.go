package cloud

// GENERATED SDK for cloud API

// Price with it's currency and textual representation
type Price struct {
	CurrencyCode CurrencyCodeEnum `json:"currencyCode,omitempty"`
	Text         string           `json:"text,omitempty"`
	Value        float64          `json:"value,omitempty"`
}
