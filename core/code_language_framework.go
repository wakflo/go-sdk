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

type CodeFramework string

const (
	CodeFrameworkNode   CodeFramework = "NODEJS"
	CodeFrameworkGoLang CodeFramework = "GOLANG"
	CodeFrameworkDeno   CodeFramework = "DENO"
)

func (CodeFramework) SQLTypeName() string {
	return "code_framework_type"
}

// Values returns a slice of all String values of the enum.
func (CodeFramework) Values() []string {
	return []string{
		"NODEJS",
		"GOLANG",
		"DENO",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j CodeFramework) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j CodeFramework) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("CodeFramework(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern CodeFramework(%d) instead.
func (_j CodeFramework) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("CodeFramework(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for CodeFramework.
func (_j CodeFramework) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeFramework. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for CodeFramework.
func (_j *CodeFramework) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("CodeFramework cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeFrameworkFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeFramework", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for CodeFramework.
func (_j CodeFramework) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for CodeFramework.
func (_j *CodeFramework) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of CodeFramework: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("CodeFramework cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeFrameworkFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeFramework", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for CodeFramework.
func (_j CodeFramework) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeFramework. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CodeFramework.
func (_j *CodeFramework) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("CodeFramework should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("CodeFramework cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeFrameworkFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeFramework", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for CodeFramework.
func (_j *CodeFramework) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of CodeFramework: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("CodeFramework cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeFrameworkFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeFramework", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for CodeFramework.
func (_j CodeFramework) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeFramework. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CodeFramework.
func (_j *CodeFramework) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("CodeFramework cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeFrameworkFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeFramework", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for CodeFramework.
func (_j CodeFramework) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeFramework. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CodeFramework.
func (_j *CodeFramework) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("CodeFramework cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeFrameworkFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeFramework", str)
	}
	return nil
}

// CodeFrameworkFromString determines the enum value with an exact case match.
func CodeFrameworkFromString(raw string) (CodeFramework, bool) {
	v, ok := _CodeFrameworkStringToValueMap[raw]
	if !ok {
		return CodeFrameworkNode, false
	}
	return v, true
}

// CodeFrameworkFromStringIgnoreCase determines the enum value with a case-insensitive match.
func CodeFrameworkFromStringIgnoreCase(raw string) (CodeFramework, bool) {
	v, ok := CodeFrameworkFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _CodeFrameworkLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_CodeFrameworkStringToValueMap = map[string]CodeFramework{
		"NODEJS": CodeFrameworkNode,
		"GOLANG": CodeFrameworkGoLang,
		"DENO":   CodeFrameworkDeno,
	}
	_CodeFrameworkLowerStringToValueMap = map[string]CodeFramework{
		"NODEJS": CodeFrameworkNode,
		"GOLANG": CodeFrameworkGoLang,
		"DENO":   CodeFrameworkDeno,
	}
)
