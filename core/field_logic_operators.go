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
	"fmt"
	"io"
	"slices"
	"strconv"
)

type FieldLogicOperators string

const (
	Eq      FieldLogicOperators = "=="
	EEq                         = "==="
	NEq                         = "!="
	NNEq                        = "!=="
	Not                         = "!"
	NNot                        = "!!"
	Or                          = "or"
	And                         = "and"
	Lt                          = "<"
	Gt                          = ">"
	GEq                         = ">="
	LEq                         = "<="
	Between                     = "between"
	Max                         = "max"
	Min                         = "min"
	Plus                        = "+"
	Minus                       = "-"
	Divide                      = "/"
	Modulo                      = "%"
	Map                         = "map"
	Reduce                      = "reduce"
	Filter                      = "filter"
	All                         = "all"
	//None                        = "none"
	Some   = "some"
	Merge  = "merge"
	In     = "in"
	Cat    = "cat"
	Substr = "substr"
)

func (FieldLogicOperators) SqlTypeName() string {
	return "auth_type"
}

// Values returns a slice of all String values of the enum.
func (FieldLogicOperators) Values() []string {
	return []string{
		"==",
		"===",
		"!=",
		"!==",
		"!",
		"!!",
		"or",
		"and",
		"<",
		">",
		">=",
		"<=",
		"between",
		"max",
		"min",
		"+",
		"-",
		"/",
		"%",
		"map",
		"reduce",
		"filter",
		"all",
		"none",
		"some",
		"merge",
		"in",
		"cat",
		"substr",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j FieldLogicOperators) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j FieldLogicOperators) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("FieldLogicOperators(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern FieldLogicOperators(%d) instead.
func (_j FieldLogicOperators) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("FieldLogicOperators(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for FieldLogicOperators.
func (_j FieldLogicOperators) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FieldLogicOperators. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for FieldLogicOperators.
func (_j *FieldLogicOperators) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return fmt.Errorf("FieldLogicOperators cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldLogicOperatorsFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FieldLogicOperators", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for FieldLogicOperators.
func (_j FieldLogicOperators) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for FieldLogicOperators.
func (_j *FieldLogicOperators) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of FieldLogicOperators: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return fmt.Errorf("FieldLogicOperators cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldLogicOperatorsFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FieldLogicOperators", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for FieldLogicOperators.
func (_j FieldLogicOperators) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FieldLogicOperators. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FieldLogicOperators.
func (_j *FieldLogicOperators) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("FieldLogicOperators should be a string, got %q", data)
	}
	if len(str) == 0 {
		return fmt.Errorf("FieldLogicOperators cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldLogicOperatorsFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FieldLogicOperators", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for FieldLogicOperators.
func (_j *FieldLogicOperators) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of FieldLogicOperators: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return fmt.Errorf("FieldLogicOperators cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldLogicOperatorsFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FieldLogicOperators", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for FieldLogicOperators.
func (_j FieldLogicOperators) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FieldLogicOperators. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for FieldLogicOperators.
func (_j *FieldLogicOperators) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return fmt.Errorf("FieldLogicOperators cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldLogicOperatorsFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FieldLogicOperators", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for FieldLogicOperators.
func (_j FieldLogicOperators) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as FieldLogicOperators. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for FieldLogicOperators.
func (_j *FieldLogicOperators) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return fmt.Errorf("FieldLogicOperators cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = FieldLogicOperatorsFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a FieldLogicOperators", str)
	}
	return nil
}

// FieldLogicOperatorsFromString determines the enum value with an exact case match.
func FieldLogicOperatorsFromString(raw string) (FieldLogicOperators, bool) {
	v, ok := _FieldLogicOperatorsStringToValueMap[raw]
	if !ok {
		return Eq, false
	}
	return v, true
}

// FieldLogicOperatorsFromStringIgnoreCase determines the enum value with a case-insensitive match.
func FieldLogicOperatorsFromStringIgnoreCase(raw string) (FieldLogicOperators, bool) {
	v, ok := FieldLogicOperatorsFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _FieldLogicOperatorsLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_FieldLogicOperatorsStringToValueMap = map[string]FieldLogicOperators{
		"==":      Eq,
		"===":     EEq,
		"!=":      NEq,
		"!==":     NNEq,
		"!":       Not,
		"!!":      NNot,
		"or":      Or,
		"and":     And,
		"<":       Lt,
		">":       Gt,
		">=":      GEq,
		"<=":      LEq,
		"between": Between,
		"max":     Max,
		"min":     Min,
		"+":       Plus,
		"-":       Minus,
		"/":       Divide,
		"%":       Modulo,
		"map":     Map,
		"reduce":  Reduce,
		"filter":  Filter,
		"all":     All,
		//"none":    None,
		"some":   Some,
		"merge":  Merge,
		"in":     In,
		"cat":    Cat,
		"substr": Substr,
	}
	_FieldLogicOperatorsLowerStringToValueMap = map[string]FieldLogicOperators{
		"==":      Eq,
		"===":     EEq,
		"!=":      NEq,
		"!==":     NNEq,
		"!":       Not,
		"!!":      NNot,
		"or":      Or,
		"and":     And,
		"<":       Lt,
		">":       Gt,
		">=":      GEq,
		"<=":      LEq,
		"between": Between,
		"max":     Max,
		"min":     Min,
		"+":       Plus,
		"-":       Minus,
		"/":       Divide,
		"%":       Modulo,
		"map":     Map,
		"reduce":  Reduce,
		"filter":  Filter,
		"all":     All,
		//"none":    None,
		"some":   Some,
		"merge":  Merge,
		"in":     In,
		"cat":    Cat,
		"substr": Substr,
	}
)
