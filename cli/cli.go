package cli

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

const (
	errorMessage = "❌ Erro: %v\n"
)

type CLI struct {
	projectPath string
	info        *structure.ProjectInfo
}

func NewCLI(projectPath string) *CLI {
	return &CLI{
		projectPath: projectPath,
	}
}

func (cli *CLI) Run() error {
	// Detectar tipo de projeto
	fmt.Println("🔍 Detectando tipo de projeto...")
	projectType := detector.DetectProjectType(cli.projectPath)

	if projectType == detector.ProjectTypeUnknown {
		fmt.Println("❌ Tipo de projeto não reconhecido.")
		fmt.Println("   Certifique-se de que existe um dos seguintes arquivos:")
		fmt.Println("   - go.mod (para projetos Go)")
		fmt.Println("   - package.json (para projetos Node.js)")
		fmt.Println("   - pom.xml (para projetos Java/Maven)")
		return fmt.Errorf("tipo de projeto não reconhecido")
	}

	fmt.Printf("✅ Projeto detectado: %s\n", projectType)

	// Parsear informações do projeto
	fmt.Println("📖 Parseando informações do projeto...")
	info, err := parser.ParseProject(cli.projectPath, projectType)
	if err != nil {
		return fmt.Errorf("erro ao parsear projeto: %v", err)
	}
	cli.info = info

	// Menu de opções
	return cli.showMenu()
}

func (cli *CLI) showMenu() error {
	for {
		fmt.Println("\n🚀 Kraken - Menu Principal")
		fmt.Println("========================")
		fmt.Println("1. Gerar documentos com IA externa")
		fmt.Println("2. Gerar documentos com IA da IDE")
		fmt.Println("3. Configurar provedor de IA")
		fmt.Println("4. Sair")
		fmt.Print("Escolha uma opção: ")

		choice := cli.getInput()

		switch choice {
		case "1":
			err := cli.generateDocumentsWithExternalAI()
			if err != nil {
				fmt.Printf(errorMessage, err)
			}
		case "2":
			err := cli.generateDocumentsWithIDEAI()
			if err != nil {
				fmt.Printf(errorMessage, err)
			}
		case "3":
			err := cli.configureAI()
			if err != nil {
				fmt.Printf(errorMessage, err)
			}
		case "4":
			fmt.Println("👋 Até logo!")
			return nil
		default:
			fmt.Println("❌ Opção inválida. Tente novamente.")
		}
	}
}

func (cli *CLI) getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (cli *CLI) generateDocumentsWithExternalAI() error {
	fmt.Println("\n🤖 Gerando documentos com IA externa...")

	// Ler configuração de IA
	config, provider, err := cli.loadAIConfig()
	if err != nil {
		fmt.Printf("❌ Erro na configuração de IA: %v\n", err)
		fmt.Println("💡 Use a opção 3 para configurar o provedor de IA")
		return err
	}

	// Analisar endpoints
	fmt.Println("🔍 Analisando endpoints do projeto...")
	endpoints, err := analyzer.AnalyzeEndpoints(cli.projectPath)
	if err != nil {
		fmt.Printf("⚠️  Aviso: Erro ao analisar endpoints: %v\n", err)
		endpoints = []structure.Endpoint{}
	}

	fmt.Printf("📊 Encontrados %d endpoints documentados\n", len(endpoints))
	cli.info.Endpoints = endpoints

	// Gerar documentação completa contextualizada
	fmt.Println("� Gerando documentação completa contextualizada...")
	contextualGen := prd.NewContextualPRDGenerator(cli.info)
	contextualGen.SetAIProvider(provider, config)

	err = contextualGen.GenerateCompleteDocumentation()
	if err != nil {
		return fmt.Errorf("erro ao gerar documentação completa: %v", err)
	}

	return nil
}

func (cli *CLI) generateDocumentsWithIDEAI() error {
	fmt.Println("\n🤖 Gerando documentos com IA da IDE...")

	// Analisar endpoints
	fmt.Println("🔍 Analisando endpoints do projeto...")
	endpoints, err := analyzer.AnalyzeEndpoints(cli.projectPath)
	if err != nil {
		fmt.Printf("⚠️  Aviso: Erro ao analisar endpoints: %v\n", err)
		endpoints = []structure.Endpoint{}
	}

	fmt.Printf("📊 Encontrados %d endpoints documentados\n", len(endpoints))
	cli.info.Endpoints = endpoints

	// Gerar documentação técnica
	fmt.Println("📝 Gerando documentação técnica...")
	generatedFiles, err := generator.GenerateAllRouteDocumentation(cli.info, cli.projectPath)
	if err != nil {
		return fmt.Errorf("erro ao gerar documentação: %v", err)
	}

	// Gerar PRD com IA da IDE
	fmt.Println("📋 Gerando PRD com IA da IDE...")
	ideGenerator, err := prd.NewIDEPRDGenerator(cli.info)
	if err != nil {
		fmt.Printf("❌ Erro ao criar gerador IDE: %v\n", err)
		if err.Error() == "IDE não detectada ou não suportada" {
			fmt.Println("💡 Dicas para detectar a IDE:")
			fmt.Println("   - Windsurf: Verifique se o diretório .windsurf existe")
			fmt.Println("   - Cursor: Verifique se o diretório .cursor existe")
			fmt.Println("   - VS Code: Verifique se o diretório .vscode existe")
			fmt.Println("   - IntelliJ: Verifique se o diretório .idea existe")
		}
		return err
	}

	err = ideGenerator.GeneratePRDWithIDE()
	if err != nil {
		return fmt.Errorf("erro ao gerar PRD com IA da IDE: %v", err)
	}

	fmt.Printf("✅ Documentos gerados com sucesso!\n")
	fmt.Printf("📄 %d arquivos de documentação criados:\n", len(generatedFiles))
	for _, file := range generatedFiles {
		fmt.Printf("   - %s\n", file)
	}
	fmt.Printf("📄 PRD criado em docs/kraken/\n")

	return nil
}

func (cli *CLI) configureAI() error {
	fmt.Println("\n⚙️ Configuração de Provedor de IA")
	fmt.Println("================================")
	fmt.Println("Provedores disponíveis:")
	fmt.Println("1. OpenAI (GPT-4)")
	fmt.Println("2. Anthropic (Claude)")
	fmt.Println("3. Ollama (Local)")
	fmt.Println("4. Google Gemini")
	fmt.Print("Escolha o provedor: ")

	choice := cli.getInput()

	var provider ai.AIProvider
	var config map[string]string

	switch choice {
	case "1":
		provider = ai.ProviderOpenAI
		fmt.Print("API Key da OpenAI: ")
		apiKey := cli.getInput()
		config = map[string]string{"api_key": apiKey}

	case "2":
		provider = ai.ProviderAnthropic
		fmt.Print("API Key da Anthropic: ")
		apiKey := cli.getInput()
		config = map[string]string{"api_key": apiKey}

	case "3":
		provider = ai.ProviderOllama
		fmt.Print("URL do Ollama (padrão: http://localhost:11434): ")
		baseURL := cli.getInput()
		if baseURL == "" {
			baseURL = "http://localhost:11434"
		}
		fmt.Print("Modelo (padrão: llama2): ")
		model := cli.getInput()
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
		apiKey := cli.getInput()
		config = map[string]string{"api_key": apiKey}

	default:
		return fmt.Errorf("provedor inválido")
	}

	// Salvar configuração
	err := cli.saveAIConfig(provider, config)
	if err != nil {
		return fmt.Errorf("erro ao salvar configuração: %v", err)
	}

	fmt.Println("✅ Configuração salva com sucesso!")
	return nil
}

func (cli *CLI) loadAIConfig() (map[string]string, ai.AIProvider, error) {
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

func (cli *CLI) saveAIConfig(provider ai.AIProvider, config map[string]string) error {
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
