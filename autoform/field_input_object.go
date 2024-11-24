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

type InputMapField struct {
	*BaseComponentField
}

func NewInputMapField() *InputMapField {
	c := &InputMapField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Object)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeWrapper)
	c.builder.WithDescription("input")
	c.builder.WithTitle("Input")

	return c
}

func (b *InputMapField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *InputMapField) SetProperties(properties map[string]*sdkcore.AutoFormSchema) *InputMapField {
	var required []string
	order := make([]string, 0, len(properties))

	for key, schema := range properties {
		order = append(order, key)
		if schema.IsRequired {
			required = append(required, key)
		}
	}

	b.builder.WithProperties(properties)
	b.builder.WithRequired(required)
	b.builder.WithOrder(order)
	return b
}

func (b *InputMapField) SetOrder(order []string) *InputMapField {
	b.builder.WithOrder(order)
	return b
}

func (b *InputMapField) SetRequired(required bool) *InputMapField {
	b.Required = required
	b.builder = b.builder.WithFieldRequired(required)
	return b
}

func (b *InputMapField) SetDisabled(disabled bool) *InputMapField {
	b.builder = b.builder.WithFieldDisabled(disabled)
	return b
}

func (b *InputMapField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *InputMapField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *InputMapField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *InputMapField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *InputMapField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *InputMapField {
	b.builder.WithAllOf(schemas)
	return b
}

func (b *InputMapField) SetPlaceholder(placeholder string) *InputMapField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *InputMapField) SetLabel(label string) *InputMapField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *InputMapField) SetHint(hint string) *InputMapField {
	b.builder.schema.UIProps.Hint = hint
	return b
}
