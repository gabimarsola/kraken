package ide

import (
	"fmt"
	"kraken/src/structure"
)

type IDEManager struct {
	ideInfo IDEInfo
	client  IDEAIClient
}

func NewIDEManager() (*IDEManager, error) {
	ideInfo := DetectIDE()

	if ideInfo.Type == IDEUnknown {
		return nil, fmt.Errorf("IDE não detectada ou não suportada")
	}

	// Sempre permitir uso, mesmo sem IA nativa
	var client IDEAIClient
	switch ideInfo.Type {
	case IDEWindsurf:
		client = NewFixedIDEAIClient(ideInfo.Workspace, IDEWindsurf)
	case IDECursor:
		client = NewFixedIDEAIClient(ideInfo.Workspace, IDECursor)
	case IDEVSCode:
		client = NewFixedIDEAIClient(ideInfo.Workspace, IDEVSCode)
	case IDEIntelliJ:
		client = NewFixedIDEAIClient(ideInfo.Workspace, IDEIntelliJ)
	default:
		return nil, fmt.Errorf("IDE %s não possui suporte implementado", ideInfo.Name)
	}

	return &IDEManager{
		ideInfo: ideInfo,
		client:  client,
	}, nil
}

func (m *IDEManager) GetIDEInfo() IDEInfo {
	return m.ideInfo
}

func (m *IDEManager) GeneratePRD(projectInfo *structure.ProjectInfo, endpoints []string) (*PRDData, error) {
	projectInfoStr := m.formatProjectInfo(projectInfo)

	response, err := m.client.GeneratePRD(projectInfoStr, endpoints)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar PRD com IA da IDE: %v", err)
	}

	// Converter resposta para estrutura PRDData
	prdData, err := m.parsePRDResponse(response)
	if err != nil {
		return nil, fmt.Errorf("erro ao processar resposta da IA: %v", err)
	}

	return prdData, nil
}

func (m *IDEManager) GenerateDocumentation(endpoints []string) (string, error) {
	return m.client.GenerateDocumentation(endpoints)
}

func (m *IDEManager) formatProjectInfo(info *structure.ProjectInfo) string {
	projectStr := fmt.Sprintf("Nome: %s\n", info.Name)
	projectStr += fmt.Sprintf("Tipo: %s\n", info.ProjectType)
	projectStr += fmt.Sprintf("Versão: %s\n", info.Version)
	projectStr += fmt.Sprintf("Descrição: %s\n", info.Description)

	if len(info.Endpoints) > 0 {
		projectStr += fmt.Sprintf("Total de endpoints: %d\n", len(info.Endpoints))
	}

	return projectStr
}

func (m *IDEManager) parsePRDResponse(response string) (*PRDData, error) {
	// Implementar parsing JSON similar ao do AI manager
	// Por enquanto, retorna dados simulados
	return &PRDData{
		Title:        "Sistema de Autenticação",
		Introduction: "Implementação de sistema completo de autenticação",
		Objectives: []string{
			"Proporcionar login seguro",
			"Gestão de perfis",
			"Segurança de dados",
		},
		UserStories: []UserStoryData{
			{
				ID:          "US001",
				Title:       "Cadastro de Usuário",
				Description: "Usuário pode criar nova conta",
				AcceptanceCriteria: []string{
					"Email válido",
					"Senha forte",
					"Verificação por email",
				},
			},
		},
		FunctionalReqs: []FunctionalReqData{
			{
				ID:          "FR001",
				Title:       "Validação de Email",
				Description: "Validar formato e unicidade",
				Priority:    "high",
			},
		},
		OutOfScope: []string{
			"Login social",
			"2FA",
		},
		DesignConsiderations: []string{
			"Interface responsiva",
			"Design acessível",
		},
		TechConsiderations: []string{
			"bcrypt para senhas",
			"JWT para sessões",
		},
		SuccessMetrics: []string{
			"Taxa de sucesso > 95%",
			"Tempo de cadastro < 2min",
		},
		OpenQuestions: []string{
			"Política de expiração?",
			"Rate limiting?",
		},
	}, nil
}

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
