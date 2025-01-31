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
	"errors"
	"fmt"
	"testing"

	sdkcore "github.com/wakflo/go-sdk/core"
	"golang.org/x/exp/maps"
)

type SpiderTestRuntime struct {
	instance *ConnectorPlugin
	t        *testing.T
}

func NewSpiderTest(t *testing.T, instance *ConnectorPlugin) *SpiderTestRuntime {
	return &SpiderTestRuntime{
		t:        t,
		instance: instance,
	}
}

func (s *SpiderTestRuntime) GetConfig() *ConnectorMetadata {
	return s.instance.Metadata
}

func (s *SpiderTestRuntime) Triggers() []*sdkcore.Trigger {
	return maps.Values(s.instance.TriggerMetadata)
}

func (s *SpiderTestRuntime) Operations() []*sdkcore.Action {
	return maps.Values(s.instance.OperationsMetadata)
}

func (s *SpiderTestRuntime) OperationConfig(name string) *sdkcore.Action {
	return s.instance.OperationsMetadata[name]
}

func (s *SpiderTestRuntime) TriggerConfig(name string) *sdkcore.Trigger {
	return s.instance.TriggerMetadata[name]
}

func (s *SpiderTestRuntime) RunOperation(name string, run *RunContext) (JSON, error) {
	ops, err := s.getOperationByName(name)
	if err != nil {
		return nil, err
	}
	return ops.Run(run)
}

func (s *SpiderTestRuntime) RunTrigger(name string, run *RunContext) (JSON, error) {
	trigger, err := s.getTriggerByName(name)
	if err != nil {
		return nil, err
	}
	return trigger.Run(run)
}

func (s *SpiderTestRuntime) RunAuth() error {
	return nil
}

func (s *SpiderTestRuntime) Close() error {
	// return s.db.Close()
	return nil
}

func (s *SpiderTestRuntime) getOperationByName(name string) (IOperation, error) {
	op, ok := s.instance.Operations[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("missing operation of name = %v", op.GetInfo().Name))
	}

	return op, nil
}

func (s *SpiderTestRuntime) getTriggerByName(name string) (ITrigger, error) {
	op, ok := s.instance.Triggers[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("missing trigger of name = %v", op.GetInfo().Name))
	}

	return op, nil
}
