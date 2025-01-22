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

// MULTI_SELECT_DROPDOWN = 'MULTI_SELECT_DROPDOWN',
// STATIC_MULTI_SELECT_DROPDOWN = 'STATIC_MULTI_SELECT_DROPDOWN',
// DYNAMIC = 'DYNAMIC',
// CUSTOM_AUTH = 'CUSTOM_AUTH',

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

type AutoFormFieldType string

const (
	AutoFormFieldTypeShortText      AutoFormFieldType = "short_text"
	AutoFormFieldTypeLongText       AutoFormFieldType = "long_text"
	AutoFormFieldTypeMarkdown       AutoFormFieldType = "markdown"
	AutoFormFieldTypeSelect         AutoFormFieldType = "select"
	AutoFormFieldTypeStaticDropdown AutoFormFieldType = "static_dropdown"
	AutoFormFieldTypeNumber         AutoFormFieldType = "number"
	AutoFormFieldTypeCheckbox       AutoFormFieldType = "checkbox"
	AutoFormFieldTypeOauth2         AutoFormFieldType = "oauth"
	AutoFormFieldTypeSecretAuth     AutoFormFieldType = "secret"
	AutoFormFieldTypeCustomAuth     AutoFormFieldType = "custom_auth"
	AutoFormFieldTypeArray          AutoFormFieldType = "array"
	AutoFormFieldTypeObject         AutoFormFieldType = "fieldset"
	AutoFormFieldTypeBasicAuth      AutoFormFieldType = "basic_auth"
	AutoFormFieldTypeJSON           AutoFormFieldType = "json"
	AutoFormFieldTypeDateTime       AutoFormFieldType = "datetime"
	AutoFormFieldTypeFile           AutoFormFieldType = "file"
	AutoFormFieldTypeFileString     AutoFormFieldType = "file_string"
	AutoFormFieldTypeBoolean        AutoFormFieldType = "boolean"
	AutoFormFieldTypeDynamic        AutoFormFieldType = "dynamic"
	AutoFormFieldTypeCode           AutoFormFieldType = "code"
	AutoFormFieldTypeRichText       AutoFormFieldType = "richtext"
	AutoFormFieldTypeBranch         AutoFormFieldType = "branch"
	AutoFormFieldTypeWrapper        AutoFormFieldType = "wrapper"
	AutoFormFieldTypeCodeFramework  AutoFormFieldType = "framework"
)

func (AutoFormFieldType) SQLTypeName() string {
	return "auth_type"
}

// Values returns a slice of all String values of the enum.
func (AutoFormFieldType) Values() []string {
	return []string{
		"short_text",
		"long_text",
		"markdown",
		"select",
		"static_dropdown",
		"number",
		"checkbox",
		"oauth",
		"secret",
		"custom_auth",
		"array",
		"fieldset",
		"basic_auth",
		"json",
		"datetime",
		"file",
		"file_string",
		"boolean",
		"dynamic",
		"code",
		"richtext",
		"branch",
		"wrapper",
		"framework",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j AutoFormFieldType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j AutoFormFieldType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("AutoFormFieldType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern AutoFormFieldType(%d) instead.
func (_j AutoFormFieldType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("AutoFormFieldType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for AutoFormFieldType.
func (_j AutoFormFieldType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AutoFormFieldType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for AutoFormFieldType.
func (_j *AutoFormFieldType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("AutoFormFieldType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AutoFormFieldType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for AutoFormFieldType.
func (_j AutoFormFieldType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for AutoFormFieldType.
func (_j *AutoFormFieldType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AutoFormFieldType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("AutoFormFieldType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AutoFormFieldType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for AutoFormFieldType.
func (_j AutoFormFieldType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AutoFormFieldType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoFormFieldType.
func (_j *AutoFormFieldType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("AutoFormFieldType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("AutoFormFieldType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AutoFormFieldType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for AutoFormFieldType.
func (_j *AutoFormFieldType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AutoFormFieldType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("AutoFormFieldType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AutoFormFieldType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for AutoFormFieldType.
func (_j AutoFormFieldType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AutoFormFieldType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for AutoFormFieldType.
func (_j *AutoFormFieldType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("AutoFormFieldType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AutoFormFieldType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for AutoFormFieldType.
func (_j AutoFormFieldType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AutoFormFieldType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for AutoFormFieldType.
func (_j *AutoFormFieldType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("AutoFormFieldType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AutoFormFieldType", str)
	}
	return nil
}

// FieldTypeFromString determines the enum value with an exact case match.
func FieldTypeFromString(raw string) (AutoFormFieldType, bool) {
	v, ok := _FieldTypeStringToValueMap[raw]
	if !ok {
		return AutoFormFieldTypeShortText, false
	}
	return v, true
}

// FieldTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func FieldTypeFromStringIgnoreCase(raw string) (AutoFormFieldType, bool) {
	v, ok := FieldTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _FieldTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_FieldTypeStringToValueMap = map[string]AutoFormFieldType{
		"short_text":     AutoFormFieldTypeShortText,
		"long_text":      AutoFormFieldTypeLongText,
		"markdown":       AutoFormFieldTypeMarkdown,
		"select":         AutoFormFieldTypeSelect,
		"staticDropdown": AutoFormFieldTypeStaticDropdown,
		"number":         AutoFormFieldTypeNumber,
		"checkbox":       AutoFormFieldTypeCheckbox,
		"oauth":          AutoFormFieldTypeOauth2,
		"secret":         AutoFormFieldTypeSecretAuth,
		"custom_auth":    AutoFormFieldTypeCustomAuth,
		"array":          AutoFormFieldTypeArray,
		"fieldset":       AutoFormFieldTypeObject,
		"basic_auth":     AutoFormFieldTypeBasicAuth,
		"json":           AutoFormFieldTypeJSON,
		"datetime":       AutoFormFieldTypeDateTime,
		"file":           AutoFormFieldTypeFile,
		"file_string":    AutoFormFieldTypeFileString,
		"boolean":        AutoFormFieldTypeBoolean,
		"dynamic":        AutoFormFieldTypeDynamic,
		"code":           AutoFormFieldTypeCode,
		"richtext":       AutoFormFieldTypeRichText,
		"branch":         AutoFormFieldTypeBranch,
		"wrapper":        AutoFormFieldTypeWrapper,
		"framework":      AutoFormFieldTypeCodeFramework,
	}
	_FieldTypeLowerStringToValueMap = map[string]AutoFormFieldType{
		"short_text":     AutoFormFieldTypeShortText,
		"long_text":      AutoFormFieldTypeLongText,
		"markdown":       AutoFormFieldTypeMarkdown,
		"select":         AutoFormFieldTypeSelect,
		"staticDropdown": AutoFormFieldTypeStaticDropdown,
		"number":         AutoFormFieldTypeNumber,
		"checkbox":       AutoFormFieldTypeCheckbox,
		"oauth":          AutoFormFieldTypeOauth2,
		"secret":         AutoFormFieldTypeSecretAuth,
		"custom_auth":    AutoFormFieldTypeCustomAuth,
		"array":          AutoFormFieldTypeArray,
		"fieldset":       AutoFormFieldTypeObject,
		"basic_auth":     AutoFormFieldTypeBasicAuth,
		"json":           AutoFormFieldTypeJSON,
		"datetime":       AutoFormFieldTypeDateTime,
		"file":           AutoFormFieldTypeFile,
		"file_string":    AutoFormFieldTypeFileString,
		"boolean":        AutoFormFieldTypeBoolean,
		"dynamic":        AutoFormFieldTypeDynamic,
		"code":           AutoFormFieldTypeCode,
		"richtext":       AutoFormFieldTypeRichText,
		"branch":         AutoFormFieldTypeBranch,
		"wrapper":        AutoFormFieldTypeWrapper,
		"framework":      AutoFormFieldTypeCodeFramework,
	}
)
