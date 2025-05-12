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

// IntegrationLanguage represents programming languages supported for plugins.
type IntegrationLanguage string

const (
	IntegrationLanguageJavaScript IntegrationLanguage = "javascript"
	IntegrationLanguageTypeScript IntegrationLanguage = "typescript"
	IntegrationLanguagePython     IntegrationLanguage = "python"
	IntegrationLanguageGo         IntegrationLanguage = "go"
	IntegrationLanguageRuby       IntegrationLanguage = "ruby"
	IntegrationLanguageJava       IntegrationLanguage = "java"
	IntegrationLanguageCSharp     IntegrationLanguage = "csharp"
)

func (IntegrationLanguage) SQLTypeName() string {
	return "plugin_language"
}

// Values returns a slice of all String values of the enum.
func (IntegrationLanguage) Values() []string {
	return []string{
		"javascript",
		"typescript",
		"python",
		"go",
		"ruby",
		"java",
		"csharp",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j IntegrationLanguage) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j IntegrationLanguage) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("IntegrationLanguage(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
func (_j IntegrationLanguage) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("IntegrationLanguage(%v)", string(_j))
	}
	return string(_j)
}

// MarshalJSON implements the json.Marshaler interface for IntegrationLanguage.
func (_j IntegrationLanguage) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as IntegrationLanguage. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for IntegrationLanguage.
func (_j *IntegrationLanguage) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("IntegrationLanguage should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("IntegrationLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = IntegrationLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a IntegrationLanguage", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for IntegrationLanguage.
func (_j IntegrationLanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for IntegrationLanguage.
func (_j *IntegrationLanguage) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of IntegrationLanguage: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("IntegrationLanguage cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = IntegrationLanguageFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a IntegrationLanguage", str)
	}
	return nil
}

// IntegrationLanguageFromString determines the enum value with an exact case match.
func IntegrationLanguageFromString(raw string) (IntegrationLanguage, bool) {
	v, ok := _IntegrationLanguageStringToValueMap[raw]
	if !ok {
		return IntegrationLanguageJavaScript, false
	}
	return v, true
}

var _IntegrationLanguageStringToValueMap = map[string]IntegrationLanguage{
	"javascript": IntegrationLanguageJavaScript,
	"typescript": IntegrationLanguageTypeScript,
	"python":     IntegrationLanguagePython,
	"go":         IntegrationLanguageGo,
	"ruby":       IntegrationLanguageRuby,
	"java":       IntegrationLanguageJava,
	"csharp":     IntegrationLanguageCSharp,
}

// IntegrationPlatform represents the type of a plugin.
type IntegrationPlatform string

const (
	IntegrationPlatformNative IntegrationPlatform = "native"
	IntegrationPlatformWASM   IntegrationPlatform = "wasm"
)

func (IntegrationPlatform) SQLTypeName() string {
	return "plugin_type"
}

// Values returns a slice of all String values of the enum.
func (IntegrationPlatform) Values() []string {
	return []string{
		"native",
		"wasm",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j IntegrationPlatform) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j IntegrationPlatform) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("IntegrationPlatform(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
func (_j IntegrationPlatform) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("IntegrationPlatform(%v)", string(_j))
	}
	return string(_j)
}

// MarshalJSON implements the json.Marshaler interface for IntegrationPlatform.
func (_j IntegrationPlatform) MarshalJSON() ([]byte, error) {
	if !_j.IsValid() {
		// Use default value instead of returning an error
		return json.Marshal(IntegrationPlatformNative.String())
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for IntegrationPlatform.
func (_j *IntegrationPlatform) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("IntegrationPlatform should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("IntegrationPlatform cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = IntegrationPlatformFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a IntegrationPlatform", str)
	}
	return nil
}

// IntegrationPlatformFromString determines the enum value with an exact case match.
func IntegrationPlatformFromString(raw string) (IntegrationPlatform, bool) {
	v, ok := _IntegrationPlatformStringToValueMap[raw]
	if !ok {
		return IntegrationPlatformNative, false
	}
	return v, true
}

var _IntegrationPlatformStringToValueMap = map[string]IntegrationPlatform{
	"native": IntegrationPlatformNative,
	"wasm":   IntegrationPlatformWASM,
}
