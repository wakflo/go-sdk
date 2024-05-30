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

package pluginenums

import (
	"fmt"
	"io"
)

// PluginLanguage enum definition.
type PluginLanguage string

const (
	Rust   PluginLanguage = "rust"
	Golang PluginLanguage = "golang"
	Python PluginLanguage = "python"
	Zig    PluginLanguage = "zig"
)

// Values provides list valid values for Enum.
func (PluginLanguage) Values() (kinds []string) {
	for _, s := range []PluginLanguage{Rust, Golang, Python, Zig} {
		kinds = append(kinds, string(s))
	}
	return
}

func (e PluginLanguage) String() string {
	switch e {
	case Python:
		return "python"
	case Zig:
		return "zig"
	case Golang:
		return "golang"
	case Rust:
		return "rust"
	default:
		return "rust"
	}
}

func (e PluginLanguage) IsValid() bool {
	switch e {
	case Python:
	case Zig:
	case Golang:
	case Rust:
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
		return fmt.Errorf("enums must be strings")
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
		return fmt.Errorf("enums must be strings")
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
