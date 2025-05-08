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

type StateErrorReason string

const (
	StateErrorReasonUnknown  StateErrorReason = "UNKNOWN"
	StateErrorReasonTimeout  StateErrorReason = "TIMEOUT"
	StateErrorReasonError    StateErrorReason = "ERROR"
	StateErrorReasonSkipped  StateErrorReason = "SKIPPED"
	StateErrorReasonAborted  StateErrorReason = "ABORTED"
	StateErrorReasonCanceled StateErrorReason = "CANCELED"
)

func (StateErrorReason) SQLTypeName() string {
	return "state_error_reason"
}

// Values returns a slice of all String values of the enum.
func (StateErrorReason) Values() []string {
	return []string{
		"UNKNOWN",
		"TIMEOUT",
		"ERROR",
		"SKIPPED",
		"ABORTED",
		"CANCELED",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j StateErrorReason) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j StateErrorReason) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("StateErrorReason(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern StateErrorReason(%d) instead.
func (_j StateErrorReason) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("StateErrorReason(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for StateErrorReason.
func (_j StateErrorReason) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StateErrorReason. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for StateErrorReason.
func (_j *StateErrorReason) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("StateErrorReason cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StateErrorReasonFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StateErrorReason", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for StateErrorReason.
func (_j StateErrorReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for StateErrorReason.
func (_j *StateErrorReason) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of StateErrorReason: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("StateErrorReason cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StateErrorReasonFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StateErrorReason", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for StateErrorReason.
func (_j StateErrorReason) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StateErrorReason. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for StateErrorReason.
func (_j *StateErrorReason) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("StateErrorReason should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("StateErrorReason cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StateErrorReasonFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StateErrorReason", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for StateErrorReason.
func (_j *StateErrorReason) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of StateErrorReason: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("StateErrorReason cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StateErrorReasonFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StateErrorReason", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for StateErrorReason.
func (_j StateErrorReason) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StateErrorReason. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for StateErrorReason.
func (_j *StateErrorReason) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("StateErrorReason cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StateErrorReasonFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StateErrorReason", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for StateErrorReason.
func (_j StateErrorReason) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StateErrorReason. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for StateErrorReason.
func (_j *StateErrorReason) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("StateErrorReason cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StateErrorReasonFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StateErrorReason", str)
	}
	return nil
}

// StateErrorReasonFromString determines the enum value with an exact case match.
func StateErrorReasonFromString(raw string) (StateErrorReason, bool) {
	v, ok := _StateErrorReasonStringToValueMap[raw]
	if !ok {
		return StateErrorReasonUnknown, false
	}
	return v, true
}

// StateErrorReasonFromStringIgnoreCase determines the enum value with a case-insensitive match.
func StateErrorReasonFromStringIgnoreCase(raw string) (StateErrorReason, bool) {
	v, ok := StateErrorReasonFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _StateErrorReasonLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_StateErrorReasonStringToValueMap = map[string]StateErrorReason{
		"UNKNOWN":  StateErrorReasonUnknown,
		"TIMEOUT":  StateErrorReasonTimeout,
		"ERROR":    StateErrorReasonError,
		"SKIPPED":  StateErrorReasonSkipped,
		"ABORTED":  StateErrorReasonAborted,
		"CANCELED": StateErrorReasonCanceled,
	}
	_StateErrorReasonLowerStringToValueMap = map[string]StateErrorReason{
		"UNKNOWN":  StateErrorReasonUnknown,
		"TIMEOUT":  StateErrorReasonTimeout,
		"ERROR":    StateErrorReasonError,
		"SKIPPED":  StateErrorReasonSkipped,
		"ABORTED":  StateErrorReasonAborted,
		"CANCELED": StateErrorReasonCanceled,
	}
)
