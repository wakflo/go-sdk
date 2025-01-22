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

type FlowVersionState string

const (
	FlowVersionStateLocked FlowVersionState = "LOCKED"
	FlowVersionStateDraft  FlowVersionState = "DRAFT"
)

func (FlowVersionState) SQLTypeName() string {
	return "flow_version_state"
}

// Values returns a slice of all String values of the enum.
func (FlowVersionState) Values() []string {
	return []string{
		"LOCKED",
		"DRAFT",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j FlowVersionState) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j FlowVersionState) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("FlowVersionState(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern FlowVersionState(%d) instead.
func (_j FlowVersionState) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("FlowVersionState(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for FlowVersionState.
func (_j FlowVersionState) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowVersionState. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for FlowVersionState.
func (_j *FlowVersionState) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("FlowVersionState cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowVersionStateFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowVersionState", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for FlowVersionState.
func (_j FlowVersionState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for FlowVersionState.
func (_j *FlowVersionState) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of FlowVersionState: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("FlowVersionState cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowVersionStateFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowVersionState", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for FlowVersionState.
func (_j FlowVersionState) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowVersionState. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FlowVersionState.
func (_j *FlowVersionState) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("FlowVersionState should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("FlowVersionState cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowVersionStateFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowVersionState", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for FlowVersionState.
func (_j *FlowVersionState) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of FlowVersionState: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("FlowVersionState cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowVersionStateFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowVersionState", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for FlowVersionState.
func (_j FlowVersionState) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowVersionState. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for FlowVersionState.
func (_j *FlowVersionState) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("FlowVersionState cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowVersionStateFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowVersionState", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for FlowVersionState.
func (_j FlowVersionState) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowVersionState. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for FlowVersionState.
func (_j *FlowVersionState) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("FlowVersionState cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowVersionStateFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowVersionState", str)
	}
	return nil
}

// FlowVersionStateFromString determines the enum value with an exact case match.
func FlowVersionStateFromString(raw string) (FlowVersionState, bool) {
	v, ok := _FlowVersionStateStringToValueMap[raw]
	if !ok {
		return FlowVersionStateDraft, false
	}
	return v, true
}

// FlowVersionStateFromStringIgnoreCase determines the enum value with a case-insensitive match.
func FlowVersionStateFromStringIgnoreCase(raw string) (FlowVersionState, bool) {
	v, ok := FlowVersionStateFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _FlowVersionStateLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_FlowVersionStateStringToValueMap = map[string]FlowVersionState{
		"LOCKED": FlowVersionStateLocked,
		"DRAFT":  FlowVersionStateDraft,
	}
	_FlowVersionStateLowerStringToValueMap = map[string]FlowVersionState{
		"LOCKED": FlowVersionStateLocked,
		"DRAFT":  FlowVersionStateDraft,
	}
)
