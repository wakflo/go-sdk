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

type BranchExecutionType string

const (
	BranchExecutionTypeFirstMatch BranchExecutionType = "FIRST_MATCH"
	BranchExecutionTypeAllMatches BranchExecutionType = "ALL_MATCHES"
)

func (BranchExecutionType) SQLTypeName() string {
	return "edge_type"
}

// Values returns a slice of all String values of the enum.
func (BranchExecutionType) Values() []string {
	return []string{
		"FIRST_MATCH",
		"ALL_MATCHES",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j BranchExecutionType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j BranchExecutionType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("BranchExecutionType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern BranchExecutionType(%d) instead.
func (_j BranchExecutionType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("BranchExecutionType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for BranchExecutionType.
func (_j BranchExecutionType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as BranchExecutionType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for BranchExecutionType.
func (_j *BranchExecutionType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("BranchExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchExecutionType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for BranchExecutionType.
func (_j BranchExecutionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for BranchExecutionType.
func (_j *BranchExecutionType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of BranchExecutionType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("BranchExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchExecutionType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for BranchExecutionType.
func (_j BranchExecutionType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as BranchExecutionType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for BranchExecutionType.
func (_j *BranchExecutionType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("BranchExecutionType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("BranchExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchExecutionType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for BranchExecutionType.
func (_j *BranchExecutionType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of BranchExecutionType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("BranchExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchExecutionType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for BranchExecutionType.
func (_j BranchExecutionType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as BranchExecutionType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for BranchExecutionType.
func (_j *BranchExecutionType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("BranchExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchExecutionType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for BranchExecutionType.
func (_j BranchExecutionType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as BranchExecutionType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for BranchExecutionType.
func (_j *BranchExecutionType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("BranchExecutionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = BranchExecutionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a BranchExecutionType", str)
	}
	return nil
}

// BranchExecutionTypeFromString determines the enum value with an exact case match.
func BranchExecutionTypeFromString(raw string) (BranchExecutionType, bool) {
	v, ok := _BranchExecutionTypeStringToValueMap[raw]
	if !ok {
		return BranchExecutionTypeFirstMatch, false
	}
	return v, true
}

// BranchExecutionTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func BranchExecutionTypeFromStringIgnoreCase(raw string) (BranchExecutionType, bool) {
	v, ok := BranchExecutionTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _BranchExecutionTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_BranchExecutionTypeStringToValueMap = map[string]BranchExecutionType{
		"FIRST_MATCH": BranchExecutionTypeFirstMatch,
		"ALL_MATCHES": BranchExecutionTypeAllMatches,
	}
	_BranchExecutionTypeLowerStringToValueMap = map[string]BranchExecutionType{
		"FIRST_MATCH": BranchExecutionTypeFirstMatch,
		"ALL_MATCHES": BranchExecutionTypeAllMatches,
	}
)
