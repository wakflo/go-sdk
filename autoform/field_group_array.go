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

type GroupArrayField struct {
	*BaseComponentField
}

func NewGroupArrayField() *GroupArrayField {
	c := &GroupArrayField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Array)
	c.builder.WithFieldType(sdkcore.GroupArrayType)

	return c
}

func (b *GroupArrayField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *GroupArrayField) SetDefaultValue(defaultValue interface{}) *GroupArrayField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *GroupArrayField) SetItems(item *sdkcore.AutoFormSchema) *GroupArrayField {
	b.builder.WithItems(item)
	return b
}

func (b *GroupArrayField) SetUnique(unique bool) *GroupArrayField {
	b.builder.WithUniqueItems(unique)
	return b
}

// rest
func (b *GroupArrayField) SetDescription(desc string) *GroupArrayField {
	b.builder.WithDescription(desc)
	return b
}

func (b *GroupArrayField) SetDisplayName(title string) *GroupArrayField {
	b.builder.WithTitle(title)
	return b
}

func (b *GroupArrayField) SetRequired(required bool) *GroupArrayField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *GroupArrayField) SetDisabled(disabled bool) *GroupArrayField {
	b.builder.schema.Presentation.Disabled = disabled
	b.builder.schema.Disabled = disabled
	return b
}

func (b *GroupArrayField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *GroupArrayField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *GroupArrayField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *GroupArrayField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *GroupArrayField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *GroupArrayField {
	b.builder.WithAllOf(schemas)
	return b
}
