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

type FlowStatus string

const (
	FlowStatusEnabled  FlowStatus = "ENABLED"
	FlowStatusDisabled FlowStatus = "DISABLED"
)

func (FlowStatus) SQLTypeName() string {
	return "flow_status"
}

// Values returns a slice of all String values of the enum.
func (FlowStatus) Values() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j FlowStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j FlowStatus) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("FlowStatus(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern FlowStatus(%d) instead.
func (_j FlowStatus) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("FlowStatus(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for FlowStatus.
func (_j FlowStatus) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for FlowStatus.
func (_j *FlowStatus) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("FlowStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStatus", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for FlowStatus.
func (_j FlowStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for FlowStatus.
func (_j *FlowStatus) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of FlowStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("FlowStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStatus", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for FlowStatus.
func (_j FlowStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowStatus. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FlowStatus.
func (_j *FlowStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("FlowStatus should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("FlowStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStatus", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for FlowStatus.
func (_j *FlowStatus) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of FlowStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("FlowStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStatus", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for FlowStatus.
func (_j FlowStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for FlowStatus.
func (_j *FlowStatus) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("FlowStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStatus", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for FlowStatus.
func (_j FlowStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowStatus. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for FlowStatus.
func (_j *FlowStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("FlowStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStatus", str)
	}
	return nil
}

// FlowStatusFromString determines the enum value with an exact case match.
func FlowStatusFromString(raw string) (FlowStatus, bool) {
	v, ok := _FlowStatusStringToValueMap[raw]
	if !ok {
		return FlowStatusDisabled, false
	}
	return v, true
}

// FlowStatusFromStringIgnoreCase determines the enum value with a case-insensitive match.
func FlowStatusFromStringIgnoreCase(raw string) (FlowStatus, bool) {
	v, ok := FlowStatusFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _FlowStatusLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_FlowStatusStringToValueMap = map[string]FlowStatus{
		"ENABLED":  FlowStatusEnabled,
		"DISABLED": FlowStatusDisabled,
	}
	_FlowStatusLowerStringToValueMap = map[string]FlowStatus{
		"ENABLED":  FlowStatusEnabled,
		"DISABLED": FlowStatusDisabled,
	}
)
