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
	"errors"
	"fmt"
	"io"
)

// PluginLanguage enum definition.
type PluginLanguage string

const (
	PluginLanguageRust   PluginLanguage = "rust"
	PluginLanguageGolang PluginLanguage = "golang"
	PluginLanguagePython PluginLanguage = "python"
	PluginLanguageZig    PluginLanguage = "zig"
)

// Values provides list valid values for Enum.
func (PluginLanguage) Values() (kinds []string) {
	for _, s := range []PluginLanguage{PluginLanguageRust, PluginLanguageGolang, PluginLanguagePython, PluginLanguageZig} {
		kinds = append(kinds, string(s))
	}
	return
}

func (e PluginLanguage) String() string {
	switch e {
	case PluginLanguagePython:
		return "python"
	case PluginLanguageZig:
		return "zig"
	case PluginLanguageGolang:
		return "golang"
	case PluginLanguageRust:
		return "rust"
	default:
		return "rust"
	}
}

func (e PluginLanguage) IsValid() bool {
	switch e {
	case PluginLanguagePython:
	case PluginLanguageZig:
	case PluginLanguageGolang:
	case PluginLanguageRust:
		return true
	default:
		return false
	}
	return false
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (e PluginLanguage) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	rsp := PluginLanguage(str)
	if !rsp.IsValid() {
		return fmt.Errorf("%s is not a valid trigger-type", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (e PluginLanguage) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(e.String()))
}

// PluginCompiler enum definition.
type PluginCompiler string

const (
	Wasm   PluginCompiler = "wasm"
	Wasi   PluginCompiler = "wasi"
	Wasix  PluginCompiler = "wasix"
	TinyGo PluginCompiler = "tinygo"
)

// Values provides list valid values for Enum.
func (PluginCompiler) Values() (kinds []string) {
	for _, s := range []PluginCompiler{Wasm, Wasi, Wasix, TinyGo} {
		kinds = append(kinds, string(s))
	}
	return
}

func (e PluginCompiler) String() string {
	switch e {
	case Wasm:
		return "wasm"
	case Wasi:
		return "wasi"
	case Wasix:
		return "wasix"
	case TinyGo:
		return "tinygo"
	default:
		return "wasi"
	}
}

func (e PluginCompiler) IsValid() bool {
	switch e {
	case Wasm:
	case Wasi:
	case Wasix:
	case TinyGo:
		return true
	default:
		return false
	}
	return false
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (e PluginCompiler) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	rsp := PluginCompiler(str)
	if !rsp.IsValid() {
		return fmt.Errorf("%s is not a valid trigger-type", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (e PluginCompiler) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(e.String()))
}
