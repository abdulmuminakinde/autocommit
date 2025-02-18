package groq

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/christian-gama/autocommit/internal/provider"
)

func (ai *provider.AIProvider) ValidateApiKey(modelURL, apiKey string) error {
	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	httpRequest, err := httpRequestWithAuth(http.MethodGet, modelURL, apiKey)
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
