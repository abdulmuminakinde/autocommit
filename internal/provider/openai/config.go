package openai

import (
	"context"
	"fmt"
	"net/http"

	"github.com/christian-gama/autocommit/internal/provider"
)

// OpenAIConfig is the configuration for the OpenAI API.
type OpenAIConfig struct {
	APIKey      string  `json:"apiKey"`
	Temperature float32 `json:"temperature"`
	Model       string  `json:"model"`
}

type OpenAIProvider struct {
	config     *OpenAIConfig
	httpClient provider.HTTPClient
	models     []string
}

// NewConfig creates a new Config.
func NewConfig(apiKey, model string, temperature float32) *OpenAIConfig {
	return &OpenAIConfig{
		APIKey:      apiKey,
		Temperature: temperature,
		Model:       model,
	}
}

func NewOpenAIProvider(config *OpenAIConfig, client provider.HTTPClient) *OpenAIProvider {
	return &OpenAIProvider{
		config:     config,
		httpClient: client,
		models:     AllowedModels,
	}
}

func (op *OpenAIProvider) ValidateAPIKey(ctx context.Context) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.openai.com/v1/models",
		nil,
	)
	if err != nil {
		fmt.Errorf("error creating new request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+op.config.APIKey)

	resp, err := op.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if err := handleHTTPResponse(resp); err != nil {
		return err
	}

	return nil
}

func (op *OpenAIProvider) GetModels() []string {
	return op.models
}

func (op *OpenAIProvider) DefaultConfig() interface{} {
	return &OpenAIConfig{
		Model:       AllowedModels[0],
		Temperature: 0.7,
	}
}
