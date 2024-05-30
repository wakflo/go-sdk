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

type MultiSelectField struct {
	*BaseComponentField
}

func NewMultiSelectField() *MultiSelectField {
	c := &MultiSelectField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Array)
	c.builder.WithFieldType(sdkcore.MultiSelectType)

	return c
}

func (b *MultiSelectField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *MultiSelectField) SetDefaultValue(defaultValue interface{}) *MultiSelectField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *MultiSelectField) SetUnique(unique bool) *MultiSelectField {
	b.builder.WithUniqueItems(unique)
	return b
}

func (b *MultiSelectField) SetOptions(schemas []*sdkcore.AutoFormSchema) *MultiSelectField {
	item := NewDefaultBaseComponentField().SetAnyOf(schemas).Build()
	b.builder.WithItems(item)
	return b
}

// rest
func (b *MultiSelectField) SetDescription(desc string) *MultiSelectField {
	b.builder.WithDescription(desc)
	return b
}

func (b *MultiSelectField) SetDisplayName(title string) *MultiSelectField {
	b.builder.WithTitle(title)
	return b
}

func (b *MultiSelectField) SetRequired(required bool) *MultiSelectField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *MultiSelectField) SetDisabled(disabled bool) *MultiSelectField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.Presentation.Disabled = disabled
	return b
}

type SelectField struct {
	*BaseComponentField
}

func NewSelectField() *MultiSelectField {
	c := &MultiSelectField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.String)
	c.builder.WithFieldType(sdkcore.DropdownType)

	return c
}

func (b *SelectField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *SelectField) SetDefaultValue(defaultValue interface{}) *SelectField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *SelectField) SetUnique(unique bool) *SelectField {
	b.builder.WithUniqueItems(unique)
	return b
}

func (b *SelectField) SetOptions(schemas []*sdkcore.AutoFormSchema) *SelectField {
	b.builder.WithOneOf(schemas)
	return b
}

// rest
func (b *SelectField) SetDescription(desc string) *SelectField {
	b.builder.WithDescription(desc)
	return b
}

func (b *SelectField) SetDisplayName(title string) *SelectField {
	b.builder.WithTitle(title)
	return b
}

func (b *SelectField) SetRequired(required bool) *SelectField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	return b
}

func (b *SelectField) SetDisabled(disabled bool) *SelectField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.Presentation.Disabled = disabled
	return b
}
