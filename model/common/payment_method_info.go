package common

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/types"

// PaymentMethodInfo represents payment method information
type PaymentMethodInfo struct {
	// Category is the payment category
	Category types.PaymentCategory `json:"category,omitempty"`

	// ID is the specific payment method: WECHAT (WeChat)/ALIPAY (Alipay) etc. For card payments, usually only category needs to be specified
	ID string `json:"id,omitempty"`

	// SubID is the sub payment method. When category = CARD, must not be specified. Currently only when category = EBT and id = EBT may SubID be specified; allowed values: SNAP, VOUCHER, BENEFIT (see EBT sub payment method)
	SubID string `json:"subId,omitempty"`
}

