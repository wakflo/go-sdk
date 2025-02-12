package sdk

import "github.com/wakflo/go-sdk/core"

type ScheduledOutput = core.ScheduleTriggerCriteria
type WebhookOutput = core.WebhookTriggerCriteria

func OutputScheduled(criteria ScheduledOutput) ScheduledOutput {
	return ScheduledOutput{
		CronExpression: criteria.CronExpression,
		StartTime:      criteria.StartTime,
		EndTime:        criteria.EndTime,
		TimeZone:       criteria.TimeZone,
		Enabled:        true,
	}
}

func OutputWebhook(criteria WebhookOutput) WebhookOutput {
	return WebhookOutput{
		Endpoint:         criteria.Endpoint,
		AuthEnabled:      criteria.AuthEnabled,
		Headers:          criteria.Headers,
		HttpMethod:       criteria.HttpMethod,
		QueryParams:      criteria.QueryParams,
		ValidationSecret: criteria.ValidationSecret,
		Enabled:          true,
	}
}
