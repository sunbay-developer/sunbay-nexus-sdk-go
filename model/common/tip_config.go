package common

import "fmt"

// TipSuggestions represents a tip suggestion option.
type TipSuggestions struct {
	// Name is the display name for the tip option.
	Name string `json:"name,omitempty"`

	// FeeMode is the fee mode for tip suggestions. Possible values: RATE, AMOUNT
	FeeMode string `json:"feeMode"`

	// Values is the list of suggested tip values (percentages for RATE mode, amounts in cents for AMOUNT mode)
	Values []int `json:"values"`
}

// TipConfig represents tip configuration for a transaction.
type TipConfig struct {
	// OnScreenTip indicates whether to show tip on screen
	OnScreenTip bool `json:"onScreenTip"`

	// TipMode is the tip mode. Possible values: ON_SALE, AFTER_SALE
	TipMode string `json:"tipMode"`

	// TipWithTax indicates whether tip should be calculated with tax included
	TipWithTax bool `json:"tipWithTax"`

	// Suggestions is the list of tip suggestion options.
	// Supports at most 3 items.
	Suggestions []TipSuggestions `json:"suggestions,omitempty"`
}

// Validate checks whether the tip configuration is valid.
func (c *TipConfig) Validate() error {
	if c == nil {
		return nil
	}
	if len(c.Suggestions) > 3 {
		return fmt.Errorf("suggestions supports at most 3 items")
	}
	return nil
}
