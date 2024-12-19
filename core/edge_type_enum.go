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

type EdgeType string

const (
	EdgeTypeStep        EdgeType = "step"
	EdgeTypePlaceholder EdgeType = "placeholder"
	EdgeTypeLoop        EdgeType = "loop"
)

func (EdgeType) SQLTypeName() string {
	return "edge_type"
}

// Values returns a slice of all String values of the enum.
func (EdgeType) Values() []string {
	return []string{
		"step",
		"placeholder",
		"loop",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j EdgeType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j EdgeType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("EdgeType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern EdgeType(%d) instead.
func (_j EdgeType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("EdgeType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for EdgeType.
func (_j EdgeType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as EdgeType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for EdgeType.
func (_j *EdgeType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("EdgeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = EdgeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a EdgeType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for EdgeType.
func (_j EdgeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for EdgeType.
func (_j *EdgeType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of EdgeType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("EdgeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = EdgeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a EdgeType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for EdgeType.
func (_j EdgeType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as EdgeType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for EdgeType.
func (_j *EdgeType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("EdgeType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("EdgeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = EdgeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a EdgeType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for EdgeType.
func (_j *EdgeType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of EdgeType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("EdgeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = EdgeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a EdgeType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for EdgeType.
func (_j EdgeType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as EdgeType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for EdgeType.
func (_j *EdgeType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("EdgeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = EdgeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a EdgeType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for EdgeType.
func (_j EdgeType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as EdgeType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for EdgeType.
func (_j *EdgeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("EdgeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = EdgeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a EdgeType", str)
	}
	return nil
}

// EdgeTypeFromString determines the enum value with an exact case match.
func EdgeTypeFromString(raw string) (EdgeType, bool) {
	v, ok := _EdgeTypeStringToValueMap[raw]
	if !ok {
		return EdgeTypeStep, false
	}
	return v, true
}

// EdgeTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func EdgeTypeFromStringIgnoreCase(raw string) (EdgeType, bool) {
	v, ok := EdgeTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _EdgeTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_EdgeTypeStringToValueMap = map[string]EdgeType{
		"step":        EdgeTypeStep,
		"placeholder": EdgeTypePlaceholder,
		"loop":        EdgeTypeLoop,
	}
	_EdgeTypeLowerStringToValueMap = map[string]EdgeType{
		"step":        EdgeTypeStep,
		"placeholder": EdgeTypePlaceholder,
		"loop":        EdgeTypeLoop,
	}
)
