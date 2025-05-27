package core

import (
	"encoding/json"
	"fmt"
)

// ToJSONRaw converts any Go value to a raw JSON representation
func ToJSONRaw(data any) (json.RawMessage, error) {
	// Marshal the data to JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	// Convert to json.RawMessage type
	return jsonBytes, nil
}

// ToJSONRawPretty converts any Go value to a raw JSON representation with indentation
func ToJSONRawPretty(data any) (json.RawMessage, error) {
	// Marshal the data to JSON bytes with indentation
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	// Convert to json.RawMessage type
	return json.RawMessage(jsonBytes), nil
}

// AnyToRawJSON converts any Go value to a json.RawMessage
func AnyToRawJSON(data any) (json.RawMessage, error) {
	// For nil input, return an empty JSON object
	if data == nil {
		return json.RawMessage("null"), nil
	}

	// Check if data is already a json.RawMessage
	if rawJSON, ok := data.(json.RawMessage); ok {
		return rawJSON, nil
	}

	// If it's a string, check if it's already valid JSON
	if str, ok := data.(string); ok {
		if json.Valid([]byte(str)) {
			return json.RawMessage(str), nil
		}
	}

	// Otherwise, marshal the data to JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	return jsonBytes, nil
}

// RawJSONToAny converts a json.RawMessage to an any type
func RawJSONToAny(raw json.RawMessage) (any, error) {
	// For empty or nil input, return nil
	if raw == nil || len(raw) == 0 {
		return nil, nil
	}

	// Create a variable to hold the unmarshaled data
	var result any

	// Unmarshal the raw JSON into the any variable
	err := json.Unmarshal(raw, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return result, nil
}

func PrettyPrint(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", prettyJSON)
}
