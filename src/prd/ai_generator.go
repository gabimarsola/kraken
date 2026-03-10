package prd

import (
	"fmt"
	"kraken/src/ai"
	"kraken/src/structure"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type AIPRDGenerator struct {
	projectInfo *structure.ProjectInfo
	aiManager   *ai.AIManager
}

func NewAIPRDGenerator(projectInfo *structure.ProjectInfo, aiProvider ai.AIProvider, config map[string]string) (*AIPRDGenerator, error) {
	aiManager, err := ai.NewAIManager(aiProvider, config)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar AI manager: %v", err)
	}

	return &AIPRDGenerator{
		projectInfo: projectInfo,
		aiManager:   aiManager,
	}, nil
}

func (g *AIPRDGenerator) GeneratePRDFromProject() error {
	fmt.Println("🤖 Gerando PRD automaticamente com IA...")

	// Preparar informações do projeto
	projectInfoStr := g.formatProjectInfo()

	// Extrair endpoints
	endpoints := g.extractEndpoints()

	// Gerar PRD com IA
	prdData, err := g.aiManager.GeneratePRD(projectInfoStr, endpoints)
	if err != nil {
		return fmt.Errorf("erro ao gerar PRD com IA: %v", err)
	}

	// Converter para estrutura interna
	prd := g.convertToPRD(prdData)

	// Gerar arquivo
	return g.generatePRDFile(prd)
}

func (g *AIPRDGenerator) formatProjectInfo() string {
	info := fmt.Sprintf("Nome: %s\n", g.projectInfo.Name)
	info += fmt.Sprintf("Tipo: %s\n", g.projectInfo.ProjectType)
	info += fmt.Sprintf("Versão: %s\n", g.projectInfo.Version)
	info += fmt.Sprintf("Descrição: %s\n", g.projectInfo.Description)

	if len(g.projectInfo.Endpoints) > 0 {
		info += fmt.Sprintf("\nTotal de endpoints: %d\n", len(g.projectInfo.Endpoints))
	}

	return info
}

func (g *AIPRDGenerator) extractEndpoints() []string {
	var endpoints []string

	for _, endpoint := range g.projectInfo.Endpoints {
		endpointStr := fmt.Sprintf("%s %s - %s", endpoint.Method, endpoint.Path, endpoint.Description)
		if endpoint.Summary != "" {
			endpointStr += fmt.Sprintf(" (%s)", endpoint.Summary)
		}
		endpoints = append(endpoints, endpointStr)
	}

	return endpoints
}

func (g *AIPRDGenerator) convertToPRD(data *ai.PRDData) structure.PRD {
	prd := structure.PRD{
		Title:                data.Title,
		Introduction:         data.Introduction,
		Objectives:           data.Objectives,
		FunctionalReqs:       make([]structure.FunctionalRequirement, 0),
		UserStories:          make([]structure.UserStory, 0),
		OutOfScope:           data.OutOfScope,
		DesignConsiderations: data.DesignConsiderations,
		TechConsiderations:   data.TechConsiderations,
		SuccessMetrics:       data.SuccessMetrics,
		OpenQuestions:        data.OpenQuestions,
	}

	// Converter user stories
	for _, us := range data.UserStories {
		userStory := structure.UserStory{
			ID:                 us.ID,
			Title:              us.Title,
			Description:        us.Description,
			AcceptanceCriteria: us.AcceptanceCriteria,
		}
		prd.UserStories = append(prd.UserStories, userStory)
	}

	// Converter requisitos funcionais
	for _, fr := range data.FunctionalReqs {
		funcReq := structure.FunctionalRequirement{
			ID:          fr.ID,
			Title:       fr.Title,
			Description: fr.Description,
			Priority:    fr.Priority,
		}
		prd.FunctionalReqs = append(prd.FunctionalReqs, funcReq)
	}

	return prd
}

func (g *AIPRDGenerator) generatePRDFile(prd structure.PRD) error {
	// Criar diretório docs/kraken se não existir
	docsDir := filepath.Join(".", "docs", "kraken")
	err := os.MkdirAll(docsDir, 0755)
	if err != nil {
		return fmt.Errorf("erro ao criar diretório docs/kraken: %v", err)
	}

	// Gerar nome do arquivo
	safeTitle := strings.ToLower(prd.Title)
	safeTitle = strings.ReplaceAll(safeTitle, " ", "-")
	safeTitle = strings.ReplaceAll(safeTitle, "ã", "a")
	safeTitle = strings.ReplaceAll(safeTitle, "ão", "ao")
	safeTitle = strings.ReplaceAll(safeTitle, "é", "e")
	safeTitle = strings.ReplaceAll(safeTitle, "ê", "e")
	safeTitle = strings.ReplaceAll(safeTitle, "í", "i")
	safeTitle = strings.ReplaceAll(safeTitle, "ó", "o")
	safeTitle = strings.ReplaceAll(safeTitle, "ú", "u")
	safeTitle = strings.ReplaceAll(safeTitle, "ç", "c")

	fileName := fmt.Sprintf("prd-%s.md", safeTitle)
	filePath := filepath.Join(docsDir, fileName)

	// Criar template
	tmpl, err := template.New("prd").Parse(structure.PRDTemplate)
	if err != nil {
		return fmt.Errorf("erro ao criar template PRD: %v", err)
	}

	// Criar arquivo
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo PRD %s: %v", fileName, err)
	}
	defer file.Close()

	// Preparar dados para o template
	data := structure.PRDData{
		PRD:         prd,
		ProjectInfo: g.projectInfo,
	}

	// Executar template
	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("erro ao executar template PRD: %v", err)
	}

	fmt.Printf("✅ PRD gerado com sucesso pela IA!\n")
	fmt.Printf("📄 Arquivo criado: %s\n", filePath)

	return nil
}
