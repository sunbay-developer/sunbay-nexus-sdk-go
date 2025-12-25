package enums

// EntryMode represents entry mode
type EntryMode string

const (
	// EntryModeManual is the manual entry mode
	EntryModeManual EntryMode = "MANUAL"

	// EntryModeSwipe is the swipe card mode
	EntryModeSwipe EntryMode = "SWIPE"

	// EntryModeFallbackSwipe is the fallback swipe mode
	EntryModeFallbackSwipe EntryMode = "FALLBACK_SWIPE"

	// EntryModeContact is the contact chip mode
	EntryModeContact EntryMode = "CONTACT"

	// EntryModeContactless is the contactless mode
	EntryModeContactless EntryMode = "CONTACTLESS"
)

// String returns the string representation
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

