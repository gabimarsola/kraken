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

type GitPRDGenerator struct {
	projectInfo *structure.ProjectInfo
	gitAnalyzer *analyzer.GitAnalyzer
	ideManager  *ide.IDEManager
}

func NewGitPRDGenerator(projectInfo *structure.ProjectInfo) (*GitPRDGenerator, error) {
	gitAnalyzer := analyzer.NewGitAnalyzer(".")

	// Tentar criar IDE manager para usar IA se disponível
	ideManager, err := ide.NewIDEManager()
	if err != nil {
		// Se não tiver IDE, continua sem IA
		ideManager = nil
	}

	return &GitPRDGenerator{
		projectInfo: projectInfo,
		gitAnalyzer: gitAnalyzer,
		ideManager:  ideManager,
	}, nil
}

func (g *GitPRDGenerator) GeneratePRDFromChanges() error {
	fmt.Println("🔍 Analisando alterações do projeto...")

	// Analisar mudanças no Git
	analysis, err := g.gitAnalyzer.AnalyzeChanges()
	if err != nil {
		return fmt.Errorf("erro ao analisar Git: %v", err)
	}

	if len(analysis.Changes) == 0 {
		fmt.Println("ℹ️ Nenhuma alteração encontrada para analisar")
		return nil
	}

	fmt.Printf("📊 Encontradas %d alterações\n", len(analysis.Changes))
	fmt.Printf("🌿 Branch: %s\n", analysis.Branch)

	if analysis.IsClean {
		fmt.Println("📝 Analisando último commit")
	} else {
		fmt.Println("⚠️ Analisando alterações não commitadas")
	}

	// Gerar PRD baseado nas alterações
	prd, err := g.generatePRDFromGitAnalysis(analysis)
	if err != nil {
		return fmt.Errorf("erro ao gerar PRD: %v", err)
	}

	// Gerar arquivo
	return g.generatePRDFile(prd, analysis)
}

func (g *GitPRDGenerator) generatePRDFromGitAnalysis(analysis *analyzer.GitAnalysis) (structure.PRD, error) {
	// Extrair informações das alterações
	changesInfo := g.extractChangesInfo(analysis)

	// Gerar título baseado nas alterações
	title := g.generateTitle(analysis)

	// Se tiver IDE com IA, usar para gerar PRD
	if g.ideManager != nil {
		return g.generatePRDWithIDE(analysis, changesInfo)
	}

	// Senão, gerar PRD baseado nas alterações
	return g.generatePRDFromChanges(analysis, changesInfo, title), nil
}

func (g *GitPRDGenerator) generatePRDWithIDE(analysis *analyzer.GitAnalysis, changesInfo string) (structure.PRD, error) {
	fmt.Println("🤖 Usando IA da IDE para analisar alterações...")

	// Extrair endpoints das alterações
	endpoints := g.extractEndpointsFromChanges(analysis.Changes)

	// Gerar PRD com IA
	prdData, err := g.ideManager.GeneratePRD(g.projectInfo, endpoints)
	if err != nil {
		fmt.Printf("⚠️ Erro ao usar IA da IDE: %v\n", err)
		// Fallback para geração manual
		title := g.generateTitle(analysis)
		return g.generatePRDFromChanges(analysis, changesInfo, title), nil
	}

	// Converter para estrutura interna
	return g.convertIDEPRDToStructure(prdData), nil
}

func (g *GitPRDGenerator) generatePRDFromChanges(analysis *analyzer.GitAnalysis, changesInfo, title string) structure.PRD {
	fmt.Println("📝 Gerando PRD baseado nas alterações...")

	prd := structure.PRD{
		Title:                title,
		Introduction:         g.generateIntroduction(analysis, changesInfo),
		Objectives:           g.generateObjectives(analysis),
		UserStories:          g.generateUserStories(analysis),
		FunctionalReqs:       g.generateFunctionalReqs(analysis),
		OutOfScope:           g.generateOutOfScope(analysis),
		DesignConsiderations: g.generateDesignConsiderations(analysis),
		TechConsiderations:   g.generateTechConsiderations(analysis),
		SuccessMetrics:       g.generateSuccessMetrics(analysis),
		OpenQuestions:        g.generateOpenQuestions(analysis),
	}

	return prd
}

func (g *GitPRDGenerator) extractChangesInfo(analysis *analyzer.GitAnalysis) string {
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

func (g *GitPRDGenerator) extractGoFunctions(content string) []string {
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

func (g *GitPRDGenerator) generateTitle(analysis *analyzer.GitAnalysis) string {
	if len(analysis.Changes) == 0 {
		return "Atualização do Projeto"
	}

	// Analisar padrões nos arquivos para gerar título
	hasAPI := false
	hasAuth := false
	hasDB := false
	hasUI := false

	for _, change := range analysis.Changes {
		lowerPath := strings.ToLower(change.FilePath)

		if strings.Contains(lowerPath, "api") || strings.Contains(lowerPath, "endpoint") {
			hasAPI = true
		}
		if strings.Contains(lowerPath, "auth") || strings.Contains(lowerPath, "login") || strings.Contains(lowerPath, "user") {
			hasAuth = true
		}
		if strings.Contains(lowerPath, "db") || strings.Contains(lowerPath, "database") || strings.Contains(lowerPath, "model") {
			hasDB = true
		}
		if strings.Contains(lowerPath, "ui") || strings.Contains(lowerPath, "frontend") || strings.Contains(lowerPath, "view") {
			hasUI = true
		}
	}

	// Gerar título baseado nas funcionalidades
	var titleParts []string

	if hasAuth {
		titleParts = append(titleParts, "Autenticação")
	}
	if hasAPI {
		titleParts = append(titleParts, "API")
	}
	if hasDB {
		titleParts = append(titleParts, "Dados")
	}
	if hasUI {
		titleParts = append(titleParts, "Interface")
	}

	if len(titleParts) == 0 {
		return "Atualização do Sistema"
	}

	return fmt.Sprintf("Sistema de %s", strings.Join(titleParts, " e "))
}

func (g *GitPRDGenerator) generateIntroduction(analysis *analyzer.GitAnalysis, changesInfo string) string {
	if analysis.IsClean {
		return fmt.Sprintf("Atualização do sistema baseada no último commit (%s) que implementa melhorias e novas funcionalidades identificadas na análise de código.", analysis.LastCommit[:7])
	}

	return fmt.Sprintf("Desenvolvimento em andamento com %d arquivos modificados. Esta atualização foca em implementar as funcionalidades identificadas nas alterações recentes do projeto.", len(analysis.Changes))
}

func (g *GitPRDGenerator) generateObjectives(analysis *analyzer.GitAnalysis) []string {
	objectives := []string{
		"Implementar as alterações identificadas no código",
		"Manter compatibilidade com o sistema existente",
		"Garantir qualidade e performance do código",
	}

	// Adicionar objetivos específicos baseados nas alterações
	for _, change := range analysis.Changes {
		if strings.Contains(strings.ToLower(change.FilePath), "test") {
			objectives = append(objectives, "Aumentar cobertura de testes")
		}
		if strings.Contains(strings.ToLower(change.FilePath), "auth") {
			objectives = append(objectives, "Melhorar segurança da autenticação")
		}
		if strings.Contains(strings.ToLower(change.FilePath), "api") {
			objectives = append(objectives, "Expandir funcionalidades da API")
		}
	}

	return objectives
}

func (g *GitPRDGenerator) generateUserStories(analysis *analyzer.GitAnalysis) []structure.UserStory {
	var stories []structure.UserStory
	storyID := 1

	for _, change := range analysis.Changes {
		if change.Status == "Deleted" {
			continue
		}

		story := structure.UserStory{
			ID:                 fmt.Sprintf("US%03d", storyID),
			Title:              g.generateStoryTitle(change),
			Description:        g.generateStoryDescription(change),
			AcceptanceCriteria: g.generateStoryCriteria(change),
		}

		stories = append(stories, story)
		storyID++
	}

	return stories
}

func (g *GitPRDGenerator) generateStoryTitle(change analyzer.GitChange) string {
	fileName := filepath.Base(change.FilePath)

	if strings.HasSuffix(fileName, ".go") {
		return fmt.Sprintf("Implementação de %s", strings.TrimSuffix(fileName, ".go"))
	}

	return fmt.Sprintf("Atualização de %s", fileName)
}

func (g *GitPRDGenerator) generateStoryDescription(change analyzer.GitChange) string {
	return fmt.Sprintf("Como usuário do sistema, eu quero que as funcionalidades do arquivo %s sejam implementadas/atualizadas para melhorar a experiência e funcionalidade do sistema.", change.FilePath)
}

func (g *GitPRDGenerator) generateStoryCriteria(change analyzer.GitChange) []string {
	criteria := []string{
		"O código deve compilar sem erros",
		"As funcionalidades devem funcionar conforme esperado",
		"O código deve seguir os padrões do projeto",
	}

	if strings.HasSuffix(change.FilePath, "_test.go") {
		criteria = append(criteria, "Todos os testes devem passar")
		criteria = append(criteria, "Cobertura de testes adequada")
	}

	return criteria
}

func (g *GitPRDGenerator) generateFunctionalReqs(analysis *analyzer.GitAnalysis) []structure.FunctionalRequirement {
	var reqs []structure.FunctionalRequirement
	reqID := 1

	for _, change := range analysis.Changes {
		if change.Status == "Deleted" {
			continue
		}

		req := structure.FunctionalRequirement{
			ID:          fmt.Sprintf("FR%03d", reqID),
			Title:       fmt.Sprintf("Implementar %s", filepath.Base(change.FilePath)),
			Description: fmt.Sprintf("Desenvolver as funcionalidades especificadas no arquivo %s conforme as alterações identificadas.", change.FilePath),
			Priority:    g.getPriority(change),
		}

		reqs = append(reqs, req)
		reqID++
	}

	return reqs
}

func (g *GitPRDGenerator) getPriority(change analyzer.GitChange) string {
	// Arquivos de teste e config têm prioridade menor
	if strings.Contains(change.FilePath, "test") || strings.Contains(change.FilePath, "config") {
		return "low"
	}

	// Arquivos principais têm alta prioridade
	if strings.Contains(change.FilePath, "main") || strings.Contains(change.FilePath, "api") {
		return "high"
	}

	return "medium"
}

func (g *GitPRDGenerator) generateOutOfScope(analysis *analyzer.GitAnalysis) []string {
	return []string{
		"Modificação em arquivos não relacionados às alterações",
		"Mudanças na estrutura de banco de dados não planejadas",
		"Alterações em APIs de terceiros",
	}
}

func (g *GitPRDGenerator) generateDesignConsiderations(analysis *analyzer.GitAnalysis) []string {
	considerations := []string{
		"Manter consistência com o código existente",
		"Seguir os padrões de arquitetura do projeto",
	}

	// Adicionar considerações específicas
	for _, change := range analysis.Changes {
		if strings.HasSuffix(change.FilePath, ".go") {
			considerations = append(considerations, "Seguir as convenções da linguagem Go")
		}
		if strings.Contains(strings.ToLower(change.FilePath), "api") {
			considerations = append(considerations, "Manter compatibilidade backward da API")
		}
	}

	return considerations
}

func (g *GitPRDGenerator) generateTechConsiderations(analysis *analyzer.GitAnalysis) []string {
	considerations := []string{
		"Garantir performance das novas funcionalidades",
		"Implementar logging adequado",
		"Tratamento de erros consistente",
	}

	if !analysis.IsClean {
		considerations = append(considerations, "Preparar código para commit")
	}

	return considerations
}

func (g *GitPRDGenerator) generateSuccessMetrics(analysis *analyzer.GitAnalysis) []string {
	return []string{
		"Todos os testes passando",
		"Código compilando sem warnings",
		"Performance mantida ou melhorada",
		"Zero regressões identificadas",
	}
}

func (g *GitPRDGenerator) generateOpenQuestions(analysis *analyzer.GitAnalysis) []string {
	questions := []string{
		"As alterações afetam outras partes do sistema?",
		"É necessário atualizar a documentação?",
		"Há dependências que precisam ser atualizadas?",
	}

	if !analysis.IsClean {
		questions = append(questions, "Todas as alterações estão prontas para commit?")
	}

	return questions
}

func (g *GitPRDGenerator) extractEndpointsFromChanges(changes []analyzer.GitChange) []string {
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

func (g *GitPRDGenerator) convertIDEPRDToStructure(prdData *ide.PRDData) structure.PRD {
	prd := structure.PRD{
		Title:                prdData.Title,
		Introduction:         prdData.Introduction,
		Objectives:           prdData.Objectives,
		FunctionalReqs:       make([]structure.FunctionalRequirement, 0),
		UserStories:          make([]structure.UserStory, 0),
		OutOfScope:           prdData.OutOfScope,
		DesignConsiderations: prdData.DesignConsiderations,
		TechConsiderations:   prdData.TechConsiderations,
		SuccessMetrics:       prdData.SuccessMetrics,
		OpenQuestions:        prdData.OpenQuestions,
	}

	// Converter user stories
	for _, us := range prdData.UserStories {
		userStory := structure.UserStory{
			ID:                 us.ID,
			Title:              us.Title,
			Description:        us.Description,
			AcceptanceCriteria: us.AcceptanceCriteria,
		}
		prd.UserStories = append(prd.UserStories, userStory)
	}

	// Converter requisitos funcionais
	for _, fr := range prdData.FunctionalReqs {
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

func (g *GitPRDGenerator) generatePRDFile(prd structure.PRD, analysis *analyzer.GitAnalysis) error {
	// Criar diretório docs/kraken se não existir
	docsDir := filepath.Join(".", "docs", "kraken")
	err := os.MkdirAll(docsDir, 0755)
	if err != nil {
		return fmt.Errorf("erro ao criar diretório docs/kraken: %v", err)
	}

	// Gerar nome do arquivo baseado no título e no commit/branch
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

	// Adicionar sufixo baseado no tipo de análise
	suffix := "changes"
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

	fmt.Printf("✅ PRD gerado com sucesso!\n")
	fmt.Printf("📄 Arquivo criado: %s\n", filePath)

	// Mostrar resumo das alterações analisadas
	fmt.Printf("\n📊 Resumo da Análise:\n")
	fmt.Printf("🌿 Branch: %s\n", analysis.Branch)
	fmt.Printf("📁 Arquivos: %d\n", len(analysis.Changes))
	if analysis.LastCommit != "" {
		fmt.Printf("🔗 Commit: %s\n", analysis.LastCommit[:7])
	}

	return nil
}
