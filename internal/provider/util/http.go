package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// ValidateApiKey validates the API key for the AI provider API. It does so by making a request to the models endpoint - if it fails, the API key is invalid.
func ValidateAPIKey(
	ctx context.Context,
	client *http.Client,
	url, apiKey, headerName string,
) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set(headerName, "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid API key: received status %d", resp.StatusCode)
	}

	return nil
}

type httpErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   any    `json:"param"`
	Type    string `json:"type"`
}

type httpErrorResponse struct {
	Error httpErrorDetail `json:"error"`
}

func handleHTTPResponse(response *http.Response) error {
	if response.StatusCode < 400 {
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var parsedResponse httpErrorResponse
	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		return fmt.Errorf("failed to parse response body: %w", err)
	}

	return errors.New(parsedResponse.Error.Message)
}
