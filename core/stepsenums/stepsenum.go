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

package stepsenum

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

var (
	ErrNoValidEnum = errors.New("not a valid enum")
)

type StepType string

const (
	Normal    StepType = "normal"
	Branch             = "branch"
	Boolean            = "boolean"
	Loop               = "loop"
	Condition          = "condition"
	Start              = "start"
	End                = "end"
)

func (StepType) SqlTypeName() string {
	return "connector_step_type"
}

// Values returns a slice of all String values of the enum.
func (StepType) Values() []string {
	return []string{
		"normal",
		"branch",
		"boolean",
		"loop",
		"condition",
		"start",
		"end",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j StepType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j StepType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("StepType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern StepType(%d) instead.
func (_j StepType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("StepType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for StepType.
func (_j StepType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StepType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for StepType.
func (_j *StepType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return fmt.Errorf("StepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for StepType.
func (_j StepType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for StepType.
func (_j *StepType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of StepType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return fmt.Errorf("StepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for StepType.
func (_j StepType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StepType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for StepType.
func (_j *StepType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("StepType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return fmt.Errorf("StepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for StepType.
func (_j *StepType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of StepType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return fmt.Errorf("StepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for StepType.
func (_j StepType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StepType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for StepType.
func (_j *StepType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return fmt.Errorf("StepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for StepType.
func (_j StepType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StepType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for StepType.
func (_j *StepType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return fmt.Errorf("StepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepType", str)
	}
	return nil
}

// StepTypeFromString determines the enum value with an exact case match.
func StepTypeFromString(raw string) (StepType, bool) {
	v, ok := _StepTypeStringToValueMap[raw]
	if !ok {
		return Normal, false
	}
	return v, true
}

// StepTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func StepTypeFromStringIgnoreCase(raw string) (StepType, bool) {
	v, ok := StepTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _StepTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_StepTypeStringToValueMap = map[string]StepType{
		"branch":    Branch,
		"boolean":   Boolean,
		"normal":    Normal,
		"loop":      Loop,
		"condition": Condition,
		"start":     Start,
		"end":       End,
	}
	_StepTypeLowerStringToValueMap = map[string]StepType{
		"branch":    Branch,
		"boolean":   Boolean,
		"normal":    Normal,
		"loop":      Loop,
		"condition": Condition,
		"start":     Start,
		"end":       End,
	}
)
