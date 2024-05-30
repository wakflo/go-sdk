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
	"sync"
)

type BaseComponentField struct {
	schema   *sdkcore.AutoFormSchema
	builder  *SchemaBuilder
	Required bool
}

func NewBaseComponentField() *BaseComponentField {
	c := &BaseComponentField{schema: &sdkcore.AutoFormSchema{}, builder: NewSchemaBuilder()}
	return c
}

type DefaultBaseComponentField struct {
	mu       sync.Mutex
	schema   *sdkcore.AutoFormSchema
	builder  *SchemaBuilder
	Required bool
}

func NewDefaultBaseComponentField() *DefaultBaseComponentField {
	c := &DefaultBaseComponentField{schema: &sdkcore.AutoFormSchema{}, builder: NewSchemaBuilder()}
	return c
}

func (b *DefaultBaseComponentField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *DefaultBaseComponentField) SetDescription(desc string) *DefaultBaseComponentField {
	b.builder.WithDescription(desc)
	return b
}

func (b *DefaultBaseComponentField) SetDisplayName(title string) *DefaultBaseComponentField {
	b.builder.WithTitle(title)
	return b
}

func (b *DefaultBaseComponentField) SetRequired(required bool) *DefaultBaseComponentField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	return b
}

func (b *DefaultBaseComponentField) SetDisabled(disabled bool) *DefaultBaseComponentField {
	b.builder.schema.Presentation.Disabled = disabled
	return b
}

func (b *DefaultBaseComponentField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *DefaultBaseComponentField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *DefaultBaseComponentField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *DefaultBaseComponentField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *DefaultBaseComponentField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *DefaultBaseComponentField {
	b.builder.WithAllOf(schemas)
	return b
}
