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

type PauseType string

const (
	DelayPause   PauseType = "DELAY"
	WebhookPause PauseType = "WEBHOOK"
)

func (PauseType) SQLTypeName() string {
	return "run_pause_type"
}

// Values returns a slice of all String values of the enum.
func (PauseType) Values() []string {
	return []string{
		"DELAY",
		"WEBHOOK",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j PauseType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j PauseType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("PauseType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern PauseType(%d) instead.
func (_j PauseType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("PauseType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for PauseType.
func (_j PauseType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as PauseType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for PauseType.
func (_j *PauseType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("PauseType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = PauseTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a PauseType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for PauseType.
func (_j PauseType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for PauseType.
func (_j *PauseType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of PauseType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("PauseType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = PauseTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a PauseType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for PauseType.
func (_j PauseType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as PauseType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for PauseType.
func (_j *PauseType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("PauseType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("PauseType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = PauseTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a PauseType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for PauseType.
func (_j *PauseType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of PauseType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("PauseType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = PauseTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a PauseType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for PauseType.
func (_j PauseType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as PauseType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for PauseType.
func (_j *PauseType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("PauseType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = PauseTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a PauseType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for PauseType.
func (_j PauseType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as PauseType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for PauseType.
func (_j *PauseType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("PauseType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = PauseTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a PauseType", str)
	}
	return nil
}

// PauseTypeFromString determines the enum value with an exact case match.
func PauseTypeFromString(raw string) (PauseType, bool) {
	v, ok := _PauseTypeStringToValueMap[raw]
	if !ok {
		return DelayPause, false
	}
	return v, true
}

// PauseTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func PauseTypeFromStringIgnoreCase(raw string) (PauseType, bool) {
	v, ok := PauseTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _PauseTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_PauseTypeStringToValueMap = map[string]PauseType{
		"DELAY":   DelayPause,
		"WEBHOOK": WebhookPause,
	}
	_PauseTypeLowerStringToValueMap = map[string]PauseType{
		"DELAY":   DelayPause,
		"WEBHOOK": WebhookPause,
	}
)
