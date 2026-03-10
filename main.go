package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"kraken/src/ai"
	"kraken/src/analyzer"
	"kraken/src/detector"
	"kraken/src/generator"
	"kraken/src/parser"
	"kraken/src/prd"
	"kraken/src/structure"
	"os"
	"strings"
)

func main() {
	projectPath := "."

	fmt.Println("🔍 Detectando tipo de projeto...")
	projectType := detector.DetectProjectType(projectPath)

	if projectType == detector.ProjectTypeUnknown {
		fmt.Println("❌ Tipo de projeto não reconhecido.")
		fmt.Println("   Certifique-se de que existe um dos seguintes arquivos:")
		fmt.Println("   - go.mod (para projetos Go)")
		fmt.Println("   - package.json (para projetos Node.js)")
		fmt.Println("   - pom.xml (para projetos Java/Maven)")
		return
	}

	fmt.Printf("✅ Projeto detectado: %s\n", projectType)

	fmt.Println("📖 Parseando informações do projeto...")
	info, err := parser.ParseProject(projectPath, projectType)
	if err != nil {
		fmt.Printf("❌ Erro ao parsear projeto: %v\n", err)
		return
	}

	// Menu de opções
	for {
		fmt.Println("\n🚀 Kraken - Menu Principal")
		fmt.Println("========================")
		fmt.Println("1. Gerar documentos com IA externa")
		fmt.Println("2. Gerar documentos com IA da IDE")
		fmt.Println("3. Configurar provedor de IA")
		fmt.Println("4. Sair")
		fmt.Print("Escolha uma opção: ")

		choice := getInput()

		switch choice {
		case "1":
			generateDocumentsWithExternalAI(info, projectPath)
		case "2":
			generateDocumentsWithIDEAI(info, projectPath)
		case "3":
			configureAI()
		case "4":
			fmt.Println("👋 Até logo!")
			return
		default:
			fmt.Println("❌ Opção inválida. Tente novamente.")
		}
	}
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func generateDocumentsWithExternalAI(info *structure.ProjectInfo, projectPath string) {
	fmt.Println("\n🤖 Gerando documentos com IA externa...")

	// Ler configuração de IA
	config, provider, err := loadAIConfig()
	if err != nil {
		fmt.Printf("❌ Erro na configuração de IA: %v\n", err)
		fmt.Println("� Use a opção 3 para configurar o provedor de IA")
		return
	}

	// Analisar endpoints
	fmt.Println("� Analisando endpoints do projeto...")
	endpoints, err := analyzer.AnalyzeEndpoints(projectPath)
	if err != nil {
		fmt.Printf("⚠️  Aviso: Erro ao analisar endpoints: %v\n", err)
		endpoints = []structure.Endpoint{}
	}

	fmt.Printf("📊 Encontrados %d endpoints documentados\n", len(endpoints))
	info.Endpoints = endpoints

	// Gerar documentação técnica
	fmt.Println("📝 Gerando documentação técnica...")
	generatedFiles, err := generator.GenerateAllRouteDocumentation(info, projectPath)
	if err != nil {
		fmt.Printf("❌ Erro ao gerar documentação: %v\n", err)
		return
	}

	// Gerar PRD com IA
	fmt.Println("📋 Gerando PRD com IA...")
	aiGenerator, err := prd.NewAIPRDGenerator(info, provider, config)
	if err != nil {
		fmt.Printf("❌ Erro ao criar gerador AI: %v\n", err)
		return
	}

	err = aiGenerator.GeneratePRDFromProject()
	if err != nil {
		fmt.Printf("❌ Erro ao gerar PRD com IA: %v\n", err)
		return
	}

	fmt.Printf("✅ Documentos gerados com sucesso!\n")
	fmt.Printf("📄 %d arquivos de documentação criados:\n", len(generatedFiles))
	for _, file := range generatedFiles {
		fmt.Printf("   - %s\n", file)
	}
	fmt.Printf("📄 PRD criado em docs/kraken/\n")
}

func generateDocumentsWithIDEAI(info *structure.ProjectInfo, projectPath string) {
	fmt.Println("\n🤖 Gerando documentos com IA da IDE...")

	// Analisar endpoints
	fmt.Println("🔍 Analisando endpoints do projeto...")
	endpoints, err := analyzer.AnalyzeEndpoints(projectPath)
	if err != nil {
		fmt.Printf("⚠️  Aviso: Erro ao analisar endpoints: %v\n", err)
		endpoints = []structure.Endpoint{}
	}

	fmt.Printf("📊 Encontrados %d endpoints documentados\n", len(endpoints))
	info.Endpoints = endpoints

	// Gerar documentação técnica
	fmt.Println("📝 Gerando documentação técnica...")
	generatedFiles, err := generator.GenerateAllRouteDocumentation(info, projectPath)
	if err != nil {
		fmt.Printf("❌ Erro ao gerar documentação: %v\n", err)
		return
	}

	// Gerar PRD com IA da IDE
	fmt.Println("📋 Gerando PRD com IA da IDE...")
	ideGenerator, err := prd.NewIDEPRDGenerator(info)
	if err != nil {
		fmt.Printf("❌ Erro ao criar gerador IDE: %v\n", err)
		if err.Error() == "IDE não detectada ou não suportada" {
			fmt.Println("💡 Dicas para detectar a IDE:")
			fmt.Println("   - Windsurf: Verifique se o diretório .windsurf existe")
			fmt.Println("   - Cursor: Verifique se o diretório .cursor existe")
			fmt.Println("   - VS Code: Verifique se o diretório .vscode existe")
			fmt.Println("   - IntelliJ: Verifique se o diretório .idea existe")
		}
		return
	}

	err = ideGenerator.GeneratePRDWithIDE()
	if err != nil {
		fmt.Printf("❌ Erro ao gerar PRD com IA da IDE: %v\n", err)
		return
	}

	fmt.Printf("✅ Documentos gerados com sucesso!\n")
	fmt.Printf("📄 %d arquivos de documentação criados:\n", len(generatedFiles))
	for _, file := range generatedFiles {
		fmt.Printf("   - %s\n", file)
	}
	fmt.Printf("📄 PRD criado em docs/kraken/\n")
}

func configureAI() {
	fmt.Println("\n⚙️ Configuração de Provedor de IA")
	fmt.Println("================================")
	fmt.Println("Provedores disponíveis:")
	fmt.Println("1. OpenAI (GPT-4)")
	fmt.Println("2. Anthropic (Claude)")
	fmt.Println("3. Ollama (Local)")
	fmt.Println("4. Google Gemini")
	fmt.Print("Escolha o provedor: ")

	choice := getInput()

	var provider ai.AIProvider
	var config map[string]string

	switch choice {
	case "1":
		provider = ai.ProviderOpenAI
		fmt.Print("API Key da OpenAI: ")
		apiKey := getInput()
		config = map[string]string{"api_key": apiKey}

	case "2":
		provider = ai.ProviderAnthropic
		fmt.Print("API Key da Anthropic: ")
		apiKey := getInput()
		config = map[string]string{"api_key": apiKey}

	case "3":
		provider = ai.ProviderOllama
		fmt.Print("URL do Ollama (padrão: http://localhost:11434): ")
		baseURL := getInput()
		if baseURL == "" {
			baseURL = "http://localhost:11434"
		}
		fmt.Print("Modelo (padrão: llama2): ")
		model := getInput()
		if model == "" {
			model = "llama2"
		}
		config = map[string]string{
			"base_url": baseURL,
			"model":    model,
		}

	case "4":
		provider = ai.ProviderGemini
		fmt.Print("API Key da Gemini: ")
		apiKey := getInput()
		config = map[string]string{"api_key": apiKey}

	default:
		fmt.Println("❌ Provedor inválido")
		return
	}

	// Salvar configuração
	err := saveAIConfig(provider, config)
	if err != nil {
		fmt.Printf("❌ Erro ao salvar configuração: %v\n", err)
		return
	}

	fmt.Println("✅ Configuração salva com sucesso!")
}

func loadAIConfig() (map[string]string, ai.AIProvider, error) {
	file, err := os.ReadFile("ai_config.json")
	if err != nil {
		return nil, "", fmt.Errorf("arquivo de configuração não encontrado")
	}

	var configData struct {
		Provider string            `json:"provider"`
		Config   map[string]string `json:"config"`
	}

	err = json.Unmarshal(file, &configData)
	if err != nil {
		return nil, "", fmt.Errorf("erro ao ler configuração")
	}

	var provider ai.AIProvider
	switch configData.Provider {
	case "openai":
		provider = ai.ProviderOpenAI
	case "anthropic":
		provider = ai.ProviderAnthropic
	case "ollama":
		provider = ai.ProviderOllama
	case "gemini":
		provider = ai.ProviderGemini
	default:
		return nil, "", fmt.Errorf("provedor desconhecido")
	}

	return configData.Config, provider, nil
}

func saveAIConfig(provider ai.AIProvider, config map[string]string) error {
	configData := struct {
		Provider string            `json:"provider"`
		Config   map[string]string `json:"config"`
	}{
		Provider: string(provider),
		Config:   config,
	}

	jsonData, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("ai_config.json", jsonData, 0644)
}

func enhanceDocumentationWithAI(info *structure.ProjectInfo, generatedFiles []string) {
	fmt.Println("\n🤖 Aprimorando documentação com IA...")

	// Ler configuração de IA
	config, provider, err := loadAIConfig()
	if err != nil {
		fmt.Printf("❌ Erro na configuração de IA: %v\n", err)
		fmt.Println("💡 Use a opção 5 para configurar o provedor de IA")
		return
	}

	// Criar gerenciador de IA
	aiManager, err := ai.NewAIManager(provider, config)
	if err != nil {
		fmt.Printf("❌ Erro ao criar gerenciador de IA: %v\n", err)
		return
	}

	// Preparar endpoints para a IA
	var endpointStrings []string
	for _, endpoint := range info.Endpoints {
		endpointStr := fmt.Sprintf("%s %s - %s", endpoint.Method, endpoint.Path, endpoint.Description)
		if endpoint.Summary != "" {
			endpointStr += fmt.Sprintf(" (%s)", endpoint.Summary)
		}
		endpointStrings = append(endpointStrings, endpointStr)
	}

	// Gerar documentação aprimorada
	fmt.Println("📝 Gerando documentação aprimorada...")
	enhancedDoc, err := aiManager.GenerateDocumentation(endpointStrings)
	if err != nil {
		fmt.Printf("❌ Erro ao gerar documentação aprimorada: %v\n", err)
		return
	}

	// Salvar documentação aprimorada
	enhancedFile := "docs/kraken/API-DOC-ENHANCED.md"
	err = os.WriteFile(enhancedFile, []byte(enhancedDoc), 0644)
	if err != nil {
		fmt.Printf("❌ Erro ao salvar documentação aprimorada: %v\n", err)
		return
	}

	fmt.Printf("✅ Documentação aprimorada criada com sucesso!\n")
	fmt.Printf("📄 Arquivo: %s\n", enhancedFile)
}
