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

type ObjectField struct {
	*BaseComponentField
}

func NewObjectField() *ObjectField {
	c := &ObjectField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Object)
	c.builder.WithFieldType(sdkcore.ObjectType)

	return c
}

func (b *ObjectField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *ObjectField) SetDefaultValue(defaultValue interface{}) *ObjectField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *ObjectField) SetProperties(properties map[string]*sdkcore.AutoFormSchema) *ObjectField {
	var required []string
	var order []string

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

// rest
func (b *ObjectField) SetOrder(order []string) *ObjectField {
	b.builder.WithOrder(order)
	return b
}

func (b *ObjectField) SetDescription(desc string) *ObjectField {
	b.builder.WithDescription(desc)
	return b
}

func (b *ObjectField) SetDisplayName(title string) *ObjectField {
	b.builder.WithTitle(title)
	return b
}

func (b *ObjectField) SetRequired(required bool) *ObjectField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *ObjectField) SetDisabled(disabled bool) *ObjectField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.Presentation.Disabled = disabled
	return b
}

func (b *ObjectField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *ObjectField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *ObjectField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *ObjectField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *ObjectField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *ObjectField {
	b.builder.WithAllOf(schemas)
	return b
}
