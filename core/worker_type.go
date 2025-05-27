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

type WorkerStatus string

const (
	WorkerStatusActive      WorkerStatus = "ACTIVE"
	WorkerStatusInActive    WorkerStatus = "INACTIVE"
	WorkerStatusFaulted     WorkerStatus = "FAULTED"
	WorkerStatusMaintenance WorkerStatus = "MAINTENANCE"
)

func (WorkerStatus) SQLTypeName() string {
	return "worker_status"
}

// Values returns a slice of all String values of the enum.
func (WorkerStatus) Values() []string {
	return []string{
		"ACTIVE", "INACTIVE", "FAULTED", "MAINTENANCE",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j WorkerStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j WorkerStatus) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("WorkerStatus(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern WorkerStatus(%d) instead.
func (_j WorkerStatus) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("WorkerStatus(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for WorkerStatus.
func (_j WorkerStatus) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as WorkerStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for WorkerStatus.
func (_j *WorkerStatus) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("WorkerStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkerStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a WorkerStatus", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for WorkerStatus.
func (_j WorkerStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for WorkerStatus.
func (_j *WorkerStatus) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of WorkerStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("WorkerStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkerStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a WorkerStatus", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for WorkerStatus.
func (_j WorkerStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as WorkerStatus. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for WorkerStatus.
func (_j *WorkerStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("WorkerStatus should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("WorkerStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkerStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a WorkerStatus", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for WorkerStatus.
func (_j *WorkerStatus) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of WorkerStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("WorkerStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkerStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a WorkerStatus", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for WorkerStatus.
func (_j WorkerStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as WorkerStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for WorkerStatus.
func (_j *WorkerStatus) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("WorkerStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkerStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a WorkerStatus", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for WorkerStatus.
func (_j WorkerStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as WorkerStatus. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for WorkerStatus.
func (_j *WorkerStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("WorkerStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkerStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a WorkerStatus", str)
	}
	return nil
}

// WorkerStatusFromString determines the enum value with an exact case match.
func WorkerStatusFromString(raw string) (WorkerStatus, bool) {
	v, ok := _WorkerStatusStringToValueMap[raw]
	if !ok {
		return WorkerStatusActive, false
	}
	return v, true
}

// WorkerStatusFromStringIgnoreCase determines the enum value with a case-insensitive match.
func WorkerStatusFromStringIgnoreCase(raw string) (WorkerStatus, bool) {
	v, ok := WorkerStatusFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _WorkerStatusLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_WorkerStatusStringToValueMap = map[string]WorkerStatus{
		"ACTIVE":      WorkerStatusActive,
		"INACTIVE":    WorkerStatusInActive,
		"FAULTED":     WorkerStatusFaulted,
		"MAINTENANCE": WorkerStatusMaintenance,
	}
	_WorkerStatusLowerStringToValueMap = map[string]WorkerStatus{
		"ACTIVE":      WorkerStatusActive,
		"INACTIVE":    WorkerStatusInActive,
		"FAULTED":     WorkerStatusFaulted,
		"MAINTENANCE": WorkerStatusMaintenance,
	}
)
