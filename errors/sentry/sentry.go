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

package sentry

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
)

type SentryAlerter struct {
	client *sentry.Client
}

func noIntegrations(ints []sentry.Integration) []sentry.Integration {
	return []sentry.Integration{}
}

type SentryAlerterOpts struct {
	DSN         string
	Environment string
}

func NewSentryAlerter(opts *SentryAlerterOpts) (*SentryAlerter, error) {
	sentryClient, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:              opts.DSN,
		AttachStacktrace: true,
		Integrations:     noIntegrations,
		Environment:      opts.Environment,
	})
	if err != nil {
		return nil, err
	}

	return &SentryAlerter{
		client: sentryClient,
	}, nil
}

func (s *SentryAlerter) SendAlert(ctx context.Context, err error, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}

	scope := sentry.NewScope()

	for key, val := range data {
		scope.SetTag(key, fmt.Sprintf("%v", val))
	}

	s.client.CaptureException(
		err,
		&sentry.EventHint{
			Data: data,
		},
		scope,
	)
}
