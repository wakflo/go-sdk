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

type CodeEditorLanguage string

const (
	Javascript CodeEditorLanguage = "javascript"
	GoLang     CodeEditorLanguage = "go"
	Typescript CodeEditorLanguage = "typescript"
	Lua        CodeEditorLanguage = "lua"
)

func (CodeEditorLanguage) SQLTypeName() string {
	return "code_editor_language_type"
}

// Values returns a slice of all String values of the enum.
func (CodeEditorLanguage) Values() []string {
	return []string{
		"javascript",
		"go",
		"typescript",
		"lua",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j CodeEditorLanguage) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j CodeEditorLanguage) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("CodeEditorLanguage(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern CodeEditorLanguage(%d) instead.
func (_j CodeEditorLanguage) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("CodeEditorLanguage(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for CodeEditorLanguage.
func (_j CodeEditorLanguage) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeEditorLanguage. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for CodeEditorLanguage.
func (_j *CodeEditorLanguage) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("CodeEditorLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeEditorLanguage", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for CodeEditorLanguage.
func (_j CodeEditorLanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for CodeEditorLanguage.
func (_j *CodeEditorLanguage) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of CodeEditorLanguage: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("CodeEditorLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeEditorLanguage", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for CodeEditorLanguage.
func (_j CodeEditorLanguage) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeEditorLanguage. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CodeEditorLanguage.
func (_j *CodeEditorLanguage) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("CodeEditorLanguage should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("CodeEditorLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeEditorLanguage", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for CodeEditorLanguage.
func (_j *CodeEditorLanguage) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of CodeEditorLanguage: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("CodeEditorLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeEditorLanguage", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for CodeEditorLanguage.
func (_j CodeEditorLanguage) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeEditorLanguage. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CodeEditorLanguage.
func (_j *CodeEditorLanguage) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("CodeEditorLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeEditorLanguage", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for CodeEditorLanguage.
func (_j CodeEditorLanguage) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeEditorLanguage. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CodeEditorLanguage.
func (_j *CodeEditorLanguage) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("CodeEditorLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeEditorLanguage", str)
	}
	return nil
}

// CodeEditorLanguageFromString determines the enum value with an exact case match.
func CodeEditorLanguageFromString(raw string) (CodeEditorLanguage, bool) {
	v, ok := _CodeEditorLanguageStringToValueMap[raw]
	if !ok {
		return Javascript, false
	}
	return v, true
}

// CodeEditorLanguageFromStringIgnoreCase determines the enum value with a case-insensitive match.
func CodeEditorLanguageFromStringIgnoreCase(raw string) (CodeEditorLanguage, bool) {
	v, ok := CodeEditorLanguageFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _CodeEditorLanguageLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_CodeEditorLanguageStringToValueMap = map[string]CodeEditorLanguage{
		"javascript": Javascript,
		"go":         GoLang,
		"typescript": Typescript,
		"lua":        Lua,
	}
	_CodeEditorLanguageLowerStringToValueMap = map[string]CodeEditorLanguage{
		"javascript": Javascript,
		"go":         GoLang,
		"typescript": Typescript,
		"lua":        Lua,
	}
)
