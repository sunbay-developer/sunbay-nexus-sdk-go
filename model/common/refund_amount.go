package common

// RefundAmount represents refund amount information
// Supports: orderAmount, tipAmount, taxAmount, surchargeAmount, cashbackAmount
type RefundAmount struct {
	// OrderAmount is the order amount (required)
	OrderAmount *float64 `json:"orderAmount"`

	// TipAmount is the tip amount (optional, must be greater than or equal to 0)
	TipAmount *float64 `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount (optional, must be greater than or equal to 0)
	TaxAmount *float64 `json:"taxAmount,omitempty"`

	// SurchargeAmount is the surcharge amount (optional, must be greater than or equal to 0)
	// Note: Some processors may require surcharge to be refunded proportionally. Please contact technical support for detailed policies.
	SurchargeAmount *float64 `json:"surchargeAmount,omitempty"`

	// CashbackAmount is the cashback amount (optional, must be greater than or equal to 0)
	CashbackAmount *float64 `json:"cashbackAmount,omitempty"`

	// PricingCurrency is the pricing currency (ISO 4217, required)
	PricingCurrency string `json:"pricingCurrency"`
}

