package common

import "fmt"

// TipSuggestions represents a tip suggestion option.
type TipSuggestions struct {
	// Names are the display names for the tip options.
	Names []string `json:"names,omitempty"`

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
	// Supports at most 3 items; each suggestion's names and values lists must
	// also be at most 3 items and must have the same length.
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
	for i, suggestion := range c.Suggestions {
		if len(suggestion.Names) > 3 {
			return fmt.Errorf("suggestions[%d].names supports at most 3 items", i)
		}
		if len(suggestion.Values) > 3 {
			return fmt.Errorf("suggestions[%d].values supports at most 3 items", i)
		}
		if len(suggestion.Names) != len(suggestion.Values) {
			return fmt.Errorf("suggestions[%d].names and suggestions[%d].values must have the same length", i, i)
		}
	}
	return nil
}
