package prd

import (
	"fmt"
	"kraken/src/analyzer"
	"kraken/src/ide"
	"kraken/src/structure"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type IDEPRDGenerator struct {
	projectInfo *structure.ProjectInfo
	ideManager  *ide.IDEManager
}

func NewIDEPRDGenerator(projectInfo *structure.ProjectInfo) (*IDEPRDGenerator, error) {
	ideManager, err := ide.NewIDEManager()
	if err != nil {
		return nil, fmt.Errorf("erro ao criar IDE manager: %v", err)
	}

	return &IDEPRDGenerator{
		projectInfo: projectInfo,
		ideManager:  ideManager,
	}, nil
}

func (g *IDEPRDGenerator) GeneratePRDWithIDE() error {
	ideInfo := g.ideManager.GetIDEInfo()
	fmt.Printf("🤖 Detectado IDE: %s\n", ideInfo.Name)
	fmt.Printf("📍 Workspace: %s\n", ideInfo.Workspace)
	fmt.Println("� Analisando alterações do projeto...")
	fmt.Println("� Gerando PRD com IA integrada da IDE baseado nas alterações...")

	// Primeiro analisar alterações do Git
	gitAnalyzer := analyzer.NewGitAnalyzer(".")
	analysis, err := gitAnalyzer.AnalyzeChanges()
	if err != nil {
		fmt.Printf("⚠️ Erro ao analisar Git: %v\n", err)
		fmt.Println("🔄 Usando análise completa do projeto...")
		return g.generatePRDFromFullProject()
	}

	if len(analysis.Changes) == 0 {
		fmt.Println("ℹ️ Nenhuma alteração encontrada, usando análise completa do projeto...")
		return g.generatePRDFromFullProject()
	}

	fmt.Printf("📊 Encontradas %d alterações\n", len(analysis.Changes))
	fmt.Printf("🌿 Branch: %s\n", analysis.Branch)

	if analysis.IsClean {
		fmt.Println("📝 Analisando último commit")
	} else {
		fmt.Println("⚠️ Analisando alterações não commitadas")
	}

	// Gerar PRD com IA da IDE baseado nas alterações
	prdData, err := g.generatePRDFromGitChanges(analysis)
	if err != nil {
		fmt.Printf("❌ Erro ao gerar PRD com IA da IDE: %v\n", err)
		return err
	}

	// Converter para estrutura interna
	prd := g.convertToPRD(prdData)

	// Gerar arquivo
	return g.generatePRDFileWithGit(prd, analysis)
}

func (g *IDEPRDGenerator) generatePRDFileWithGit(prd structure.PRD, analysis *analyzer.GitAnalysis) error {
	// Criar diretório docs/kraken se não existir
	docsDir := filepath.Join(".", "docs", "kraken")
	err := os.MkdirAll(docsDir, 0755)
	if err != nil {
		return fmt.Errorf("erro ao criar diretório docs/kraken: %v", err)
	}

	// Gerar nome do arquivo baseado no título e no Git
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

	// Adicionar sufixo baseado no Git
	suffix := "git"
	if analysis.IsClean && analysis.LastCommit != "" {
		suffix = fmt.Sprintf("commit-%s", analysis.LastCommit[:7])
	} else {
		suffix = fmt.Sprintf("branch-%s", analysis.Branch)
	}

	fileName := fmt.Sprintf("prd-%s-%s.md", safeTitle, suffix)
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

	ideInfo := g.ideManager.GetIDEInfo()
	fmt.Printf("✅ PRD gerado com sucesso usando IA do %s!\n", ideInfo.Name)
	fmt.Printf("📄 Arquivo criado: %s\n", filePath)
	fmt.Printf("🌿 Baseado em: %s\n", suffix)

	return nil
}

func (g *IDEPRDGenerator) extractEndpoints() []string {
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

func (g *IDEPRDGenerator) convertToPRD(data *ide.PRDData) structure.PRD {
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

func (g *IDEPRDGenerator) generatePRDFile(prd structure.PRD) error {
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

	ideInfo := g.ideManager.GetIDEInfo()
	fmt.Printf("✅ PRD gerado com sucesso usando IA do %s!\n", ideInfo.Name)
	fmt.Printf("📄 Arquivo criado: %s\n", filePath)

	return nil
}

func (g *IDEPRDGenerator) generatePRDFromFullProject() error {
	fmt.Println("🔍 Analisando projeto completo...")

	// Extrair endpoints
	endpoints := g.extractEndpoints()

	// Gerar PRD com IA da IDE
	prdData, err := g.ideManager.GeneratePRD(g.projectInfo, endpoints)
	if err != nil {
		return fmt.Errorf("erro ao gerar PRD com IA da IDE: %v", err)
	}

	// Converter para estrutura interna
	prd := g.convertToPRD(prdData)

	// Gerar arquivo
	return g.generatePRDFile(prd)
}

func (g *IDEPRDGenerator) generatePRDFromGitChanges(analysis *analyzer.GitAnalysis) (*ide.PRDData, error) {
	// Extrair endpoints das alterações
	endpoints := g.extractEndpointsFromChanges(analysis.Changes)

	// Gerar PRD com IA da IDE usando contexto do Git
	prdData, err := g.ideManager.GeneratePRD(g.projectInfo, endpoints)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar PRD com IA da IDE: %v", err)
	}

	return prdData, nil
}

func (g *IDEPRDGenerator) extractChangesInfo(analysis *analyzer.GitAnalysis) string {
	var info strings.Builder

	info.WriteString(fmt.Sprintf("Análise de %d arquivos alterados:\n\n", len(analysis.Changes)))

	for _, change := range analysis.Changes {
		info.WriteString(fmt.Sprintf("📁 %s (%s)\n", change.FilePath, change.Status))

		if change.Status == "Added" || change.Status == "Modified" {
			// Analisar conteúdo do arquivo
			if strings.HasSuffix(change.FilePath, ".go") {
				functions := g.extractGoFunctions(change.Content)
				if len(functions) > 0 {
					info.WriteString(fmt.Sprintf("   🔧 Funções: %s\n", strings.Join(functions, ", ")))
				}
			}

			// Contar linhas
			lines := strings.Split(change.Content, "\n")
			info.WriteString(fmt.Sprintf("   📏 %d linhas\n", len(lines)))
		}

		if change.Message != "" {
			info.WriteString(fmt.Sprintf("   💬 %s\n", change.Message))
		}
		info.WriteString("\n")
	}

	return info.String()
}

func (g *IDEPRDGenerator) extractEndpointsFromChanges(changes []analyzer.GitChange) []string {
	var endpoints []string

	for _, change := range changes {
		if strings.HasSuffix(change.FilePath, ".go") {
			// Extrair endpoints do código Go (simplificado)
			lines := strings.Split(change.Content, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "router.") || strings.Contains(line, "gin.") {
					// Adicionar como endpoint detectado
					endpoints = append(endpoints, fmt.Sprintf("Endpoint em %s: %s", change.FilePath, line))
				}
			}
		}
	}

	return endpoints
}

func (g *IDEPRDGenerator) extractGoFunctions(content string) []string {
	var functions []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "func ") {
			// Extrair nome da função
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				funcName := parts[1]
				// Remover parênteses e parâmetros
				if idx := strings.Index(funcName, "("); idx != -1 {
					funcName = funcName[:idx]
				}
				// Remover receiver se existir
				if strings.Contains(funcName, ")") {
					if idx := strings.Index(funcName, ")"); idx != -1 && idx+1 < len(funcName) {
						funcName = funcName[idx+1:]
					}
				}
				functions = append(functions, funcName)
			}
		}
	}

	return functions
}
