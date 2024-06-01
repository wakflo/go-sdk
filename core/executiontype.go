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
	"io"
	"slices"
	"strconv"
)

type ExecutionType string

const (
	Begin  ExecutionType = "BEGIN"
	Resume ExecutionType = "RESUME"
)

func (ExecutionType) SQLTypeName() string {
	return "trigger_hook_type"
}

// Values returns a slice of all String values of the enum.
func (ExecutionType) Values() []string {
	return []string{
		"BEGIN",
		"RESUME",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j ExecutionType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j ExecutionType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("ExecutionType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern ExecutionType(%d) instead.
func (_j ExecutionType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("ExecutionType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for ExecutionType.
func (_j ExecutionType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ExecutionType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for ExecutionType.
func (_j *ExecutionType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for ExecutionType.
func (_j ExecutionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for ExecutionType.
func (_j *ExecutionType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ExecutionType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for ExecutionType.
func (_j ExecutionType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ExecutionType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExecutionType.
func (_j *ExecutionType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("ExecutionType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("ExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for ExecutionType.
func (_j *ExecutionType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ExecutionType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for ExecutionType.
func (_j ExecutionType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ExecutionType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ExecutionType.
func (_j *ExecutionType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for ExecutionType.
func (_j ExecutionType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ExecutionType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ExecutionType.
func (_j *ExecutionType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("ExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionType", str)
	}
	return nil
}

// ExecutionTypeFromString determines the enum value with an exact case match.
func ExecutionTypeFromString(raw string) (ExecutionType, bool) {
	v, ok := _ExecutionTypeStringToValueMap[raw]
	if !ok {
		return Begin, false
	}
	return v, true
}

// ExecutionTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func ExecutionTypeFromStringIgnoreCase(raw string) (ExecutionType, bool) {
	v, ok := ExecutionTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _ExecutionTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_ExecutionTypeStringToValueMap = map[string]ExecutionType{
		"BEGIN":  Begin,
		"RESUME": Resume,
	}
	_ExecutionTypeLowerStringToValueMap = map[string]ExecutionType{
		"BEGIN":  Begin,
		"RESUME": Resume,
	}
)
