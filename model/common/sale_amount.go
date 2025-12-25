package common

// SaleAmount represents sale transaction amount information
// Supports: orderAmount, tipAmount, taxAmount, surchargeAmount, cashbackAmount
type SaleAmount struct {
	// OrderAmount is the order amount (required)
	OrderAmount *float64 `json:"orderAmount"`

	// TipAmount is the tip amount (optional)
	TipAmount *float64 `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount (optional)
	TaxAmount *float64 `json:"taxAmount,omitempty"`

	// SurchargeAmount is the surcharge amount (optional)
	SurchargeAmount *float64 `json:"surchargeAmount,omitempty"`

	// CashbackAmount is the cashback amount (optional)
	CashbackAmount *float64 `json:"cashbackAmount,omitempty"`

	// PricingCurrency is the pricing currency (ISO 4217, required)
	PricingCurrency string `json:"pricingCurrency"`
}

