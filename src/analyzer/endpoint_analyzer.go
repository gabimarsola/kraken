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
				strings.Contains(path, ".git") {
				return filepath.SkipDir
			}
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		endpoints := registry.ExtractFromFile(path, content)

		for _, endpoint := range endpoints {
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
