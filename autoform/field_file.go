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
	"io"

	sdkcore "github.com/wakflo/go-sdk/core"
)

type FileField struct {
	*BaseComponentField
}

func NewFileField() *FileField {
	c := &FileField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.String)
	c.builder.WithFieldType(sdkcore.FileStringType)

	return c
}

func (b *FileField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *FileField) SetDefaultValue(defaultValue interface{}) *FileField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *FileField) SetDescription(desc string) *FileField {
	b.builder.WithDescription(desc)
	return b
}

func (b *FileField) SetDisplayName(title string) *FileField {
	b.builder.WithTitle(title)
	return b
}

func (b *FileField) SetRequired(required bool) *FileField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *FileField) SetDisabled(disabled bool) *FileField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.Presentation.Disabled = disabled
	return b
}

type File struct {
	// FileExtension: Output only. The final component of
	// `fullFileExtension`. This is only available for files with binary
	// content in Google Drive.
	Extension string `json:"extension,omitempty"`

	// FileExtension: Output only. The final component of
	// `fullFileExtension`. This is only available for files with binary
	// content in Google Drive.
	Mime string `json:"mime,omitempty"`

	// Name: The name of the file. This is not necessarily unique within a
	// folder. Note that for immutable items such as the top level folders
	// of shared drives, My Drive root folder, and Application Data folder
	// the name is constant.
	Name string `json:"name,omitempty"`

	// Size: Output only. Size in bytes of blobs and first party editor
	// files. Won't be populated for files that have no size, like shortcuts
	// and folders.
	Size int64 `json:"size,omitempty,string"`

	Data io.Reader `json:"data"`
}
