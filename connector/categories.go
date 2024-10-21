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

var ErrNoValidEnum = errors.New("not a valid enum")

type ConnectorGroup string

const (
	ConnectorGroupApps    ConnectorGroup = "apps"
	ConnectorGroupCore    ConnectorGroup = "core"
	ConnectorGroupAI      ConnectorGroup = "ai"
	ConnectorGroupScripts ConnectorGroup = "scripts"
	ConnectorGroupTools   ConnectorGroup = "tools"
)

func (ConnectorGroup) SQLTypeName() string {
	return "connector_category"
}

// Values returns a slice of all String values of the enum.
func (ConnectorGroup) Values() []string {
	return []string{
		"apps",
		"core",
		"ai",
		"scripts",
		"tools",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j ConnectorGroup) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j ConnectorGroup) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("ConnectorGroup(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern ConnectorGroup(%d) instead.
func (_j ConnectorGroup) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("ConnectorGroup(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for ConnectorGroup.
func (_j ConnectorGroup) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorGroup. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for ConnectorGroup.
func (_j *ConnectorGroup) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ConnectorGroup cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorGroup", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for ConnectorGroup.
func (_j ConnectorGroup) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for ConnectorGroup.
func (_j *ConnectorGroup) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ConnectorGroup: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ConnectorGroup cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorGroup", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for ConnectorGroup.
func (_j ConnectorGroup) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorGroup. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConnectorGroup.
func (_j *ConnectorGroup) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("ConnectorGroup should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("ConnectorGroup cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorGroup", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for ConnectorGroup.
func (_j *ConnectorGroup) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ConnectorGroup: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ConnectorGroup cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorGroup", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for ConnectorGroup.
func (_j ConnectorGroup) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorGroup. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ConnectorGroup.
func (_j *ConnectorGroup) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ConnectorGroup cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorGroup", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for ConnectorGroup.
func (_j ConnectorGroup) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ConnectorGroup. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ConnectorGroup.
func (_j *ConnectorGroup) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("ConnectorGroup cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ConnectorCategoryFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ConnectorGroup", str)
	}
	return nil
}

// ConnectorCategoryFromString determines the enum value with an exact case match.
func ConnectorCategoryFromString(raw string) (ConnectorGroup, bool) {
	v, ok := _ConnectorCategoryStringToValueMap[raw]
	if !ok {
		return ConnectorGroupApps, false
	}
	return v, true
}

// ConnectorCategoryFromStringIgnoreCase determines the enum value with a case-insensitive match.
func ConnectorCategoryFromStringIgnoreCase(raw string) (ConnectorGroup, bool) {
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
	_ConnectorCategoryStringToValueMap = map[string]ConnectorGroup{
		"apps":    ConnectorGroupApps,
		"core":    ConnectorGroupCore,
		"ai":      ConnectorGroupAI,
		"scripts": ConnectorGroupScripts,
		"tools":   ConnectorGroupTools,
	}
	_ConnectorCategoryLowerStringToValueMap = map[string]ConnectorGroup{
		"apps":    ConnectorGroupApps,
		"core":    ConnectorGroupCore,
		"ai":      ConnectorGroupAI,
		"scripts": ConnectorGroupScripts,
		"tools":   ConnectorGroupTools,
	}
)
