package errors

import "fmt"

// BusinessError represents a business error
// Used for API business errors and parameter validation errors
type BusinessError struct {
	code    string
	message string
	traceID string
}

// NewBusinessError creates a business error
func NewBusinessError(code, message, traceID string) *BusinessError {
	return &BusinessError{
		code:    code,
		message: message,
		traceID: traceID,
	}
}

// Error implements the error interface
func (e *BusinessError) Error() string {
	if e.code != "" {
		return fmt.Sprintf("BusinessError{code='%s', message='%s', traceID='%s'}", e.code, e.message, e.traceID)
	}
	return fmt.Sprintf("BusinessError{message='%s'}", e.message)
}

// Code returns the error code
func (e *BusinessError) Code() string {
	return e.code
}

// Message returns the error message
func (e *BusinessError) Message() string {
	return e.message
}

// TraceID returns the trace ID
func (e *BusinessError) TraceID() string {
	return e.traceID
}

