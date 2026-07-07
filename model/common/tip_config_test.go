package common

import (
	"encoding/json"
	"testing"
)

func TestTipConfigValidate(t *testing.T) {
	t.Run("allows up to three suggestions", func(t *testing.T) {
		cfg := &TipConfig{
			Suggestions: []TipSuggestions{
				{Names: []string{"No Tip"}, FeeMode: "AMOUNT", Values: []int{0}},
				{Names: []string{"15%"}, FeeMode: "RATE", Values: []int{15}},
				{Names: []string{"20%"}, FeeMode: "RATE", Values: []int{20}},
			},
		}
		if err := cfg.Validate(); err != nil {
			t.Fatalf("Validate() returned error: %v", err)
		}
	})

	t.Run("rejects more than three suggestions", func(t *testing.T) {
		cfg := &TipConfig{
			Suggestions: []TipSuggestions{
				{Names: []string{"A"}},
				{Names: []string{"B"}},
				{Names: []string{"C"}},
				{Names: []string{"D"}},
			},
		}
		if err := cfg.Validate(); err == nil {
			t.Fatal("Validate() expected error, got nil")
		}
	})

	t.Run("rejects mismatched names and values", func(t *testing.T) {
		cfg := &TipConfig{
			Suggestions: []TipSuggestions{
				{Names: []string{"A", "B"}, Values: []int{10}},
			},
		}
		if err := cfg.Validate(); err == nil {
			t.Fatal("Validate() expected error, got nil")
		}
	})

	t.Run("rejects more than three names or values", func(t *testing.T) {
		cfg := &TipConfig{
			Suggestions: []TipSuggestions{
				{
					Names:  []string{"A", "B", "C", "D"},
					Values: []int{1, 2, 3, 4},
				},
			},
		}
		if err := cfg.Validate(); err == nil {
			t.Fatal("Validate() expected error, got nil")
		}
	})
}

func TestTipConfigMarshalIncludesSuggestionNames(t *testing.T) {
	cfg := &TipConfig{
		OnScreenTip: true,
		TipMode:     "ON_SALE",
		TipWithTax:  true,
		Suggestions: []TipSuggestions{
			{Names: []string{"15%"}, FeeMode: "RATE", Values: []int{15}},
		},
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("json.Marshal() returned error: %v", err)
	}

	want := `{"onScreenTip":true,"tipMode":"ON_SALE","tipWithTax":true,"suggestions":[{"names":["15%"],"feeMode":"RATE","values":[15]}]}`
	if string(data) != want {
		t.Fatalf("unexpected JSON:\nwant %s\ngot  %s", want, string(data))
	}
}
