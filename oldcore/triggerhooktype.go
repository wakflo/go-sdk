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

type TriggerHookType string

const (
	TriggerHookTypeRun        TriggerHookType = "RUN"
	TriggerHookTypeTest       TriggerHookType = "TEST"
	TriggerHookTypeOnEnable   TriggerHookType = "ON_ENABLED"
	TriggerHookTypeOnDisabled TriggerHookType = "ON_DISABLED"
	TriggerHookTypeRenew      TriggerHookType = "RENEW"
	TriggerHookTypeOptions    TriggerHookType = "OPTIONS"
)

func (TriggerHookType) SQLTypeName() string {
	return "trigger_hook_type"
}

// Values returns a slice of all String values of the enum.
func (TriggerHookType) Values() []string {
	return []string{
		"RUN",
		"TEST",
		"ON_ENABLED",
		"ON_DISABLED",
		"RENEW",
		"OPTIONS",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j TriggerHookType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j TriggerHookType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("TriggerHookType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern TriggerHookType(%d) instead.
func (_j TriggerHookType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("TriggerHookType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for TriggerHookType.
func (_j TriggerHookType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerHookType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for TriggerHookType.
func (_j *TriggerHookType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("TriggerHookType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerHookTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerHookType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for TriggerHookType.
func (_j TriggerHookType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for TriggerHookType.
func (_j *TriggerHookType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of TriggerHookType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("TriggerHookType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerHookTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerHookType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for TriggerHookType.
func (_j TriggerHookType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerHookType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TriggerHookType.
func (_j *TriggerHookType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("TriggerHookType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("TriggerHookType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerHookTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerHookType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for TriggerHookType.
func (_j *TriggerHookType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of TriggerHookType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("TriggerHookType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerHookTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerHookType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for TriggerHookType.
func (_j TriggerHookType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerHookType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for TriggerHookType.
func (_j *TriggerHookType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("TriggerHookType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerHookTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerHookType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for TriggerHookType.
func (_j TriggerHookType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerHookType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for TriggerHookType.
func (_j *TriggerHookType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("TriggerHookType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerHookTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerHookType", str)
	}
	return nil
}

// TriggerHookTypeFromString determines the enum value with an exact case match.
func TriggerHookTypeFromString(raw string) (TriggerHookType, bool) {
	v, ok := _TriggerHookTypeStringToValueMap[raw]
	if !ok {
		return TriggerHookTypeRun, false
	}
	return v, true
}

// TriggerHookTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func TriggerHookTypeFromStringIgnoreCase(raw string) (TriggerHookType, bool) {
	v, ok := TriggerHookTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _TriggerHookTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_TriggerHookTypeStringToValueMap = map[string]TriggerHookType{
		"RUN":         TriggerHookTypeRun,
		"TEST":        TriggerHookTypeTest,
		"ON_ENABLED":  TriggerHookTypeOnEnable,
		"ON_DISABLED": TriggerHookTypeOnDisabled,
		"RENEW":       TriggerHookTypeRenew,
		"OPTIONS":     TriggerHookTypeOptions,
	}
	_TriggerHookTypeLowerStringToValueMap = map[string]TriggerHookType{
		"RUN":         TriggerHookTypeRun,
		"TEST":        TriggerHookTypeTest,
		"ON_ENABLED":  TriggerHookTypeOnEnable,
		"ON_DISABLED": TriggerHookTypeOnDisabled,
		"RENEW":       TriggerHookTypeRenew,
		"OPTIONS":     TriggerHookTypeOptions,
	}
)
