package common

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/enums"

// PaymentMethodInfo represents payment method information
type PaymentMethodInfo struct {
	// Category is the payment category
	Category enums.PaymentCategory `json:"category,omitempty"`

	// ID is the specific payment method: WECHAT (WeChat)/ALIPAY (Alipay) etc. For card payments, usually only category needs to be specified
	ID string `json:"id,omitempty"`
}

