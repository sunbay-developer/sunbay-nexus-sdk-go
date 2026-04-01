package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// CreateCheckoutSessionResponse is the data payload for create checkout session success.
type CreateCheckoutSessionResponse struct {
	common.BaseResponse

	// CheckoutURL is the hosted payment page URL to redirect the customer to
	CheckoutURL string `json:"checkoutUrl,omitempty"`

	// ExpiresAt is the session expiry time (e.g. ISO 8601); session lifetime is 30 minutes from success
	ExpiresAt string `json:"expiresAt,omitempty"`
}
