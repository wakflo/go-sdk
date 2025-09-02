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

package context

import (
	"context"
	"io"

	"github.com/rs/xid"
)

type FileOutput struct {
	ID         xid.ID `json:"id"`
	ContentURL string `json:"contentUrl"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
}

type FileResource interface {
	GetFileAsBytes(ctx context.Context, fileId string) ([]byte, error)
	GetFile(ctx context.Context, fileId string) (io.ReadCloser, error)
	UploadFile(ctx context.Context, name string, content io.Reader) (*FileOutput, error)
}
