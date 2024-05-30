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

type BranchField struct {
	*BaseComponentField
}

func NewBranchField() *BranchField {
	c := &BranchField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Array)
	c.builder.WithFieldType(sdkcore.BranchType)

	return c
}

func (b *BranchField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *BranchField) SetDefaultValue(defaultValue interface{}) *BranchField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *BranchField) SetItems(item *sdkcore.AutoFormSchema) *BranchField {
	b.builder.WithItems(item)
	return b
}

func (b *BranchField) SetUnique(unique bool) *BranchField {
	b.builder.WithUniqueItems(unique)
	return b
}

// rest
func (b *BranchField) SetDescription(desc string) *BranchField {
	b.builder.WithDescription(desc)
	return b
}

func (b *BranchField) SetDisplayName(title string) *BranchField {
	b.builder.WithTitle(title)
	return b
}

func (b *BranchField) SetRequired(required bool) *BranchField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *BranchField) SetDisabled(disabled bool) *BranchField {
	b.builder.schema.Presentation.Disabled = disabled
	b.builder.schema.Disabled = disabled
	return b
}

func (b *BranchField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *BranchField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *BranchField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *BranchField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *BranchField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *BranchField {
	b.builder.WithAllOf(schemas)
	return b
}
