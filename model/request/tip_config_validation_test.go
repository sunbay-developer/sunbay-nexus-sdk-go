package request

import (
	"testing"

	"github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"
)

func TestRequestTipConfigValidation(t *testing.T) {
	t.Run("sale request", func(t *testing.T) {
		req := &SaleRequest{
			TipConfig: &common.TipConfig{
				Suggestions: []common.TipSuggestions{
					{Names: []string{"A"}},
					{Names: []string{"B"}},
					{Names: []string{"C"}},
					{Names: []string{"D"}},
				},
			},
		}
		if err := req.Validate(); err == nil {
			t.Fatal("Validate() expected error, got nil")
		}
	})

	t.Run("post auth request", func(t *testing.T) {
		req := &PostAuthRequest{
			TipConfig: &common.TipConfig{
				Suggestions: []common.TipSuggestions{
					{Names: []string{"A"}},
					{Names: []string{"B"}},
					{Names: []string{"C"}},
					{Names: []string{"D"}},
				},
			},
		}
		if err := req.Validate(); err == nil {
			t.Fatal("Validate() expected error, got nil")
		}
	})
}
