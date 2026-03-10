package ai

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type AIManager struct {
	client LLMClient
}

type AIProvider string

const (
	ProviderOpenAI    AIProvider = "openai"
	ProviderAnthropic AIProvider = "anthropic"
	ProviderOllama    AIProvider = "ollama"
	ProviderGemini    AIProvider = "gemini"
)

type PRDData struct {
	Title                string              `json:"title"`
	Introduction         string              `json:"introduction"`
	Objectives           []string            `json:"objectives"`
	UserStories          []UserStoryData     `json:"userStories"`
	FunctionalReqs       []FunctionalReqData `json:"functionalReqs"`
	OutOfScope           []string            `json:"outOfScope"`
	DesignConsiderations []string            `json:"designConsiderations"`
	TechConsiderations   []string            `json:"techConsiderations"`
	SuccessMetrics       []string            `json:"successMetrics"`
	OpenQuestions        []string            `json:"openQuestions"`
}

type UserStoryData struct {
	ID                 string   `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	AcceptanceCriteria []string `json:"acceptanceCriteria"`
}

type FunctionalReqData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

func NewAIManager(provider AIProvider, config map[string]string) (*AIManager, error) {
	var client LLMClient

	switch provider {
	case ProviderOpenAI:
		apiKey, ok := config["api_key"]
		if !ok {
			apiKey = os.Getenv("OPENAI_API_KEY")
		}
		if apiKey == "" {
			return nil, fmt.Errorf("API key da OpenAI não encontrada")
		}
		client = NewOpenAIClient(apiKey)

	case ProviderAnthropic:
		apiKey, ok := config["api_key"]
		if !ok {
			apiKey = os.Getenv("ANTHROPIC_API_KEY")
		}
		if apiKey == "" {
			return nil, fmt.Errorf("API key da Anthropic não encontrada")
		}
		client = NewAnthropicClient(apiKey)

	case ProviderOllama:
		baseURL := config["base_url"]
		model := config["model"]
		client = NewOllamaClient(baseURL, model)

	case ProviderGemini:
		apiKey, ok := config["api_key"]
		if !ok {
			apiKey = os.Getenv("GEMINI_API_KEY")
		}
		if apiKey == "" {
			return nil, fmt.Errorf("API key da Gemini não encontrada")
		}
		client = NewGeminiClient(apiKey)

	default:
		return nil, fmt.Errorf("provedor AI não suportado: %s", provider)
	}

	return &AIManager{client: client}, nil
}

func (m *AIManager) GeneratePRD(projectInfo string, endpoints []string) (*PRDData, error) {
	response, err := m.client.GeneratePRD(projectInfo, endpoints)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar PRD: %v", err)
	}

	// Limpar resposta JSON
	response = strings.TrimSpace(response)
	if strings.HasPrefix(response, "```json") {
		response = strings.TrimPrefix(response, "```json")
		response = strings.TrimSuffix(response, "```")
		response = strings.TrimSpace(response)
	} else if strings.HasPrefix(response, "```") {
		response = strings.TrimPrefix(response, "```")
		response = strings.TrimSuffix(response, "```")
		response = strings.TrimSpace(response)
	}

	var prdData PRDData
	err = json.Unmarshal([]byte(response), &prdData)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON da resposta: %v\nResposta bruta: %s", err, response)
	}

	return &prdData, nil
}

func (m *AIManager) GenerateDocumentation(endpoints []string) (string, error) {
	return m.client.GenerateDocumentation(endpoints)
}
