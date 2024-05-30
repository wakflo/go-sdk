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

package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

var (
	ErrNoValidEnum = errors.New("not a valid enum")
)

type ConnectorCategory string

const (
	Apps    ConnectorCategory = "apps"
	Core                      = "core"
	AI                        = "ai"
	Scripts                   = "scripts"
	Tools                     = "tools"
)

func (ConnectorCategory) SqlTypeName() string {
	return "connector_category"
}

// Values returns a slice of all String values of the enum.
func (ConnectorCategory) Values() []string {
	return []string{
		"apps",
		"core",
		"ai",
		"scripts",
		"tools",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j ConnectorCategory) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j ConnectorCategory) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("ConnectorCategory(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern ConnectorCategory(%d) instead.
func (_j ConnectorCategory) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("ConnectorCategory(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for ConnectorCategory.
func (_j ConnectorCategory) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorCategory. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for ConnectorCategory.
func (_j *ConnectorCategory) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return fmt.Errorf("ConnectorCategory cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorCategory", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for ConnectorCategory.
func (_j ConnectorCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for ConnectorCategory.
func (_j *ConnectorCategory) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ConnectorCategory: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return fmt.Errorf("ConnectorCategory cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorCategory", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for ConnectorCategory.
func (_j ConnectorCategory) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorCategory. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConnectorCategory.
func (_j *ConnectorCategory) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("ConnectorCategory should be a string, got %q", data)
	}
	if len(str) == 0 {
		return fmt.Errorf("ConnectorCategory cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorCategory", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for ConnectorCategory.
func (_j *ConnectorCategory) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ConnectorCategory: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return fmt.Errorf("ConnectorCategory cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorCategory", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for ConnectorCategory.
func (_j ConnectorCategory) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorCategory. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ConnectorCategory.
func (_j *ConnectorCategory) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return fmt.Errorf("ConnectorCategory cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorCategory", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for ConnectorCategory.
func (_j ConnectorCategory) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorCategory. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ConnectorCategory.
func (_j *ConnectorCategory) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return fmt.Errorf("ConnectorCategory cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorCategory", str)
	}
	return nil
}

// ConnectorCategoryFromString determines the enum value with an exact case match.
func ConnectorCategoryFromString(raw string) (ConnectorCategory, bool) {
	v, ok := _ConnectorCategoryStringToValueMap[raw]
	if !ok {
		return Apps, false
	}
	return v, true
}

// ConnectorCategoryFromStringIgnoreCase determines the enum value with a case-insensitive match.
func ConnectorCategoryFromStringIgnoreCase(raw string) (ConnectorCategory, bool) {
	v, ok := ConnectorCategoryFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _ConnectorCategoryLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_ConnectorCategoryStringToValueMap = map[string]ConnectorCategory{
		"apps":    Apps,
		"core":    Core,
		"ai":      AI,
		"scripts": Scripts,
		"tools":   Tools,
	}
	_ConnectorCategoryLowerStringToValueMap = map[string]ConnectorCategory{
		"apps":    Apps,
		"core":    Core,
		"ai":      AI,
		"scripts": Scripts,
		"tools":   Tools,
	}
)
