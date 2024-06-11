package autoform

import (
	sdkcore "github.com/wakflo/go-sdk/core"
)

type CustomAuthField struct {
	*BaseComponentField
	props map[string]sdkcore.AutoFormSchema
}

func NewCustomAuthField(authURL string, tokenURL *string, scopes []string) *CustomAuthField {
	c := &CustomAuthField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.Object)
	c.builder.WithFieldType(sdkcore.CustomAuthType)
	c.builder.WithDescription("Custom Connection")
	c.builder.WithTitle("Custom Connection")

	required := false
	c.Required = required
	c.builder.schema.Presentation.Required = required
	c.builder.schema.IsRequired = required

	c.builder.schema.AuthURL = &authURL
	c.builder.schema.TokenURL = tokenURL
	c.builder.schema.Scope = scopes

	return c
}

func (b *CustomAuthField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *CustomAuthField) SetDescription(desc string) *CustomAuthField {
	b.builder.WithDescription(desc)
	return b
}

func (b *CustomAuthField) SetRequired(required bool) *CustomAuthField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *CustomAuthField) SetFields(fields map[string]*sdkcore.AutoFormSchema) *CustomAuthField {
	b.builder.WithProperties(fields)
	return b
}