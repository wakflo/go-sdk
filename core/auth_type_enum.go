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

var ErrNoValidEnum = errors.New("not a valid enum")

type AuthType string

const (
	None   AuthType = "none"
	Basic  AuthType = "basic"
	Secret AuthType = "secret"
	APIKey AuthType = "api_key"
	OAuth2 AuthType = "oauth2"
	Custom AuthType = "custom"
)

func (AuthType) SQLTypeName() string {
	return "auth_type"
}

// Values returns a slice of all String values of the enum.
func (AuthType) Values() []string {
	return []string{
		"basic",
		"secret",
		"api_key",
		"oauth2",
		"none",
		"custom",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j AuthType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j AuthType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("AuthType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern AuthType(%d) instead.
func (_j AuthType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("AuthType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for AuthType.
func (_j AuthType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for AuthType.
func (_j *AuthType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for AuthType.
func (_j AuthType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for AuthType.
func (_j *AuthType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AuthType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for AuthType.
func (_j AuthType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AuthType.
func (_j *AuthType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("AuthType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for AuthType.
func (_j *AuthType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AuthType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for AuthType.
func (_j AuthType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for AuthType.
func (_j *AuthType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for AuthType.
func (_j AuthType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for AuthType.
func (_j *AuthType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// AuthTypeFromString determines the enum value with an exact case match.
func AuthTypeFromString(raw string) (AuthType, bool) {
	v, ok := _AuthTypeStringToValueMap[raw]
	if !ok {
		return None, false
	}
	return v, true
}

// AuthTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func AuthTypeFromStringIgnoreCase(raw string) (AuthType, bool) {
	v, ok := AuthTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _AuthTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_AuthTypeStringToValueMap = map[string]AuthType{
		"basic":   Basic,
		"secret":  Secret,
		"api_key": APIKey,
		"oauth2":  OAuth2,
		"none":    None,
		"custom":  Custom,
	}
	_AuthTypeLowerStringToValueMap = map[string]AuthType{
		"basic":   Basic,
		"secret":  Secret,
		"api_key": APIKey,
		"oauth2":  OAuth2,
		"none":    None,
		"custom":  Custom,
	}
)
