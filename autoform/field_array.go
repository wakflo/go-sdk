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
	c.builder.WithFieldType(sdkcore.ArrayType)

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
	if item.Type == sdkcore.Integer || item.Type == sdkcore.String || item.Type == sdkcore.Number {
		if b.builder.schema.Presentation.Extras == nil {
			b.builder.schema.Presentation.Extras = map[string]any{}
		}

		b.builder.schema.Presentation.Extras["arrayType"] = item.Type
	}

	b.builder.WithItems(item)
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
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *ArrayField) SetDisabled(disabled bool) *ArrayField {
	b.builder.schema.Presentation.Disabled = disabled
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
