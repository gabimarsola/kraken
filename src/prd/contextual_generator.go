package prd

import (
	"fmt"
	"kraken/src/ai"
	"kraken/src/structure"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type ContextualPRDGenerator struct {
	projectInfo *structure.ProjectInfo
	aiProvider  ai.AIProvider
	aiConfig    map[string]string
}

func NewContextualPRDGenerator(projectInfo *structure.ProjectInfo) *ContextualPRDGenerator {
	return &ContextualPRDGenerator{
		projectInfo: projectInfo,
	}
}

func (g *ContextualPRDGenerator) SetAIProvider(provider ai.AIProvider, config map[string]string) {
	g.aiProvider = provider
	g.aiConfig = config
}

func (g *ContextualPRDGenerator) GenerateCompleteDocumentation() error {
	fmt.Println("🚀 Gerando Documentação Completa Contextualizada")
	fmt.Println("==============================================")

	// 1. Analisar contexto completo do projeto
	fmt.Println("📊 Analisando contexto do projeto...")
	context, err := g.analyzeProjectContext()
	if err != nil {
		return fmt.Errorf("erro ao analisar contexto: %v", err)
	}

	// 2. Analisar histórico Git
	fmt.Println("🌿 Analisando histórico do projeto...")
	gitContext, err := g.analyzeGitContext()
	if err != nil {
		fmt.Printf("⚠️  Aviso: Erro ao analisar Git: %v\n", err)
		gitContext = &GitContext{} // Usar contexto vazio
	}

	// 3. Gerar PRD contextualizado
	fmt.Println("� Gerando PRD contextualizado...")
	prdContent, err := g.generateContextualPRD(context)
	if err != nil {
		return fmt.Errorf("erro ao gerar PRD: %v", err)
	}

	// 4. Gerar documento de versão COMPLETO (incluindo API)
	fmt.Println("� Gerando documento de versão completo...")
	versionContent, err := g.generateCompleteVersionDocument(context, gitContext)
	if err != nil {
		return fmt.Errorf("erro ao gerar documento de versão: %v", err)
	}

	// 5. Salvar arquivos (só PRD e versão)
	err = g.saveDocuments(prdContent, versionContent, context)
	if err != nil {
		return fmt.Errorf("erro ao salvar documentos: %v", err)
	}

	fmt.Println("\n✅ Documentação completa gerada com sucesso!")
	return nil
}

type ProjectContext struct {
	ProjectName   string
	ProjectType   string
	Version       string
	Description   string
	Dependencies  []string
	Endpoints     []structure.Endpoint
	RecentChanges []string
	Architecture  string
	Patterns      []string
	BusinessLogic string
	Technologies  []string
	Structure     string
	LastModified  time.Time
}

type GitContext struct {
	CurrentBranch    string
	LastCommit       string
	LastCommitMsg    string
	LastMerge        string
	UncommittedFiles []string
	RecentCommits    []GitCommit
	Changes          []string
}

type GitCommit struct {
	Hash    string
	Message string
	Author  string
	Date    time.Time
	Files   []string
}

func (g *ContextualPRDGenerator) analyzeProjectContext() (*ProjectContext, error) {
	context := &ProjectContext{
		ProjectName:  g.projectInfo.Name,
		ProjectType:  g.projectInfo.ProjectType,
		Version:      g.getProjectVersion(),
		Endpoints:    g.projectInfo.Endpoints,
		Dependencies: g.extractDependencies(),
		Technologies: g.extractTechnologies(),
		Structure:    g.analyzeProjectStructure(),
		LastModified: time.Now(),
	}

	// Analisar arquivos para entender o contexto
	context.Description = g.generateProjectDescription(context)
	context.Architecture = g.analyzeArchitecture(context)
	context.BusinessLogic = g.analyzeBusinessLogic(context)
	context.Patterns = g.identifyPatterns(context)
	context.RecentChanges = g.identifyRecentChanges(context)

	return context, nil
}

func (g *ContextualPRDGenerator) getProjectVersion() string {
	// Tentar extrair versão do package.json, go.mod, etc.
	if g.projectInfo.ProjectType == "nodejs" {
		if content, err := os.ReadFile("package.json"); err == nil {
			// Parse simples para extrair versão
			contentStr := string(content)
			if strings.Contains(contentStr, "\"version\"") {
				lines := strings.Split(contentStr, "\n")
				for _, line := range lines {
					if strings.Contains(line, "\"version\"") {
						parts := strings.Split(line, "\"")
						for i, part := range parts {
							if part == "version" && i+2 < len(parts) {
								return parts[i+2]
							}
						}
					}
				}
			}
		}
	}
	return "1.0.0"
}

func (g *ContextualPRDGenerator) extractDependencies() []string {
	var deps []string

	if g.projectInfo.ProjectType == "nodejs" {
		if content, err := os.ReadFile("package.json"); err == nil {
			contentStr := string(content)
			lines := strings.Split(contentStr, "\n")
			inDependencies := false

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "\"dependencies\"") {
					inDependencies = true
					continue
				}
				if inDependencies && strings.Contains(line, "}") {
					break
				}
				if inDependencies && strings.Contains(line, "\"") && !strings.Contains(line, "//") {
					parts := strings.Split(line, "\"")
					if len(parts) >= 3 {
						deps = append(deps, parts[1])
					}
				}
			}
		}
	}

	return deps
}

func (g *ContextualPRDGenerator) extractTechnologies() []string {
	var techs []string

	// Baseado no tipo do projeto
	switch g.projectInfo.ProjectType {
	case "nodejs":
		techs = append(techs, "Node.js", "Express.js", "JavaScript")
	case "go":
		techs = append(techs, "Go", "Gin")
	case "java":
		techs = append(techs, "Java", "Spring Boot")
	}

	// Adicionar dependências como tecnologias
	for _, dep := range g.extractDependencies() {
		if !strings.Contains(dep, "@") && len(dep) > 2 {
			techs = append(techs, dep)
		}
	}

	return techs
}

func (g *ContextualPRDGenerator) analyzeProjectStructure() string {
	var structure strings.Builder

	structure.WriteString("## Estrutura do Projeto\n\n")

	// Analisar diretórios principais
	if dirs, err := os.ReadDir("."); err == nil {
		structure.WriteString("### Diretórios Principais\n\n")
		for _, dir := range dirs {
			if dir.IsDir() && !strings.HasPrefix(dir.Name(), ".") && dir.Name() != "node_modules" {
				structure.WriteString(fmt.Sprintf("- **%s/**\n", dir.Name))
			}
		}
	}

	if dirs, err := os.ReadDir("src"); err == nil {
		structure.WriteString("\n### Estrutura src/\n\n")
		for _, dir := range dirs {
			if dir.IsDir() {
				structure.WriteString(fmt.Sprintf("- **src/%s/**\n", dir.Name()))
			}
		}
	}

	return structure.String()
}

func (g *ContextualPRDGenerator) generateProjectDescription(context *ProjectContext) string {
	var desc strings.Builder

	desc.WriteString(fmt.Sprintf("# %s\n\n", context.ProjectName))
	desc.WriteString(fmt.Sprintf("**Tipo:** %s\n", context.ProjectType))
	desc.WriteString(fmt.Sprintf("**Versão:** %s\n", context.Version))
	desc.WriteString(fmt.Sprintf("**Tecnologias:** %s\n\n", strings.Join(context.Technologies, ", ")))

	desc.WriteString("## Descrição Geral\n\n")
	desc.WriteString(fmt.Sprintf("Projeto %s desenvolvido em %s utilizando as seguintes tecnologias principais: %s.\n\n",
		context.ProjectName, context.ProjectType, strings.Join(context.Technologies, ", ")))

	if len(context.Endpoints) > 0 {
		desc.WriteString(fmt.Sprintf("A aplicação expõe %d endpoints para interação via API REST.\n\n", len(context.Endpoints)))
	}

	return desc.String()
}

func (g *ContextualPRDGenerator) analyzeArchitecture(context *ProjectContext) string {
	var arch strings.Builder

	arch.WriteString("## Arquitetura do Sistema\n\n")
	arch.WriteString("### Padrão Arquitetural\n\n")

	switch context.ProjectType {
	case "nodejs":
		arch.WriteString("O projeto segue uma arquitetura em camadas típica de aplicações Node.js:\n\n")
		arch.WriteString("- **Camada de Roteamento** (routes/): Define os endpoints da API\n")
		arch.WriteString("- **Camada de Controladores** (controllers/): Lógica de negócio\n")
		arch.WriteString("- **Camada de Serviços** (services/): Integrações e regras de negócio\n")
		arch.WriteString("- **Camada de Modelos** (models/): Estrutura de dados\n")
		arch.WriteString("- **Middlewares**: Interceptadores de requisição\n\n")
	case "go":
		arch.WriteString("Arquitetura baseada em pacotes Go:\n\n")
		arch.WriteString("- **handlers/**: Controladores HTTP\n")
		arch.WriteString("- **services/**: Lógica de negócio\n")
		arch.WriteString("- **models/**: Estruturas de dados\n")
		arch.WriteString("- **middleware/**: Interceptadores\n\n")
	}

	arch.WriteString("### Fluxo de Requisição\n\n")
	arch.WriteString("1. Cliente envia requisição para a API\n")
	arch.WriteString("2. Middleware processa autenticação/validação\n")
	arch.WriteString("3. Router direciona para o controller apropriado\n")
	arch.WriteString("4. Controller executa lógica de negócio\n")
	arch.WriteString("5. Resposta é retornada ao cliente\n\n")

	return arch.String()
}

func (g *ContextualPRDGenerator) analyzeBusinessLogic(context *ProjectContext) string {
	var logic strings.Builder

	logic.WriteString("## Lógica de Negócio\n\n")

	if len(context.Endpoints) > 0 {
		logic.WriteString("### Principais Funcionalidades\n\n")

		// Agrupar endpoints por funcionalidade
		grouped := make(map[string][]structure.Endpoint)
		for _, endpoint := range context.Endpoints {
			parts := strings.Split(strings.Trim(endpoint.Path, "/"), "/")
			if len(parts) > 0 {
				group := parts[0]
				if group == "" {
					group = "root"
				}
				grouped[group] = append(grouped[group], endpoint)
			}
		}

		for group, endpoints := range grouped {
			logic.WriteString(fmt.Sprintf("#### %s\n\n", strings.Title(group)))
			for _, endpoint := range endpoints {
				logic.WriteString(fmt.Sprintf("- **%s %s**: %s\n",
					endpoint.Method, endpoint.Path, endpoint.Description))
			}
			logic.WriteString("\n")
		}
	}

	return logic.String()
}

func (g *ContextualPRDGenerator) identifyPatterns(context *ProjectContext) []string {
	var patterns []string

	// Identificar padrões baseados na estrutura
	if strings.Contains(context.Structure, "routes") {
		patterns = append(patterns, "MVC (Model-View-Controller)")
	}
	if strings.Contains(context.Structure, "middleware") {
		patterns = append(patterns, "Middleware Pattern")
	}
	if strings.Contains(context.Structure, "services") {
		patterns = append(patterns, "Service Layer Pattern")
	}

	// Padrões baseados em dependências
	for _, dep := range context.Dependencies {
		if strings.Contains(dep, "express") {
			patterns = append(patterns, "RESTful API")
		}
		if strings.Contains(dep, "mongoose") || strings.Contains(dep, "sequelize") {
			patterns = append(patterns, "ORM/ODM Pattern")
		}
		if strings.Contains(dep, "jwt") {
			patterns = append(patterns, "JWT Authentication")
		}
	}

	return patterns
}

func (g *ContextualPRDGenerator) identifyRecentChanges(context *ProjectContext) []string {
	var changes []string

	// Em uma implementação real, isso analisaria o Git
	// Por enquanto, vamos adicionar algumas mudanças simuladas
	changes = append(changes, "Implementação de novos endpoints de API")
	changes = append(changes, "Atualização de dependências")
	changes = append(changes, "Melhorias de performance")

	return changes
}

func (g *ContextualPRDGenerator) analyzeGitContext() (*GitContext, error) {
	context := &GitContext{}

	// Verificar se estamos em um repositório Git
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		return context, fmt.Errorf("não é um repositório Git")
	}

	// Obter branch atual
	if branch, err := g.runGitCommand("git", "rev-parse", "--abbrev-ref", "HEAD"); err == nil {
		context.CurrentBranch = strings.TrimSpace(branch)
	}

	// Obter último commit
	if lastCommit, err := g.runGitCommand("git", "log", "-1", "--format=%H"); err == nil {
		context.LastCommit = strings.TrimSpace(lastCommit)
	}

	// Obter mensagem do último commit
	if lastMsg, err := g.runGitCommand("git", "log", "-1", "--format=%s"); err == nil {
		context.LastCommitMsg = strings.TrimSpace(lastMsg)
	}

	// Obter último merge
	if lastMerge, err := g.runGitCommand("git", "log", "--merges", "-1", "--format=%s"); err == nil {
		context.LastMerge = strings.TrimSpace(lastMerge)
	}

	// Obter arquivos não commitados
	if uncommitted, err := g.runGitCommand("git", "status", "--porcelain"); err == nil {
		lines := strings.Split(strings.TrimSpace(uncommitted), "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					context.UncommittedFiles = append(context.UncommittedFiles, parts[1])
				}
			}
		}
	}

	// Obter commits recentes (últimos 5)
	if recentCommits, err := g.runGitCommand("git", "log", "-5", "--format=%H|%s|%an|%ai"); err == nil {
		lines := strings.Split(strings.TrimSpace(recentCommits), "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				parts := strings.Split(line, "|")
				if len(parts) >= 4 {
					date, _ := time.Parse("2006-01-02 15:04:05 -0700", parts[3])
					commit := GitCommit{
						Hash:    parts[0],
						Message: parts[1],
						Author:  parts[2],
						Date:    date,
					}
					context.RecentCommits = append(context.RecentCommits, commit)
				}
			}
		}
	}

	// Obter mudanças desde o último commit
	if changes, err := g.runGitCommand("git", "diff", "--name-only", "HEAD~1", "HEAD"); err == nil {
		lines := strings.Split(strings.TrimSpace(changes), "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				context.Changes = append(context.Changes, strings.TrimSpace(line))
			}
		}
	}

	return context, nil
}

func (g *ContextualPRDGenerator) runGitCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (g *ContextualPRDGenerator) generateContextualPRD(context *ProjectContext) (string, error) {
	if g.aiProvider == "" {
		// Usar novo template PRD profissional
		ticketID := GenerateTicketID(context)
		template := NewPRDTemplate(ticketID, context.ProjectName)
		return template.generateBasicPRD(context, &GitContext{}), nil
	}

	// Preparar contexto para a IA com novo template
	ticketID := GenerateTicketID(context)
	template := NewPRDTemplate(ticketID, context.ProjectName)

	return template.GeneratePRD(context, &GitContext{}, g.aiProvider, g.aiConfig)
}

func (g *ContextualPRDGenerator) generateVersionDocument(context *ProjectContext) (string, error) {
	if g.aiProvider == "" {
		return g.generateBasicVersionDoc(context), nil
	}

	prompt := g.buildVersionPrompt(context)

	aiManager, err := ai.NewAIManager(g.aiProvider, g.aiConfig)
	if err != nil {
		return g.generateBasicVersionDoc(context), nil
	}

	response, err := aiManager.GenerateDocumentation([]string{prompt})
	if err != nil {
		return g.generateBasicVersionDoc(context), nil
	}

	return response, nil
}

func (g *ContextualPRDGenerator) generateCompleteVersionDocument(context *ProjectContext, gitContext *GitContext) (string, error) {
	if g.aiProvider == "" {
		return g.generateBasicCompleteVersionDoc(context, gitContext), nil
	}

	prompt := g.buildCompleteVersionPrompt(context, gitContext)

	aiManager, err := ai.NewAIManager(g.aiProvider, g.aiConfig)
	if err != nil {
		return g.generateBasicCompleteVersionDoc(context, gitContext), nil
	}

	response, err := aiManager.GenerateDocumentation([]string{prompt})
	if err != nil {
		return g.generateBasicCompleteVersionDoc(context, gitContext), nil
	}

	return response, nil
}

func (g *ContextualPRDGenerator) generateEndpointsDocumentation(context *ProjectContext) (string, error) {
	var doc strings.Builder

	doc.WriteString("# 📚 Documentação Técnica da API\n\n")
	doc.WriteString(fmt.Sprintf("**Projeto:** %s\n", context.ProjectName))
	doc.WriteString(fmt.Sprintf("**Versão:** %s\n", context.Version))
	doc.WriteString(fmt.Sprintf("**Data:** %s\n\n", context.LastModified.Format("02/01/2006")))

	if len(context.Endpoints) > 0 {
		doc.WriteString("## 📋 Endpoints Disponíveis\n\n")

		// Agrupar por funcionalidade
		grouped := make(map[string][]structure.Endpoint)
		for _, endpoint := range context.Endpoints {
			parts := strings.Split(strings.Trim(endpoint.Path, "/"), "/")
			if len(parts) > 0 {
				group := parts[0]
				if group == "" {
					group = "root"
				}
				grouped[group] = append(grouped[group], endpoint)
			}
		}

		for group, endpoints := range grouped {
			doc.WriteString(fmt.Sprintf("### %s\n\n", strings.Title(group)))

			for _, endpoint := range endpoints {
				doc.WriteString(fmt.Sprintf("#### %s %s\n\n", endpoint.Method, endpoint.Path))
				doc.WriteString(fmt.Sprintf("**Descrição:** %s\n\n", endpoint.Description))

				if endpoint.Summary != "" {
					doc.WriteString(fmt.Sprintf("**Funcionamento:** %s\n\n", endpoint.Summary))
				}

				if len(endpoint.Parameters) > 0 {
					doc.WriteString("**Parâmetros:**\n")
					for _, param := range endpoint.Parameters {
						doc.WriteString(fmt.Sprintf("- `%s` (%s): %s\n", param.Name, param.Type, param.Description))
					}
					doc.WriteString("\n")
				}

				doc.WriteString("**Exemplo de Requisição:**\n")
				doc.WriteString("```bash\n")
				doc.WriteString(fmt.Sprintf("curl -X %s 'http://localhost:3000%s'", endpoint.Method, endpoint.Path))
				doc.WriteString("\n```\n\n")

				doc.WriteString("---\n\n")
			}
		}
	}

	doc.WriteString("## 🔧 Tecnologias Utilizadas\n\n")
	for _, tech := range context.Technologies {
		doc.WriteString(fmt.Sprintf("- %s\n", tech))
	}
	doc.WriteString("\n")

	doc.WriteString("## 📦 Dependências Principais\n\n")
	for _, dep := range context.Dependencies {
		doc.WriteString(fmt.Sprintf("- %s\n", dep))
	}

	return doc.String(), nil
}

func (g *ContextualPRDGenerator) buildPRDPrompt(context *ProjectContext) string {
	var prompt strings.Builder

	prompt.WriteString("Como especialista em produtos de software para DESENVOLVEDORES JÚNIOR, analise este contexto completo e crie um PRD detalhado:\n\n")

	prompt.WriteString("## CONTEXTO DO PROJETO\n\n")
	prompt.WriteString(fmt.Sprintf("**Nome:** %s\n", context.ProjectName))
	prompt.WriteString(fmt.Sprintf("**Tipo:** %s\n", context.ProjectType))
	prompt.WriteString(fmt.Sprintf("**Versão:** %s\n", context.Version))
	prompt.WriteString(fmt.Sprintf("**Tecnologias:** %s\n", strings.Join(context.Technologies, ", ")))
	prompt.WriteString(fmt.Sprintf("**Dependências:** %s\n", strings.Join(context.Dependencies, ", ")))

	prompt.WriteString("\n## ESTRUTURA DO PROJETO\n\n")
	prompt.WriteString(context.Structure)

	prompt.WriteString("\n## ARQUITETURA\n\n")
	prompt.WriteString(context.Architecture)

	prompt.WriteString("\n## LÓGICA DE NEGÓCIO\n\n")
	prompt.WriteString(context.BusinessLogic)

	prompt.WriteString("\n## PADRÕES IDENTIFICADOS\n\n")
	for _, pattern := range context.Patterns {
		prompt.WriteString(fmt.Sprintf("- %s\n", pattern))
	}

	prompt.WriteString("\n## ENDPOINTS\n\n")
	for _, endpoint := range context.Endpoints {
		prompt.WriteString(fmt.Sprintf("- %s %s: %s\n", endpoint.Method, endpoint.Path, endpoint.Description))
	}

	prompt.WriteString("\n## MUDANÇAS RECENTES\n\n")
	for _, change := range context.RecentChanges {
		prompt.WriteString(fmt.Sprintf("- %s\n", change))
	}

	prompt.WriteString("\n## REGRAS PARA O PRD (Focado em Desenvolvedores Júnior)\n")
	prompt.WriteString("1. Crie um título descritivo e claro da funcionalidade principal\n")
	prompt.WriteString("2. Escreva uma introdução detalhada explicando o propósito de forma simples\n")
	prompt.WriteString("3. Defina 3-5 objetivos principais com linguagem clara e direta\n")
	prompt.WriteString("4. Crie 2-3 histórias de usuário com critérios de aceitação explícitos\n")
	prompt.WriteString("5. Liste 3-5 requisitos funcionais com descrições detalhadas\n")
	prompt.WriteString("6. Para cada requisito, explique: O QUE é, POR QUE é necessário, e COMO deve funcionar\n")
	prompt.WriteString("7. Defina o que está fora do escopo de forma clara\n")
	prompt.WriteString("8. Adicione considerações de design e técnicas com explicações para iniciantes\n")
	prompt.WriteString("9. Sugira métricas de sucesso claras e mensuráveis\n")
	prompt.WriteString("10. Liste questões importantes com explicações contextuais\n")

	prompt.WriteString("\n## IMPORTANTE\n")
	prompt.WriteString("Escreva como se estivesse explicando para um desenvolvedor júnior. Seja explícito, detalhado e evite assumir conhecimento prévio. Use exemplos quando possível.\n")
	prompt.WriteString("Baseie o PRD na realidade do projeto analisado, não invente funcionalidades que não existem.\n")

	return prompt.String()
}

func (g *ContextualPRDGenerator) buildVersionPrompt(context *ProjectContext) string {
	var prompt strings.Builder

	prompt.WriteString("Como especialista técnico, analise este contexto e gere uma documentação detalhada da versão atual:\n\n")

	prompt.WriteString("## INFORMAÇÕES DO PROJETO\n\n")
	prompt.WriteString(fmt.Sprintf("**Projeto:** %s\n", context.ProjectName))
	prompt.WriteString(fmt.Sprintf("**Versão:** %s\n", context.Version))
	prompt.WriteString(fmt.Sprintf("**Tipo:** %s\n", context.ProjectType))
	prompt.WriteString(fmt.Sprintf("**Data:** %s\n", context.LastModified.Format("02/01/2006")))

	prompt.WriteString("\n## TECNOLOGIAS E DEPENDÊNCIAS\n\n")
	prompt.WriteString("**Stack Principal:**\n")
	for _, tech := range context.Technologies {
		prompt.WriteString(fmt.Sprintf("- %s\n", tech))
	}

	prompt.WriteString("\n**Dependências:**\n")
	for _, dep := range context.Dependencies {
		prompt.WriteString(fmt.Sprintf("- %s\n", dep))
	}

	prompt.WriteString("\n## ESTRUTURA E ARQUITETURA\n\n")
	prompt.WriteString(context.Architecture)

	prompt.WriteString("\n## FUNCIONALIDADES IMPLEMENTADAS\n\n")
	prompt.WriteString(context.BusinessLogic)

	prompt.WriteString("\n## MUDANÇAS DESTA VERSÃO\n\n")
	for _, change := range context.RecentChanges {
		prompt.WriteString(fmt.Sprintf("- %s\n", change))
	}

	prompt.WriteString("\n## REQUISITOS PARA A DOCUMENTAÇÃO\n\n")
	prompt.WriteString("1. **Descrição do Funcionamento**: Explique detalhadamente como o sistema funciona\n")
	prompt.WriteString("2. **Explicação para Desenvolvedores Júnior**: Use linguagem clara, evite jargões\n")
	prompt.WriteString("3. **Detalhes Técnicos**: Inclua informações importantes sobre implementação\n")
	prompt.WriteString("4. **Exemplos Práticos**: Forneça exemplos de uso e configuração\n")
	prompt.WriteString("5. **Considerações de Deploy**: Informações importantes para deploy\n")
	prompt.WriteString("6. **Boas Práticas**: Dicas para manutenção e desenvolvimento\n")

	prompt.WriteString("\n## FORMATO DESEJADO\n\n")
	prompt.WriteString("Use Markdown bem estruturado com:\n")
	prompt.WriteString("- Títulos hierárquicos claros\n")
	prompt.WriteString("- Listas organizadas\n")
	prompt.WriteString("- Blocos de código para exemplos\n")
	prompt.WriteString("- Tabelas quando apropriado\n")
	prompt.WriteString("- Ênfase em pontos importantes\n")

	return prompt.String()
}

func (g *ContextualPRDGenerator) extractEndpointsInfo(context *ProjectContext) []string {
	var endpoints []string

	for _, endpoint := range context.Endpoints {
		info := fmt.Sprintf("%s %s - %s", endpoint.Method, endpoint.Path, endpoint.Description)
		if endpoint.Summary != "" {
			info += fmt.Sprintf(" (%s)", endpoint.Summary)
		}
		endpoints = append(endpoints, info)
	}

	return endpoints
}

func (g *ContextualPRDGenerator) generateBasicPRD(context *ProjectContext) string {
	var prd strings.Builder

	prd.WriteString(fmt.Sprintf("# PRD: %s\n\n", context.ProjectName))
	prd.WriteString("## Introdução\n\n")
	prd.WriteString(fmt.Sprintf("Este documento descreve os requisitos e especificações para o projeto %s, ", context.ProjectName))
	prd.WriteString(fmt.Sprintf("uma aplicação %s desenvolvida com as tecnologias %s.\n\n",
		context.ProjectType, strings.Join(context.Technologies, ", ")))

	prd.WriteString("## Objetivos\n\n")
	prd.WriteString("1. Prover uma API robusta e escalável\n")
	prd.WriteString("2. Garantir a manutenibilidade do código\n")
	prd.WriteString("3. Oferecer uma boa experiência para desenvolvedores\n\n")

	prd.WriteString("## Histórias de Usuário\n\n")
	prd.WriteString("### US001 - Como desenvolvedor, quero consumir a API facilmente\n")
	prd.WriteString("**Critérios de Aceitação:**\n")
	prd.WriteString("- A API deve responder consistentemente\n")
	prd.WriteString("- A documentação deve estar clara\n")
	prd.WriteString("- Os endpoints devem seguir padrões REST\n\n")

	prd.WriteString("## Requisitos Funcionais\n\n")
	for _, endpoint := range context.Endpoints {
		prd.WriteString(fmt.Sprintf("### %s %s\n", endpoint.Method, endpoint.Path))
		prd.WriteString(fmt.Sprintf("**Descrição:** %s\n\n", endpoint.Description))
	}

	return prd.String()
}

func (g *ContextualPRDGenerator) generateBasicVersionDoc(context *ProjectContext) string {
	var doc strings.Builder

	version := strings.ReplaceAll(context.Version, ".", "")

	doc.WriteString(fmt.Sprintf("# Versão %s - Documentação Técnica\n\n", version))
	doc.WriteString(fmt.Sprintf("**Projeto:** %s\n", context.ProjectName))
	doc.WriteString(fmt.Sprintf("**Data:** %s\n\n", context.LastModified.Format("02/01/2006")))

	doc.WriteString("## Descrição do Funcionamento\n\n")
	doc.WriteString(fmt.Sprintf("O %s é uma aplicação %s que oferece uma API REST para ", context.ProjectName, context.ProjectType))
	doc.WriteString("interação via HTTP. A aplicação segue uma arquitetura em camadas, separando ")
	doc.WriteString("responsabilidades entre roteamento, controle de negócio e persistência de dados.\n\n")

	doc.WriteString("## Explicação Detalhada\n\n")
	doc.WriteString("### Arquitetura\n")
	doc.WriteString("O sistema foi projetado seguindo os melhores práticas de desenvolvimento ")
	doc.WriteString("de software, com separação clara de responsabilidades e baixo acoplamento ")
	doc.WriteString("entre os componentes.\n\n")

	doc.WriteString("### Funcionamento\n")
	doc.WriteString("Quando uma requisição é recebida pela API, ela passa pelos seguintes estágios:\n")
	doc.WriteString("1. **Roteamento**: A requisição é direcionada ao endpoint apropriado\n")
	doc.WriteString("2. **Validação**: Middleware validam autenticação e formato dos dados\n")
	doc.WriteString("3. **Processamento**: O controller executa a lógica de negócio\n")
	doc.WriteString("4. **Resposta**: O resultado é retornado no formato JSON\n\n")

	doc.WriteString("## Tecnologias Utilizadas\n\n")
	for _, tech := range context.Technologies {
		doc.WriteString(fmt.Sprintf("- **%s**: Tecnologia principal da stack\n", tech))
	}

	doc.WriteString("\n## Dependências\n\n")
	for _, dep := range context.Dependencies {
		doc.WriteString(fmt.Sprintf("- %s\n", dep))
	}

	return doc.String()
}

func (g *ContextualPRDGenerator) saveAllDocuments(prdContent, versionContent, endpointsDoc string, context *ProjectContext) error {
	// Criar diretório docs/kraken
	docsDir := filepath.Join(".", "docs", "kraken")
	err := os.MkdirAll(docsDir, 0755)
	if err != nil {
		return fmt.Errorf("erro ao criar diretório docs/kraken: %v", err)
	}

	// Gerar nome base para os arquivos
	safeName := strings.ToLower(context.ProjectName)
	safeName = strings.ReplaceAll(safeName, " ", "-")
	safeName = strings.ReplaceAll(safeName, "_", "-")

	version := strings.ReplaceAll(context.Version, ".", "")

	// Salvar PRD
	prdFile := filepath.Join(docsDir, fmt.Sprintf("prd-%s.md", safeName))
	err = os.WriteFile(prdFile, []byte(prdContent), 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar PRD: %v", err)
	}

	// Salvar documento de versão
	versionFile := filepath.Join(docsDir, fmt.Sprintf("versao%s-doc.md", version))
	err = os.WriteFile(versionFile, []byte(versionContent), 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar documento de versão: %v", err)
	}

	// Salvar documentação de endpoints
	apiFile := filepath.Join(docsDir, fmt.Sprintf("api-%s-doc.md", safeName))
	err = os.WriteFile(apiFile, []byte(endpointsDoc), 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar documentação de API: %v", err)
	}

	// Exibir resumo
	fmt.Printf("\n📄 Arquivos gerados:\n")
	fmt.Printf("   - PRD: %s\n", prdFile)
	fmt.Printf("   - Versão: %s\n", versionFile)
	fmt.Printf("   - API: %s\n", apiFile)

	return nil
}

func (g *ContextualPRDGenerator) saveDocuments(prdContent, versionContent string, context *ProjectContext) error {
	// Criar diretório docs/kraken
	docsDir := filepath.Join(".", "docs", "kraken")
	err := os.MkdirAll(docsDir, 0755)
	if err != nil {
		return fmt.Errorf("erro ao criar diretório docs/kraken: %v", err)
	}

	// Gerar nome base para os arquivos
	safeName := strings.ToLower(context.ProjectName)
	safeName = strings.ReplaceAll(safeName, " ", "-")
	safeName = strings.ReplaceAll(safeName, "_", "-")

	version := strings.ReplaceAll(context.Version, ".", "")

	// Salvar PRD
	prdFile := filepath.Join(docsDir, fmt.Sprintf("prd-%s.md", safeName))
	err = os.WriteFile(prdFile, []byte(prdContent), 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar PRD: %v", err)
	}

	// Salvar documento de versão
	versionFile := filepath.Join(docsDir, fmt.Sprintf("versao%s-doc.md", version))
	err = os.WriteFile(versionFile, []byte(versionContent), 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar documento de versão: %v", err)
	}

	// Exibir resumo
	fmt.Printf("\n📄 Arquivos gerados:\n")
	fmt.Printf("   - PRD: %s\n", prdFile)
	fmt.Printf("   - Versão: %s\n", versionFile)

	return nil
}

func (g *ContextualPRDGenerator) convertPRDDataToString(prdData *ai.PRDData) string {
	var prd strings.Builder

	prd.WriteString(fmt.Sprintf("# %s\n\n", prdData.Title))
	prd.WriteString("## Introdução\n\n")
	prd.WriteString(fmt.Sprintf("%s\n\n", prdData.Introduction))

	prd.WriteString("## Objetivos\n\n")
	for _, objective := range prdData.Objectives {
		prd.WriteString(fmt.Sprintf("- %s\n", objective))
	}
	prd.WriteString("\n")

	prd.WriteString("## Histórias de Usuário\n\n")
	for _, story := range prdData.UserStories {
		prd.WriteString(fmt.Sprintf("### %s - %s\n\n", story.ID, story.Title))
		prd.WriteString(fmt.Sprintf("**Descrição:** %s\n\n", story.Description))
		prd.WriteString("**Critérios de Aceitação:**\n")
		for _, criteria := range story.AcceptanceCriteria {
			prd.WriteString(fmt.Sprintf("- %s\n", criteria))
		}
		prd.WriteString("\n")
	}

	prd.WriteString("## Requisitos Funcionais\n\n")
	for _, req := range prdData.FunctionalReqs {
		prd.WriteString(fmt.Sprintf("### %s - %s\n\n", req.ID, req.Title))
		prd.WriteString(fmt.Sprintf("**Descrição:** %s\n", req.Description))
		prd.WriteString(fmt.Sprintf("**Prioridade:** %s\n\n", req.Priority))
	}

	if len(prdData.OutOfScope) > 0 {
		prd.WriteString("## Fora do Escopo\n\n")
		for _, item := range prdData.OutOfScope {
			prd.WriteString(fmt.Sprintf("- %s\n", item))
		}
		prd.WriteString("\n")
	}

	if len(prdData.DesignConsiderations) > 0 {
		prd.WriteString("## Considerações de Design\n\n")
		for _, consideration := range prdData.DesignConsiderations {
			prd.WriteString(fmt.Sprintf("- %s\n", consideration))
		}
		prd.WriteString("\n")
	}

	if len(prdData.TechConsiderations) > 0 {
		prd.WriteString("## Considerações Técnicas\n\n")
		for _, consideration := range prdData.TechConsiderations {
			prd.WriteString(fmt.Sprintf("- %s\n", consideration))
		}
		prd.WriteString("\n")
	}

	if len(prdData.SuccessMetrics) > 0 {
		prd.WriteString("## Métricas de Sucesso\n\n")
		for _, metric := range prdData.SuccessMetrics {
			prd.WriteString(fmt.Sprintf("- %s\n", metric))
		}
		prd.WriteString("\n")
	}

	if len(prdData.OpenQuestions) > 0 {
		prd.WriteString("## Questões em Aberto\n\n")
		for _, question := range prdData.OpenQuestions {
			prd.WriteString(fmt.Sprintf("- %s\n", question))
		}
		prd.WriteString("\n")
	}

	return prd.String()
}
