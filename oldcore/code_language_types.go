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

type CodeLanguage string

const (
	CodeLanguageJavascript CodeLanguage = "javascript"
	CodeLanguageGoLang     CodeLanguage = "go"
	CodeLanguageTypescript CodeLanguage = "typescript"
	CodeLanguageLua        CodeLanguage = "lua"
	CodeLanguagePGSql      CodeLanguage = "pgsql"
)

func (CodeLanguage) SQLTypeName() string {
	return "code_editor_language_type"
}

// Values returns a slice of all String values of the enum.
func (CodeLanguage) Values() []string {
	return []string{
		"javascript",
		"go",
		"typescript",
		"lua",
		"pgsql",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j CodeLanguage) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j CodeLanguage) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("CodeLanguage(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern CodeLanguage(%d) instead.
func (_j CodeLanguage) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("CodeLanguage(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for CodeLanguage.
func (_j CodeLanguage) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeLanguage. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for CodeLanguage.
func (_j *CodeLanguage) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("CodeLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeLanguage", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for CodeLanguage.
func (_j CodeLanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for CodeLanguage.
func (_j *CodeLanguage) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of CodeLanguage: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("CodeLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeLanguage", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for CodeLanguage.
func (_j CodeLanguage) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeLanguage. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CodeLanguage.
func (_j *CodeLanguage) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("CodeLanguage should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("CodeLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeLanguage", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for CodeLanguage.
func (_j *CodeLanguage) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of CodeLanguage: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("CodeLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeLanguage", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for CodeLanguage.
func (_j CodeLanguage) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeLanguage. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CodeLanguage.
func (_j *CodeLanguage) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("CodeLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeLanguage", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for CodeLanguage.
func (_j CodeLanguage) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as CodeLanguage. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CodeLanguage.
func (_j *CodeLanguage) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("CodeLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = CodeEditorLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a CodeLanguage", str)
	}
	return nil
}

// CodeEditorLanguageFromString determines the enum value with an exact case match.
func CodeEditorLanguageFromString(raw string) (CodeLanguage, bool) {
	v, ok := _CodeEditorLanguageStringToValueMap[raw]
	if !ok {
		return CodeLanguageJavascript, false
	}
	return v, true
}

// CodeEditorLanguageFromStringIgnoreCase determines the enum value with a case-insensitive match.
func CodeEditorLanguageFromStringIgnoreCase(raw string) (CodeLanguage, bool) {
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
	_CodeEditorLanguageStringToValueMap = map[string]CodeLanguage{
		"javascript": CodeLanguageJavascript,
		"go":         CodeLanguageGoLang,
		"typescript": CodeLanguageTypescript,
		"lua":        CodeLanguageLua,
		"pgsql":      CodeLanguagePGSql,
	}
	_CodeEditorLanguageLowerStringToValueMap = map[string]CodeLanguage{
		"javascript": CodeLanguageJavascript,
		"go":         CodeLanguageGoLang,
		"typescript": CodeLanguageTypescript,
		"lua":        CodeLanguageLua,
		"pgsql":      CodeLanguagePGSql,
	}
)
