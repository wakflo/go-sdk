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

type TriggerType string

const (
	TriggerTypeScheduled TriggerType = "SCHEDULED"
	TriggerTypeEvent     TriggerType = "EVENT"
	TriggerTypePolling   TriggerType = "POLLING"
	TriggerTypeWebhook   TriggerType = "WEBHOOK"
	TriggerTypeManual    TriggerType = "MANUAL"
)

func (TriggerType) SQLTypeName() string {
	return "trigger_type"
}

// Values returns a slice of all String values of the enum.
func (TriggerType) Values() []string {
	return []string{
		"SCHEDULED",
		"EVENT",
		"POLLING",
		"WEBHOOK",
		"MANUAL",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j TriggerType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j TriggerType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("TriggerType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern TriggerType(%d) instead.
func (_j TriggerType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("TriggerType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for TriggerType.
func (_j TriggerType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for TriggerType.
func (_j *TriggerType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for TriggerType.
func (_j TriggerType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for TriggerType.
func (_j *TriggerType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of TriggerType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for TriggerType.
func (_j TriggerType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TriggerType.
func (_j *TriggerType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("TriggerType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for TriggerType.
func (_j *TriggerType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of TriggerType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for TriggerType.
func (_j TriggerType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for TriggerType.
func (_j *TriggerType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for TriggerType.
func (_j TriggerType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for TriggerType.
func (_j *TriggerType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// TriggerTypeFromString determines the enum value with an exact case match.
func TriggerTypeFromString(raw string) (TriggerType, bool) {
	v, ok := _TriggerTypeStringToValueMap[raw]
	if !ok {
		return TriggerTypeScheduled, false
	}
	return v, true
}

// TriggerTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func TriggerTypeFromStringIgnoreCase(raw string) (TriggerType, bool) {
	v, ok := TriggerTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _TriggerTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_TriggerTypeStringToValueMap = map[string]TriggerType{
		"SCHEDULED": TriggerTypeScheduled,
		"EVENT":     TriggerTypeEvent,
		"POLLING":   TriggerTypePolling,
		"WEBHOOK":   TriggerTypeWebhook,
		"MANUAL":    TriggerTypeManual,
	}
	_TriggerTypeLowerStringToValueMap = map[string]TriggerType{
		"SCHEDULED": TriggerTypeScheduled,
		"EVENT":     TriggerTypeEvent,
		"POLLING":   TriggerTypePolling,
		"WEBHOOK":   TriggerTypeWebhook,
		"MANUAL":    TriggerTypeManual,
	}
)
