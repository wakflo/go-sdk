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

package core

type BranchCondition struct {
	Operator      LogicalOperator `json:"operator"`
	LeftValue     any             `json:"leftValue"`
	RightValue    any             `json:"rightValue"`
	CaseSensitive bool            `json:"caseSensitive"`
}

type BranchSettings struct {
	MaxBranches     int                 `json:"maxBranches"`
	ExecutionType   BranchExecutionType `json:"executionType"`
	DefaultBranches []FlowBranch        `json:"defaultBranches,omitempty"`
}

type FlowBranch struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Type       BranchType        `json:"type"`
	Conditions []BranchCondition `json:"conditions"`
}
