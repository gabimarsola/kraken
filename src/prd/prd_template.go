package prd

import (
	"fmt"
	"kraken/src/ai"
	"strings"
	"time"
)

// PRDTemplate representa o formato de PRD baseado em padrões de engenharia
type PRDTemplate struct {
	TicketID    string
	ProjectName string
	CreatedDate time.Time
	Author      string
	Status      string
	Priority    string
}

// NewPRDTemplate cria uma nova instância do template
func NewPRDTemplate(ticketID, projectName string) *PRDTemplate {
	return &PRDTemplate{
		TicketID:    ticketID,
		ProjectName: projectName,
		CreatedDate: time.Now(),
		Status:      "Draft",
		Priority:    "Medium",
	}
}

// GeneratePRD gera a PRD no formato profissional
func (t *PRDTemplate) GeneratePRD(context *ProjectContext, gitContext *GitContext, aiProvider ai.AIProvider, aiConfig map[string]string) (string, error) {
	if aiProvider != "" {
		return t.generateAIPRD(context, gitContext, aiProvider, aiConfig)
	}
	return t.generateBasicPRD(context, gitContext), nil
}

func (t *PRDTemplate) generateAIPRD(context *ProjectContext, gitContext *GitContext, provider ai.AIProvider, config map[string]string) (string, error) {
	prompt := t.buildAIPrompt(context, gitContext)

	aiManager, err := ai.NewAIManager(provider, config)
	if err != nil {
		return t.generateBasicPRD(context, gitContext), nil
	}

	response, err := aiManager.GenerateDocumentation([]string{prompt})
	if err != nil {
		return t.generateBasicPRD(context, gitContext), nil
	}

	return response, nil
}

func (t *PRDTemplate) buildAIPrompt(context *ProjectContext, gitContext *GitContext) string {
	var prompt strings.Builder

	prompt.WriteString(fmt.Sprintf("Como Product Manager sênior, crie uma PRD profissional em PORTUGUÊS BRASILEIRO no formato de engenharia de software para o ticket %s.\n\n", t.TicketID))

	prompt.WriteString("## CONTEXTO DO PROJETO\n\n")
	prompt.WriteString(fmt.Sprintf("**Projeto:** %s\n", context.ProjectName))
	prompt.WriteString(fmt.Sprintf("**Tipo:** %s\n", context.ProjectType))
	prompt.WriteString(fmt.Sprintf("**Versão:** %s\n", context.Version))
	prompt.WriteString(fmt.Sprintf("**Tecnologias:** %s\n", strings.Join(context.Technologies, ", ")))

	prompt.WriteString("\n## CONTEXTO GIT\n\n")
	prompt.WriteString(fmt.Sprintf("**Branch:** %s\n", gitContext.CurrentBranch))

	if gitContext.LastCommit != "" {
		shortCommit := gitContext.LastCommit
		if len(shortCommit) > 7 {
			shortCommit = shortCommit[:7]
		}
		prompt.WriteString(fmt.Sprintf("**Último Commit:** %s - %s\n", shortCommit, gitContext.LastCommitMsg))
	}

	if len(gitContext.Changes) > 0 {
		prompt.WriteString("**Arquivos Alterados:**\n")
		for _, change := range gitContext.Changes {
			prompt.WriteString(fmt.Sprintf("- %s\n", change))
		}
	}

	prompt.WriteString("\n## ENDPOINTS EXISTENTES\n\n")
	for _, endpoint := range context.Endpoints {
		prompt.WriteString(fmt.Sprintf("- %s %s: %s\n", endpoint.Method, endpoint.Path, endpoint.Description))
	}

	prompt.WriteString("\n## REQUISITOS DA PRD\n\n")
	prompt.WriteString("Crie uma PRD profissional em PORTUGUÊS BRASILEIRO seguindo este formato:\n\n")

	prompt.WriteString("1. **Header** com informações do ticket\n")
	prompt.WriteString("2. **Visão Geral** - resumo executivo\n")
	prompt.WriteString("3. **Declaração do Problema** - problema a ser resolvido\n")
	prompt.WriteString("4. **Objetivos e Metas** - metas SMART\n")
	prompt.WriteString("5. **Requisitos Funcionais** - requisitos funcionais\n")
	prompt.WriteString("6. **Requisitos Não Funcionais** - requisitos não funcionais\n")
	prompt.WriteString("7. **Considerações Técnicas** - considerações técnicas\n")
	prompt.WriteString("8. **Dependências** - dependências externas\n")
	prompt.WriteString("9. **Riscos e Mitigações** - riscos e mitigações\n")

	prompt.WriteString("\n## IMPORTANTE\n\n")
	prompt.WriteString("- Use linguagem profissional e técnica em PORTUGUÊS BRASILEIRO\n")
	prompt.WriteString("- Seja específico e mensurável\n")
	prompt.WriteString("- Inclua exemplos práticos\n")
	prompt.WriteString("- Considere o contexto atual do projeto\n")
	prompt.WriteString("- Foque em desenvolvedores júnior como audiência\n")
	prompt.WriteString("- NÃO inclua seções a partir de \"Success Metrics\" (métricas de sucesso)\n")
	prompt.WriteString("- NÃO inclua: Histórias de Usuário, Fora do Escopo, Cronograma, Stakeholders, Contexto Git\n")

	return prompt.String()
}

func (t *PRDTemplate) generateBasicPRD(context *ProjectContext, gitContext *GitContext) string {
	var prd strings.Builder

	// Header
	prd.WriteString(fmt.Sprintf("# %s - Documento de Requisitos do Produto\n\n", t.TicketID))
	prd.WriteString(fmt.Sprintf("**Projeto:** %s\n", t.ProjectName))
	prd.WriteString(fmt.Sprintf("**Status:** %s\n", t.Status))
	prd.WriteString(fmt.Sprintf("**Prioridade:** %s\n", t.Priority))
	prd.WriteString(fmt.Sprintf("**Criado:** %s\n", t.CreatedDate.Format("02/01/2006")))
	prd.WriteString(fmt.Sprintf("**Autor:** Equipe de Engenharia\n\n"))

	// Visão Geral
	prd.WriteString("## Visão Geral\n\n")
	prd.WriteString(fmt.Sprintf("Este documento descreve os requisitos para %s, uma aplicação %s construída com %s. ",
		context.ProjectName, context.ProjectType, strings.Join(context.Technologies, ", ")))
	prd.WriteString("A solução aborda necessidades atuais de negócio mantendo excelência técnica e escalabilidade.\n\n")

	// Declaração do Problema
	prd.WriteString("## Declaração do Problema\n\n")
	prd.WriteString("O sistema atual requer melhorias para aprimorar a experiência do desenvolvedor, manutenibilidade e alinhamento com práticas modernas de engenharia de software. Esta PRD aborda esses desafios através de melhorias estruturadas.\n\n")

	// Objetivos e Metas
	prd.WriteString("## Objetivos e Metas\n\n")
	prd.WriteString("### Objetivos Principais\n")
	prd.WriteString("1. **Aprimorar Experiência do Desenvolvedor**: Melhorar documentação e onboarding para desenvolvedores júnior\n")
	prd.WriteString("2. **Melhorar Manutenibilidade**: Padronizar estrutura de código e documentação\n")
	prd.WriteString("3. **Aumentar Eficiência**: Otimizar fluxos de trabalho de desenvolvimento\n\n")

	prd.WriteString("### Critérios de Sucesso\n")
	prd.WriteString("- Documentação completa de API com exemplos\n")
	prd.WriteString("- Documentação abrangente do projeto\n")
	prd.WriteString("- Rastreamento claro de versões e alterações\n\n")

	// Requisitos Funcionais
	prd.WriteString("## Requisitos Funcionais\n\n")
	prd.WriteString("### RF001 - Documentação de API\n")
	prd.WriteString("**Descrição:** Documentação completa de todos os endpoints da API\n")
	prd.WriteString("**Prioridade:** Alta\n")
	prd.WriteString("**Critérios de Aceite:**\n")
	prd.WriteString("- Todos os endpoints documentados com exemplos\n")
	prd.WriteString("- Formatos de request/response especificados\n")
	prd.WriteString("- Tratamento de erros documentado\n\n")

	prd.WriteString("### RF002 - Gerenciamento de Versão\n")
	prd.WriteString("**Descrição:** Rastrear e documentar versões do projeto\n")
	prd.WriteString("**Prioridade:** Média\n")
	prd.WriteString("**Critérios de Aceite:**\n")
	prd.WriteString("- Números de versão seguem versionamento semântico\n")
	prd.WriteString("- Logs de alterações mantidos\n")
	prd.WriteString("- Documentação de lançamento gerada\n\n")

	// Requisitos Não Funcionais
	prd.WriteString("## Requisitos Não Funcionais\n\n")
	prd.WriteString("### RNF001 - Performance\n")
	prd.WriteString("**Descrição:** Tempo de resposta aceitável para documentação\n")
	prd.WriteString("**Prioridade:** Média\n")
	prd.WriteString("**Critérios de Aceite:**\n")
	prd.WriteString("- Documentação gerada em menos de 30 segundos\n")
	prd.WriteString("- Formatos legíveis e bem estruturados\n\n")

	// Considerações Técnicas
	prd.WriteString("## Considerações Técnicas\n\n")
	prd.WriteString("### Arquitetura\n")
	prd.WriteString("A solução segue padrões estabelecidos:\n")
	prd.WriteString("- Princípios de design de API RESTful\n")
	prd.WriteString("- Arquitetura modular\n")
	prd.WriteString("- Separação de responsabilidades\n\n")

	prd.WriteString("### Tecnologias\n")
	for _, tech := range context.Technologies {
		prd.WriteString(fmt.Sprintf("- **%s**: Componente principal da stack tecnológica\n", tech))
	}
	prd.WriteString("\n")

	// Dependências
	if len(context.Dependencies) > 0 {
		prd.WriteString("## Dependências\n\n")
		for _, dep := range context.Dependencies {
			prd.WriteString(fmt.Sprintf("- %s\n", dep))
		}
		prd.WriteString("\n")
	}

	// Riscos e Mitigações
	prd.WriteString("## Riscos e Mitigações\n\n")
	prd.WriteString("### Risco 1 - Complexidade Técnica\n")
	prd.WriteString("**Descrição:** Dificuldade em manter documentação sincronizada\n")
	prd.WriteString("**Mitigação:** Automação de geração de documentação\n\n")

	prd.WriteString("### Risco 2 - Adoção pela Equipe\n")
	prd.WriteString("**Descrição:** Resistência em usar novos padrões\n")
	prd.WriteString("**Mitigação:** Treinamento e documentação clara\n\n")

	return prd.String()
}

// GenerateTicketID gera um ID de ticket baseado no contexto
func GenerateTicketID(context *ProjectContext) string {
	// Simular geração de ticket baseado em padrões
	prefix := "DIST"
	if context.ProjectType == "nodejs" {
		prefix = "FEAT"
	} else if context.ProjectType == "go" {
		prefix = "BE"
	}

	// Gerar número sequencial simples
	return fmt.Sprintf("%s-%04d", prefix, time.Now().Unix()%10000)
}
