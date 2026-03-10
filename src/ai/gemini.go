package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GeminiClient struct {
	apiKey  string
	baseURL string
	model   string
}

func NewGeminiClient(apiKey string) *GeminiClient {
	return &GeminiClient{
		apiKey:  apiKey,
		baseURL: "https://generativelanguage.googleapis.com/v1beta",
		model:   "gemini-pro",
	}
}

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content Content `json:"content"`
}

func (c *GeminiClient) GeneratePRD(projectInfo string, endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Analise as informações do projeto e gere um PRD completo e estruturado para DESENVOLVEDORES JÚNIOR:

INFORMAÇÕES DO PROJETO:
%s

ENDPOINTS IDENTIFICADOS:
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

Retorne EXATAMENTE neste formato JSON:
{
  "title": "Título da Funcionalidade Principal",
  "introduction": "Descrição detalhada para desenvolvedores júnior",
  "objectives": ["Objetivo 1", "Objetivo 2", "Objetivo 3"],
  "userStories": [
    {
      "id": "US001",
      "title": "Título da História de Usuário",
      "description": "Descrição completa da história",
      "acceptanceCriteria": ["Critério 1", "Critério 2", "Critério 3"]
    }
  ],
  "functionalReqs": [
    {
      "id": "FR001",
      "title": "Título do Requisito Funcional",
      "description": "Descrição detalhada explicando o que, por que e como",
      "priority": "high"
    }
  ],
  "outOfScope": ["Funcionalidade não incluída 1", "Funcionalidade não incluída 2"],
  "designConsiderations": ["Consideração de design 1", "Consideração de design 2"],
  "techConsiderations": ["Consideração técnica 1", "Consideração técnica 2"],
  "successMetrics": ["Métrica de sucesso 1", "Métrica de sucesso 2"],
  "openQuestions": ["Questão importante 1", "Questão importante 2"]
}

Seja específico, prático e educativo nas recomendações.
`, projectInfo, strings.Join(endpoints, "\n"))

	return c.callAPI(prompt)
}

func (c *GeminiClient) GenerateDocumentation(endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Gere documentação técnica detalhada para DESENVOLVEDORES JÚNIOR para os seguintes endpoints:

ENDPOINTS:
%s

Para cada endpoint, documente:
1. Descrição clara do funcionamento (explicado de forma simples)
2. Parâmetros necessários (tipo, obrigatório, descrição detalhada)
3. Exemplos de requisição/resposta (com valores explicados)
4. Códigos de status esperados (com explicações do que cada um significa)
5. Casos de erro (com explicações de como resolver)
6. Dicas de implementação para desenvolvedores júnior

IMPORTANTE: Escreva como se estivesse ensinando um desenvolvedor júnior. Seja explícito, detalhado e evite jargões. Use exemplos práticos e explique o porquê de cada coisa.

Use formato Markdown bem estruturado com seções claras e explicações detalhadas.
`, strings.Join(endpoints, "\n"))

	return c.callAPI(prompt)
}

func (c *GeminiClient) callAPI(prompt string) (string, error) {
	reqBody := GeminiRequest{
		Contents: []Content{
			{Type: "text", Text: prompt},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erro ao serializar: %v", err)
	}

	url := fmt.Sprintf("%s/models/%s:generateContent?key=%s", c.baseURL, c.model, c.apiKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
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

	var geminiResp GeminiResponse
	err = json.Unmarshal(body, &geminiResp)
	if err != nil {
		return "", fmt.Errorf("erro ao decodificar: %v", err)
	}

	if len(geminiResp.Candidates) == 0 {
		return "", fmt.Errorf("nenhuma resposta da API")
	}

	return geminiResp.Candidates[0].Content.Text, nil
}
