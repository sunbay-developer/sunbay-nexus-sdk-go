package enums

// AuthenticationMethod represents authentication method
type AuthenticationMethod string

const (
	// AuthenticationMethodNotAuthenticated is the not authenticated method
	AuthenticationMethodNotAuthenticated AuthenticationMethod = "NOT_AUTHENTICATED"

	// AuthenticationMethodPIN is the PIN authentication method
	AuthenticationMethodPIN AuthenticationMethod = "PIN"

	// AuthenticationMethodOfflinePIN is the offline PIN method
	AuthenticationMethodOfflinePIN AuthenticationMethod = "OFFLINE_PIN"

	// AuthenticationMethodBypass is the bypass authentication method
	AuthenticationMethodBypass AuthenticationMethod = "BY_PASS"

	// AuthenticationMethodSignature is the signature authentication method
	AuthenticationMethodSignature AuthenticationMethod = "SIGNATURE"
)

// String returns the string representation
func (a AuthenticationMethod) String() string {
	return string(a)
}

// IsValid checks if the authentication method is valid
func (a AuthenticationMethod) IsValid() bool {
	switch a {
	case AuthenticationMethodNotAuthenticated, AuthenticationMethodPIN,
		AuthenticationMethodOfflinePIN, AuthenticationMethodBypass,
		AuthenticationMethodSignature:
		return true
	default:
		return false
	}
}

