package prd

import (
	"fmt"
	"strings"
)

func (g *ContextualPRDGenerator) buildCompleteVersionPrompt(context *ProjectContext, gitContext *GitContext) string {
	var prompt strings.Builder

	prompt.WriteString("Como especialista técnico, analise este contexto completo e gere uma documentação detalhada da versão atual (incluindo documentação de API):\n\n")

	prompt.WriteString("## INFORMAÇÕES DO PROJETO\n\n")
	prompt.WriteString(fmt.Sprintf("**Projeto:** %s\n", context.ProjectName))
	prompt.WriteString(fmt.Sprintf("**Versão:** %s\n", context.Version))
	prompt.WriteString(fmt.Sprintf("**Tipo:** %s\n", context.ProjectType))
	prompt.WriteString(fmt.Sprintf("**Data:** %s\n", context.LastModified.Format("02/01/2006")))

	prompt.WriteString("\n## CONTEXTO GIT\n\n")
	prompt.WriteString(fmt.Sprintf("**Branch Atual:** %s\n", gitContext.CurrentBranch))
	prompt.WriteString(fmt.Sprintf("**Último Commit:** %s\n", gitContext.LastCommit))
	prompt.WriteString(fmt.Sprintf("**Mensagem:** %s\n", gitContext.LastCommitMsg))

	if gitContext.LastMerge != "" {
		prompt.WriteString(fmt.Sprintf("**Último Merge:** %s\n", gitContext.LastMerge))
	}

	if len(gitContext.Changes) > 0 {
		prompt.WriteString("\n**Arquivos Alterados no Último Commit:**\n")
		for _, change := range gitContext.Changes {
			prompt.WriteString(fmt.Sprintf("- %s\n", change))
		}
	}

	prompt.WriteString("\n## ENDPOINTS DA API\n\n")
	for _, endpoint := range context.Endpoints {
		prompt.WriteString(fmt.Sprintf("- %s %s: %s\n", endpoint.Method, endpoint.Path, endpoint.Description))
	}

	prompt.WriteString("\n## REQUISITOS PARA A DOCUMENTAÇÃO\n\n")
	prompt.WriteString("1. **Descrição do Funcionamento**: Explique detalhadamente como o sistema funciona\n")
	prompt.WriteString("2. **Explicação para Desenvolvedores Júnior**: Use linguagem clara, evite jargões\n")
	prompt.WriteString("3. **Documentação de API**: Inclua todos os endpoints com exemplos\n")
	prompt.WriteString("4. **Alterações da Versão**: Detalhe as mudanças desde a última versão\n")
	prompt.WriteString("5. **Contexto Git**: Explique o que mudou no último commit/merge\n")

	return prompt.String()
}

func (g *ContextualPRDGenerator) generateBasicCompleteVersionDoc(context *ProjectContext, gitContext *GitContext) string {
	var doc strings.Builder

	version := strings.ReplaceAll(context.Version, ".", "")

	doc.WriteString(fmt.Sprintf("# Versão %s - Documentação Completa\n\n", version))
	doc.WriteString(fmt.Sprintf("**Projeto:** %s\n", context.ProjectName))
	doc.WriteString(fmt.Sprintf("**Data:** %s\n\n", context.LastModified.Format("02/01/2006")))

	doc.WriteString("## 🌿 Contexto Git\n\n")
	doc.WriteString(fmt.Sprintf("**Branch Atual:** %s\n", gitContext.CurrentBranch))
	doc.WriteString(fmt.Sprintf("**Último Commit:** %s\n", gitContext.LastCommit))
	doc.WriteString(fmt.Sprintf("**Mensagem:** %s\n", gitContext.LastCommitMsg))

	if len(gitContext.Changes) > 0 {
		doc.WriteString("\n### Arquivos Alterados no Último Commit\n\n")
		for _, change := range gitContext.Changes {
			doc.WriteString(fmt.Sprintf("- %s\n", change))
		}
	}

	doc.WriteString("\n## 📋 Descrição do Funcionamento\n\n")
	doc.WriteString(fmt.Sprintf("O %s é uma aplicação %s que oferece uma API REST para ", context.ProjectName, context.ProjectType))
	doc.WriteString("interação via HTTP. A aplicação segue uma arquitetura em camadas, separando ")
	doc.WriteString("responsabilidades entre roteamento, controle de negócio e persistência de dados.\n\n")

	// Incluir documentação da API
	if len(context.Endpoints) > 0 {
		doc.WriteString("## 📚 Documentação da API\n\n")

		for _, endpoint := range context.Endpoints {
			doc.WriteString(fmt.Sprintf("#### %s %s\n\n", endpoint.Method, endpoint.Path))
			doc.WriteString(fmt.Sprintf("**Descrição:** %s\n\n", endpoint.Description))

			doc.WriteString("**Exemplo de Requisição:**\n")
			doc.WriteString("```bash\n")
			doc.WriteString(fmt.Sprintf("curl -X %s 'http://localhost:3000%s'", endpoint.Method, endpoint.Path))
			doc.WriteString("\n```\n\n")
		}
	}

	return doc.String()
}
