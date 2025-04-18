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

package oldcore

import (
	"context"
)

type DynamicFieldPaging struct {
	*GetDynamicOptionsInput
	Auth  *AuthContext `json:"auth,omitempty"`
	Input any          `json:"input,omitempty"`
	Ctx   context.Context
}

type DynamicFieldContext struct {
	*GetDynamicOptionsInput
	Auth  *AuthContext `json:"auth,omitempty"`
	Input any          `json:"input,omitempty"`
	Ctx   context.Context
}

func NewDynamicFieldContext(ctx DynamicFieldContext) *DynamicFieldContext {
	if ctx.Filter == nil {
		ctx.Filter = &DynamicOptionsFilterParams{
			Offset:     0,
			Limit:      0,
			FilterTerm: "",
		}
	}

	if ctx.Filter.Offset < 0 {
		ctx.Filter.Offset = 0
	}

	if ctx.Filter.Limit <= 0 {
		ctx.Filter.Limit = 20
	}

	return &ctx
}

func (c *DynamicFieldContext) GetContext() context.Context {
	return c.Ctx
}

func (c *DynamicFieldContext) Respond(data any, totalItems int) (*DynamicOptionsResponse, error) {
	// if c

	return &DynamicOptionsResponse{
		Metadata: OffsetPaginationMeta{
			Offset:     c.Filter.Offset,
			Limit:      c.Filter.Limit,
			TotalItems: totalItems,
			HasMore:    (c.Filter.Offset + c.Filter.Limit) < totalItems,
		},
		Items: data,
	}, nil
}

func (c *DynamicFieldContext) RespondJSON(data any, totalItems int) (JSON, error) {
	return c.Respond(data, totalItems)
}

type JSON = any
