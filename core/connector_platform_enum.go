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

type ConnectorPlatform string

const (
	ConnectorPlatformNative ConnectorPlatform = "native"
	ConnectorPlatformPlugin ConnectorPlatform = "plugin"
	ConnectorPlatformWasm   ConnectorPlatform = "wasm"
)

func (ConnectorPlatform) SQLTypeName() string {
	return "connector_platform"
}

// Values returns a slice of all String values of the enum.
func (ConnectorPlatform) Values() []string {
	return []string{
		"native",
		"plugin",
		"wasm",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j ConnectorPlatform) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j ConnectorPlatform) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("ConnectorPlatform(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern ConnectorPlatform(%d) instead.
func (_j ConnectorPlatform) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("ConnectorPlatform(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for ConnectorPlatform.
func (_j ConnectorPlatform) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorPlatform. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for ConnectorPlatform.
func (_j *ConnectorPlatform) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ConnectorPlatform cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorPlatformFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorPlatform", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for ConnectorPlatform.
func (_j ConnectorPlatform) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for ConnectorPlatform.
func (_j *ConnectorPlatform) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ConnectorPlatform: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ConnectorPlatform cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorPlatformFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorPlatform", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for ConnectorPlatform.
func (_j ConnectorPlatform) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorPlatform. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConnectorPlatform.
func (_j *ConnectorPlatform) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("ConnectorPlatform should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("ConnectorPlatform cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorPlatformFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorPlatform", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for ConnectorPlatform.
func (_j *ConnectorPlatform) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ConnectorPlatform: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ConnectorPlatform cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorPlatformFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorPlatform", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for ConnectorPlatform.
func (_j ConnectorPlatform) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorPlatform. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ConnectorPlatform.
func (_j *ConnectorPlatform) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ConnectorPlatform cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorPlatformFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorPlatform", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for ConnectorPlatform.
func (_j ConnectorPlatform) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorPlatform. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ConnectorPlatform.
func (_j *ConnectorPlatform) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("ConnectorPlatform cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorPlatformFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorPlatform", str)
	}
	return nil
}

// ConnectorPlatformFromString determines the enum value with an exact case match.
func ConnectorPlatformFromString(raw string) (ConnectorPlatform, bool) {
	v, ok := _ConnectorPlatformStringToValueMap[raw]
	if !ok {
		return ConnectorPlatformNative, false
	}
	return v, true
}

// ConnectorPlatformFromStringIgnoreCase determines the enum value with a case-insensitive match.
func ConnectorPlatformFromStringIgnoreCase(raw string) (ConnectorPlatform, bool) {
	v, ok := ConnectorPlatformFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _ConnectorPlatformLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_ConnectorPlatformStringToValueMap = map[string]ConnectorPlatform{
		"native": ConnectorPlatformNative,
		"plugin": ConnectorPlatformPlugin,
		"wasm":   ConnectorPlatformWasm,
	}
	_ConnectorPlatformLowerStringToValueMap = map[string]ConnectorPlatform{
		"native": ConnectorPlatformNative,
		"plugin": ConnectorPlatformPlugin,
		"wasm":   ConnectorPlatformWasm,
	}
)
