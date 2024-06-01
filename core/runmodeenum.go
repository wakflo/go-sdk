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

type RunModeType string

const (
	TriggerRun   RunModeType = "trigger"
	ManualRun    RunModeType = "manual"
	ScheduledRun RunModeType = "schedule"
)

func (RunModeType) SQLTypeName() string {
	return "run_mode_type"
}

// Values returns a slice of all String values of the enum.
func (RunModeType) Values() []string {
	return []string{
		"trigger",
		"manual",
		"schedule",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j RunModeType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j RunModeType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("RunModeType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern RunModeType(%d) instead.
func (_j RunModeType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("RunModeType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for RunModeType.
func (_j RunModeType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as RunModeType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for RunModeType.
func (_j *RunModeType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("RunModeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = RunModeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a RunModeType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for RunModeType.
func (_j RunModeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for RunModeType.
func (_j *RunModeType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of RunModeType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("RunModeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = RunModeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a RunModeType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for RunModeType.
func (_j RunModeType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as RunModeType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for RunModeType.
func (_j *RunModeType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("RunModeType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("RunModeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = RunModeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a RunModeType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for RunModeType.
func (_j *RunModeType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of RunModeType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("RunModeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = RunModeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a RunModeType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for RunModeType.
func (_j RunModeType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as RunModeType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for RunModeType.
func (_j *RunModeType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("RunModeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = RunModeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a RunModeType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for RunModeType.
func (_j RunModeType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as RunModeType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for RunModeType.
func (_j *RunModeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("RunModeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = RunModeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a RunModeType", str)
	}
	return nil
}

// RunModeTypeFromString determines the enum value with an exact case match.
func RunModeTypeFromString(raw string) (RunModeType, bool) {
	v, ok := _RunModeTypeStringToValueMap[raw]
	if !ok {
		return TriggerRun, false
	}
	return v, true
}

// RunModeTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func RunModeTypeFromStringIgnoreCase(raw string) (RunModeType, bool) {
	v, ok := RunModeTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _RunModeTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_RunModeTypeStringToValueMap = map[string]RunModeType{
		"trigger":  TriggerRun,
		"manual":   ManualRun,
		"schedule": ScheduledRun,
	}
	_RunModeTypeLowerStringToValueMap = map[string]RunModeType{
		"trigger":  TriggerRun,
		"manual":   ManualRun,
		"schedule": ScheduledRun,
	}
)
