package request

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// CreateCheckoutSessionRequest is the body for POST /v1/checkout/create-session (hosted payment page).
type CreateCheckoutSessionRequest struct {
	// AppID is the application ID assigned by SUNBAY
	AppID string `json:"appId"`

	// TransactionRequestID is a unique idempotency key (letters, digits, underscore, hyphen; max 32)
	TransactionRequestID string `json:"transactionRequestId"`

	// ReferenceOrderID is the merchant reference order ID (6–32 chars)
	ReferenceOrderID string `json:"referenceOrderId"`

	// MerchantID is the merchant ID assigned by SUNBAY
	MerchantID string `json:"merchantId"`

	// Amount is the amount breakdown; payable total = orderAmount + taxAmount + surchargeAmount
	Amount *common.SaleAmount `json:"amount"`

	// Description is the order description shown on checkout
	Description string `json:"description"`

	// ProductList is optional cart line items
	ProductList []common.CheckoutProductLine `json:"productList,omitempty"`

	// CollectBillingAddress requests billing address collection on checkout (default false)
	CollectBillingAddress bool `json:"collectBillingAddress,omitempty"`

	// CollectShippingAddress requests shipping address collection on checkout (default false)
	CollectShippingAddress bool `json:"collectShippingAddress,omitempty"`

	// MerchantReturnURL is the return URL after payment completes or is cancelled
	MerchantReturnURL string `json:"merchantReturnUrl,omitempty"`

	// NotifyURL is the optional async webhook URL (public HTTPS)
	NotifyURL string `json:"notifyUrl,omitempty"`
}
