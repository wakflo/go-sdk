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

// StepRunStatus represents the status of a step run.
type StepRunStatus string

// Enum values for StepRunStatus.
const (
	StepRunStatusPending   StepRunStatus = "PENDING"
	StepRunStatusPaused    StepRunStatus = "PAUSED"
	StepRunStatusRunning   StepRunStatus = "RUNNING"
	StepRunStatusSucceeded StepRunStatus = "SUCCEEDED"
	StepRunStatusFailed    StepRunStatus = "FAILED"
	StepRunStatusCancelled StepRunStatus = "CANCELLED"
)

func (StepRunStatus) SQLTypeName() string {
	return "step_run_status"
}

// Values returns a slice of all String values of the enum.
func (StepRunStatus) Values() []string {
	return []string{
		"PENDING",
		"PAUSED",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j StepRunStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j StepRunStatus) Validate() error {
	if !_j.IsValid() {
		return errors.New(fmt.Sprintf("StepRunStatus(%v) is %v", _j, ErrNoValidEnum))
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern StepRunStatus(%d) instead.
func (_j StepRunStatus) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("StepRunStatus(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for StepRunStatus.
func (_j StepRunStatus) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as StepRunStatus. %v", _j, err))
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for StepRunStatus.
func (_j *StepRunStatus) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a StepRunStatus", str))
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for StepRunStatus.
func (_j StepRunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for StepRunStatus.
func (_j *StepRunStatus) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return errors.New(fmt.Sprintf("invalid value of StepRunStatus: %[1]T(%[1]v)", value))
	}
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a StepRunStatus", str))
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for StepRunStatus.
func (_j StepRunStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as StepRunStatus. %v", _j, err))
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for StepRunStatus.
func (_j *StepRunStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return errors.New(fmt.Sprintf("StepRunStatus should be a string, got %q", data))
	}
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a StepRunStatus", str))
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for StepRunStatus.
func (_j *StepRunStatus) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return errors.New(fmt.Sprintf("invalid value of StepRunStatus: %[1]T(%[1]v)", value))
	}
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a StepRunStatus", str))
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for StepRunStatus.
func (_j StepRunStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as StepRunStatus. %v", _j, err))
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for StepRunStatus.
func (_j *StepRunStatus) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a StepRunStatus", str))
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for StepRunStatus.
func (_j StepRunStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as StepRunStatus. %v", _j, err))
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for StepRunStatus.
func (_j *StepRunStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a StepRunStatus", str))
	}
	return nil
}

// StepRunStatusFromString determines the enum value with an exact case match.
func StepRunStatusFromString(raw string) (StepRunStatus, bool) {
	v, ok := _StepRunStatusStringToValueMap[raw]
	if !ok {
		return StepRunStatusPending, false
	}
	return v, true
}

// StepRunStatusFromStringIgnoreCase determines the enum value with a case-insensitive match.
func StepRunStatusFromStringIgnoreCase(raw string) (StepRunStatus, bool) {
	v, ok := StepRunStatusFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _StepRunStatusLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_StepRunStatusStringToValueMap = map[string]StepRunStatus{
		"PENDING":   StepRunStatusPending,
		"PAUSED":    StepRunStatusPaused,
		"RUNNING":   StepRunStatusRunning,
		"SUCCEEDED": StepRunStatusSucceeded,
		"FAILED":    StepRunStatusFailed,
		"CANCELLED": StepRunStatusCancelled,
	}
	_StepRunStatusLowerStringToValueMap = map[string]StepRunStatus{
		"PENDING":   StepRunStatusPending,
		"PAUSED":    StepRunStatusPaused,
		"RUNNING":   StepRunStatusRunning,
		"SUCCEEDED": StepRunStatusSucceeded,
		"FAILED":    StepRunStatusFailed,
		"CANCELLED": StepRunStatusCancelled,
	}
)
