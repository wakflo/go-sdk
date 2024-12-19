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

type ArrayField struct {
	*BaseComponentField
}

func NewArrayField() *ArrayField {
	c := &ArrayField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Array)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeArray)

	return c
}

func (b *ArrayField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *ArrayField) SetDefaultValue(defaultValue interface{}) *ArrayField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *ArrayField) SetItems(item *sdkcore.AutoFormSchema) *ArrayField {
	if item == nil {
		return b
	}

	if item.Type == sdkcore.Object || item.Type == sdkcore.Array {
		b.builder.WithItems(item)
		return b
	}

	b.builder.WithItems(
		NewObjectField().
			SetLabel(item.Title).
			SetDescription(item.Description).
			SetPlaceholder(item.UIProps.Placeholder).
			SetProperties(map[string]*sdkcore.AutoFormSchema{
				"value": item,
			}).
			Build(),
	)

	return b
}

func (b *ArrayField) SetUnique(unique bool) *ArrayField {
	b.builder.WithUniqueItems(unique)
	return b
}

// rest
func (b *ArrayField) SetDescription(desc string) *ArrayField {
	b.builder.WithDescription(desc)
	return b
}

func (b *ArrayField) SetDisplayName(title string) *ArrayField {
	b.builder.WithTitle(title)
	return b
}

func (b *ArrayField) SetRequired(required bool) *ArrayField {
	b.Required = required
	b.builder = b.builder.WithFieldRequired(required)
	return b
}

func (b *ArrayField) SetDisabled(disabled bool) *ArrayField {
	b.builder.schema.UIProps.Disabled = disabled
	b.builder.schema.Disabled = disabled
	return b
}

func (b *ArrayField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *ArrayField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *ArrayField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *ArrayField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *ArrayField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *ArrayField {
	b.builder.WithAllOf(schemas)
	return b
}

func (b *ArrayField) SetPlaceholder(placeholder string) *ArrayField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *ArrayField) SetHidden(hidden bool) *ArrayField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}

func (b *ArrayField) SetLabel(label string) *ArrayField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *ArrayField) SetHint(hint string) *ArrayField {
	b.builder.schema.UIProps.Hint = hint
	return b
}
