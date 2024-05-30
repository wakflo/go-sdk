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
	"time"
)

type DateTimeField struct {
	*BaseComponentField
}

func NewDateTimeField() *DateTimeField {
	c := &DateTimeField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.String)
	c.builder.WithFieldType(sdkcore.DateTimeType)

	return c
}

func (b *DateTimeField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *DateTimeField) SetDefaultValue(defaultValue time.Time) *DateTimeField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *DateTimeField) SetMinimum(len time.Time) *DateTimeField {
	b.builder.WithMinimum(&len)
	return b
}

func (b *DateTimeField) SetMaximum(len time.Time) *DateTimeField {
	b.builder.WithMaximum(&len)
	return b
}

// rest
func (b *DateTimeField) SetDescription(desc string) *DateTimeField {
	b.builder.WithDescription(desc)
	return b
}

func (b *DateTimeField) SetDisplayName(title string) *DateTimeField {
	b.builder.WithTitle(title)
	return b
}

func (b *DateTimeField) SetRequired(required bool) *DateTimeField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *DateTimeField) SetDisabled(disabled bool) *DateTimeField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.Presentation.Disabled = disabled
	return b
}
