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

	if !strings.Contains(contentStr, "express") && !strings.Contains(contentStr, "router") && !strings.Contains(contentStr, "app") {
		return endpoints, nil
	}

	lines := strings.Split(contentStr, "\n")

	routePatterns := []struct {
		regex  *regexp.Regexp
		method string
	}{
		// Padrões tradicionais: router.get('/', handler)
		{regexp.MustCompile(`(?:router|app)\.get\s*\(\s*['\"]([^'\"]+)['\"]`), "GET"},
		{regexp.MustCompile(`(?:router|app)\.post\s*\(\s*['\"]([^'\"]+)['\"]`), "POST"},
		{regexp.MustCompile(`(?:router|app)\.put\s*\(\s*['\"]([^'\"]+)['\"]`), "PUT"},
		{regexp.MustCompile(`(?:router|app)\.delete\s*\(\s*['\"]([^'\"]+)['\"]`), "DELETE"},
		{regexp.MustCompile(`(?:router|app)\.patch\s*\(\s*['\"]([^'\"]+)['\"]`), "PATCH"},

		// Padrões com route(): router.route('/').get(handler)
		{regexp.MustCompile(`router\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.get`), "GET"},
		{regexp.MustCompile(`router\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.post`), "POST"},
		{regexp.MustCompile(`router\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.put`), "PUT"},
		{regexp.MustCompile(`router\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.delete`), "DELETE"},
		{regexp.MustCompile(`router\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.patch`), "PATCH"},

		// Padrões com route() em múltiplas linhas
		{regexp.MustCompile(`router\s*\.\s*route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\s*\.\s*get`), "GET"},
		{regexp.MustCompile(`router\s*\.\s*route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\s*\.\s*post`), "POST"},
		{regexp.MustCompile(`router\s*\.\s*route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\s*\.\s*put`), "PUT"},
		{regexp.MustCompile(`router\s*\.\s*route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\s*\.\s*delete`), "DELETE"},
		{regexp.MustCompile(`router\s*\.\s*route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\s*\.\s*patch`), "PATCH"},

		// Padrão app.route(): app.route('/').get(handler)
		{regexp.MustCompile(`app\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.get`), "GET"},
		{regexp.MustCompile(`app\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.post`), "POST"},
		{regexp.MustCompile(`app\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.put`), "PUT"},
		{regexp.MustCompile(`app\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.delete`), "DELETE"},
		{regexp.MustCompile(`app\.route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)\.patch`), "PATCH"},
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

	// Detectar padrões multi-linha (route().get().post())
	endpoints = append(endpoints, extractMultiLineRoutes(contentStr, lines)...)

	return endpoints, nil
}

func extractJSDescription(lines []string, currentLine int) string {
	start := currentLine - 15
	if start < 0 {
		start = 0
	}

	// Procurar por JSDoc antes da rota
	for i := start; i < currentLine; i++ {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "/**") {
			description := ""
			for j := i; j < currentLine && j < len(lines); j++ {
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

		// Procurar por comentários simples
		if strings.HasPrefix(line, "//") {
			comment := strings.TrimPrefix(line, "//")
			comment = strings.TrimSpace(comment)
			if comment != "" {
				return comment
			}
		}
	}

	// Tentar extrair do nome do controlador
	for i := currentLine; i < len(lines) && i < currentLine+5; i++ {
		line := strings.TrimSpace(lines[i])

		// Procurar por padrões como Controller.method
		if strings.Contains(line, "Controller") && strings.Contains(line, ".") {
			// Extrair nome do método
			parts := strings.Split(line, ".")
			if len(parts) >= 2 {
				methodPart := strings.Split(parts[1], ",")[0]
				methodPart = strings.Split(methodPart, ")")[0]

				// Mapear nomes comuns para descrições
				methodDescriptions := map[string]string{
					"list":                    "Lista recursos",
					"create":                  "Cria um novo recurso",
					"update":                  "Atualiza um recurso existente",
					"delete":                  "Remove um recurso",
					"show":                    "Mostra detalhes de um recurso",
					"edit":                    "Edita um recurso",
					"store":                   "Armazena um novo recurso",
					"destroy":                 "Remove um recurso permanentemente",
					"index":                   "Lista todos os recursos",
					"checkHealth":             "Verifica saúde da aplicação",
					"checkDependenciesHealth": "Verifica saúde das dependências",
				}

				if desc, exists := methodDescriptions[methodPart]; exists {
					return desc
				}
			}
		}
	}

	return ""
}

// extractMultiLineRoutes detecta padrões como router.route('/').get().post()
func extractMultiLineRoutes(contentStr string, lines []string) []structure.Endpoint {
	var endpoints []structure.Endpoint

	// Regex para detectar router.route('/', middleware).get().post()
	routeRegex := regexp.MustCompile(`router\s*\.\s*route\s*\(\s*['\"]([^'\"]+)['\"]\s*\)`)

	// Encontrar todos os blocos route()
	matches := routeRegex.FindAllStringSubmatchIndex(contentStr, -1)

	for _, match := range matches {
		if len(match) >= 4 {
			path := contentStr[match[2]:match[3]]

			// Procurar métodos após o route()
			afterRoute := contentStr[match[1]:]

			methods := []struct {
				regex  *regexp.Regexp
				method string
			}{
				{regexp.MustCompile(`\.get\s*\(`), "GET"},
				{regexp.MustCompile(`\.post\s*\(`), "POST"},
				{regexp.MustCompile(`\.put\s*\(`), "PUT"},
				{regexp.MustCompile(`\.delete\s*\(`), "DELETE"},
				{regexp.MustCompile(`\.patch\s*\(`), "PATCH"},
			}

			for _, methodPattern := range methods {
				if methodPattern.regex.MatchString(afterRoute) {
					endpoint := structure.Endpoint{
						Method:          methodPattern.method,
						Path:            path,
						Description:     "Endpoint detectado automaticamente (Express.js)",
						Summary:         "Rota Express.js com múltiplos métodos",
						Parameters:      ExtractBasicPathParameters(path),
						RequestExamples: GenerateBasicRequestExample(methodPattern.method, path),
					}
					endpoints = append(endpoints, endpoint)
				}
			}
		}
	}

	return endpoints
}
