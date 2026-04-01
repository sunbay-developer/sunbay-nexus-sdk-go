package request

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// CheckoutDirectSaleRequest is the body for POST /v1/checkout/sale (direct / wallet payment).
type CheckoutDirectSaleRequest struct {
	// AppID is the application ID assigned by SUNBAY
	AppID string `json:"appId"`

	// MerchantID is the merchant ID assigned by SUNBAY
	MerchantID string `json:"merchantId"`

	// TransactionRequestID is a merchant-generated idempotency key (max 32)
	TransactionRequestID string `json:"transactionRequestId"`

	// ReferenceOrderID is the merchant order ID (6–32 chars)
	ReferenceOrderID string `json:"referenceOrderId"`

	// Description is the order description
	Description string `json:"description"`

	// Amount is the amount breakdown; charged total = orderAmount + taxAmount + surchargeAmount
	Amount *common.SaleAmount `json:"amount"`

	// ProductList is optional line items
	ProductList []common.CheckoutProductLine `json:"productList,omitempty"`

	// PaymentMethod is GOOGLE_PAY or APPLE_PAY (see constant.CheckoutPaymentMethod*)
	PaymentMethod string `json:"paymentMethod"`

	// CardEncryptedData is the wallet token JSON string when using GOOGLE_PAY or APPLE_PAY
	CardEncryptedData string `json:"cardEncryptedData,omitempty"`

	// CustomerEmail is the buyer email
	CustomerEmail string `json:"customerEmail,omitempty"`

	// CustomerName is the buyer name
	CustomerName string `json:"customerName,omitempty"`

	// BillingAddress is optional billing address
	BillingAddress *common.CheckoutAddress `json:"billingAddress,omitempty"`

	// ShippingAddress is optional shipping address
	ShippingAddress *common.CheckoutAddress `json:"shippingAddress,omitempty"`

	// NotifyURL is the optional async webhook URL (public HTTPS)
	NotifyURL string `json:"notifyUrl,omitempty"`

	// MerchantReturnURL is the browser return URL (e.g. 3DS redirect)
	MerchantReturnURL string `json:"merchantReturnUrl,omitempty"`
}
