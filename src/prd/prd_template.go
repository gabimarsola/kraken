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

	prompt.WriteString(fmt.Sprintf("Como Product Manager sênior, crie uma PRD profissional no formato de engenharia de software para o ticket %s.\n\n", t.TicketID))

	prompt.WriteString("## CONTEXTO DO PROJETO\n\n")
	prompt.WriteString(fmt.Sprintf("**Projeto:** %s\n", context.ProjectName))
	prompt.WriteString(fmt.Sprintf("**Tipo:** %s\n", context.ProjectType))
	prompt.WriteString(fmt.Sprintf("**Versão:** %s\n", context.Version))
	prompt.WriteString(fmt.Sprintf("**Tecnologias:** %s\n", strings.Join(context.Technologies, ", ")))

	prompt.WriteString("\n## CONTEXTO GIT\n\n")
	prompt.WriteString(fmt.Sprintf("**Branch:** %s\n", gitContext.CurrentBranch))
	prompt.WriteString(fmt.Sprintf("**Último Commit:** %s - %s\n", gitContext.LastCommit[:7], gitContext.LastCommitMsg))

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
	prompt.WriteString("Crie uma PRD profissional seguindo este formato:\n\n")

	prompt.WriteString("1. **Header** com informações do ticket\n")
	prompt.WriteString("2. **Overview** - resumo executivo\n")
	prompt.WriteString("3. **Problem Statement** - problema a ser resolvido\n")
	prompt.WriteString("4. **Goals & Objectives** - metas SMART\n")
	prompt.WriteString("5. **User Stories** - histórias de usuário detalhadas\n")
	prompt.WriteString("6. **Functional Requirements** - requisitos funcionais\n")
	prompt.WriteString("7. **Non-Functional Requirements** - requisitos não funcionais\n")
	prompt.WriteString("8. **Technical Considerations** - considerações técnicas\n")
	prompt.WriteString("9. **Dependencies** - dependências externas\n")
	prompt.WriteString("10. **Risks & Mitigations** - riscos e mitigações\n")
	prompt.WriteString("11. **Success Metrics** - métricas de sucesso\n")
	prompt.WriteString("12. **Out of Scope** - fora do escopo\n")
	prompt.WriteString("13. **Timeline & Milestones** - cronograma\n")
	prompt.WriteString("14. **Stakeholders** - stakeholders\n")

	prompt.WriteString("\n## IMPORTANTE\n\n")
	prompt.WriteString("- Use linguagem profissional e técnica\n")
	prompt.WriteString("- Seja específico e mensurável\n")
	prompt.WriteString("- Inclua exemplos práticos\n")
	prompt.WriteString("- Considere o contexto atual do projeto\n")
	prompt.WriteString("- Foque em desenvolvedores júnior como audiência\n")

	return prompt.String()
}

func (t *PRDTemplate) generateBasicPRD(context *ProjectContext, gitContext *GitContext) string {
	var prd strings.Builder

	// Header
	prd.WriteString(fmt.Sprintf("# %s - Product Requirements Document\n\n", t.TicketID))
	prd.WriteString(fmt.Sprintf("**Project:** %s\n", t.ProjectName))
	prd.WriteString(fmt.Sprintf("**Status:** %s\n", t.Status))
	prd.WriteString(fmt.Sprintf("**Priority:** %s\n", t.Priority))
	prd.WriteString(fmt.Sprintf("**Created:** %s\n", t.CreatedDate.Format("2006-01-02")))
	prd.WriteString(fmt.Sprintf("**Author:** Engineering Team\n\n"))

	// Overview
	prd.WriteString("## Overview\n\n")
	prd.WriteString(fmt.Sprintf("This document outlines the requirements for %s, a %s application built with %s. ",
		context.ProjectName, context.ProjectType, strings.Join(context.Technologies, ", ")))
	prd.WriteString("The solution addresses current business needs while maintaining technical excellence and scalability.\n\n")

	// Problem Statement
	prd.WriteString("## Problem Statement\n\n")
	prd.WriteString("The current system requires enhancements to improve developer experience, maintainability, and alignment with modern software engineering practices. This PRD addresses these challenges through structured improvements.\n\n")

	// Goals & Objectives
	prd.WriteString("## Goals & Objectives\n\n")
	prd.WriteString("### Primary Goals\n")
	prd.WriteString("1. **Enhance Developer Experience**: Improve documentation and onboarding for junior developers\n")
	prd.WriteString("2. **Improve Maintainability**: Standardize code structure and documentation\n")
	prd.WriteString("3. **Increase Efficiency**: Streamline development workflows\n\n")

	prd.WriteString("### Success Criteria\n")
	prd.WriteString("- Complete API documentation with examples\n")
	prd.WriteString("- Comprehensive project documentation\n")
	prd.WriteString("- Clear version tracking and change management\n\n")

	// User Stories
	prd.WriteString("## User Stories\n\n")
	prd.WriteString("### US001 - As a Junior Developer\n")
	prd.WriteString("**I want** to understand the project structure quickly **so that** I can contribute effectively.\n\n")
	prd.WriteString("**Acceptance Criteria:**\n")
	prd.WriteString("- [ ] Complete API documentation is available\n")
	prd.WriteString("- [ ] Project setup instructions are clear\n")
	prd.WriteString("- [ ] Code examples are provided\n\n")

	prd.WriteString("### US002 - As a Team Lead\n")
	prd.WriteString("**I want** to track project changes easily **so that** I can manage releases effectively.\n\n")
	prd.WriteString("**Acceptance Criteria:**\n")
	prd.WriteString("- [ ] Version documentation is maintained\n")
	prd.WriteString("- [ ] Change history is tracked\n")
	prd.WriteString("- [ ] Release notes are generated\n\n")

	// Functional Requirements
	prd.WriteString("## Functional Requirements\n\n")
	prd.WriteString("### FR001 - API Documentation\n")
	prd.WriteString("**Description:** Complete documentation of all API endpoints\n")
	prd.WriteString("**Priority:** High\n")
	prd.WriteString("**Acceptance Criteria:**\n")
	prd.WriteString("- All endpoints documented with examples\n")
	prd.WriteString("- Request/response formats specified\n")
	prd.WriteString("- Error handling documented\n\n")

	prd.WriteString("### FR002 - Version Management\n")
	prd.WriteString("**Description:** Track and document project versions\n")
	prd.WriteString("**Priority:** Medium\n")
	prd.WriteString("**Acceptance Criteria:**\n")
	prd.WriteString("- Version numbers follow semantic versioning\n")
	prd.WriteString("- Change logs are maintained\n")
	prd.WriteString("- Release documentation is generated\n\n")

	// Technical Considerations
	prd.WriteString("## Technical Considerations\n\n")
	prd.WriteString("### Architecture\n")
	prd.WriteString("The solution follows established patterns:\n")
	prd.WriteString("- RESTful API design principles\n")
	prd.WriteString("- Modular architecture\n")
	prd.WriteString("- Separation of concerns\n\n")

	prd.WriteString("### Technologies\n")
	for _, tech := range context.Technologies {
		prd.WriteString(fmt.Sprintf("- **%s**: Primary technology stack component\n", tech))
	}
	prd.WriteString("\n")

	// Dependencies
	if len(context.Dependencies) > 0 {
		prd.WriteString("## Dependencies\n\n")
		for _, dep := range context.Dependencies {
			prd.WriteString(fmt.Sprintf("- %s\n", dep))
		}
		prd.WriteString("\n")
	}

	// Success Metrics
	prd.WriteString("## Success Metrics\n\n")
	prd.WriteString("1. **Documentation Coverage**: 100% of endpoints documented\n")
	prd.WriteString("2. **Developer Onboarding Time**: Reduced by 50%\n")
	prd.WriteString("3. **Code Quality**: Improved maintainability scores\n")
	prd.WriteString("4. **Release Efficiency**: Streamlined deployment process\n\n")

	// Out of Scope
	prd.WriteString("## Out of Scope\n\n")
	prd.WriteString("- Complete system redesign\n")
	prd.WriteString("- Database schema changes\n")
	prd.WriteString("- Third-party integrations\n")
	prd.WriteString("- Performance optimization\n\n")

	// Timeline
	prd.WriteString("## Timeline & Milestones\n\n")
	prd.WriteString("### Phase 1 (Week 1-2)\n")
	prd.WriteString("- Documentation analysis and planning\n")
	prd.WriteString("- Template development\n\n")

	prd.WriteString("### Phase 2 (Week 3-4)\n")
	prd.WriteString("- Documentation generation\n")
	prd.WriteString("- Review and refinement\n\n")

	prd.WriteString("### Phase 3 (Week 5-6)\n")
	prd.WriteString("- Final validation\n")
	prd.WriteString("- Release preparation\n\n")

	// Stakeholders
	prd.WriteString("## Stakeholders\n\n")
	prd.WriteString("- **Engineering Team**: Implementation and maintenance\n")
	prd.WriteString("- **Product Team**: Requirements validation\n")
	prd.WriteString("- **QA Team**: Quality assurance\n")
	prd.WriteString("- **DevOps Team**: Deployment and monitoring\n\n")

	// Git Context
	prd.WriteString("## Git Context\n\n")
	prd.WriteString(fmt.Sprintf("**Current Branch:** %s\n", gitContext.CurrentBranch))
	prd.WriteString(fmt.Sprintf("**Last Commit:** %s\n", gitContext.LastCommitMsg))

	if len(gitContext.Changes) > 0 {
		prd.WriteString("**Recent Changes:**\n")
		for _, change := range gitContext.Changes {
			prd.WriteString(fmt.Sprintf("- %s\n", change))
		}
	}
	prd.WriteString("\n")

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
