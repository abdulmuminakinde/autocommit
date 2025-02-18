package provider

import (
	"context"
	"net/http"
)

type AIProvider interface {
	GetName() string
	ValidateAPIKey(ctx context.Context) error
	GetModels() []string
	DefaultConfig() interface{}
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}
