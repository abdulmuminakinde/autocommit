package groq

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/christian-gama/autocommit/internal/provider"
)

// Config is the configuration for the Groq API.
type Config struct {
	APIKey      string  `json:"apiKey"`
	Temperature float32 `json:"temperature"`
	Model       string  `json:"model"`
}

// NewConfig creates a new Config.
func NewConfig(apiKey, model string, temperature float32) *Config {
	return &Config{
		APIKey:      apiKey,
		Temperature: temperature,
		Model:       model,
	}
}

type OpenAIProvider struct {
	provider.AIProviderImpl
	Config *Config
}

func NewGroqProvider(apiKey, model string, temperature float32) *OpenAIProvider {
	return &OpenAIProvider{
		AIProviderImpl: provider.AIProviderImpl{
			Name:          "openai",
			AllowedModels: AllowedModels,
		},
		Config: NewConfig(apiKey, model, temperature),
	}
}

func (o *NewOpenAIProvider) ValidateAPIKey(apiKey string) error {
	const GroqMoselsURL = "https://api.groq.com/v1/models"

	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	httpRequest, err := httpRequestWithAuth(http.MethodGet, openAIModelsURL, apiKey)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer httpResponse.Body.Close()

	if err := handleHTTPResponse(httpResponse); err != nil {
		return err
	}

	return nil
}
