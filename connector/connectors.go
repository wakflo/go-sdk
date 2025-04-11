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

package connector

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/wakflo/go-sdk/autoform"
	sdkcore "github.com/wakflo/go-sdk/core"
	"github.com/wakflo/go-sdk/internal/utils"
	"github.com/wakflo/go-sdk/validator"
)

// IConnectorPlugin is an interface that defines the methods for a connector plugin implementation.
// It provides functionality for authentication and testing.
type IConnectorPlugin interface {
	Authenticate(ctx context.Context)
	Test(ctx context.Context)
}

// ConnectorMetadata represents the metadata of a connector.
// It includes information such as the name, display name, description,
// version, authors, category, and logo of the connector.
type ConnectorMetadata struct {
	Name        string         `json:"name"`
	DisplayName string         `json:"display_name"`
	Description string         `json:"description"`
	Version     string         `json:"version"`
	Authors     []string       `json:"authors"`
	Group       ConnectorGroup `json:"group"`
	Categories  []string       `json:"categories"`
	Logo        string         `json:"logo"`
}

// ConnectorPlugin represents a plugin that connects to a service or application.
// It contains information about the plugin's creation arguments, version metadata, connector metadata,
// triggers, operations, trigger metadata, and operations metadata.
type ConnectorPlugin struct {
	CreateArgs         *CreateConnectorArgs
	VersionMetadata    *sdkcore.ConnectorVersionMetadata
	Metadata           *ConnectorMetadata
	Triggers           map[string]ITrigger
	Operations         map[string]IOperation
	TriggerMetadata    map[string]*sdkcore.Trigger
	OperationsMetadata map[string]*sdkcore.Action
	HasTrigger         bool
}

// NewConnectorPlugin creates a new ConnectorPlugin based on the provided CreateConnectorArgs.
// It validates the connector information and creates maps for triggers and operations.
// It also generates unique slugs for trigger and operation names.
// Finally, it creates the ConnectorMetadata and returns the ConnectorPlugin.
func NewConnectorPlugin(args *CreateConnectorArgs) (*ConnectorPlugin, error) {
	v := validator.NewDefaultValidator()
	err := ValidateConnectorInfo(args)
	if err != nil {
		return &ConnectorPlugin{CreateArgs: args}, err
	}

	triggers := map[string]ITrigger{}
	operations := map[string]IOperation{}
	triggerMetadata := map[string]*sdkcore.Trigger{}
	operationsMetadata := map[string]*sdkcore.Action{}
	hasTrigger := len(args.Triggers) > 0

	for _, trigger := range args.Triggers {
		info := trigger.GetInfo()
		info.Name = strings.TrimSpace(info.Name)
		key := utils.GenerateUniqueSlug(info.Name)

		// defCron := "*/2 * * * *"
		s := sdkcore.OldTriggerSettings{
			Type: info.Type,
		}

		if info.Settings != nil {
			// if info.Settings.Cron == nil {
			//	info.Settings.Cron = &defCron
			// }
		} else {
			info.Settings = &s
		}

		verr := errors.New(fmt.Sprintf("%s - [trigger:%s] ", args.Name, info.Name))
		err = v.Validate(info)
		if err != nil {
			return nil, errors.Join(verr, err)
		}

		t := &sdkcore.Trigger{
			Name:        key,
			DisplayName: info.Name,
			Description: info.Description,
			Input:       autoform.NewInputMapField().SetProperties(info.Input).Build(),
			// RequireAuth:   info.RequireAuth,
			// Auth:          info.Auth,
			SampleOutput:  info.SampleOutput,
			Settings:      info.Settings,
			Type:          info.Type,
			Documentation: info.Documentation,
			HelpText:      info.HelpText,
		}
		err = v.Validate(t)
		if err != nil {
			return nil, errors.Join(verr, err)
		}
		triggerMetadata[key] = t

		triggers[key] = trigger
	}

	for _, op := range args.Operations {
		info := op.GetInfo()
		info.Name = strings.TrimSpace(info.Name)
		err := ValidateOperationInfo(info)
		if err != nil {
			verr := errors.New(fmt.Sprintf("%s - [operation:%s] ", args.Name, info.Name))
			return nil, errors.Join(verr, err)
		}
		key := utils.GenerateUniqueSlug(info.Name)
		operationsMetadata[key] = &sdkcore.Action{
			Name:        key,
			DisplayName: info.Name,
			Description: info.Description,
			Input:       autoform.NewInputMapField().SetProperties(info.Input).Build(),
			// RequireAuth:   info.RequireAuth,
			// Auth:          info.Auth,
			SampleOutput:  info.SampleOutput,
			Documentation: info.Documentation,
			HelpText:      info.HelpText,
			Settings:      info.Settings,
		}

		operations[key] = op
	}

	metadata := &ConnectorMetadata{
		Name:        utils.GenerateUniqueSlug(args.Name),
		DisplayName: args.Name,
		Description: args.Description,
		Version:     args.Version,
		Logo:        args.Logo,
		Group:       args.Group,
		Categories:  args.Categories,
		Authors:     args.Authors,
	}

	return &ConnectorPlugin{
		CreateArgs:         args,
		Triggers:           triggers,
		Operations:         operations,
		OperationsMetadata: operationsMetadata,
		TriggerMetadata:    triggerMetadata,
		Metadata:           metadata,
		HasTrigger:         hasTrigger,
	}, nil
}

// Authenticate authenticates the connector plugin.
// It takes a context as an argument and does not return anything.
// Example usage:
//
//	ctx := context.Background()
//	plugin.Authenticate(ctx)
func (plugin *ConnectorPlugin) Authenticate(ctx context.Context) {
}

// Test is a method of the ConnectorPlugin struct that is used to test the connectivity of the connector. It takes a context.Context as a parameter.
func (plugin *ConnectorPlugin) Test(ctx context.Context) {}

// ValidateConnectorInfo validates the connector information provided
func ValidateConnectorInfo(info *CreateConnectorArgs) error {
	if info == nil {
		return errors.New("connector options missing")
	}

	if len(info.Name) < 1 {
		return errors.New("connector name missing")
	}

	if len(info.Description) < 1 {
		return errors.New("connector description missing")
	}

	if len(info.Version) < 1 {
		return errors.New("connector version missing")
	}

	return nil
}

// ValidateTriggerInfo validates the provided TriggerInfo object.
// It checks if the info is not nil and if the name and description fields are not empty.
func ValidateTriggerInfo(info *TriggerInfo) error {
	if info == nil {
		return errors.New("trigger info missing")
	}

	if len(info.Name) < 1 {
		return errors.New("trigger info name missing")
	}

	if len(info.Description) < 1 {
		return errors.New("trigger info description missing")
	}

	return nil
}

// ValidateOperationInfo validates the provided OperationInfo.
// It checks if the OperationInfo is nil, if the name is missing, and if the description is missing.
// Returns an error if any of the checks fail, or nil if the OperationInfo is valid.
func ValidateOperationInfo(info *OperationInfo) error {
	if info == nil {
		return errors.New("operation info missing")
	}

	if len(info.Name) < 1 {
		return errors.New("operation info name missing")
	}

	if len(info.Description) < 1 {
		return errors.New("operation info description missing")
	}

	return nil
}
