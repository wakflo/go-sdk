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

type NumberField struct {
	*BaseComponentField
}

func NewNumberField() *NumberField {
	c := &NumberField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Number)
	c.builder.WithFieldType(sdkcore.NumberType)

	return c
}

func (b *NumberField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *NumberField) SetDefaultValue(defaultValue interface{}) *NumberField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *NumberField) SetMinimum(len int) *NumberField {
	b.builder.WithMinimum(&len)
	return b
}

func (b *NumberField) SetMaximum(len int) *NumberField {
	b.builder.WithMaximum(&len)
	return b
}

// rest
func (b *NumberField) SetDescription(desc string) *NumberField {
	b.builder.WithDescription(desc)
	return b
}

func (b *NumberField) SetDisplayName(title string) *NumberField {
	b.builder.WithTitle(title)
	return b
}

func (b *NumberField) SetRequired(required bool) *NumberField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *NumberField) SetDisabled(disabled bool) *NumberField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.Presentation.Disabled = disabled
	return b
}

func (b *NumberField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *NumberField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *NumberField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *NumberField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *NumberField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *NumberField {
	b.builder.WithAllOf(schemas)
	return b
}
