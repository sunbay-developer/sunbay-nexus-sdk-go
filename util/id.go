package util

import (
	"crypto/rand"
	"fmt"
	"time"
)

// GenerateUUID generates a UUID v4 using standard library
// Format: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
func GenerateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// Fallback: use timestamp-based ID if crypto/rand fails
		return generateFallbackID()
	}

	// Set version (4) and variant bits
	b[6] = (b[6] & 0x0f) | 0x40 // Version 4
	b[8] = (b[8] & 0x3f) | 0x80 // Variant 10

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// generateFallbackID generates a fallback ID if crypto/rand fails
func generateFallbackID() string {
	// Use timestamp + simple counter as fallback
	now := time.Now().UnixNano()
	return fmt.Sprintf("%016x-4000-8000-8000-%012x", now, now&0xffffffffffff)
}

// GenerateRequestID generates a request ID
func GenerateRequestID() string {
	return GenerateUUID()
}

