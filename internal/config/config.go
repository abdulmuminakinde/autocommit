package config

import (
	"github.com/christian-gama/autocommit/internal/groq"
	"github.com/christian-gama/autocommit/internal/openai"
)

// AIProvider specified for type safety
type AIProvider string

const (
	OpenAI AIProvider = "openai"
	Groq   AIProvider = "groq"
)

type ConfigManager struct {
	Provider AIProvider
	OpenAI   *openai.Config
	Groq     *groq.Config
}
