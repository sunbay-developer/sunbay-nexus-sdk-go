package errors

import "fmt"

// NetworkError represents a network error
type NetworkError struct {
	message   string
	retryable bool
	cause     error
}

// NewNetworkError creates a network error
func NewNetworkError(message string, retryable bool, cause error) *NetworkError {
	return &NetworkError{
		message:   message,
		retryable: retryable,
		cause:     cause,
	}
}

// Error implements the error interface
func (e *NetworkError) Error() string {
	return fmt.Sprintf("NetworkError{message='%s', retryable=%v}", e.message, e.retryable)
}

// Unwrap returns the original error (Go 1.13+ error wrapping)
func (e *NetworkError) Unwrap() error {
	return e.cause
}

// IsRetryable returns whether the error is retryable
func (e *NetworkError) IsRetryable() bool {
	return e.retryable
}

