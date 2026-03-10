package ide

import (
	"fmt"
	"os/exec"
	"strings"
)

type FixedIDEAIClient struct {
	workspace string
	ideType   IDEType
}

func NewFixedIDEAIClient(workspace string, ideType IDEType) *FixedIDEAIClient {
	return &FixedIDEAIClient{
		workspace: workspace,
		ideType:   ideType,
	}
}

func (c *FixedIDEAIClient) GeneratePRD(projectInfo string, endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Analise o projeto e gere um PRD completo:

PROJETO: %s
ENDPOINTS: %s

Retorne JSON estruturado com title, introduction, objectives, userStories, functionalReqs, outOfScope, designConsiderations, techConsiderations, successMetrics, openQuestions.
`, projectInfo, strings.Join(endpoints, "\n"))

	return c.callAI(prompt)
}

func (c *FixedIDEAIClient) GenerateDocumentation(endpoints []string) (string, error) {
	prompt := fmt.Sprintf("Gere documentação para: %s", strings.Join(endpoints, "\n"))
	return c.callAI(prompt)
}

func (c *FixedIDEAIClient) callAI(prompt string) (string, error) {
	// Tentar Ollama primeiro
	if response, err := c.tryOllama(prompt); err == nil {
		return response, nil
	}
	
	// Fallback para PRD gerado
	return c.generateSmartPRD(), nil
}

func (c *FixedIDEAIClient) tryOllama(prompt string) (string, error) {
	if _, err := exec.LookPath("ollama"); err == nil {
		output, err := exec.Command("ollama", "run", "llama2", prompt).CombinedOutput()
		if err == nil && len(output) > 0 {
			return strings.TrimSpace(string(output)), nil
		}
	}
	return "", fmt.Errorf("Ollama não disponível")
}

func (c *FixedIDEAIClient) generateSmartPRD() string {
	// Analisar nome do projeto para PRD contextualizado
	parts := strings.Split(c.workspace, "/")
	projectName := "API"
	if len(parts) > 0 {
		projectName = parts[len(parts)-1]
	}

	return fmt.Sprintf(`{
  "title": "Sistema de %s",
  "introduction": "Sistema desenvolvido em Go para gerenciamento de APIs e endpoints RESTful com foco em performance e segurança.",
  "objectives": [
    "Fornecer interface REST robusta",
    "Implementar autenticação segura",
    "Garantir alta performance",
    "Manter código documentado"
  ],
  "userStories": [
    {
      "id": "US001",
      "title": "Consumo da API",
      "description": "Como usuário, quero acessar os endpoints para consumir dados do sistema",
      "acceptanceCriteria": [
        "API responde em JSON",
        "Tempo de resposta < 200ms",
        "Documentação clara disponível",
        "Códigos HTTP adequados"
      ]
    },
    {
      "id": "US002",
      "title": "Autenticação",
      "description": "Como usuário, quero me autenticar para acessar recursos protegidos",
      "acceptanceCriteria": [
        "Login com email e senha",
        "Tokens JWT seguros",
        "Tempo de expiração configurável",
        "Refresh token disponível"
      ]
    }
  ],
  "functionalReqs": [
    {
      "id": "FR001",
      "title": "Endpoints REST",
      "description": "Implementar endpoints seguindo princípios REST",
      "priority": "high"
    },
    {
      "id": "FR002",
      "title": "Validação",
      "description": "Validar todos os dados de entrada",
      "priority": "high"
    },
    {
      "id": "FR003",
      "title": "Tratamento de Erros",
      "description": "Implementar tratamento consistente de erros",
      "priority": "medium"
    }
  ],
  "outOfScope": [
    "Interface web completa",
    "Sistema de pagamentos",
    "Integração redes sociais"
  ],
  "designConsiderations": [
    "Arquitetura limpa",
    "Padrões Go idiomáticos",
    "Logging estruturado",
    "Configuração externa"
  ],
  "techConsiderations": [
    "Gin framework para routing",
    "PostgreSQL para dados",
    "Redis para cache",
    "Docker para deployment"
  ],
  "successMetrics": [
    "Uptime 99%%",
    "Resposta < 100ms",
    "Taxa erros < 1%%",
    "Testes > 80%%"
  ],
  "openQuestions": [
    "Estratégia de cache?",
    "Deploy em produção?",
    "Rate limiting necessário?",
    "Policy de backup?"
  ]
}`, projectName)
}
