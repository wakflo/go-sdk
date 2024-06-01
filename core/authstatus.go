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

	"github.com/wakflo/go-sdk/core/jobstatusenum"
)

type AuthStatus string

const (
	Disabled AuthStatus = "disabled"
	Active   AuthStatus = "active"
	Failed   AuthStatus = "failed"
)

func (AuthStatus) SQLTypeName() string {
	return "auth_status"
}

// Values returns a slice of all String values of the enum.
func (AuthStatus) Values() []string {
	return []string{
		"disabled",
		"active",
		"failed",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j AuthStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j AuthStatus) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("AuthStatus(%v) is %w", _j, jobstatusenum.ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern AuthStatus(%d) instead.
func (_j AuthStatus) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("AuthStatus(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for AuthStatus.
func (_j AuthStatus) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for AuthStatus.
func (_j *AuthStatus) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("AuthStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthStatus", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for AuthStatus.
func (_j AuthStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for AuthStatus.
func (_j *AuthStatus) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AuthStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("AuthStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthStatus", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for AuthStatus.
func (_j AuthStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthStatus. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AuthStatus.
func (_j *AuthStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("AuthStatus should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("AuthStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthStatus", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for AuthStatus.
func (_j *AuthStatus) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AuthStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("AuthStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthStatus", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for AuthStatus.
func (_j AuthStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for AuthStatus.
func (_j *AuthStatus) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("AuthStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthStatus", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for AuthStatus.
func (_j AuthStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthStatus. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for AuthStatus.
func (_j *AuthStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("AuthStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthStatus", str)
	}
	return nil
}

// AuthStatusFromString determines the enum value with an exact case match.
func AuthStatusFromString(raw string) (AuthStatus, bool) {
	v, ok := _AuthStatusStringToValueMap[raw]
	if !ok {
		return Disabled, false
	}
	return v, true
}

// AuthStatusFromStringIgnoreCase determines the enum value with a case-insensitive match.
func AuthStatusFromStringIgnoreCase(raw string) (AuthStatus, bool) {
	v, ok := AuthStatusFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _AuthStatusLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_AuthStatusStringToValueMap = map[string]AuthStatus{
		"disabled": Disabled,
		"active":   Active,
		"failed":   Failed,
	}
	_AuthStatusLowerStringToValueMap = map[string]AuthStatus{
		"disabled": Disabled,
		"active":   Active,
		"failed":   Failed,
	}
)
