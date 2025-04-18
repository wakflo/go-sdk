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

type BranchType string

const (
	BranchTypeDefault   BranchType = "DEFAULT"
	BranchTypeCondition BranchType = "CONDITION"
)

func (BranchType) SQLTypeName() string {
	return "branch_type"
}

// Values returns a slice of all String values of the enum.
func (BranchType) Values() []string {
	return []string{
		"DEFAULT",
		"CONDITION",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j BranchType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j BranchType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("BranchType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern BranchType(%d) instead.
func (_j BranchType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("BranchType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for BranchType.
func (_j BranchType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as BranchType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for BranchType.
func (_j *BranchType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("BranchType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for BranchType.
func (_j BranchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for BranchType.
func (_j *BranchType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of BranchType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("BranchType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for BranchType.
func (_j BranchType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as BranchType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for BranchType.
func (_j *BranchType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("BranchType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("BranchType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for BranchType.
func (_j *BranchType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of BranchType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("BranchType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for BranchType.
func (_j BranchType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as BranchType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for BranchType.
func (_j *BranchType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("BranchType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for BranchType.
func (_j BranchType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as BranchType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for BranchType.
func (_j *BranchType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("BranchType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchType", str)
	}
	return nil
}

// BranchTypeFromString determines the enum value with an exact case match.
func BranchTypeFromString(raw string) (BranchType, bool) {
	v, ok := _BranchTypeStringToValueMap[raw]
	if !ok {
		return BranchTypeCondition, false
	}
	return v, true
}

// BranchTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func BranchTypeFromStringIgnoreCase(raw string) (BranchType, bool) {
	v, ok := BranchTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _BranchTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_BranchTypeStringToValueMap = map[string]BranchType{
		"DEFAULT":   BranchTypeDefault,
		"CONDITION": BranchTypeCondition,
	}
	_BranchTypeLowerStringToValueMap = map[string]BranchType{
		"DEFAULT":   BranchTypeDefault,
		"CONDITION": BranchTypeCondition,
	}
)
