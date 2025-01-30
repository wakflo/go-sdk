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

type ConditionField struct {
	*BaseComponentField
	props map[string]sdkcore.AutoFormSchema
}

func NewConditionField() *ConditionField {
	c := &ConditionField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.Array)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeCondition)

	return c
}

func (b *ConditionField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *ConditionField) SetDescription(desc string) *ConditionField {
	b.builder.WithDescription(desc)
	return b
}

func (b *ConditionField) SetDisplayName(title string) *ConditionField {
	b.builder.WithTitle(title)
	return b
}

func (b *ConditionField) SetRequired(required bool) *ConditionField {
	b.Required = required
	b.builder.WithFieldRequired(required)
	return b
}

func (b *ConditionField) SetDisabled(disabled bool) *ConditionField {
	b.builder = b.builder.WithFieldDisabled(disabled)
	return b
}

func (b *ConditionField) SetPlaceholder(placeholder string) *ConditionField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *ConditionField) SetLabel(label string) *ConditionField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *ConditionField) SetHint(hint string) *ConditionField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *ConditionField) SetHidden(hidden bool) *ConditionField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
