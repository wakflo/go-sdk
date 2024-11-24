package autoform

import (
	"testing"
)

func TestAuthBasicField_Build(t *testing.T) {
	t.Run("build_auth_field_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.Build()

		if basicAuth.schema == nil {
			t.Fatal("schema not built properly")
		}
	})
}

func TestAuthBasicField_SetDescription(t *testing.T) {
	t.Run("set_description_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.SetDescription("new description")

		if basicAuth.builder.schema.Description != "new description" {
			t.Fatal("description was not set correctly")
		}
	})
}

func TestAuthBasicField_initProps(t *testing.T) {
	t.Run("init_props_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.initProps()

		if basicAuth.passwordField == nil || basicAuth.usernameField == nil {
			t.Fatal("fields not initialized properly")
		}
	})
}

func TestAuthBasicField_SetUsernamePlaceholder(t *testing.T) {
	t.Run("set_username_placeholder_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.SetUsernamePlaceholder("username placeholder")

		if basicAuth.usernameField.schema.UIProps.Placeholder != "username placeholder" {
			t.Fatal("username placeholder not set correctly")
		}
	})
}

func TestAuthBasicField_SetPasswordPlaceholder(t *testing.T) {
	t.Run("set_password_placeholder_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.SetPasswordPlaceholder("password placeholder")

		if basicAuth.passwordField.schema.UIProps.Placeholder != "password placeholder" {
			t.Fatal("password placeholder not set correctly")
		}
	})
}

func TestAuthBasicField_SetUsernameLabel(t *testing.T) {
	t.Run("set_username_label_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.SetUsernameLabel("username label")

		if basicAuth.usernameField.schema.UIProps.Label != "username label" {
			t.Fatal("username label not set correctly")
		}
	})
}

func TestAuthBasicField_SetPasswordLabel(t *testing.T) {
	t.Run("set_password_label_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.SetPasswordLabel("password label")

		if basicAuth.passwordField.schema.UIProps.Label != "password label" {
			t.Fatal("password label not set correctly")
		}
	})
}

func TestAuthBasicField_SetUsernameHint(t *testing.T) {
	t.Run("set_username_hint_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.SetUsernameHint("username hint")

		if basicAuth.usernameField.schema.UIProps.Hint != "username hint" {
			t.Fatal("username hint not set correctly")
		}
	})
}

func TestAuthBasicField_SetPasswordHint(t *testing.T) {
	t.Run("set_password_hint_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()
		basicAuth.SetPasswordHint("password hint")

		if basicAuth.passwordField.schema.UIProps.Hint != "password hint" {
			t.Fatal("password hint not set correctly")
		}
	})
}

func TestNewAuthBasicField(t *testing.T) {
	t.Run("new_auth_basic_field_success", func(t *testing.T) {
		basicAuth := NewAuthBasicField()

		if basicAuth == nil {
			t.Fatal("new auth field not created correctly")
			return
		}
		if basicAuth.builder.schema.UIProps.Required != true {
			t.Fatal("field required not set correctly")
		}
	})
}
