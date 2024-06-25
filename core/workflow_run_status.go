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

// WorkflowRunStatus represents the status of a workflow run.
type WorkflowRunStatus string

// Enum values for WorkflowRunStatus.
const (
	WorkflowRunStatusPending   WorkflowRunStatus = "PENDING"
	WorkflowRunStatusQueued    WorkflowRunStatus = "QUEUED"
	WorkflowRunStatusRunning   WorkflowRunStatus = "RUNNING"
	WorkflowRunStatusSucceeded WorkflowRunStatus = "SUCCEEDED"
	WorkflowRunStatusFailed    WorkflowRunStatus = "FAILED"
)

func (WorkflowRunStatus) SQLTypeName() string {
	return "workflow_run_status"
}

// Values returns a slice of all String values of the enum.
func (WorkflowRunStatus) Values() []string {
	return []string{
		"PENDING",
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j WorkflowRunStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j WorkflowRunStatus) Validate() error {
	if !_j.IsValid() {
		return errors.New(fmt.Sprintf("WorkflowRunStatus(%v) is %v", _j, ErrNoValidEnum))
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern WorkflowRunStatus(%d) instead.
func (_j WorkflowRunStatus) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("WorkflowRunStatus(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for WorkflowRunStatus.
func (_j WorkflowRunStatus) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as WorkflowRunStatus. %v", _j, err))
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for WorkflowRunStatus.
func (_j *WorkflowRunStatus) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("WorkflowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkflowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a WorkflowRunStatus", str))
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for WorkflowRunStatus.
func (_j WorkflowRunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for WorkflowRunStatus.
func (_j *WorkflowRunStatus) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return errors.New(fmt.Sprintf("invalid value of WorkflowRunStatus: %[1]T(%[1]v)", value))
	}
	if len(str) == 0 {
		return errors.New("WorkflowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkflowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a WorkflowRunStatus", str))
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for WorkflowRunStatus.
func (_j WorkflowRunStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as WorkflowRunStatus. %v", _j, err))
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for WorkflowRunStatus.
func (_j *WorkflowRunStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return errors.New(fmt.Sprintf("WorkflowRunStatus should be a string, got %q", data))
	}
	if len(str) == 0 {
		return errors.New("WorkflowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkflowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a WorkflowRunStatus", str))
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for WorkflowRunStatus.
func (_j *WorkflowRunStatus) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return errors.New(fmt.Sprintf("invalid value of WorkflowRunStatus: %[1]T(%[1]v)", value))
	}
	if len(str) == 0 {
		return errors.New("WorkflowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkflowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a WorkflowRunStatus", str))
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for WorkflowRunStatus.
func (_j WorkflowRunStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as WorkflowRunStatus. %v", _j, err))
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for WorkflowRunStatus.
func (_j *WorkflowRunStatus) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("WorkflowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkflowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a WorkflowRunStatus", str))
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for WorkflowRunStatus.
func (_j WorkflowRunStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as WorkflowRunStatus. %v", _j, err))
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for WorkflowRunStatus.
func (_j *WorkflowRunStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("WorkflowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = WorkflowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a WorkflowRunStatus", str))
	}
	return nil
}

// WorkflowRunStatusFromString determines the enum value with an exact case match.
func WorkflowRunStatusFromString(raw string) (WorkflowRunStatus, bool) {
	v, ok := _WorkflowRunStatusStringToValueMap[raw]
	if !ok {
		return WorkflowRunStatusPending, false
	}
	return v, true
}

// WorkflowRunStatusFromStringIgnoreCase determines the enum value with a case-insensitive match.
func WorkflowRunStatusFromStringIgnoreCase(raw string) (WorkflowRunStatus, bool) {
	v, ok := WorkflowRunStatusFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _WorkflowRunStatusLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_WorkflowRunStatusStringToValueMap = map[string]WorkflowRunStatus{
		"PENDING":   WorkflowRunStatusPending,
		"QUEUED":    WorkflowRunStatusQueued,
		"RUNNING":   WorkflowRunStatusRunning,
		"SUCCEEDED": WorkflowRunStatusSucceeded,
		"FAILED":    WorkflowRunStatusFailed,
	}
	_WorkflowRunStatusLowerStringToValueMap = map[string]WorkflowRunStatus{
		"PENDING":   WorkflowRunStatusPending,
		"QUEUED":    WorkflowRunStatusQueued,
		"RUNNING":   WorkflowRunStatusRunning,
		"SUCCEEDED": WorkflowRunStatusSucceeded,
		"FAILED":    WorkflowRunStatusFailed,
	}
)
