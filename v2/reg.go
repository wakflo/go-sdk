package sdk

type RegistrationMap struct {
	Versions map[string]Integration
}

type IntegrationsRegistrar = map[string]RegistrationMap
