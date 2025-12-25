package enums

// EntryMode represents entry mode
type EntryMode string

const (
	// EntryModeManual MANUAL (code: "MANUAL")
	EntryModeManual EntryMode = "MANUAL"

	// EntryModeSwipe SWIPE (code: "SWIPE")
	EntryModeSwipe EntryMode = "SWIPE"

	// EntryModeFallbackSwipe FALLBACK_SWIPE (code: "FALLBACK_SWIPE")
	EntryModeFallbackSwipe EntryMode = "FALLBACK_SWIPE"

	// EntryModeContact CONTACT (code: "CONTACT")
	EntryModeContact EntryMode = "CONTACT"

	// EntryModeContactless CONTACTLESS (code: "CONTACTLESS")
	EntryModeContactless EntryMode = "CONTACTLESS"
)

// String returns the entry mode code
func (e EntryMode) String() string {
	return string(e)
}

// IsValid checks if the entry mode is valid
func (e EntryMode) IsValid() bool {
	switch e {
	case EntryModeManual, EntryModeSwipe, EntryModeFallbackSwipe,
		EntryModeContact, EntryModeContactless:
		return true
	default:
		return false
	}
}

