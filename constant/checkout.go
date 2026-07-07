package constant

// Checkout (online) API paths under /v1
const (
	PathCheckoutCreateSession = CommonPrefix + "/checkout/create-session"
	PathCheckoutSale          = CommonPrefix + "/checkout/sale"
	PathCheckoutRefund        = CommonPrefix + "/checkout/refund"
)

// Direct checkout paymentMethod values (POST /v1/checkout/sale)
const (
	CheckoutPaymentMethodGooglePay = "GOOGLE_PAY"
	CheckoutPaymentMethodApplePay  = "APPLE_PAY"
)
