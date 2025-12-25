package enums

// PaymentCategory represents payment category
type PaymentCategory string

const (
	// PaymentCategoryCard is the card payment category
	PaymentCategoryCard PaymentCategory = "CARD"

	// PaymentCategoryCardCredit is the credit card network category
	PaymentCategoryCardCredit PaymentCategory = "CARD-CREDIT"

	// PaymentCategoryCardDebit is the debit card network category
	PaymentCategoryCardDebit PaymentCategory = "CARD-DEBIT"

	// PaymentCategoryQRMPM is the QR code merchant presented mode category
	PaymentCategoryQRMPM PaymentCategory = "QR-MPM"

	// PaymentCategoryQRCPM is the QR code customer presented mode category
	PaymentCategoryQRCPM PaymentCategory = "QR-CPM"
)

// String returns the string representation
func (p PaymentCategory) String() string {
	return string(p)
}

// Value returns the JSON value for serialization
func (p PaymentCategory) Value() string {
	return string(p)
}

// IsValid checks if the payment category is valid
func (p PaymentCategory) IsValid() bool {
	switch p {
	case PaymentCategoryCard, PaymentCategoryCardCredit, PaymentCategoryCardDebit,
		PaymentCategoryQRMPM, PaymentCategoryQRCPM:
		return true
	default:
		return false
	}
}

