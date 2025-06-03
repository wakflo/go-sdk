package core

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/rs/xid"
)

// mapToStruct maps a map[string]interface{} to a struct using reflection
func mapToStruct(input map[string]interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.Elem().Kind() != reflect.Struct {
		return false
	}

	targetStruct := targetValue.Elem()
	targetType := targetStruct.Type()

	for i := 0; i < targetStruct.NumField(); i++ {
		field := targetStruct.Field(i)
		fieldType := targetType.Field(i)

		// Get the JSON tag or use the field name
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = fieldType.Tag.Get("mapstructure")
		}
		if jsonTag == "" {
			jsonTag = strings.ToLower(fieldType.Name)
		} else {
			// Remove omitempty and other options
			jsonTag = strings.Split(jsonTag, ",")[0]
		}

		// Skip if field can't be set
		if !field.CanSet() {
			continue
		}

		// Get value from input
		if value, exists := input[jsonTag]; exists && value != nil {
			// Set the field value
			if field.Type() == reflect.TypeOf(value) {
				field.Set(reflect.ValueOf(value))
			} else {
				// Try to convert the value
				if convertedValue := convertValue(value, field.Type()); convertedValue.IsValid() {
					field.Set(convertedValue)
				}
			}
		}
	}

	return true
}

// convertValue attempts to convert a value to the target type
func convertValue(value interface{}, targetType reflect.Type) reflect.Value {
	sourceValue := reflect.ValueOf(value)

	// If types match, return as-is
	if sourceValue.Type() == targetType {
		return sourceValue
	}

	// Handle string to other types
	if sourceValue.Kind() == reflect.String {
		strValue := sourceValue.String()
		switch targetType.Kind() {
		case reflect.Bool:
			if strValue == "true" {
				return reflect.ValueOf(true)
			} else if strValue == "false" {
				return reflect.ValueOf(false)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if intVal, err := parseInt(strValue); err == nil {
				return reflect.ValueOf(intVal).Convert(targetType)
			}
		case reflect.Float32, reflect.Float64:
			if floatVal, err := parseFloat(strValue); err == nil {
				return reflect.ValueOf(floatVal).Convert(targetType)
			}
		}
	}

	// Handle number conversions
	if sourceValue.Kind() >= reflect.Int && sourceValue.Kind() <= reflect.Float64 {
		if targetType.Kind() >= reflect.Int && targetType.Kind() <= reflect.Float64 {
			return sourceValue.Convert(targetType)
		}
	}

	// Handle slice/array conversions
	if sourceValue.Kind() == reflect.Slice && targetType.Kind() == reflect.Slice {
		sourceLen := sourceValue.Len()
		targetSlice := reflect.MakeSlice(targetType, sourceLen, sourceLen)

		for i := 0; i < sourceLen; i++ {
			sourceElement := sourceValue.Index(i)
			targetElement := targetSlice.Index(i)

			if convertedElement := convertValue(sourceElement.Interface(), targetType.Elem()); convertedElement.IsValid() {
				targetElement.Set(convertedElement)
			}
		}

		return targetSlice
	}

	return reflect.Value{}
}

// parseInt parses a string to int64
func parseInt(s string) (int64, error) {
	var result int64
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}

// parseFloat parses a string to float64
func parseFloat(s string) (float64, error) {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	return result, err
}

// Now returns the current time as a string in RFC3339 format
func Now() string {
	return time.Now().Format(time.RFC3339)
}

// GenerateID generates a new unique ID
func GenerateID() string {
	return xid.New().String()
}

// FromJSON parses a JSON string into a target type
func FromJSON[T any](jsonStr string) (*T, error) {
	var result T
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ParseTextareaLines parses textarea input (string with newlines) into a slice of strings
func ParseTextareaLines(text interface{}) []string {
	var lines []string
	if str, ok := text.(string); ok {
		for _, line := range strings.Split(str, "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				lines = append(lines, line)
			}
		}
	}
	return lines
}

// ParseJSONField parses a JSON string field into a map or struct
func ParseJSONField[T any](jsonStr string) (*T, error) {
	if jsonStr == "" {
		return nil, nil
	}

	var result T
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &result, nil
}

// GetStringFromMap safely gets a string value from a map
func GetStringFromMap(m map[string]interface{}, key string, defaultValue string) string {
	if value, exists := m[key]; exists {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return defaultValue
}

// GetIntFromMap safely gets an int value from a map
func GetIntFromMap(m map[string]interface{}, key string, defaultValue int) int {
	if value, exists := m[key]; exists {
		switch v := value.(type) {
		case int:
			return v
		case int64:
			return int(v)
		case float64:
			return int(v)
		}
	}
	return defaultValue
}

// GetBoolFromMap safely gets a bool value from a map
func GetBoolFromMap(m map[string]interface{}, key string, defaultValue bool) bool {
	if value, exists := m[key]; exists {
		if b, ok := value.(bool); ok {
			return b
		}
	}
	return defaultValue
}

// MergeMap merges two maps, with the second map taking precedence
func MergeMap(base, override map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Copy base map
	for k, v := range base {
		result[k] = v
	}

	// Override with second map
	for k, v := range override {
		result[k] = v
	}

	return result
}

// Contains checks if a slice contains a specific value
func Contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Filter filters a slice based on a predicate function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map transforms a slice using a mapping function
func Map[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

// Reduce reduces a slice to a single value using a reducer function
func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

// Keys returns the keys of a map as a slice
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns the values of a map as a slice
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Ptr returns a pointer to the given value
func Ptr[T any](v T) *T {
	return &v
}

// Deref dereferences a pointer, returning the zero value if nil
func Deref[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// Coalesce returns the first non-nil value from the arguments
func Coalesce[T any](values ...*T) *T {
	for _, v := range values {
		if v != nil {
			return v
		}
	}
	return nil
}

// Ternary is a ternary operator function
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// StringSliceToInterfaceSlice converts []string to []interface{}
func StringSliceToInterfaceSlice(strings []string) []interface{} {
	interfaces := make([]interface{}, len(strings))
	for i, s := range strings {
		interfaces[i] = s
	}
	return interfaces
}

// InterfaceSliceToStringSlice converts []interface{} to []string
func InterfaceSliceToStringSlice(interfaces []interface{}) []string {
	strings := make([]string, 0, len(interfaces))
	for _, v := range interfaces {
		if s, ok := v.(string); ok {
			strings = append(strings, s)
		}
	}
	return strings
}

// SafeString safely converts any value to string
func SafeString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

// SafeInt safely converts any value to int
func SafeInt(v interface{}) int {
	switch val := v.(type) {
	case int:
		return val
	case int64:
		return int(val)
	case float64:
		return int(val)
	case string:
		if i, err := parseInt(val); err == nil {
			return int(i)
		}
	}
	return 0
}

// SafeBool safely converts any value to bool
func SafeBool(v interface{}) bool {
	switch val := v.(type) {
	case bool:
		return val
	case string:
		return val == "true" || val == "1" || val == "yes"
	case int, int64:
		return val != 0
	case float64:
		return val != 0.0
	}
	return false
}
