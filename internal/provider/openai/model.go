package openai

// System is the system configuration for the OpenAI API - the initial message and name.
type System struct {
	Message string
	Name    string
}

// NewSystem creates a new System.
func NewSystem(message, name string) *System {
	return &System{
		Message: message,
		Name:    name,
	}
}

const (
	GPT3Dot5Turbo16k = "gpt-3.5-turbo-16k"
	GPT3Dot5Turbo    = "gpt-3.5-turbo"
	GPT4             = "gpt-4"
	GPT432K          = "gpt-4-32k"
)

var AllowedModels = []string{
	GPT3Dot5Turbo,
	GPT3Dot5Turbo16k,
	GPT4,
	GPT432K,
}
