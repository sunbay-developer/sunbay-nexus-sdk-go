package common

// PostAuthAmount represents post authorization amount information
// Supports: orderAmount, tipAmount, taxAmount, surchargeAmount
// Does NOT support: cashbackAmount
type PostAuthAmount struct {
	// OrderAmount is the order amount in cents (required)
	OrderAmount *int64 `json:"orderAmount"`

	// TipAmount is the tip amount in cents (optional)
	TipAmount *int64 `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount in cents (optional)
	TaxAmount *int64 `json:"taxAmount,omitempty"`

	// SurchargeAmount is the surcharge amount in cents (optional)
	SurchargeAmount *int64 `json:"surchargeAmount,omitempty"`

	// PriceCurrency is the price currency (ISO 4217, required)
	PriceCurrency string `json:"priceCurrency"`
}
