// Copyright 2022-present Wakflo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

//nolint:mnd
package validator

import (
	errors2 "errors"
	"fmt"
	"strconv"
	"strings"

	"encore.dev/beta/errs"
	v10Validator "github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
	"github.com/wakflo/go-sdk/errors"
)

const (
	EmailErr      = "Invalid email address"
	PasswordErr   = "Invalid password. Passwords must be at least 8 characters in length, contain an upper and lowercase letter, and contain at least one number."
	UUIDErr       = "Invalid UUID reference"
	SpiderNameErr = "Hatchet names must match the regex ^[a-zA-Z0-9\\.\\-_]+$"
	ActionIDErr   = "Invalid action WorkflowID. Action IDs must be in the format <integrationId>:<verb>"
	CronErr       = "Invalid cron expression"
	DurationErr   = "Invalid duration. Durations must be in the format <number><unit>, where unit is one of: 's', 'm', 'h', 'd', 'w', 'M', 'y'"
)

// type APIErrors gen.APIErrors

// func (a *APIErrors) String() string {
//	var sb strings.Builder
//
//	sb.WriteString("Validation failed with the following errors:\n")
//
//	for i, err := range a.Errors {
//		sb.WriteString(fmt.Sprintf("%d: %s\n", i, err.Description))
//	}
//
//	return sb.String()
//}

// Validator will validate the fields for a request object to ensure that
// the request is well-formed. For example, it searches for required fields
// or verifies that fields are of a semantic type (like email)
type Validator interface {
	// Validate accepts a generic struct for validating. It returns a request
	// error that is meant to be shown to the end user as a readable string.
	Validate(s interface{}) error

	ValidateAPI(s interface{}) error
}

// DefaultValidator uses the go-playground v10 validator for verifying that
// request objects are well-formed
type DefaultValidator struct {
	v10 *v10Validator.Validate
}

// NewDefaultValidator returns a Validator constructed from the go-playground v10
// validator
func NewDefaultValidator() Validator {
	return &DefaultValidator{newValidator()}
}

func (v *DefaultValidator) ValidateAPI(s interface{}) error {
	err := v.v10.Struct(s)

	if err == nil {
		return nil
	}

	// translate all validator errors
	errList, ok := err.(v10Validator.ValidationErrors)

	if !ok {
		return errors.NewErrInternal(fmt.Errorf("could not cast err to validator.ValidationErrors, type %T", err))
	}

	if len(errList) == 0 {
		return nil
	}

	// convert all validator errors to error strings
	apiError := errs.B().Code(errs.InvalidArgument).Msg("Validation Error")
	for _, field := range errList {
		errObj := NewValidationErrObject(field)
		fieldStr := strings.ToLower(errObj.Field)
		apiError = apiError.Meta(fieldStr, field.Error())
	}

	return apiError.Err()
}

// Validate uses the go-playground v10 validator and checks struct fields against
// a `form:"<validator>"` tag.
func (v *DefaultValidator) Validate(s interface{}) error {
	err := v.v10.Struct(s)

	if err == nil {
		return nil
	}

	// translate all validator errors
	errs, ok := err.(v10Validator.ValidationErrors)

	if !ok {
		return errors.NewErrInternal(fmt.Errorf("could not cast err to validator.ValidationErrors, type %T", err))
	}

	// convert all validator errors to error strings
	errorStrs := make([]string, len(errs))

	for i, field := range errs {
		errObj := NewValidationErrObject(field)

		errorStrs[i] = getErrorStr(errObj)
	}

	return NewErrFailedRequestValidation(errorStrs...)
}

func getErrorStr(errObj *ValidationErrObject) string {
	switch strings.ToLower(errObj.Condition) {
	case "password":
		return PasswordErr
	case "email":
		return errObj.SafeExternalError(EmailErr)
	case "spidername":
		return errObj.SafeExternalError(SpiderNameErr)
	case "uuid":
		return errObj.SafeExternalError(UUIDErr)
	case "actionid":
		return errObj.SafeExternalError(ActionIDErr)
	case "cron":
		return errObj.SafeExternalError(CronErr)
	case "duration":
		return errObj.SafeExternalError(DurationErr)
	default:
		return errObj.SafeExternalError("")
	}
}

func NewErrFailedRequestValidation(valErrors ...string) error {
	var err error

	for _, valErr := range valErrors {
		err = multierror.Append(err, errors2.New(valErr))
	}

	return errors.NewError(
		400,
		"Bad Request",
		err.Error(),
		"",
	)
}

// ValidationErrObject represents an error referencing a specific field in a struct that
// must match a specific condition. This object is modeled off of the go-playground v10
// validator `FieldError` type, but can be used generically for any request validation
// issues that occur downstream.
type ValidationErrObject struct {
	// Field is the request field that has a validation error.
	Field string

	// Namespace contains a path to the field which has a validation error
	Namespace string

	// Condition is the condition that was not satisfied, resulting in the validation
	// error
	Condition string

	// Param is an optional field that shows a parameter that was not satisfied. For example,
	// the field value was not found in the set [ "value1", "value2" ], so "value1", "value2"
	// is the parameter in this case.
	Param string

	// ActualValue is the actual value of the field that failed validation.
	ActualValue interface{}
}

// NewValidationErrObject simply returns a ValidationErrObject from a go-playground v10
// validator `FieldError`
func NewValidationErrObject(fieldErr v10Validator.FieldError) *ValidationErrObject {
	return &ValidationErrObject{
		Field:       fieldErr.Field(),
		Namespace:   fieldErr.StructNamespace(),
		Condition:   fieldErr.ActualTag(),
		Param:       fieldErr.Param(),
		ActualValue: fieldErr.Value(),
	}
}

// SafeExternalError converts the ValidationErrObject to a string that is readable and safe
// to send externally. In this case, "safe" means that when the `ActualValue` field is cast
// to a string, it is type-checked so that only certain types are passed to the user. We
// don't want an upstream command accidentally setting a complex object in the request field
// that could leak sensitive information to the user. To limit this, we only support sending
// static `ActualValue` types: `string`, `int`, `[]string`, and `[]int`. Otherwise, we say that
// the actual value is "invalid type".
//
// Note: the test cases split on "," to parse out the different errors. Don't add commas to the
// safe external error.
func (obj *ValidationErrObject) SafeExternalError(suffix string) string {
	if suffix == "" {
		suffix = fmt.Sprintf("on condition '%s'", obj.Condition)
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("validation failed on field '%s': %s", obj.Namespace, suffix))

	if obj.Param != "" {
		sb.WriteString(fmt.Sprintf(" [ %s ]: got %s", obj.Param, obj.getActualValueString()))
	}

	return sb.String()
}

func (obj *ValidationErrObject) getActualValueString() string {
	// we translate to "json-readable" form for nil values, since clients may not be Golang
	if obj.ActualValue == nil {
		return "null"
	}

	// create type switch statement to make sure that we don't accidentally leak
	// data. we only want to write strings, numbers, or slices of strings/numbers.
	// different data types can be added if necessary, as long as they are checked
	switch v := obj.ActualValue.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return fmt.Sprintf("'%s'", v)
	case []string:
		return fmt.Sprintf("[ %s ]", strings.Join(v, " "))
	case []int:
		strArr := make([]string, len(v))

		for i, intItem := range v {
			strArr[i] = strconv.Itoa(intItem)
		}

		return fmt.Sprintf("[ %s ]", strings.Join(strArr, " "))
	default:
		return "invalid type"
	}
}
