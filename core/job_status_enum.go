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

type JobStatus string

const (
	JobStatusQueued     JobStatus = "queued"
	JobStatusRunning    JobStatus = "running"
	JobStatusCompleted  JobStatus = "completed"
	JobStatusCancelling JobStatus = "cancelling"
	JobStatusCanceled   JobStatus = "canceled"
	JobStatusFailed     JobStatus = "failed"
	JobStatusPaused     JobStatus = "paused"
)

func (JobStatus) SQLTypeName() string {
	return "job_status"
}

// Values returns a slice of all String values of the enum.
func (JobStatus) Values() []string {
	return []string{
		"queued",
		"running",
		"completed",
		"cancelling",
		"canceled",
		"failed",
		"paused",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j JobStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j JobStatus) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("JobStatus(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern JobStatus(%d) instead.
func (_j JobStatus) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("JobStatus(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for JobStatus.
func (_j JobStatus) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as JobStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for JobStatus.
func (_j *JobStatus) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("JobStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = JobStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a JobStatus", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for JobStatus.
func (_j JobStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for JobStatus.
func (_j *JobStatus) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of JobStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("JobStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = JobStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a JobStatus", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for JobStatus.
func (_j JobStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as JobStatus. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for JobStatus.
func (_j *JobStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("JobStatus should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("JobStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = JobStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a JobStatus", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for JobStatus.
func (_j *JobStatus) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of JobStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("JobStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = JobStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a JobStatus", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for JobStatus.
func (_j JobStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as JobStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for JobStatus.
func (_j *JobStatus) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("JobStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = JobStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a JobStatus", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for JobStatus.
func (_j JobStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as JobStatus. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for JobStatus.
func (_j *JobStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("JobStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = JobStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a JobStatus", str)
	}
	return nil
}

// JobStatusFromString determines the enum value with an exact case match.
func JobStatusFromString(raw string) (JobStatus, bool) {
	v, ok := _JobStatusStringToValueMap[raw]
	if !ok {
		return JobStatusQueued, false
	}
	return v, true
}

// JobStatusFromStringIgnoreCase determines the enum value with a case-insensitive match.
func JobStatusFromStringIgnoreCase(raw string) (JobStatus, bool) {
	v, ok := JobStatusFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _JobStatusLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_JobStatusStringToValueMap = map[string]JobStatus{
		"queued":     JobStatusQueued,
		"running":    JobStatusRunning,
		"completed":  JobStatusCompleted,
		"cancelling": JobStatusCancelling,
		"canceled":   JobStatusCanceled,
		"failed":     JobStatusFailed,
		"paused":     JobStatusPaused,
	}
	_JobStatusLowerStringToValueMap = map[string]JobStatus{
		"queued":     JobStatusQueued,
		"running":    JobStatusRunning,
		"completed":  JobStatusCompleted,
		"cancelling": JobStatusCancelling,
		"canceled":   JobStatusCanceled,
		"failed":     JobStatusFailed,
		"paused":     JobStatusPaused,
	}
)
