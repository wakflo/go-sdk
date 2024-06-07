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

type ConnectorType string

const (
	ConnectorTypeBranch  ConnectorType = "branch"
	ConnectorTypeBoolean ConnectorType = "boolean"
	ConnectorTypeNormal  ConnectorType = "normal"
	ConnectorTypeLoop    ConnectorType = "loop"
)

func (ConnectorType) SQLTypeName() string {
	return "auth_type"
}

// Values returns a slice of all String values of the enum.
func (ConnectorType) Values() []string {
	return []string{
		"branch",
		"boolean",
		"normal",
		"loop",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j ConnectorType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j ConnectorType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("ConnectorType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern ConnectorType(%d) instead.
func (_j ConnectorType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("ConnectorType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for ConnectorType.
func (_j ConnectorType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for ConnectorType.
func (_j *ConnectorType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ConnectorType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for ConnectorType.
func (_j ConnectorType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for ConnectorType.
func (_j *ConnectorType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ConnectorType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ConnectorType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for ConnectorType.
func (_j ConnectorType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConnectorType.
func (_j *ConnectorType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("ConnectorType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("ConnectorType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for ConnectorType.
func (_j *ConnectorType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ConnectorType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ConnectorType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for ConnectorType.
func (_j ConnectorType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ConnectorType.
func (_j *ConnectorType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ConnectorType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for ConnectorType.
func (_j ConnectorType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ConnectorType.
func (_j *ConnectorType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("ConnectorType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorType", str)
	}
	return nil
}

// ConnectorTypeFromString determines the enum value with an exact case match.
func ConnectorTypeFromString(raw string) (ConnectorType, bool) {
	v, ok := _ConnectorTypeStringToValueMap[raw]
	if !ok {
		return ConnectorTypeNormal, false
	}
	return v, true
}

// ConnectorTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func ConnectorTypeFromStringIgnoreCase(raw string) (ConnectorType, bool) {
	v, ok := ConnectorTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _ConnectorTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_ConnectorTypeStringToValueMap = map[string]ConnectorType{
		"branch":  ConnectorTypeBranch,
		"boolean": ConnectorTypeBoolean,
		"normal":  ConnectorTypeNormal,
		"loop":    ConnectorTypeLoop,
	}
	_ConnectorTypeLowerStringToValueMap = map[string]ConnectorType{
		"branch":  ConnectorTypeBranch,
		"boolean": ConnectorTypeBoolean,
		"normal":  ConnectorTypeNormal,
		"loop":    ConnectorTypeLoop,
	}
)
