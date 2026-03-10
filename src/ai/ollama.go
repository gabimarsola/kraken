package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type OllamaClient struct {
	baseURL string
	model   string
}

func NewOllamaClient(baseURL, model string) *OllamaClient {
	if baseURL == "" {
		baseURL = "http://localhost:11434"
	}
	if model == "" {
		model = "llama2"
	}
	return &OllamaClient{
		baseURL: baseURL,
		model:   model,
	}
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func (c *OllamaClient) GeneratePRD(projectInfo string, endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Como especialista em produtos de software para DESENVOLVEDORES JÚNIOR, analise estas informações e crie um PRD completo:

PROJETO:
%s

ENDPOINTS:
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

Retorne APENAS o JSON no formato exato:
{
  "title": "Título da Funcionalidade",
  "introduction": "Introdução detalhada para desenvolvedores júnior",
  "objectives": ["obj1", "obj2"],
  "userStories": [
    {
      "id": "US001",
      "title": "Título da História",
      "description": "Descrição detalhada",
      "acceptanceCriteria": ["crit1", "crit2"]
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
  "outOfScope": ["item1"],
  "designConsiderations": ["design1"],
  "techConsiderations": ["tech1"],
  "successMetrics": ["metric1"],
  "openQuestions": ["question1"]
}
`, projectInfo, strings.Join(endpoints, "\n"))

	return c.callAPI(prompt)
}

func (c *OllamaClient) GenerateDocumentation(endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Gere documentação técnica detalhada para DESENVOLVEDORES JÚNIOR para os seguintes endpoints:

%s

Para cada endpoint, inclua:
1. Descrição clara do funcionamento (explicado de forma simples)
2. Parâmetros necessários (tipo, obrigatório, descrição detalhada)
3. Exemplos de requisição/resposta (com valores explicados)
4. Códigos de status esperados (com explicações do que cada um significa)
5. Casos de erro (com explicações de como resolver)
6. Dicas de implementação para desenvolvedores júnior

IMPORTANTE: Escreva como se estivesse ensinando um desenvolvedor júnior. Seja explícito, detalhado e evite jargões. Use exemplos práticos e explique o porquê de cada coisa.

Use formato Markdown organizado com seções claras e explicações detalhadas.
`, strings.Join(endpoints, "\n"))

	return c.callAPI(prompt)
}

func (c *OllamaClient) callAPI(prompt string) (string, error) {
	reqBody := OllamaRequest{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erro ao serializar: %v", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/api/generate", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro na requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %v", err)
	}

	var ollamaResp OllamaResponse
	err = json.Unmarshal(body, &ollamaResp)
	if err != nil {
		return "", fmt.Errorf("erro ao decodificar: %v", err)
	}

	return ollamaResp.Response, nil
}
