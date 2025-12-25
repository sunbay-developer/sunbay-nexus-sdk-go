package util

import (
	"encoding/json"
	"fmt"
)

// ToJSON converts an object to JSON string
func ToJSON(obj interface{}) string {
	if obj == nil {
		return ""
	}
	data, err := json.Marshal(obj)
	if err != nil {
		// Return error JSON instead of panic
		return fmt.Sprintf(`{"error":"Failed to serialize object to JSON: %s"}`, err.Error())
	}
	return string(data)
}

// FromJSON parses JSON string to object
func FromJSON(jsonStr string, obj interface{}) error {
	if jsonStr == "" {
		return nil
	}
	return json.Unmarshal([]byte(jsonStr), obj)
}

// GetObjectMapper gets JSON encoder
// Uses Go's standard encoding/json package
func GetObjectMapper() *json.Encoder {
	return json.NewEncoder(nil)
}

