package prd

import (
	"bufio"
	"fmt"
	"kraken/src/structure"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type PRDGenerator struct {
	projectInfo *structure.ProjectInfo
}

func NewPRDGenerator(projectInfo *structure.ProjectInfo) *PRDGenerator {
	return &PRDGenerator{
		projectInfo: projectInfo,
	}
}

func (g *PRDGenerator) CreateInteractivePRD() error {
	fmt.Println("📝 Criador de PRD Interativo")
	fmt.Println("============================")
	fmt.Println()

	prd := structure.PRD{}

	// Coletar informações básicas
	prd.Title = g.getInput("Qual o título da funcionalidade? ")

	fmt.Println("\n📋 Descreva brevemente a funcionalidade que você quer implementar:")
	intro := g.getMultiLineInput()
	prd.Introduction = intro

	fmt.Println("\n🎯 Quais são os objetivos principais desta funcionalidade?")
	prd.Objectives = g.getMultiLineInputList()

	fmt.Println("\n👥 Quem é o usuário-alvo desta funcionalidade?")
	targetUser := g.getInput()
	fmt.Printf("Usuário-alvo: %s\n", targetUser)

	fmt.Println("\n📖 Vamos criar as histórias de usuário. Pressione Enter para parar.")
	prd.UserStories = g.getUserStories()

	fmt.Println("\n⚙️ Quais são os requisitos funcionais?")
	prd.FunctionalReqs = g.getFunctionalRequirements()

	fmt.Println("\n🚫 O que está fora do escopo desta funcionalidade?")
	prd.OutOfScope = g.getMultiLineInputList()

	fmt.Println("\n🎨 Quais são as considerações de design?")
	prd.DesignConsiderations = g.getMultiLineInputList()

	fmt.Println("\n🔧 Quais são as considerações técnicas?")
	prd.TechConsiderations = g.getMultiLineInputList()

	fmt.Println("\n📊 Como mediremos o sucesso desta funcionalidade?")
	prd.SuccessMetrics = g.getMultiLineInputList()

	fmt.Println("\n❓ Quais questões ainda estão em aberto?")
	prd.OpenQuestions = g.getMultiLineInputList()

	// Gerar arquivo PRD
	return g.generatePRDFile(prd)
}

func (g *PRDGenerator) getInput(prompt ...string) string {
	if len(prompt) > 0 {
		fmt.Print(prompt[0])
	} else {
		fmt.Print("> ")
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (g *PRDGenerator) getMultiLineInput(prompt ...string) string {
	if len(prompt) > 0 {
		fmt.Println(prompt[0])
	}
	fmt.Println("Digite o texto (pressione Enter duas vezes para finalizar):")

	var lines []string
	reader := bufio.NewReader(os.Stdin)
	emptyLineCount := 0

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			emptyLineCount++
			if emptyLineCount >= 2 {
				break
			}
		} else {
			emptyLineCount = 0
			lines = append(lines, line)
		}
	}

	return strings.Join(lines, "\n")
}

func (g *PRDGenerator) getMultiLineInputList() []string {
	fmt.Println("Digite os itens (um por linha, pressione Enter duas vezes para finalizar):")

	var items []string
	reader := bufio.NewReader(os.Stdin)
	emptyLineCount := 0

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			emptyLineCount++
			if emptyLineCount >= 2 {
				break
			}
		} else {
			emptyLineCount = 0
			items = append(items, line)
		}
	}

	return items
}

func (g *PRDGenerator) getUserStories() []structure.UserStory {
	var stories []structure.UserStory
	storyCount := 1

	for {
		fmt.Printf("\nHistória %d (pressione Enter para pular):\n", storyCount)

		title := g.getInput("Título: ")
		if title == "" {
			break
		}

		description := g.getMultiLineInput("Descrição:")

		fmt.Println("Critérios de aceitação:")
		criteria := g.getMultiLineInputList()

		story := structure.UserStory{
			ID:                 fmt.Sprintf("US%03d", storyCount),
			Title:              title,
			Description:        description,
			AcceptanceCriteria: criteria,
		}

		stories = append(stories, story)
		storyCount++
	}

	return stories
}

func (g *PRDGenerator) getFunctionalRequirements() []structure.FunctionalRequirement {
	var reqs []structure.FunctionalRequirement
	reqCount := 1

	for {
		fmt.Printf("\nRequisito %d (pressione Enter para pular):\n", reqCount)

		title := g.getInput("Título: ")
		if title == "" {
			break
		}

		description := g.getMultiLineInput("Descrição:")
		priority := g.getInput("Prioridade (high/medium/low): ")
		if priority == "" {
			priority = "medium"
		}

		req := structure.FunctionalRequirement{
			ID:          fmt.Sprintf("FR%03d", reqCount),
			Title:       title,
			Description: description,
			Priority:    priority,
		}

		reqs = append(reqs, req)
		reqCount++
	}

	return reqs
}

func (g *PRDGenerator) generatePRDFile(prd structure.PRD) error {
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

	fmt.Printf("\n✅ PRD gerado com sucesso!\n")
	fmt.Printf("📄 Arquivo criado: %s\n", filePath)

	return nil
}
