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

type FlowStepType string

const (
	FlowStepTypeEmpty       FlowStepType = "EMPTY"
	FlowStepTypeStepTrigger FlowStepType = "STEP_TRIGGER"
	FlowStepTypeStep        FlowStepType = "STEP"
	FlowStepTypeLoop        FlowStepType = "LOOP"
	FlowStepTypeStepRouter  FlowStepType = "ROUTER"
)

func (FlowStepType) SQLTypeName() string {
	return "flow_step_type"
}

// Values returns a slice of all String values of the enum.
func (FlowStepType) Values() []string {
	return []string{
		"EMPTY",
		"STEP_TRIGGER",
		"STEP",
		"LOOP",
		"ROUTER",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j FlowStepType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j FlowStepType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("FlowStepType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern FlowStepType(%d) instead.
func (_j FlowStepType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("FlowStepType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for FlowStepType.
func (_j FlowStepType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowStepType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for FlowStepType.
func (_j *FlowStepType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("FlowStepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStepType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for FlowStepType.
func (_j FlowStepType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for FlowStepType.
func (_j *FlowStepType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of FlowStepType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("FlowStepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStepType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for FlowStepType.
func (_j FlowStepType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowStepType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FlowStepType.
func (_j *FlowStepType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("FlowStepType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("FlowStepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStepType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for FlowStepType.
func (_j *FlowStepType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of FlowStepType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("FlowStepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStepType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for FlowStepType.
func (_j FlowStepType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowStepType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for FlowStepType.
func (_j *FlowStepType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("FlowStepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStepType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for FlowStepType.
func (_j FlowStepType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FlowStepType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for FlowStepType.
func (_j *FlowStepType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("FlowStepType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FlowStepTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FlowStepType", str)
	}
	return nil
}

// FlowStepTypeFromString determines the enum value with an exact case match.
func FlowStepTypeFromString(raw string) (FlowStepType, bool) {
	v, ok := _FlowStepTypeStringToValueMap[raw]
	if !ok {
		return FlowStepTypeEmpty, false
	}
	return v, true
}

// FlowStepTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func FlowStepTypeFromStringIgnoreCase(raw string) (FlowStepType, bool) {
	v, ok := FlowStepTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _FlowStepTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_FlowStepTypeStringToValueMap = map[string]FlowStepType{
		"EMPTY":        FlowStepTypeEmpty,
		"STEP_TRIGGER": FlowStepTypeStepTrigger,
		"STEP":         FlowStepTypeStep,
		"LOOP":         FlowStepTypeLoop,
		"ROUTER":       FlowStepTypeStepRouter,
	}
	_FlowStepTypeLowerStringToValueMap = map[string]FlowStepType{
		"EMPTY":        FlowStepTypeEmpty,
		"STEP_TRIGGER": FlowStepTypeStepTrigger,
		"STEP":         FlowStepTypeStep,
		"LOOP":         FlowStepTypeLoop,
		"ROUTER":       FlowStepTypeStepRouter,
	}
)
