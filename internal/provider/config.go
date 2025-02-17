package provider

// AIProviderImpl is a struct that implements the AIProvider interface
type AIProviderImpl struct {
	Name          string
	AllowedModels []string
}

type AIProvider interface {
	ValidateAPIKey(string) error
}
