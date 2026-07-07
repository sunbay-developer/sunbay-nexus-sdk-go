package common

import (
	"encoding/json"
	"testing"
)

func TestTipConfigValidate(t *testing.T) {
	t.Run("allows up to three suggestions", func(t *testing.T) {
		cfg := &TipConfig{
			Suggestions: []TipSuggestions{
				{Name: "No Tip", FeeMode: "AMOUNT", Values: []int{0}},
				{Name: "15%", FeeMode: "RATE", Values: []int{15}},
				{Name: "20%", FeeMode: "RATE", Values: []int{20}},
			},
		}
		if err := cfg.Validate(); err != nil {
			t.Fatalf("Validate() returned error: %v", err)
		}
	})

	t.Run("rejects more than three suggestions", func(t *testing.T) {
		cfg := &TipConfig{
			Suggestions: []TipSuggestions{
				{Name: "A"},
				{Name: "B"},
				{Name: "C"},
				{Name: "D"},
			},
		}
		if err := cfg.Validate(); err == nil {
			t.Fatal("Validate() expected error, got nil")
		}
	})
}

func TestTipConfigMarshalIncludesSuggestionName(t *testing.T) {
	cfg := &TipConfig{
		OnScreenTip: true,
		TipMode:     "ON_SALE",
		TipWithTax:  true,
		Suggestions: []TipSuggestions{
			{Name: "15%", FeeMode: "RATE", Values: []int{15}},
		},
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("json.Marshal() returned error: %v", err)
	}

	want := `{"onScreenTip":true,"tipMode":"ON_SALE","tipWithTax":true,"suggestions":[{"name":"15%","feeMode":"RATE","values":[15]}]}`
	if string(data) != want {
		t.Fatalf("unexpected JSON:\nwant %s\ngot  %s", want, string(data))
	}
}
