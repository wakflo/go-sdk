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
	// StepRunStatusPending indicates a step is queued to run but hasn't started
	StepRunStatusPending StepRunStatus = "PENDING"

	// StepRunStatusPaused indicates a step is paused and waiting for manual continuation
	StepRunStatusPaused StepRunStatus = "PAUSED"

	// StepRunStatusRunning indicates a step is currently executing
	StepRunStatusRunning StepRunStatus = "RUNNING"

	// StepRunStatusSucceeded indicates a step has completed successfully
	StepRunStatusSucceeded StepRunStatus = "SUCCEEDED"

	// StepRunStatusFailed indicates a step has failed
	StepRunStatusFailed StepRunStatus = "FAILED"

	// StepRunStatusCancelled indicates a step was manually cancelled
	StepRunStatusCancelled StepRunStatus = "CANCELLED"

	// StepRunStatusSkipped indicates a step was skipped due to conditions or branching
	StepRunStatusSkipped StepRunStatus = "SKIPPED"

	// StepRunStatusTimeout indicates a step exceeded its allowed execution time
	StepRunStatusTimeout StepRunStatus = "TIMEOUT"

	// StepRunStatusWaiting indicates a step is waiting for an external event or condition
	StepRunStatusWaiting StepRunStatus = "WAITING"

	// StepRunStatusBlocked indicates a step is blocked by a dependency or condition
	StepRunStatusBlocked StepRunStatus = "BLOCKED"

	// StepRunStatusApproved indicates a step has been manually approved
	StepRunStatusApproved StepRunStatus = "APPROVED"

	// StepRunStatusRejected indicates a step was manually rejected
	StepRunStatusRejected StepRunStatus = "REJECTED"

	StepRunStatusRetrying StepRunStatus = "RETRY"
)

// SQLTypeName returns the SQL type name for serialization
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
		"SKIPPED",
		"TIMEOUT",
		"WAITING",
		"BLOCKED",
		"APPROVED",
		"REJECTED",
		"RETRY",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j StepRunStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j StepRunStatus) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("StepRunStatus(%v) is %w", _j, ErrNoValidEnum)
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
		return nil, fmt.Errorf("cannot marshal value %q as StepRunStatus: %w", _j, err)
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
		return fmt.Errorf("value %q does not represent a StepRunStatus", str)
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
		return fmt.Errorf("invalid value of StepRunStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepRunStatus", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for StepRunStatus.
func (_j StepRunStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StepRunStatus: %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for StepRunStatus.
func (_j *StepRunStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("StepRunStatus should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepRunStatus", str)
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
		return fmt.Errorf("invalid value of StepRunStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("StepRunStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = StepRunStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a StepRunStatus", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for StepRunStatus.
func (_j StepRunStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StepRunStatus: %w", _j, err)
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
		return fmt.Errorf("value %q does not represent a StepRunStatus", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for StepRunStatus.
func (_j StepRunStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as StepRunStatus: %w", _j, err)
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
		return fmt.Errorf("value %q does not represent a StepRunStatus", str)
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

// Maps for looking up enum values from strings
var (
	_StepRunStatusStringToValueMap = map[string]StepRunStatus{
		"PENDING":   StepRunStatusPending,
		"PAUSED":    StepRunStatusPaused,
		"RUNNING":   StepRunStatusRunning,
		"SUCCEEDED": StepRunStatusSucceeded,
		"FAILED":    StepRunStatusFailed,
		"CANCELLED": StepRunStatusCancelled,
		"SKIPPED":   StepRunStatusSkipped,
		"TIMEOUT":   StepRunStatusTimeout,
		"WAITING":   StepRunStatusWaiting,
		"BLOCKED":   StepRunStatusBlocked,
		"APPROVED":  StepRunStatusApproved,
		"REJECTED":  StepRunStatusRejected,
		"RETRY":     StepRunStatusRetrying,
	}
	_StepRunStatusLowerStringToValueMap = map[string]StepRunStatus{
		"pending":   StepRunStatusPending,
		"paused":    StepRunStatusPaused,
		"running":   StepRunStatusRunning,
		"succeeded": StepRunStatusSucceeded,
		"failed":    StepRunStatusFailed,
		"cancelled": StepRunStatusCancelled,
		"skipped":   StepRunStatusSkipped,
		"timeout":   StepRunStatusTimeout,
		"waiting":   StepRunStatusWaiting,
		"blocked":   StepRunStatusBlocked,
		"approved":  StepRunStatusApproved,
		"rejected":  StepRunStatusRejected,
		"retry":     StepRunStatusRetrying,
	}
)

// IsActive returns true if the status indicates the step is still active
func (s StepRunStatus) IsActive() bool {
	return s == StepRunStatusPending || s == StepRunStatusRunning ||
		s == StepRunStatusPaused || s == StepRunStatusWaiting || s == StepRunStatusRetrying ||
		s == StepRunStatusBlocked
}

// IsComplete returns true if the status indicates the step has finished
func (s StepRunStatus) IsComplete() bool {
	return s == StepRunStatusSucceeded || s == StepRunStatusFailed ||
		s == StepRunStatusCancelled || s == StepRunStatusSkipped ||
		s == StepRunStatusTimeout || s == StepRunStatusRejected
}

// IsSuccessful returns true if the status indicates successful completion
func (s StepRunStatus) IsSuccessful() bool {
	return s == StepRunStatusSucceeded || s == StepRunStatusApproved
}

func (s StepRunStatus) IsFailed() bool {
	return s == StepRunStatusFailed || s == StepRunStatusRejected
}

func (s StepRunStatus) IsTerminalStatus() bool {
	return s == StepRunStatusSucceeded || s == StepRunStatusFailed || s == StepRunStatusCancelled || s == StepRunStatusSkipped || s == StepRunStatusTimeout || s == StepRunStatusRejected
}

// StatusColor returns a color associated with this status for UI display
func (s StepRunStatus) StatusColor() string {
	switch s {
	case StepRunStatusSucceeded, StepRunStatusApproved:
		return "green"
	case StepRunStatusFailed, StepRunStatusRejected, StepRunStatusTimeout:
		return "red"
	case StepRunStatusRunning, StepRunStatusPending, StepRunStatusWaiting, StepRunStatusRetrying:
		return "blue"
	case StepRunStatusPaused, StepRunStatusBlocked:
		return "yellow"
	case StepRunStatusCancelled, StepRunStatusSkipped:
		return "gray"
	default:
		return "gray"
	}
}

// StatusIcon returns an icon associated with this status for UI display
func (s StepRunStatus) StatusIcon() string {
	switch s {
	case StepRunStatusSucceeded, StepRunStatusApproved:
		return "check-circle"
	case StepRunStatusFailed, StepRunStatusRejected:
		return "x-circle"
	case StepRunStatusRunning:
		return "play-circle"
	case StepRunStatusPending, StepRunStatusWaiting:
		return "clock"
	case StepRunStatusPaused:
		return "pause-circle"
	case StepRunStatusBlocked:
		return "lock"
	case StepRunStatusCancelled:
		return "slash"
	case StepRunStatusSkipped:
		return "skip-forward"
	case StepRunStatusTimeout:
		return "alert-circle"
	case StepRunStatusRetrying:
		return "alert-circle"
	default:
		return "circle"
	}
}
