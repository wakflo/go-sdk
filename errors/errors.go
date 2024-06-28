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

//nolint:mnd
package errors

import (
	"fmt"
)

type DetailedError struct {
	// a custom Hatchet error code
	// example: 1400
	Code uint `json:"code"`

	// a reason for this error
	Reason string `json:"reason"`

	// a description for this error
	// example: A descriptive error message
	Description string `json:"description"`

	// a link to the documentation for this error, if it exists
	// example: github.com/armado
	DocsLink string `json:"docs_link"`
}

func (e DetailedError) Error() string {
	errStr := fmt.Sprintf("error %d: %s", e.Code, e.Description)

	if e.DocsLink != "" {
		errStr = fmt.Sprintf("%s, see %s", errStr, e.DocsLink)
	}

	return errStr
}

func NewError(code uint, reason, description, docsLink string) *DetailedError {
	return &DetailedError{
		Code:        code,
		Reason:      reason,
		Description: description,
		DocsLink:    docsLink,
	}
}

func NewErrInternal(err error) *DetailedError {
	return NewError(500, "Internal Server Error", err.Error(), "")
}

func NewErrForbidden(err error) *DetailedError {
	return NewError(403, "Forbidden", err.Error(), "")
}
