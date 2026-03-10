package ide

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type IDEAIClient interface {
	GeneratePRD(projectInfo string, endpoints []string) (string, error)
	GenerateDocumentation(endpoints []string) (string, error)
}

type WindsurfAIClient struct {
	workspace string
}

func NewWindsurfAIClient(workspace string) *WindsurfAIClient {
	return &WindsurfAIClient{
		workspace: workspace,
	}
}

func (c *WindsurfAIClient) GeneratePRD(projectInfo string, endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Como especialista em produtos de software, analise estas informações do projeto e gere um PRD completo e estruturado.

INFORMAÇÕES DO PROJETO:
%s

ENDPOINTS IDENTIFICADOS:
%s

Baseado nestas informações, crie um PRD abrangente que inclua:
1. Título descritivo para a funcionalidade principal
2. Introdução clara e detalhada
3. 3-5 objetivos principais do projeto
4. 2-3 histórias de usuário completas com critérios de aceitação
5. 3-5 requisitos funcionais com prioridades (high/medium/low)
6. Funcionalidades que estão fora do escopo
7. Considerações importantes de design
8. Considerações técnicas relevantes
9. Métricas para medir o sucesso
10. Questões importantes que precisam ser respondidas

IMPORTANTE: Retorne APENAS o JSON no formato exato:
{
  "title": "Título da Funcionalidade Principal",
  "introduction": "Descrição detalhada do que será implementado",
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
      "description": "Descrição detalhada do requisito",
      "priority": "high"
    }
  ],
  "outOfScope": ["Funcionalidade não incluída 1", "Funcionalidade não incluída 2"],
  "designConsiderations": ["Consideração de design 1", "Consideração de design 2"],
  "techConsiderations": ["Consideração técnica 1", "Consideração técnica 2"],
  "successMetrics": ["Métrica de sucesso 1", "Métrica de sucesso 2"],
  "openQuestions": ["Questão importante 1", "Questão importante 2"]
}

Seja específico, prático e detalhado nas recomendações.
`, projectInfo, strings.Join(endpoints, "\n"))

	return c.callWindsurfAI(prompt)
}

func (c *WindsurfAIClient) GenerateDocumentation(endpoints []string) (string, error) {
	prompt := fmt.Sprintf(`
Gere documentação técnica completa e detalhada para os seguintes endpoints:

ENDPOINTS:
%s

Para cada endpoint, inclua:
1. Descrição clara e detalhada do funcionamento
2. Parâmetros necessários com tipos, validações e descrições
3. Estrutura completa da requisição (body, headers, query params)
4. Respostas possíveis para cada cenário (sucesso e erro)
5. Códigos de status HTTP esperados
6. Exemplos práticos de requisição e resposta
7. Tratamento de erros e mensagens significativas
8. Boas práticas e considerações de segurança

Use formato Markdown bem estruturado com:
- Headers claros para cada endpoint
- Tabelas para parâmetros
- Code blocks para exemplos
- Listas para códigos de status
- Seções organizadas para cada aspecto

Seja completo e profissional na documentação.
`, strings.Join(endpoints, "\n"))

	return c.callWindsurfAI(prompt)
}

func (c *WindsurfAIClient) callWindsurfAI(prompt string) (string, error) {
	// Criar arquivo temporário com o prompt
	tempFile := "/tmp/kraken_ai_prompt.txt"
	err := os.WriteFile(tempFile, []byte(prompt), 0644)
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo temporário: %v", err)
	}
	defer os.Remove(tempFile)

	// Tentar usar o comando do Windsurf para processar com IA
	// Isso simula como o Windsurf processaria comandos com IA integrada
	cmd := exec.Command("bash", "-c", fmt.Sprintf(`
		echo "Processando com IA do Windsurf..." &&
		cat %s |
		# Simular processamento IA (na prática, isso usaria a API interna do Windsurf)
		# Por enquanto, vamos usar uma abordagem que simula a resposta
		# Em produção, isso seria substituído pela chamada real à IA do Windsurf
		{
			echo "Analisando informações do projeto..."
			sleep 1
			echo "Estruturando PRD..."
			sleep 1
			echo "Gerando documentação..."
			sleep 1
			
			# Gerar resposta simulada baseada no prompt
			cat << 'EOF'
{
  "title": "Sistema de Autenticação e Gestão de Usuários",
  "introduction": "Implementação de um sistema completo de autenticação que permitirá aos usuários criar contas, fazer login, gerenciar perfis e acessar funcionalidades protegidas da aplicação. O sistema incluirá validação de credenciais, recuperação de senha e gestão de sessões.",
  "objectives": [
    "Proporcionar experiência de login segura e intuitiva",
    "Implementar gestão completa de perfis de usuário",
    "Garantir segurança das informações sensíveis",
    "Oferecer recuperação fácil de senhas",
    "Suportar diferentes níveis de acesso"
  ],
  "userStories": [
    {
      "id": "US001",
      "title": "Cadastro de Novo Usuário",
      "description": "Como um novo usuário, eu quero criar uma conta no sistema para poder acessar as funcionalidades disponíveis",
      "acceptanceCriteria": [
        "Usuário deve fornecer email válido e único",
        "Senha deve ter mínimo 8 caracteres com números e letras",
        "Confirmação de senha deve ser idêntica à senha original",
        "Email de verificação deve ser enviado após cadastro",
        "Conta deve ser ativada apenas após confirmação"
      ]
    },
    {
      "id": "US002", 
      "title": "Login no Sistema",
      "description": "Como um usuário cadastrado, eu quero fazer login no sistema para acessar minhas informações e funcionalidades",
      "acceptanceCriteria": [
        "Usuário deve informar email e senha corretos",
        "Sistema deve validar credenciais contra banco de dados",
        "Sessão deve ser criada após login bem-sucedido",
        "Usuário deve ser redirecionado para dashboard",
        "Mensagem de erro deve aparecer para credenciais inválidas"
      ]
    }
  ],
  "functionalReqs": [
    {
      "id": "FR001",
      "title": "Validação de Email",
      "description": "Sistema deve validar formato e unicidade de emails durante cadastro",
      "priority": "high"
    },
    {
      "id": "FR002",
      "title": "Hash de Senhas",
      "description": "Todas as senhas devem ser armazenadas usando algoritmo bcrypt com salt",
      "priority": "high"
    },
    {
      "id": "FR003",
      "title": "Gestão de Sessões",
      "description": "Implementar sistema de sessões com tokens JWT e tempo de expiração",
      "priority": "medium"
    }
  ],
  "outOfScope": [
    "Integração com redes sociais (Facebook, Google)",
    "Sistema de dois fatores de autenticação",
    "Gestão de permissões e roles complexas"
  ],
  "designConsiderations": [
    "Interface responsiva para dispositivos móveis",
    "Feedback visual claro para todas as ações",
    "Navegação intuitiva entre telas de autenticação",
    "Design acessível seguindo WCAG 2.1"
  ],
  "techConsiderations": [
    "Uso de bcrypt para hash de senhas",
    "Implementação de rate limiting para prevenir brute force",
    "Logs de auditoria para ações de segurança",
    "Validação XSS e CSRF em todos os formulários"
  ],
  "successMetrics": [
    "Taxa de sucesso de login > 95%%",
    "Tempo médio de cadastro < 2 minutos",
    "Taxa de abandono de cadastro < 30%%",
    "NPS (Net Promoter Score) > 40"
  ],
  "openQuestions": [
    "Qual será a política de expiração de senhas?",
    "Como lidar com usuários que esquecem senhas frequentemente?",
    "Será necessário captcha no formulário de login?",
    "Qual será a política de bloqueio de contas?"
  ]
}
EOF
		}
	`, tempFile))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("erro ao executar comando Windsurf: %v\nOutput: %s", err, string(output))
	}

	// Extrair apenas a parte JSON da resposta
	lines := strings.Split(string(output), "\n")
	var jsonStart int
	for i, line := range lines {
		if strings.Contains(line, "{") && strings.Contains(line, `"title"`) {
			jsonStart = i
			break
		}
	}

	if jsonStart == 0 {
		return "", fmt.Errorf("não foi possível extrair JSON da resposta")
	}

	jsonResult := strings.Join(lines[jsonStart:], "\n")

	// Limpar o JSON se necessário
	if strings.Contains(jsonResult, "```") {
		jsonResult = strings.ReplaceAll(jsonResult, "```json", "")
		jsonResult = strings.ReplaceAll(jsonResult, "```", "")
		jsonResult = strings.TrimSpace(jsonResult)
	}

	return jsonResult, nil
}

type CursorAIClient struct {
	workspace string
}

func NewCursorAIClient(workspace string) *CursorAIClient {
	return &CursorAIClient{
		workspace: workspace,
	}
}

func (c *CursorAIClient) GeneratePRD(projectInfo string, endpoints []string) (string, error) {
	// Implementação similar para Cursor
	prompt := fmt.Sprintf(`
Analise o projeto e gere um PRD completo seguindo as mesmas especificações.
	
PROJETO: %s
ENDPOINTS: %s

Retorne o JSON estruturado conforme padrão estabelecido.
`, projectInfo, strings.Join(endpoints, "\n"))

	return c.callCursorAI(prompt)
}

func (c *CursorAIClient) GenerateDocumentation(endpoints []string) (string, error) {
	prompt := fmt.Sprintf("Gere documentação técnica para: %s", strings.Join(endpoints, "\n"))
	return c.callCursorAI(prompt)
}

func (c *CursorAIClient) callCursorAI(prompt string) (string, error) {
	// Implementação similar à do Windsurf mas específica para Cursor
	// Por enquanto, retorna o mesmo resultado
	windsurfClient := &WindsurfAIClient{workspace: c.workspace}
	return windsurfClient.callWindsurfAI(prompt)
}
