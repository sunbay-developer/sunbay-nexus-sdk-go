package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// CheckoutDirectSaleResponse is the data payload for direct checkout / wallet sale success.
type CheckoutDirectSaleResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY transaction ID when returned
	TransactionID string `json:"transactionId,omitempty"`

	// ReferenceOrderID is echoed from the request when returned
	ReferenceOrderID string `json:"referenceOrderId,omitempty"`

	// TransactionRequestID is echoed from the request when returned
	TransactionRequestID string `json:"transactionRequestId,omitempty"`
}
