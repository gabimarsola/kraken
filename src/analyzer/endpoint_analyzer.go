package analyzer

import (
	"kraken/src/extractor"
	"kraken/src/structure"
	"os"
	"path/filepath"
	"strings"
)

func AnalyzeEndpoints(projectPath string) ([]structure.Endpoint, error) {
	registry := extractor.NewExtractorRegistry()

	registry.Register(&extractor.GoRouterExtractor{})
	registry.Register(&extractor.TSNestJSExtractor{})
	registry.Register(&extractor.JavaSpringExtractor{})
	registry.Register(&extractor.JSExpressExtractor{})

	var allEndpoints []structure.Endpoint
	seen := make(map[string]bool)

	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if strings.Contains(path, "node_modules") ||
				strings.Contains(path, "vendor") ||
				strings.Contains(path, ".git") ||
				strings.Contains(path, "coverage") ||
				strings.Contains(path, "dist") ||
				strings.Contains(path, "build") {
				return filepath.SkipDir
			}
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		endpoints := registry.ExtractFromFile(path, content)

		// Construir caminho completo baseado na estrutura de diretórios
		for _, endpoint := range endpoints {
			fullPath := constructFullPath(path, projectPath, endpoint.Path)

			endpoint.Path = fullPath

			key := endpoint.Method + " " + endpoint.Path
			if !seen[key] {
				seen[key] = true
				allEndpoints = append(allEndpoints, endpoint)
			}
		}

		return nil
	})

	return allEndpoints, err
}

// constructFullPath constrói o caminho completo da rota baseado na estrutura de diretórios
func constructFullPath(filePath, projectPath, routePath string) string {
	// Remover o caminho do projeto do caminho do arquivo
	relativePath := strings.TrimPrefix(filePath, projectPath)
	relativePath = strings.TrimPrefix(relativePath, "/")

	// Se estiver em src/routes/, extrair o nome do módulo
	if strings.Contains(relativePath, "src/routes/") {
		routeParts := strings.Split(relativePath, "src/routes/")
		if len(routeParts) > 1 {
			modulePath := strings.Split(routeParts[1], "/")[0]

			// Se não for o arquivo index.js, usar o nome do diretório
			if !strings.Contains(routeParts[1], "index.js") {
				// Extrair o nome do diretório pai
				parts := strings.Split(routeParts[1], "/")
				if len(parts) > 1 {
					modulePath = parts[0]
				}
			}

			// Construir caminho completo
			if modulePath != "" && routePath != "/" {
				return "/" + modulePath + routePath
			} else if modulePath != "" {
				return "/" + modulePath
			}
		}
	}

	return routePath
}
