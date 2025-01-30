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

package sdk

import (
	sdkcore "github.com/wakflo/go-sdk/core"
)

type CreateConnectorArgs struct {
	Name        string
	Description string
	Version     string
	Logo        string
	Group       ConnectorGroup
	Categories  []string
	Triggers    []ITrigger
	Operations  []IOperation
	Authors     []string

	// Documentation represents the field used to store the connector's documentation in markdown.
	Documentation string

	Type *sdkcore.ActionType
}

func CreateConnector(args *CreateConnectorArgs) (*ConnectorPlugin, error) {
	return NewConnectorPlugin(args)
}
