package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LLMClient interface {
	GeneratePRD(projectInfo string, endpoints []string) (string, error)
	GenerateDocumentation(endpoints []string) (string, error)
}

type OpenAIClient struct {
	apiKey  string
	baseURL string
	model   string
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		apiKey:  apiKey,
		baseURL: "https://api.openai.com/v1",
		model:   "gpt-4",
	}
}

type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func (c *OpenAIClient) GeneratePRD(projectInfo string, endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Baseado nas seguintes informações do projeto, gere um PRD completo e estruturado para DESENVOLVEDORES JÚNIOR:

INFORMAÇÕES DO PROJETO:
%s

ENDPOINTS ENCONTRADOS:
%s

REGRAS PARA O PRD (Focado em Desenvolvedores Júnior):
1. Crie um título descritivo e claro da funcionalidade
2. Escreva uma introdução detalhada explicando o propósito de forma simples
3. Defina 3-5 objetivos principais com linguagem clara e direta
4. Crie 2-3 histórias de usuário com critérios de aceitação explícitos e não ambíguos
5. Liste 3-5 requisitos funcionais com descrições detalhadas, evitando jargões técnicos
6. Para cada requisito, explique: O QUE é, POR QUE é necessário, e COMO deve funcionar
7. Defina o que está fora do escopo de forma clara
8. Adicione considerações de design e técnicas com explicações para iniciantes
9. Sugira métricas de sucesso claras e mensuráveis
10. Liste questões importantes com explicações contextuais

IMPORTANTE: Escreva como se estivesse explicando para um desenvolvedor júnior. Seja explícito, detalhado e evite assumir conhecimento prévio. Use exemplos quando possível.

FORMATO DE SAÍDA:
Retorne APENAS o JSON estruturado neste formato:
{
  "title": "Título da Funcionalidade",
  "introduction": "Introdução detalhada para desenvolvedores júnior",
  "objectives": ["objetivo1", "objetivo2", ...],
  "userStories": [
    {
      "id": "US001",
      "title": "Título da História",
      "description": "Descrição detalhada",
      "acceptanceCriteria": ["critério1", "critério2", ...]
    }
  ],
  "functionalReqs": [
    {
      "id": "FR001", 
      "title": "Título do Requisito",
      "description": "Descrição detalhada explicando o que, por que e como",
      "priority": "high"
    }
  ],
  "outOfScope": ["item1", "item2", ...],
  "designConsiderations": ["consideração1", ...],
  "techConsiderations": ["consideração1", ...],
  "successMetrics": ["métrica1", ...],
  "openQuestions": ["questão1", ...]
}
`, projectInfo, strings.Join(endpoints, "\n"))

	return c.callAPI(prompt)
}

func (c *OpenAIClient) GenerateDocumentation(endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Gere documentação técnica detalhada para DESENVOLVEDORES JÚNIOR para os seguintes endpoints:

ENDPOINTS:
%s

Para cada endpoint, inclua:
1. Descrição clara do funcionamento (explicado de forma simples)
2. Parâmetros necessários (tipo, obrigatório, descrição detalhada)
3. Exemplos de requisição/resposta (com valores explicados)
4. Códigos de status esperados (com explicações do que cada um significa)
5. Casos de erro (com explicações de como resolver)
6. Dicas de implementação para desenvolvedores júnior

IMPORTANTE: Escreva como se estivesse ensinando um desenvolvedor júnior. Seja explícito, detalhado e evite jargões. Use exemplos práticos e explique o porquê de cada coisa.

FORMATO: Markdown estruturado com seções claras e explicações detalhadas
`, strings.Join(endpoints, "\n"))

	return c.callAPI(prompt)
}

func (c *OpenAIClient) callAPI(prompt string) (string, error) {
	reqBody := OpenAIRequest{
		Model: c.model,
		Messages: []Message{
			{Role: "system", Content: "Você é um especialista em documentação técnica e Product Requirements Documents para desenvolvedores júnior. Sua missão é explicar conceitos complexos de forma clara, detalhada e acessível, evitando jargões e assumindo pouco conhecimento prévio. Seja sempre explícito e use exemplos práticos."},
			{Role: "user", Content: prompt},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erro ao serializar requisição: %v", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro na requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %v", err)
	}

	var openaiResp OpenAIResponse
	err = json.Unmarshal(body, &openaiResp)
	if err != nil {
		return "", fmt.Errorf("erro ao decodificar resposta: %v", err)
	}

	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta da API")
	}

	return openaiResp.Choices[0].Message.Content, nil
}
