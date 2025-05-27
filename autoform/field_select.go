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

type SelectField struct {
	*BaseComponentField
}

func NewSelectField() *SelectField {
	c := &SelectField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.String)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeSelect)

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

func (b *SelectField) SeMultiSelect(enable bool) *SelectField {
	b.builder.schema.UIProps.Multiple = enable
	return b
}

func (b *SelectField) SetRequired(required bool) *SelectField {
	b.Required = required
	b.builder = b.builder.WithFieldRequired(required)
	return b
}

func (b *SelectField) SetDisabled(disabled bool) *SelectField {
	b.builder = b.builder.WithFieldDisabled(disabled)
	return b
}

func (b *SelectField) SetPlaceholder(placeholder string) *SelectField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *SelectField) SetLabel(label string) *SelectField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *SelectField) SetHint(hint string) *SelectField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *SelectField) SetHidden(hidden bool) *SelectField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
