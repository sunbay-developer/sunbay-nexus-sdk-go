package response

import (
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/enums"
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"
)

// QueryResponse represents a query transaction response
type QueryResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID
	TransactionID string `json:"transactionId"`

	// TransactionRequestID is the transaction request ID
	TransactionRequestID string `json:"transactionRequestId"`

	// ReferenceOrderID is the reference order ID (only returned for transactions with reference order ID such as sale, authorization, forced authorization)
	ReferenceOrderID string `json:"referenceOrderId,omitempty"`

	// TransactionStatus is the transaction status
	TransactionStatus enums.TransactionStatus `json:"transactionStatus"`

	// TransactionType is the transaction type
	TransactionType enums.TransactionType `json:"transactionType,omitempty"`

	// Amount is the transaction amount details
	Amount *common.Amount `json:"amount,omitempty"`

	// CreateTime is the transaction creation time, format: yyyy-MM-DDTHH:mm:ss+TIMEZONE (ISO 8601)
	CreateTime string `json:"createTime,omitempty"`

	// CompleteTime is the transaction completion time, format: yyyy-MM-DDTHH:mm:ss+TIMEZONE (ISO 8601)
	CompleteTime string `json:"completeTime,omitempty"`

	// MaskedPAN is the masked card number (first 6 digits + **** + last 4 digits)
	MaskedPAN string `json:"maskedPan,omitempty"`

	// CardNetworkType is the card network type
	CardNetworkType enums.CardNetworkType `json:"cardNetworkType,omitempty"`

	// PaymentMethodID is the payment method ID
	PaymentMethodID string `json:"paymentMethodId,omitempty"`

	// SubPaymentMethodID is the sub payment method ID
	SubPaymentMethodID string `json:"subPaymentMethodId,omitempty"`

	// BatchNo is the batch number
	BatchNo string `json:"batchNo,omitempty"`

	// VoucherNo is the voucher number
	VoucherNo string `json:"voucherNo,omitempty"`

	// STAN is the system trace number
	STAN string `json:"stan,omitempty"`

	// RRN is the reference number
	RRN string `json:"rrn,omitempty"`

	// AuthCode is the authorization code
	AuthCode string `json:"authCode,omitempty"`

	// EntryMode is the entry mode
	EntryMode enums.EntryMode `json:"entryMode,omitempty"`

	// AuthenticationMethod is the authentication method
	AuthenticationMethod enums.AuthenticationMethod `json:"authenticationMethod,omitempty"`

	// TransactionResultCode is the transaction result code
	TransactionResultCode string `json:"transactionResultCode,omitempty"`

	// TransactionResultMsg is the transaction result message
	TransactionResultMsg string `json:"transactionResultMsg,omitempty"`

	// TerminalSN is the terminal serial number
	TerminalSN string `json:"terminalSn,omitempty"`

	// Description is the product description
	Description string `json:"description,omitempty"`

	// Attach is additional data, returned as-is
	Attach string `json:"attach,omitempty"`
}

