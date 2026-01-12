package common

// AuthAmount represents authorization amount information
// Supports: orderAmount, priceCurrency only
// Used for: Auth, ForcedAuth, IncrementalAuth
type AuthAmount struct {
	// OrderAmount is the order amount in cents (required)
	OrderAmount *int64 `json:"orderAmount"`

	// PriceCurrency is the price currency (ISO 4217, required)
	PriceCurrency string `json:"priceCurrency"`
}
