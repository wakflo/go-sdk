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

package authenums

//go:generate go run -mod=mod github.com/mvrahden/go-enumer -serializers=json,binary,graphql,json,sql,text,yaml -transform=kebab -support=ent

//go:enum
type AuthHandlerType uint

const (
	None AuthHandlerType = iota
	Inbuilt
	Custom
)

func (AuthHandlerType) SQLTypeName() string {
	return "auth_handler_type"
}

//go:enum
type AuthType uint

const (
	Basic AuthType = iota
	Digest
	OAuth
	APIKey
	BearerToken
)

//go:enum
type AuthTokenStyle uint

const (
	Header AuthTokenStyle = iota
	QueryParams
	Cookie
)
