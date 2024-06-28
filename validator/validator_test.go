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

package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type nameResource struct {
	DisplayName string `validate:"spiderName"`
}

func TestValidatorInvalidName(t *testing.T) {
	v := newValidator()

	err := v.Struct(&nameResource{
		DisplayName: "&&!!",
	})

	require.ErrorContains(t, err, "validation for 'DisplayName' failed on the 'spiderName' tag", "should throw error on invalid name")
}

func TestValidatorValidName(t *testing.T) {
	v := newValidator()

	err := v.Struct(&nameResource{
		DisplayName: "test-name",
	})

	require.NoError(t, err, "no error")
}

type cronResource struct {
	Cron string `validate:"cron"`
}

func TestValidatorValidCron(t *testing.T) {
	v := newValidator()

	err := v.Struct(&cronResource{
		Cron: "*/5 * * * *",
	})

	require.NoError(t, err, "no error")
}

func TestValidatorInvalidCron(t *testing.T) {
	v := newValidator()

	err := v.Struct(&cronResource{
		Cron: "*/5 * * *",
	})

	require.ErrorContains(t, err, "validation for 'Cron' failed on the 'cron' tag", "should throw error on invalid cron")
}

func TestValidatorValidDuration(t *testing.T) {
	v := newValidator()

	err := v.Struct(&struct {
		Duration string `validate:"duration"`
	}{
		Duration: "5s",
	})

	require.NoError(t, err, "no error")
}

func TestValidatorInvalidDuration(t *testing.T) {
	v := newValidator()

	err := v.Struct(&struct {
		Duration string `validate:"duration"`
	}{
		Duration: "5",
	})

	require.ErrorContains(t, err, "validation for 'Duration' failed on the 'duration' tag", "should throw error on invalid duration")
}
