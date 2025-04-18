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

package oldcore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

type ExecutionMode string

const (
	ExecutionModeLive  ExecutionMode = "LIVE"
	ExecutionModeTest  ExecutionMode = "TEST"
	ExecutionModeDebug ExecutionMode = "DEBUG"
)

func (ExecutionMode) SQLTypeName() string {
	return "execution_mode"
}

// Values returns a slice of all String values of the enum.
func (ExecutionMode) Values() []string {
	return []string{
		"LIVE",
		"TEST",
		"DEBUG",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j ExecutionMode) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j ExecutionMode) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("ExecutionMode(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern ExecutionMode(%d) instead.
func (_j ExecutionMode) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("ExecutionMode(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for ExecutionMode.
func (_j ExecutionMode) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ExecutionMode. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for ExecutionMode.
func (_j *ExecutionMode) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ExecutionMode cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionModeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionMode", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for ExecutionMode.
func (_j ExecutionMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for ExecutionMode.
func (_j *ExecutionMode) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ExecutionMode: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ExecutionMode cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionModeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionMode", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for ExecutionMode.
func (_j ExecutionMode) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ExecutionMode. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExecutionMode.
func (_j *ExecutionMode) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("ExecutionMode should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("ExecutionMode cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionModeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionMode", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for ExecutionMode.
func (_j *ExecutionMode) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ExecutionMode: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ExecutionMode cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionModeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionMode", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for ExecutionMode.
func (_j ExecutionMode) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ExecutionMode. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ExecutionMode.
func (_j *ExecutionMode) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ExecutionMode cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionModeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionMode", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for ExecutionMode.
func (_j ExecutionMode) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ExecutionMode. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ExecutionMode.
func (_j *ExecutionMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("ExecutionMode cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ExecutionModeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ExecutionMode", str)
	}
	return nil
}

// ExecutionModeFromString determines the enum value with an exact case match.
func ExecutionModeFromString(raw string) (ExecutionMode, bool) {
	v, ok := _ExecutionModeStringToValueMap[raw]
	if !ok {
		return ExecutionModeDebug, false
	}
	return v, true
}

// ExecutionModeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func ExecutionModeFromStringIgnoreCase(raw string) (ExecutionMode, bool) {
	v, ok := ExecutionModeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _ExecutionModeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_ExecutionModeStringToValueMap = map[string]ExecutionMode{
		"LIVE":  ExecutionModeLive,
		"TEST":  ExecutionModeTest,
		"DEBUG": ExecutionModeDebug,
	}
	_ExecutionModeLowerStringToValueMap = map[string]ExecutionMode{
		"LIVE":  ExecutionModeLive,
		"TEST":  ExecutionModeTest,
		"DEBUG": ExecutionModeDebug,
	}
)
