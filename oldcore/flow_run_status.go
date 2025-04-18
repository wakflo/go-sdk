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

// FlowRunStatus represents the status of a flow run.
type FlowRunStatus string

// Enum values for FlowRunStatus.
const (
	FlowRunStatusReady         FlowRunStatus = "READY"
	FlowRunStatusQueued        FlowRunStatus = "QUEUED"
	FlowRunStatusRunning       FlowRunStatus = "RUNNING"
	FlowRunStatusSucceeded     FlowRunStatus = "SUCCEEDED"
	FlowRunStatusFailed        FlowRunStatus = "FAILED"
	FlowRunStatusPaused        FlowRunStatus = "PAUSED"
	FlowRunStatusInternalError FlowRunStatus = "INTERNAL_ERROR"
	FlowRunStatusQuotaExceeded FlowRunStatus = "QUOTA_EXCEEDED"
	FlowRunStatusTimeout       FlowRunStatus = "TIMEOUT"
	FlowRunStatusTerminated    FlowRunStatus = "TERMINATED"
)

func (FlowRunStatus) SQLTypeName() string {
	return "workflow_run_status"
}

// Values returns a slice of all String values of the enum.
func (FlowRunStatus) Values() []string {
	return []string{
		"READY",
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"PAUSED",
		"INTERNAL_ERROR",
		"QUOTA_EXCEEDED",
		"TIMEOUT",
		"TERMINATED",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j FlowRunStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j FlowRunStatus) Validate() error {
	if !_j.IsValid() {
		return errors.New(fmt.Sprintf("FlowRunStatus(%v) is %v", _j, ErrNoValidEnum))
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern FlowRunStatus(%d) instead.
func (_j FlowRunStatus) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("FlowRunStatus(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for FlowRunStatus.
func (_j FlowRunStatus) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as FlowRunStatus. %v", _j, err))
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for FlowRunStatus.
func (_j *FlowRunStatus) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("FlowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a FlowRunStatus", str))
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for FlowRunStatus.
func (_j FlowRunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for FlowRunStatus.
func (_j *FlowRunStatus) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return errors.New(fmt.Sprintf("invalid value of FlowRunStatus: %[1]T(%[1]v)", value))
	}
	if len(str) == 0 {
		return errors.New("FlowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a FlowRunStatus", str))
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for FlowRunStatus.
func (_j FlowRunStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as FlowRunStatus. %v", _j, err))
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FlowRunStatus.
func (_j *FlowRunStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return errors.New(fmt.Sprintf("FlowRunStatus should be a string, got %q", data))
	}
	if len(str) == 0 {
		return errors.New("FlowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a FlowRunStatus", str))
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for FlowRunStatus.
func (_j *FlowRunStatus) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return errors.New(fmt.Sprintf("invalid value of FlowRunStatus: %[1]T(%[1]v)", value))
	}
	if len(str) == 0 {
		return errors.New("FlowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a FlowRunStatus", str))
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for FlowRunStatus.
func (_j FlowRunStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as FlowRunStatus. %v", _j, err))
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for FlowRunStatus.
func (_j *FlowRunStatus) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("FlowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a FlowRunStatus", str))
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for FlowRunStatus.
func (_j FlowRunStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as FlowRunStatus. %v", _j, err))
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for FlowRunStatus.
func (_j *FlowRunStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("FlowRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowRunStatusFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a FlowRunStatus", str))
	}
	return nil
}

// FlowRunStatusFromString determines the enum value with an exact case match.
func FlowRunStatusFromString(raw string) (FlowRunStatus, bool) {
	v, ok := _FlowRunStatusStringToValueMap[raw]
	if !ok {
		return FlowRunStatusReady, false
	}
	return v, true
}

// FlowRunStatusFromStringIgnoreCase determines the enum value with a case-insensitive match.
func FlowRunStatusFromStringIgnoreCase(raw string) (FlowRunStatus, bool) {
	v, ok := FlowRunStatusFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _FlowRunStatusLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_FlowRunStatusStringToValueMap = map[string]FlowRunStatus{
		"READY":          FlowRunStatusReady,
		"QUEUED":         FlowRunStatusQueued,
		"RUNNING":        FlowRunStatusRunning,
		"SUCCEEDED":      FlowRunStatusSucceeded,
		"FAILED":         FlowRunStatusFailed,
		"PAUSED":         FlowRunStatusPaused,
		"INTERNAL_ERROR": FlowRunStatusInternalError,
		"QUOTA_EXCEEDED": FlowRunStatusQuotaExceeded,
		"TIMEOUT":        FlowRunStatusTimeout,
		"TERMINATED":     FlowRunStatusTerminated,
	}
	_FlowRunStatusLowerStringToValueMap = map[string]FlowRunStatus{
		"READY":          FlowRunStatusReady,
		"QUEUED":         FlowRunStatusQueued,
		"RUNNING":        FlowRunStatusRunning,
		"SUCCEEDED":      FlowRunStatusSucceeded,
		"FAILED":         FlowRunStatusFailed,
		"PAUSED":         FlowRunStatusPaused,
		"INTERNAL_ERROR": FlowRunStatusInternalError,
		"QUOTA_EXCEEDED": FlowRunStatusQuotaExceeded,
		"TIMEOUT":        FlowRunStatusTimeout,
		"TERMINATED":     FlowRunStatusTerminated,
	}
)
