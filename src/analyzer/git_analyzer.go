package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type GitAnalyzer struct {
	projectPath string
}

type GitChange struct {
	FilePath    string
	Status      string // Modified, Added, Deleted, Renamed
	Content     string
	OldContent  string // para comparação
	Timestamp   time.Time
	Author      string
	Message     string // commit message
}

type GitAnalysis struct {
	Changes      []GitChange
	LastCommit   string
	Branch       string
	IsClean      bool
	ModifiedFiles []string
}

func NewGitAnalyzer(projectPath string) *GitAnalyzer {
	return &GitAnalyzer{
		projectPath: projectPath,
	}
}

func (g *GitAnalyzer) AnalyzeChanges() (*GitAnalysis, error) {
	// Verificar se é um repositório git
	if !g.isGitRepository() {
		return nil, fmt.Errorf("não é um repositório git")
	}

	analysis := &GitAnalysis{}

	// Obter informações básicas
	branch, err := g.getCurrentBranch()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter branch: %v", err)
	}
	analysis.Branch = branch

	// Verificar se há alterações não commitadas
	isClean, err := g.isWorkingTreeClean()
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar status: %v", err)
	}
	analysis.IsClean = isClean

	// Analisar alterações
	if !isClean {
		// Há alterações não commitadas
		changes, err := g.getUnstagedChanges()
		if err != nil {
			return nil, fmt.Errorf("erro ao obter alterações não commitadas: %v", err)
		}
		analysis.Changes = changes
	} else {
		// Não há alterações, analisar último commit
		lastCommit, err := g.getLastCommitChanges()
		if err != nil {
			return nil, fmt.Errorf("erro ao obter último commit: %v", err)
		}
		analysis.Changes = lastCommit
	}

	// Obter último commit hash
	lastCommitHash, err := g.getLastCommitHash()
	if err == nil {
		analysis.LastCommit = lastCommitHash
	}

	return analysis, nil
}

func (g *GitAnalyzer) isGitRepository() bool {
	_, err := os.Stat(fmt.Sprintf("%s/.git", g.projectPath))
	return err == nil
}

func (g *GitAnalyzer) getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = g.projectPath
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (g *GitAnalyzer) isWorkingTreeClean() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = g.projectPath
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return len(strings.TrimSpace(string(output))) == 0, nil
}

func (g *GitAnalyzer) getUnstagedChanges() ([]GitChange, error) {
	var changes []GitChange

	// Obter status dos arquivos
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = g.projectPath
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 3 {
			continue
		}

		status := line[:2]
		filePath := line[3:]

		// Ignorar arquivos que não queremos analisar
		if g.shouldIgnoreFile(filePath) {
			continue
		}

		change := GitChange{
			FilePath: filePath,
			Status:   g.parseStatus(status),
		}

		// Obter conteúdo do arquivo
		if change.Status != "Deleted" {
			content, err := g.getFileContent(filePath)
			if err == nil {
				change.Content = content
			}
		}

		changes = append(changes, change)
	}

	return changes, nil
}

func (g *GitAnalyzer) getLastCommitChanges() ([]GitChange, error) {
	var changes []GitChange

	// Obter arquivos do último commit
	cmd := exec.Command("git", "diff", "--name-status", "HEAD~1", "HEAD")
	cmd.Dir = g.projectPath
	output, err := cmd.Output()
	if err != nil {
		// Se não há commit anterior, retorna vazio
		return changes, nil
	}

	// Obter mensagem do último commit
	commitMsg, _ := g.getLastCommitMessage()

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 2 {
			continue
		}

		status := line[:1]
		filePath := line[2:]

		// Ignorar arquivos que não queremos analisar
		if g.shouldIgnoreFile(filePath) {
			continue
		}

		change := GitChange{
			FilePath: filePath,
			Status:   g.parseStatus(status),
			Message:  commitMsg,
		}

		// Obter conteúdo atual do arquivo
		if change.Status != "Deleted" {
			content, err := g.getFileContent(filePath)
			if err == nil {
				change.Content = content
			}
		}

		changes = append(changes, change)
	}

	return changes, nil
}

func (g *GitAnalyzer) getLastCommitHash() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = g.projectPath
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (g *GitAnalyzer) getLastCommitMessage() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=%B")
	cmd.Dir = g.projectPath
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (g *GitAnalyzer) getFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(fmt.Sprintf("%s/%s", g.projectPath, filePath))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (g *GitAnalyzer) parseStatus(status string) string {
	switch status {
	case "M", " M":
		return "Modified"
	case "A", " A":
		return "Added"
	case "D", " D":
		return "Deleted"
	case "R", " R":
		return "Renamed"
	case "C", " C":
		return "Copied"
	case "??":
		return "Untracked"
	default:
		return "Unknown"
	}
}

func (g *GitAnalyzer) shouldIgnoreFile(filePath string) bool {
	// Ignorar arquivos comuns que não são relevantes para PRDs
	ignorePatterns := []string{
		".git",
		"node_modules",
		"vendor",
		".vscode",
		".windsurf",
		".cursor",
		"*.log",
		"*.tmp",
		"dist/",
		"build/",
		"*.exe",
		"kraken",
		"test_",
		"_test.go",
		"README.md",
		".md",
	}

	for _, pattern := range ignorePatterns {
		if strings.Contains(filePath, pattern) {
			return true
		}
	}

	return false
}

func (g *GitAnalyzer) GetChangesSummary() string {
	analysis, err := g.AnalyzeChanges()
	if err != nil {
		return fmt.Sprintf("Erro ao analisar: %v", err)
	}

	var summary strings.Builder
	summary.WriteString(fmt.Sprintf("📊 Análise Git - Branch: %s\n", analysis.Branch))
	
	if analysis.IsClean {
		summary.WriteString("✅ Working tree limpo\n")
		if analysis.LastCommit != "" {
			summary.WriteString(fmt.Sprintf("📝 Analisando último commit: %s\n", analysis.LastCommit[:7]))
		}
	} else {
		summary.WriteString("⚠️ Há alterações não commitadas\n")
	}

	summary.WriteString(fmt.Sprintf("📁 %d arquivos alterados:\n\n", len(analysis.Changes)))

	for _, change := range analysis.Changes {
		emoji := g.getStatusEmoji(change.Status)
		summary.WriteString(fmt.Sprintf("%s %s - %s\n", emoji, change.Status, change.FilePath))
		
		// Adicionar resumo do conteúdo se for arquivo de código
		if g.isCodeFile(change.FilePath) && len(change.Content) > 0 {
			lines := strings.Split(change.Content, "\n")
			summary.WriteString(fmt.Sprintf("   📝 %d linhas de código\n", len(lines)))
			
			// Mostrar funções/clases principais
			if strings.HasSuffix(change.FilePath, ".go") {
				functions := g.extractGoFunctions(change.Content)
				if len(functions) > 0 {
					summary.WriteString(fmt.Sprintf("   🔧 Funções: %s\n", strings.Join(functions, ", ")))
				}
			}
		}
		
		if change.Message != "" {
			summary.WriteString(fmt.Sprintf("   💬 %s\n", change.Message))
		}
		summary.WriteString("\n")
	}

	return summary.String()
}

func (g *GitAnalyzer) getStatusEmoji(status string) string {
	switch status {
	case "Added":
		return "➕"
	case "Modified":
		return "✏️"
	case "Deleted":
		return "🗑️"
	case "Renamed":
		return "🔄"
	default:
		return "📄"
	}
}

func (g *GitAnalyzer) isCodeFile(filePath string) bool {
	codeExtensions := []string{
		".go", ".js", ".ts", ".py", ".java", ".cpp", ".c", ".h",
		".cs", ".php", ".rb", ".swift", ".kt", ".rs", ".dart",
	}

	for _, ext := range codeExtensions {
		if strings.HasSuffix(filePath, ext) {
			return true
		}
	}
	return false
}

func (g *GitAnalyzer) extractGoFunctions(content string) []string {
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
