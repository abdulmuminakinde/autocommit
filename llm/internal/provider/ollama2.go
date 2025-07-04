package provider

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/v2/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Ollama2 struct{}

func (o Ollama2) New(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(o.Name())
	if !ok {
		return nil, fmt.Errorf("no Ollama LLM provider found")
	}

	if err := os.Setenv("OLLAMA_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set OLLAMA_API_KEY: %w", err)
	}

	return ollama.New(ollama.WithModel(llm.Model))
}

func (o Ollama2) Name() string {
	return "Ollama2"
}

func (o Ollama2) Models() []string {
	return []string{
		"gemma:1b",
		"gemma:4b",
		"gemma:12b",
		"gemma:27b",
		"qwen3:0.6b",
		"qwen3:1.7b",
		"qwen3:4b",
		"qwen3:8b",
		"qwen3:14b",
		"qwen3:30b",
		"qwen3:32b",
		"qwen3:235b",
		"deepseek-r1:1.5b",
		"deepseek-r1:7b",
		"deepseek-r1:8b",
		"deepseek-r1:14b",
		"deepseek-r1:32b",
		"deepseek-r1:70b",
		"deepseek-r1:671b",
		"llama4",
		"llama3.3",
	}
}
