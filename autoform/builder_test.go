// Filename: builder_test.go
package autoform

import (
	"testing"

	sdkcore "github.com/wakflo/go-sdk/core"
)

func TestNewSchemaBuilder(t *testing.T) {
	builder := NewSchemaBuilder()

	if builder.schema.Title != "" {
		t.Errorf("schema title is not empty")
	}
}

func TestBuilder(t *testing.T) {
	builder := NewSchemaBuilder()
	builder.WithTitle("test title")

	if builder.schema.Title != "test title" {
		t.Errorf("unexpected schema title: got %v, want %v",
			builder.schema.Title, "test title")
	}

	builder.WithType("string")

	if builder.schema.Type != ("string") {
		t.Errorf("unexpected schema type: got %v, want %v",
			builder.schema.Type, sdkcore.AutoFormType("string"))
	}

	builder.WithFieldType("number")

	if builder.schema.UIControl != ("number") {
		t.Errorf("unexpected schema form field type: got %v, want %v",
			builder.schema.UIControl, sdkcore.AutoFormFieldType("number"))
	}

	builder.WithFieldRequired(true)

	if builder.schema.UIProps.Required != true {
		t.Errorf("unexpected schema form field 'required' value: got %v, want %v",
			builder.schema.UIProps.Required, true)
	}

	// and so on for all the other methods, make sure to cover all corner cases,
	// such as empty strings, zero values, etc.
}
