package extractor

import (
	"kraken/src/structure"
	"regexp"
	"strings"
)

type JSExpressExtractor struct{}

func (e *JSExpressExtractor) Supports(filePath string) bool {
	return strings.HasSuffix(filePath, ".js") && !strings.HasSuffix(filePath, ".spec.js") && !strings.HasSuffix(filePath, ".test.js")
}

func (e *JSExpressExtractor) Extract(filePath string, content []byte) ([]structure.Endpoint, error) {
	var endpoints []structure.Endpoint
	contentStr := string(content)

	if !strings.Contains(contentStr, "express") && !strings.Contains(contentStr, "router") {
		return endpoints, nil
	}

	lines := strings.Split(contentStr, "\n")

	routePatterns := []struct {
		regex  *regexp.Regexp
		method string
	}{
		{regexp.MustCompile(`(?:router|app)\.get\s*\(\s*['\"]([^'\"]+)['\"]`), "GET"},
		{regexp.MustCompile(`(?:router|app)\.post\s*\(\s*['\"]([^'\"]+)['\"]`), "POST"},
		{regexp.MustCompile(`(?:router|app)\.put\s*\(\s*['\"]([^'\"]+)['\"]`), "PUT"},
		{regexp.MustCompile(`(?:router|app)\.delete\s*\(\s*['\"]([^'\"]+)['\"]`), "DELETE"},
		{regexp.MustCompile(`(?:router|app)\.patch\s*\(\s*['\"]([^'\"]+)['\"]`), "PATCH"},
	}

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		for _, pattern := range routePatterns {
			if matches := pattern.regex.FindStringSubmatch(line); len(matches) > 1 {
				path := matches[1]

				endpoint := structure.Endpoint{
					Method:          pattern.method,
					Path:            path,
					Description:     extractJSDescription(lines, i),
					Summary:         AnalyzeMethodCode(lines, i, "javascript"),
					Parameters:      ExtractBasicPathParameters(path),
					RequestExamples: GenerateBasicRequestExample(pattern.method, path),
				}

				if endpoint.Description == "" {
					endpoint.Description = "Endpoint detectado automaticamente (Express.js)"
				}

				endpoints = append(endpoints, endpoint)
				break
			}
		}
	}

	return endpoints, nil
}

func extractJSDescription(lines []string, currentLine int) string {
	start := currentLine - 10
	if start < 0 {
		start = 0
	}

	for i := start; i < currentLine; i++ {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "/**") {
			description := ""
			for j := i; j < currentLine; j++ {
				commentLine := strings.TrimSpace(lines[j])
				commentLine = strings.TrimPrefix(commentLine, "/**")
				commentLine = strings.TrimPrefix(commentLine, "/*")
				commentLine = strings.TrimPrefix(commentLine, "*")
				commentLine = strings.TrimSuffix(commentLine, "*/")
				commentLine = strings.TrimSpace(commentLine)

				if commentLine != "" && !strings.HasPrefix(commentLine, "@") {
					if description != "" {
						description += " "
					}
					description += commentLine
				}
			}
			if description != "" {
				return description
			}
		}

		if strings.HasPrefix(line, "//") {
			comment := strings.TrimPrefix(line, "//")
			comment = strings.TrimSpace(comment)
			if comment != "" {
				return comment
			}
		}
	}

	return ""
}
