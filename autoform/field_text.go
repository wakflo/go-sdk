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

package autoform

import (
	sdkcore "github.com/wakflo/go-sdk/core"
)

// BaseTextField is a type that represents a text field component in a form. It is a composition of BaseComponentField, and inherits its properties and methods.
// BaseTextField type has the following fields:
// - schema: AutoFormSchema
// - builder: *SchemaBuilder
// - Required: bool
// The methods available for BaseTextField type are:
// - Build(): AutoFormSchema
// - SetDescription(desc string) *SchemaBuilder
// - SetDisplayName(title string) *SchemaBuilder
// - SetRequired(required bool) *SchemaBuilder
// Example usage of BaseTextField:
// Create a new BaseTextField instance:
//
//	func newBaseTextField() *BaseTextField {
//	  c := &BaseTextField{
//	    BaseComponentField: NewBaseComponentFieldBuilder(),
//	  }
//	  c.builder.WithType(String)
//	  return c
//	}
//
// Build the AutoFormSchema for BaseTextField:
//
//	func (b *BaseTextField) Build() AutoFormSchema {
//	  b.schema = b.builder.Build()
//	  return b.schema
//	}
//
// Set a default value for BaseTextField:
//
//	func (b *BaseTextField) SetDefaultValue(defaultValue string) *SchemaBuilder {
//	  return b.builder.WithDefault(defaultValue)
//	}
//
// Set the minimum length for BaseTextField:
//
//	func (b *BaseTextField) SetMinLength(len int) *SchemaBuilder {
//	  return b.builder.WithMinLength(&len)
//	}
//
// Set the maximum length for BaseTextField:
//
//	func (b *BaseTextField) SetMaxLength(len int) *SchemaBuilder {
//	  return b.builder.WithMaxLength(&len)
//	}
//
// BaseTextField can be extended by other types, such as ShortTextField, which embeds BaseTextField.
type BaseTextField struct {
	*BaseComponentField
}

func newBaseTextField() *BaseTextField {
	c := &BaseTextField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.String)

	return c
}

func (b *BaseTextField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *BaseTextField) SetDefaultValue(defaultValue string) *BaseTextField {
	b.builder.WithDefault(defaultValue)
	b.builder.schema.UIProps.InitialValue = defaultValue
	return b
}

func (b *BaseTextField) SetMinLength(len int) *BaseTextField {
	b.builder.WithMinLength(&len)
	return b
}

func (b *BaseTextField) SetMaxLength(len int) *BaseTextField {
	b.builder.WithMaxLength(&len)
	return b
}

// ----- new

func (b *BaseTextField) SetPlaceholder(placeholder string) *BaseTextField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *BaseTextField) SetLabel(label string) *BaseTextField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *BaseTextField) SetHint(hint string) *BaseTextField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

// rest
func (b *BaseTextField) SetDescription(desc string) *BaseTextField {
	b.builder.WithDescription(desc)
	return b
}

func (b *BaseTextField) SetDisplayName(title string) *BaseTextField {
	b.builder.WithTitle(title)
	return b
}

func (b *BaseTextField) SetRequired(required bool) *BaseTextField {
	b.Required = required
	b.builder = b.builder.WithFieldRequired(required)
	return b
}

func (b *BaseTextField) SetDisabled(disabled bool) *BaseTextField {
	b.builder = b.builder.WithFieldDisabled(disabled)
	return b
}

func (b *BaseTextField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *BaseTextField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *BaseTextField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *BaseTextField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *BaseTextField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *BaseTextField {
	b.builder.WithAllOf(schemas)
	return b
}

func (b *BaseTextField) SetHidden(hidden bool) *BaseTextField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}

// ShortTextField is a type that represents a short text field. It inherits from BaseTextField.
type ShortTextField struct {
	*BaseTextField
}

// NewShortTextField creates a new instance of ShortTextField.
// It initializes the BaseTextField using the newBaseTextField function
// and sets the AutoFormFieldType to ShortTextType using the builder.
func NewShortTextField() *ShortTextField {
	c := &ShortTextField{
		BaseTextField: newBaseTextField(),
	}
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeShortText)

	return c
}

// LongTextField is a type that represents a long text field component in a form. It is an extension of BaseTextField, inheriting its properties and methods.
// LongTextField type does not have any additional fields or methods compared to BaseTextField.
type LongTextField struct {
	*BaseTextField
}

// NewLongTextField creates a new instance of LongTextField.
// It initializes the BaseTextField using the newBaseTextField function
// and sets the AutoFormFieldType to LongTextType using the builder
func NewLongTextField() *LongTextField {
	c := &LongTextField{
		BaseTextField: newBaseTextField(),
	}
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeLongText)

	return c
}
