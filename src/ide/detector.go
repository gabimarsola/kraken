package ide

import (
	"os"
	"strings"
)

type IDEType string

const (
	IDEWindsurf IDEType = "windsurf"
	IDEVSCode   IDEType = "vscode"
	IDECursor   IDEType = "cursor"
	IDEIntelliJ IDEType = "intellij"
	IDEUnknown  IDEType = "unknown"
)

type IDEInfo struct {
	Type      IDEType
	Name      string
	Version   string
	Workspace string
	HasAI     bool
}

func DetectIDE() IDEInfo {
	// Detectar Windsurf
	if windsurfInfo := detectWindsurf(); windsurfInfo.Type != IDEUnknown {
		return windsurfInfo
	}

	// Detectar VS Code
	if vscodeInfo := detectVSCode(); vscodeInfo.Type != IDEUnknown {
		return vscodeInfo
	}

	// Detectar Cursor
	if cursorInfo := detectCursor(); cursorInfo.Type != IDEUnknown {
		return cursorInfo
	}

	// Detectar IntelliJ
	if intellijInfo := detectIntelliJ(); intellijInfo.Type != IDEUnknown {
		return intellijInfo
	}

	return IDEInfo{Type: IDEUnknown}
}

func detectWindsurf() IDEInfo {
	// Verificar variáveis de ambiente do Windsurf
	if wsPath := os.Getenv("WINDSURF_PATH"); wsPath != "" {
		return IDEInfo{
			Type:      IDEWindsurf,
			Name:      "Windsurf",
			Workspace: wsPath,
			HasAI:     true,
		}
	}

	// Verificar arquivos de configuração do Windsurf
	if _, err := os.Stat(".windsurf"); err == nil {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDEWindsurf,
			Name:      "Windsurf",
			Workspace: workspace,
			HasAI:     true,
		}
	}

	// Verificar se estamos rodando dentro do Windsurf
	if terminal := os.Getenv("TERM_PROGRAM"); strings.Contains(terminal, "Windsurf") {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDEWindsurf,
			Name:      "Windsurf",
			Workspace: workspace,
			HasAI:     true,
		}
	}

	return IDEInfo{Type: IDEUnknown}
}

func detectVSCode() IDEInfo {
	// Verificar variáveis de ambiente do VS Code
	if vscodePath := os.Getenv("VSCODE_PATH"); vscodePath != "" {
		return IDEInfo{
			Type:      IDEVSCode,
			Name:      "VS Code",
			Workspace: vscodePath,
			HasAI:     false, // VS Code precisa de extensões
		}
	}

	// Verificar arquivos de configuração do VS Code
	if _, err := os.Stat(".vscode"); err == nil {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDEVSCode,
			Name:      "VS Code",
			Workspace: workspace,
			HasAI:     false,
		}
	}

	// Verificar se estamos rodando dentro do VS Code
	if terminal := os.Getenv("TERM_PROGRAM"); strings.Contains(terminal, "vscode") {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDEVSCode,
			Name:      "VS Code",
			Workspace: workspace,
			HasAI:     false,
		}
	}

	return IDEInfo{Type: IDEUnknown}
}

func detectCursor() IDEInfo {
	// Verificar variáveis de ambiente do Cursor
	if cursorPath := os.Getenv("CURSOR_PATH"); cursorPath != "" {
		return IDEInfo{
			Type:      IDECursor,
			Name:      "Cursor",
			Workspace: cursorPath,
			HasAI:     true,
		}
	}

	// Verificar arquivos de configuração do Cursor
	if _, err := os.Stat(".cursor"); err == nil {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDECursor,
			Name:      "Cursor",
			Workspace: workspace,
			HasAI:     true,
		}
	}

	// Verificar se estamos rodando dentro do Cursor
	if terminal := os.Getenv("TERM_PROGRAM"); strings.Contains(terminal, "cursor") {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDECursor,
			Name:      "Cursor",
			Workspace: workspace,
			HasAI:     true,
		}
	}

	return IDEInfo{Type: IDEUnknown}
}

func detectIntelliJ() IDEInfo {
	// Verificar variáveis de ambiente do IntelliJ
	if intellijPath := os.Getenv("IDEA_PATH"); intellijPath != "" {
		return IDEInfo{
			Type:      IDEIntelliJ,
			Name:      "IntelliJ IDEA",
			Workspace: intellijPath,
			HasAI:     true, // IntelliJ Ultimate tem AI integrada
		}
	}

	// Verificar arquivos de configuração do IntelliJ
	if _, err := os.Stat(".idea"); err == nil {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDEIntelliJ,
			Name:      "IntelliJ IDEA",
			Workspace: workspace,
			HasAI:     true,
		}
	}

	// Verificar se estamos rodando dentro do IntelliJ
	if terminal := os.Getenv("TERMINAL_EMULATOR"); strings.Contains(terminal, "intellij") {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDEIntelliJ,
			Name:      "IntelliJ IDEA",
			Workspace: workspace,
			HasAI:     true,
		}
	}

	// Verificar variável específica do IntelliJ
	if intellijPath := os.Getenv("INTELLIJ_IDEA_ENVIRONMENT"); intellijPath != "" {
		workspace, _ := os.Getwd()
		return IDEInfo{
			Type:      IDEIntelliJ,
			Name:      "IntelliJ IDEA",
			Workspace: workspace,
			HasAI:     true,
		}
	}

	return IDEInfo{Type: IDEUnknown}
}

func GetIDEAIProvider(ideInfo IDEInfo) string {
	switch ideInfo.Type {
	case IDEWindsurf:
		return "windsurf"
	case IDECursor:
		return "cursor"
	case IDEVSCode:
		return "vscode"
	case IDEIntelliJ:
		return "intellij"
	default:
		return ""
	}
}
