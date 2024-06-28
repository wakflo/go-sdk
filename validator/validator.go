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
	"encoding/json"
	"regexp"
	"time"
	"unicode"

	"github.com/Masterminds/semver/v3"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var NameRegex = regexp.MustCompile("^[a-zA-Z0-9\\.\\-_]+$") //nolint:gosimple

var CronRegex = regexp.MustCompile(`(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})`) //nolint:gosimple

func newValidator() *validator.Validate {
	validate := validator.New()

	_ = validate.RegisterValidation("spiderName", func(fl validator.FieldLevel) bool {
		return NameRegex.MatchString(fl.Field().String())
	})

	_ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return passwordValidation(fl.Field().String())
	})

	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		return IsValidUUID(fl.Field().String())
	})

	_ = validate.RegisterValidation("cron", func(fl validator.FieldLevel) bool {
		return CronRegex.MatchString(fl.Field().String())
	})

	// _ = validate.RegisterValidation("actionId", func(fl validator.FieldLevel) bool {
	//	action, err := types.ParseActionID(fl.Field().String())
	//	if err != nil {
	//		return false
	//	}
	//
	//	return action.Service != "" && action.Verb != ""
	// })

	_ = validate.RegisterValidation("semver", func(fl validator.FieldLevel) bool {
		_, err := semver.NewVersion(fl.Field().String())

		return err == nil
	})

	_ = validate.RegisterValidation("json", func(fl validator.FieldLevel) bool {
		return isValidJSON(fl.Field().String())
	})

	_ = validate.RegisterValidation("duration", func(fl validator.FieldLevel) bool {
		_, err := time.ParseDuration(fl.Field().String())

		return err == nil
	})

	return validate
}

func passwordValidation(pw string) bool {
	pwLen := len(pw)
	var hasNumber, hasUpper, hasLower bool

	for _, char := range pw {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		}
	}

	return hasNumber && hasUpper && hasLower && pwLen >= 8 && pwLen <= 32
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func isValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
