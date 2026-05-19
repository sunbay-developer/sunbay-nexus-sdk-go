package common

// TipSuggestions represents tip suggestion configuration
type TipSuggestions struct {
	// FeeMode is the fee mode for tip suggestions. Possible values: RATE, AMOUNT
	FeeMode string `json:"feeMode"`

	// Values is the list of suggested tip values (percentages for RATE mode, amounts in cents for AMOUNT mode)
	Values []int `json:"values"`
}

// TipConfig represents tip configuration for a transaction
type TipConfig struct {
	// OnScreenTip indicates whether to show tip on screen
	OnScreenTip bool `json:"onScreenTip"`

	// TipMode is the tip mode. Possible values: ON_SALE, AFTER_SALE
	TipMode string `json:"tipMode"`

	// TipWithTax indicates whether tip should be calculated with tax included
	TipWithTax bool `json:"tipWithTax"`

	// Suggestions is the tip suggestion configuration
	Suggestions *TipSuggestions `json:"suggestions,omitempty"`
}
