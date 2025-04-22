// Copyright 2022-present Wakflo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// JSONObject is a type alias for map[string]any.
type JSONObject = map[string]any

// ToJSONMap converts an interface of type `any` to a map[string]any.
// Returns the map and a boolean indicating success.
func ToJSONMap(input any) (JSONObject, bool) {
	if input == nil {
		return nil, false
	}

	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, false
	}

	var result map[string]any
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, false
	}

	return result, true
}

// ToJSON converts an interface of type `any` to a map[string]any.
// Returns the map and a boolean indicating success.
func ToJSON(input any) (JSON, bool) {
	if input == nil {
		return nil, false
	}

	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, false
	}

	var result any
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, false
	}

	return result, true
}

// JSON is a generic type for JSON values
type JSON interface{}

// GetString retrieves a string value from a JSON object at the specified path
func GetString(data JSON, path string) (string, error) {
	value, err := GetValue(data, path)
	if err != nil {
		return "", err
	}

	switch v := value.(type) {
	case string:
		return v, nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case int:
		return strconv.Itoa(v), nil
	case bool:
		return strconv.FormatBool(v), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("value at path %s is not a string: %T", path, value)
	}
}

// GetInt retrieves an integer value from a JSON object at the specified path
func GetInt(data JSON, path string) (int, error) {
	value, err := GetValue(data, path)
	if err != nil {
		return 0, err
	}

	switch v := value.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("value at path %s is not an integer: %T", path, value)
	}
}

// GetFloat retrieves a float value from a JSON object at the specified path
func GetFloat(data JSON, path string) (float64, error) {
	value, err := GetValue(data, path)
	if err != nil {
		return 0, err
	}

	switch v := value.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("value at path %s is not a float: %T", path, value)
	}
}

// GetBool retrieves a boolean value from a JSON object at the specified path
func GetBool(data JSON, path string) (bool, error) {
	value, err := GetValue(data, path)
	if err != nil {
		return false, err
	}

	switch v := value.(type) {
	case bool:
		return v, nil
	case int:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case string:
		return strconv.ParseBool(v)
	case nil:
		return false, nil
	default:
		return false, fmt.Errorf("value at path %s is not a boolean: %T", path, value)
	}
}

// GetObject retrieves an object (map) from a JSON object at the specified path
func GetObject(data JSON, path string) (map[string]interface{}, error) {
	value, err := GetValue(data, path)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	obj, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("value at path %s is not an object: %T", path, value)
	}

	return obj, nil
}

// GetArray retrieves an array from a JSON object at the specified path
func GetArray(data JSON, path string) ([]interface{}, error) {
	value, err := GetValue(data, path)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	arr, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("value at path %s is not an array: %T", path, value)
	}

	return arr, nil
}

// GetValue retrieves a value from a JSON object at the specified path
func GetValue(data JSON, path string) (interface{}, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}

	// Handle empty path
	if path == "" || path == "." {
		return data, nil
	}

	// Split the path into parts
	parts := strings.Split(path, ".")

	// Start with the root object
	var current interface{} = data

	// Traverse the path
	for i, part := range parts {
		// Handle array indices in the path (e.g., "items[0].name")
		indexStart := strings.Index(part, "[")
		indexEnd := strings.Index(part, "]")

		if indexStart > 0 && indexEnd > indexStart {
			// Extract the property name and array index
			propertyName := part[:indexStart]
			indexStr := part[indexStart+1 : indexEnd]
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return nil, fmt.Errorf("invalid array index %s at path %s", indexStr, strings.Join(parts[:i+1], "."))
			}

			// Get the object or map
			obj, ok := current.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("value at path %s is not an object: %T", strings.Join(parts[:i], "."), current)
			}

			// Get the array from the object
			arr, ok := obj[propertyName].([]interface{})
			if !ok {
				return nil, fmt.Errorf("value at path %s is not an array: %T", strings.Join(parts[:i+1], "."), obj[propertyName])
			}

			// Check if the index is in bounds
			if index < 0 || index >= len(arr) {
				return nil, fmt.Errorf("array index %d out of bounds at path %s", index, strings.Join(parts[:i+1], "."))
			}

			// Set the current value to the array element
			current = arr[index]
		} else {
			// Regular property access
			obj, ok := current.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("value at path %s is not an object: %T", strings.Join(parts[:i], "."), current)
			}

			// Check if the property exists
			current, ok = obj[part]
			if !ok {
				return nil, fmt.Errorf("property %s not found at path %s", part, strings.Join(parts[:i+1], "."))
			}
		}
	}

	return current, nil
}

// SetValue sets a value in a JSON object at the specified path
func SetValue(data map[string]interface{}, path string, value interface{}) error {
	if data == nil {
		return errors.New("data is nil")
	}

	// Handle empty path
	if path == "" || path == "." {
		return errors.New("cannot set value at root path")
	}

	// Split the path into parts
	parts := strings.Split(path, ".")

	// Navigate to the parent object
	parent := data
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]

		// Handle array indices in the path (e.g., "items[0].name")
		indexStart := strings.Index(part, "[")
		indexEnd := strings.Index(part, "]")

		if indexStart > 0 && indexEnd > indexStart {
			// Extract the property name and array index
			propertyName := part[:indexStart]
			indexStr := part[indexStart+1 : indexEnd]
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return fmt.Errorf("invalid array index %s at path %s", indexStr, strings.Join(parts[:i+1], "."))
			}

			// Get or create the array
			arrObj, ok := parent[propertyName]
			if !ok {
				// Create the array if it doesn't exist
				arrObj = make([]interface{}, index+1)
				parent[propertyName] = arrObj
			}

			arr, ok := arrObj.([]interface{})
			if !ok {
				return fmt.Errorf("value at path %s is not an array: %T", strings.Join(parts[:i+1], "."), arrObj)
			}

			// Ensure the array is big enough
			if index >= len(arr) {
				newArr := make([]interface{}, index+1)
				copy(newArr, arr)
				arr = newArr
				parent[propertyName] = arr
			}

			// If we're at the second-to-last part, set the value in the array
			if i == len(parts)-2 {
				arr[index] = value
				return nil
			}

			// Otherwise, ensure the array element is an object
			if arr[index] == nil {
				arr[index] = make(map[string]interface{})
			}

			objElem, ok := arr[index].(map[string]interface{})
			if !ok {
				return fmt.Errorf("value at path %s is not an object: %T", strings.Join(parts[:i+1], "."), arr[index])
			}

			parent = objElem
		} else {
			// Regular property access
			nextObj, ok := parent[part]
			if !ok {
				// Create the next object if it doesn't exist
				nextObj = make(map[string]interface{})
				parent[part] = nextObj
			}

			nextParent, ok := nextObj.(map[string]interface{})
			if !ok {
				return fmt.Errorf("value at path %s is not an object: %T", strings.Join(parts[:i+1], "."), nextObj)
			}

			parent = nextParent
		}
	}

	// Set the value in the parent object
	lastPart := parts[len(parts)-1]

	// Handle array indices in the last part
	indexStart := strings.Index(lastPart, "[")
	indexEnd := strings.Index(lastPart, "]")

	if indexStart > 0 && indexEnd > indexStart {
		// Extract the property name and array index
		propertyName := lastPart[:indexStart]
		indexStr := lastPart[indexStart+1 : indexEnd]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return fmt.Errorf("invalid array index %s at path %s", indexStr, path)
		}

		// Get or create the array
		arrObj, ok := parent[propertyName]
		if !ok {
			// Create the array if it doesn't exist
			arrObj = make([]interface{}, index+1)
			parent[propertyName] = arrObj
		}

		arr, ok := arrObj.([]interface{})
		if !ok {
			return fmt.Errorf("value at path %s is not an array: %T", strings.Join(parts[:len(parts)-1], "."), arrObj)
		}

		// Ensure the array is big enough
		if index >= len(arr) {
			newArr := make([]interface{}, index+1)
			copy(newArr, arr)
			arr = newArr
			parent[propertyName] = arr
		}

		// Set the value in the array
		arr[index] = value
	} else {
		// Regular property access
		parent[lastPart] = value
	}

	return nil
}

// RemoveValue removes a value from a JSON object at the specified path
func RemoveValue(data map[string]interface{}, path string) error {
	if data == nil {
		return errors.New("data is nil")
	}

	// Handle empty path
	if path == "" || path == "." {
		return errors.New("cannot remove root path")
	}

	// Split the path into parts
	parts := strings.Split(path, ".")

	// Navigate to the parent object
	parent := data
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]

		// Handle array indices in the path (e.g., "items[0].name")
		indexStart := strings.Index(part, "[")
		indexEnd := strings.Index(part, "]")

		if indexStart > 0 && indexEnd > indexStart {
			// Extract the property name and array index
			propertyName := part[:indexStart]
			indexStr := part[indexStart+1 : indexEnd]
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return fmt.Errorf("invalid array index %s at path %s", indexStr, strings.Join(parts[:i+1], "."))
			}

			// Get the array
			arrObj, ok := parent[propertyName]
			if !ok {
				return fmt.Errorf("property %s not found at path %s", propertyName, strings.Join(parts[:i+1], "."))
			}

			arr, ok := arrObj.([]interface{})
			if !ok {
				return fmt.Errorf("value at path %s is not an array: %T", strings.Join(parts[:i+1], "."), arrObj)
			}

			// Check if the index is in bounds
			if index < 0 || index >= len(arr) {
				return fmt.Errorf("array index %d out of bounds at path %s", index, strings.Join(parts[:i+1], "."))
			}

			// Get the object at the array index
			objElem, ok := arr[index].(map[string]interface{})
			if !ok {
				return fmt.Errorf("value at path %s is not an object: %T", strings.Join(parts[:i+1], "."), arr[index])
			}

			parent = objElem
		} else {
			// Regular property access
			nextObj, ok := parent[part]
			if !ok {
				return fmt.Errorf("property %s not found at path %s", part, strings.Join(parts[:i+1], "."))
			}

			nextParent, ok := nextObj.(map[string]interface{})
			if !ok {
				return fmt.Errorf("value at path %s is not an object: %T", strings.Join(parts[:i+1], "."), nextObj)
			}

			parent = nextParent
		}
	}

	// Remove the value from the parent object
	lastPart := parts[len(parts)-1]

	// Handle array indices in the last part
	indexStart := strings.Index(lastPart, "[")
	indexEnd := strings.Index(lastPart, "]")

	if indexStart > 0 && indexEnd > indexStart {
		// Extract the property name and array index
		propertyName := lastPart[:indexStart]
		indexStr := lastPart[indexStart+1 : indexEnd]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return fmt.Errorf("invalid array index %s at path %s", indexStr, path)
		}

		// Get the array
		arrObj, ok := parent[propertyName]
		if !ok {
			return fmt.Errorf("property %s not found at path %s", propertyName, strings.Join(parts[:len(parts)-1], "."))
		}

		arr, ok := arrObj.([]interface{})
		if !ok {
			return fmt.Errorf("value at path %s is not an array: %T", strings.Join(parts[:len(parts)-1], "."), arrObj)
		}

		// Check if the index is in bounds
		if index < 0 || index >= len(arr) {
			return fmt.Errorf("array index %d out of bounds at path %s", index, path)
		}

		// Remove the element from the array (by setting to null)
		arr[index] = nil
	} else {
		// Regular property access - delete the property
		delete(parent, lastPart)
	}

	return nil
}

// DeepCopy creates a deep copy of a JSON value
func DeepCopy(data interface{}) interface{} {
	if data == nil {
		return nil
	}

	// Use json.Marshal and json.Unmarshal to create a deep copy
	bytes, err := json.Marshal(data)
	if err != nil {
		// If marshaling fails, try reflection-based deep copy
		return reflectDeepCopy(data)
	}

	var result interface{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		// If unmarshaling fails, try reflection-based deep copy
		return reflectDeepCopy(data)
	}

	return result
}

// // DeepCopy creates a deep copy of a JSON object or array
// func DeepCopy(value interface{}) interface{} {
// 	if value == nil {
// 		return nil
// 	}
//
// 	// Marshal and unmarshal to create a deep copy
// 	data, err := json.Marshal(value)
// 	if err != nil {
// 		// Return a shallow copy on error
// 		return shallowCopy(value)
// 	}
//
// 	var result interface{}
// 	if err := json.Unmarshal(data, &result); err != nil {
// 		// Return a shallow copy on error
// 		return shallowCopy(value)
// 	}
//
// 	return result
// }

// reflectDeepCopy creates a deep copy using reflection
func reflectDeepCopy(data interface{}) interface{} {
	if data == nil {
		return nil
	}

	value := reflect.ValueOf(data)

	// Handle simple types
	switch value.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String:
		return data
	}

	// Handle complex types
	switch value.Kind() {
	case reflect.Slice:
		if value.IsNil() {
			return nil
		}

		sliceLen := value.Len()
		sliceType := value.Type()
		newSlice := reflect.MakeSlice(sliceType, sliceLen, sliceLen)

		for i := 0; i < sliceLen; i++ {
			elem := value.Index(i)
			newElem := reflect.ValueOf(reflectDeepCopy(elem.Interface()))

			// Handle nil elements
			if newElem.IsValid() {
				newSlice.Index(i).Set(newElem)
			}
		}

		return newSlice.Interface()

	case reflect.Map:
		if value.IsNil() {
			return nil
		}

		mapType := value.Type()
		newMap := reflect.MakeMap(mapType)

		for _, key := range value.MapKeys() {
			val := value.MapIndex(key)
			newVal := reflect.ValueOf(reflectDeepCopy(val.Interface()))

			// Handle nil values
			if newVal.IsValid() {
				newMap.SetMapIndex(key, newVal)
			}
		}

		return newMap.Interface()

	case reflect.Ptr:
		if value.IsNil() {
			return nil
		}

		// Create a new pointer of the same type
		ptrType := value.Type()
		newPtr := reflect.New(ptrType.Elem())

		// Deep copy the pointed-to value
		elem := value.Elem()
		newElem := reflect.ValueOf(reflectDeepCopy(elem.Interface()))

		// Set the new pointer to point to the copied value
		newPtr.Elem().Set(newElem)

		return newPtr.Interface()

	case reflect.Struct:
		// Create a new struct of the same type
		structType := value.Type()
		newStruct := reflect.New(structType).Elem()

		// Copy each field
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			newField := reflect.ValueOf(reflectDeepCopy(field.Interface()))

			// Handle nil fields
			if newField.IsValid() {
				newStruct.Field(i).Set(newField)
			}
		}

		return newStruct.Interface()

	default:
		// For unsupported types, return as is
		return data
	}
}

// MergeJSON merges source JSON into target JSON
func MergeJSON(target, source map[string]interface{}) map[string]interface{} {
	if target == nil {
		return DeepCopy(source).(map[string]interface{})
	}

	if source == nil {
		return target
	}

	// Iterate through source and merge into target
	for key, sourceValue := range source {
		// Get the existing value from target
		targetValue, exists := target[key]

		if !exists {
			// Key doesn't exist in target, copy the value
			target[key] = DeepCopy(sourceValue)
			continue
		}

		// Handle maps recursively
		sourceMap, sourceIsMap := sourceValue.(map[string]interface{})
		targetMap, targetIsMap := targetValue.(map[string]interface{})

		if sourceIsMap && targetIsMap {
			// Recursively merge the maps
			target[key] = MergeJSON(targetMap, sourceMap)
			continue
		}

		// Handle arrays
		sourceArray, sourceIsArray := sourceValue.([]interface{})
		targetArray, targetIsArray := targetValue.([]interface{})

		if sourceIsArray && targetIsArray {
			// Concatenate the arrays
			newArray := make([]interface{}, len(targetArray)+len(sourceArray))
			copy(newArray, targetArray)
			copy(newArray[len(targetArray):], DeepCopy(sourceArray).([]interface{}))
			target[key] = newArray
			continue
		}

		// For other types, overwrite the value
		target[key] = DeepCopy(sourceValue)
	}

	return target
}

// GetPathValue gets a value from a nested structure using dot notation path
func GetPathValue(obj interface{}, path string) (interface{}, bool) {
	if obj == nil {
		return nil, false
	}

	if path == "" {
		return obj, true
	}

	// Split the path into segments
	segments := strings.Split(path, ".")
	current := obj
	var ok bool

	// Traverse the path
	for _, segment := range segments {
		// Process array indexing, which can be in format: users[0].name or users.0.name
		arrayIndexRegex := regexp.MustCompile(`^(.*?)\[(\d+)\]$`)
		matches := arrayIndexRegex.FindStringSubmatch(segment)

		if len(matches) == 3 {
			// Format: users[0]
			field := matches[1]
			indexStr := matches[2]

			// Get the object at the current segment if it's not empty
			if field != "" {
				var ok bool
				current, ok = getObjectField(current, field)
				if !ok {
					return nil, false
				}
			}

			// Parse the index
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return nil, false
			}

			// Handle the array
			current, ok = getArrayItem(current, index)
			if !ok {
				return nil, false
			}
		} else if indexNum, err := strconv.Atoi(segment); err == nil {
			// Format: users.0
			// Pure numeric segment means array index
			current, ok = getArrayItem(current, indexNum)
			if !ok {
				return nil, false
			}
		} else {
			// Regular field access
			var ok bool
			current, ok = getObjectField(current, segment)
			if !ok {
				return nil, false
			}
		}
	}

	return current, true
}

// getObjectField extracts a field from an object
func getObjectField(obj interface{}, field string) (interface{}, bool) {
	// Handle nil case
	if obj == nil {
		return nil, false
	}

	// Try to access as map[string]interface{}
	if objMap, ok := obj.(map[string]interface{}); ok {
		val, exists := objMap[field]
		return val, exists
	}

	// Try to access as map[string]any
	if objMap, ok := obj.(map[string]any); ok {
		val, exists := objMap[field]
		return val, exists
	}

	// Try to access as JSONObject
	if objMap, ok := obj.(JSONObject); ok {
		val, exists := objMap[field]
		return val, exists
	}

	// Try to marshal/unmarshal for struct types
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, false
	}

	var objMap map[string]interface{}
	if err := json.Unmarshal(data, &objMap); err != nil {
		return nil, false
	}

	val, exists := objMap[field]
	return val, exists
}

// getArrayItem extracts an item from an array by index
func getArrayItem(obj interface{}, index int) (interface{}, bool) {
	// Handle nil case
	if obj == nil {
		return nil, false
	}

	// Try to access as []interface{}
	if arr, ok := obj.([]interface{}); ok {
		if index < 0 || index >= len(arr) {
			return nil, false
		}
		return arr[index], true
	}

	// Try to access as []any
	if arr, ok := obj.([]any); ok {
		if index < 0 || index >= len(arr) {
			return nil, false
		}
		return arr[index], true
	}

	// Try to marshal/unmarshal for other array types
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, false
	}

	var arr []interface{}
	if err := json.Unmarshal(data, &arr); err != nil {
		return nil, false
	}

	if index < 0 || index >= len(arr) {
		return nil, false
	}

	return arr[index], true
}

// SetPathValue sets a value in a nested structure using dot notation path
func SetPathValue(obj map[string]interface{}, path string, value interface{}) (map[string]interface{}, bool) {
	if obj == nil {
		return nil, false
	}

	if path == "" {
		return obj, true
	}

	// Create a copy of the object to avoid modifying the original
	result := make(map[string]interface{})
	for k, v := range obj {
		result[k] = v
	}

	// Split the path into segments
	segments := strings.Split(path, ".")
	current := result

	// Traverse and create the path as needed
	for i, segment := range segments {
		// Process array indexing
		arrayIndexRegex := regexp.MustCompile(`^(.*?)\[(\d+)\]$`)
		matches := arrayIndexRegex.FindStringSubmatch(segment)

		if len(matches) == 3 {
			// Format: users[0]
			field := matches[1]
			indexStr := matches[2]

			// Parse the index
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return obj, false
			}

			// Get or create the array at the current segment
			var arr []interface{}
			if field != "" {
				tmp, exists := current[field]
				if !exists {
					// Create new array
					arr = []interface{}{}
					current[field] = arr
				} else if arrTmp, ok := tmp.([]interface{}); ok {
					// Use existing array
					arr = arrTmp
				} else {
					// Not an array, convert to array with single item
					arr = []interface{}{tmp}
					current[field] = arr
				}
			} else {
				return obj, false // Can't have empty field name
			}

			// Ensure array has sufficient capacity
			for len(arr) <= index {
				arr = append(arr, nil)
			}

			// Update the array in the current object
			current[field] = arr

			// If this is the last segment, set the value
			if i == len(segments)-1 {
				arr[index] = value
			} else {
				// Otherwise, create or get the next object
				if arr[index] == nil {
					arr[index] = make(map[string]interface{})
				}

				// Check if it's a map we can continue with
				nextCurrent, ok := arr[index].(map[string]interface{})
				if !ok {
					// Not a map, convert to map
					nextCurrent = make(map[string]interface{})
					arr[index] = nextCurrent
				}

				current = nextCurrent
			}
		} else if indexNum, err := strconv.Atoi(segment); err == nil && i > 0 {
			// Format: users.0
			// Pure numeric segment means array index, but not for the first segment
			prevSegment := segments[i-1]

			// Get the array from the previous segment
			tmp, exists := current[prevSegment]
			if !exists {
				// Create new array
				arr := []interface{}{}
				current[prevSegment] = arr

				// Ensure array has sufficient capacity
				for len(arr) <= indexNum {
					arr = append(arr, nil)
				}

				// If this is the last segment, set the value
				if i == len(segments)-1 {
					arr[indexNum] = value
				} else {
					// Create next object
					arr[indexNum] = make(map[string]interface{})

					// Move to next object
					current = arr[indexNum].(map[string]interface{})
				}
			} else if arr, ok := tmp.([]interface{}); ok {
				// Use existing array

				// Ensure array has sufficient capacity
				for len(arr) <= indexNum {
					arr = append(arr, nil)
				}
				current[prevSegment] = arr

				// If this is the last segment, set the value
				if i == len(segments)-1 {
					arr[indexNum] = value
				} else {
					// Create or get next object
					if arr[indexNum] == nil {
						arr[indexNum] = make(map[string]interface{})
					}

					// Check if it's a map
					nextCurrent, ok := arr[indexNum].(map[string]interface{})
					if !ok {
						// Not a map, convert to map
						nextCurrent = make(map[string]interface{})
						arr[indexNum] = nextCurrent
					}

					current = nextCurrent
				}
			} else {
				return obj, false // Not an array
			}
		} else {
			// Regular field access
			if i == len(segments)-1 {
				// Last segment, set the value
				current[segment] = value
			} else {
				// Not last segment, get or create the next object
				nextCurrent, exists := current[segment]
				if !exists {
					nextCurrent = make(map[string]interface{})
					current[segment] = nextCurrent
				}

				// Check if it's a map we can continue with
				nextMap, ok := nextCurrent.(map[string]interface{})
				if !ok {
					// Not a map, convert to map
					nextMap = make(map[string]interface{})
					current[segment] = nextMap
				}

				current = nextMap
			}
		}
	}

	return result, true
}

// MergeMaps merges multiple maps into a single map
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// FilterMap returns a new map containing only the specified keys
func FilterMap(m map[string]interface{}, keys []string) map[string]interface{} {
	result := make(map[string]interface{})

	for _, key := range keys {
		if val, exists := m[key]; exists {
			result[key] = val
		}
	}

	return result
}

// ExcludeFromMap returns a new map excluding the specified keys
func ExcludeFromMap(m map[string]interface{}, keys []string) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range m {
		exclude := false
		for _, key := range keys {
			if k == key {
				exclude = true
				break
			}
		}
		if !exclude {
			result[k] = v
		}
	}

	return result
}

// shallowCopy creates a shallow copy of a value
func shallowCopy(value interface{}) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(v))
		copy(result, v)
		return result
	default:
		return v
	}
}

// FormatJSONValue formats a JSON value as a string for display
func FormatJSONValue(value interface{}) string {
	if value == nil {
		return "null"
	}

	switch v := value.(type) {
	case string:
		return fmt.Sprintf("%q", v)
	case float64:
		if v == float64(int(v)) {
			return fmt.Sprintf("%d", int(v))
		}
		return fmt.Sprintf("%g", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		data, err := json.Marshal(value)
		if err != nil {
			return fmt.Sprintf("%v", value)
		}
		return string(data)
	}
}

// ParseJSONPath parses a JSON path and validates it
func ParseJSONPath(path string) ([]string, error) {
	if path == "" {
		return []string{}, nil
	}

	segments := strings.Split(path, ".")
	for i, segment := range segments {
		if segment == "" {
			return nil, fmt.Errorf("empty segment at position %d in path '%s'", i, path)
		}
	}

	return segments, nil
}
