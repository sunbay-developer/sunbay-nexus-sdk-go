package types

// AuthenticationMethod represents authentication method
type AuthenticationMethod string

const (
	// AuthenticationMethodNotAuthenticated NOT_AUTHENTICATED (code: "NOT_AUTHENTICATED")
	AuthenticationMethodNotAuthenticated AuthenticationMethod = "NOT_AUTHENTICATED"

	// AuthenticationMethodPIN PIN (code: "PIN")
	AuthenticationMethodPIN AuthenticationMethod = "PIN"

	// AuthenticationMethodOfflinePIN OFFLINE_PIN (code: "OFFLINE_PIN")
	AuthenticationMethodOfflinePIN AuthenticationMethod = "OFFLINE_PIN"

	// AuthenticationMethodBypass BY_PASS (code: "BY_PASS")
	AuthenticationMethodBypass AuthenticationMethod = "BY_PASS"

	// AuthenticationMethodSignature SIGNATURE (code: "SIGNATURE")
	AuthenticationMethodSignature AuthenticationMethod = "SIGNATURE"
)

// String returns the authentication method code
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

