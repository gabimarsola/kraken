package extractor

import (
	"kraken/src/structure"
	"regexp"
	"strings"
)

type JavaSpringExtractor struct{}

func (e *JavaSpringExtractor) Supports(filePath string) bool {
	return strings.HasSuffix(filePath, ".java")
}

func (e *JavaSpringExtractor) Extract(filePath string, content []byte) ([]structure.Endpoint, error) {
	var endpoints []structure.Endpoint
	contentStr := string(content)

	if !strings.Contains(contentStr, "@RestController") && !strings.Contains(contentStr, "@Controller") {
		return endpoints, nil
	}

	lines := strings.Split(contentStr, "\n")

	basePath := extractSpringBasePath(lines)

	methodPatterns := []struct {
		annotation string
		method     string
	}{
		{"@GetMapping", "GET"},
		{"@PostMapping", "POST"},
		{"@PutMapping", "PUT"},
		{"@DeleteMapping", "DELETE"},
		{"@PatchMapping", "PATCH"},
	}

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		for _, pattern := range methodPatterns {
			if strings.Contains(line, pattern.annotation) {
				path := extractSpringPath(line)
				method := pattern.method

				if method == "" {
					method = extractRequestMethod(line)
					if method == "" {
						method = "GET"
					}
				}

				fullPath := ""
				if basePath != "" && path != "" {
					if strings.HasPrefix(path, "/") {
						fullPath = basePath + path
					} else {
						fullPath = basePath + "/" + path
					}
				} else if basePath != "" {
					fullPath = basePath
				} else if path != "" {
					fullPath = path
				} else {
					fullPath = "/"
				}

				fullPath = strings.ReplaceAll(fullPath, "//", "/")
				if !strings.HasPrefix(fullPath, "/") {
					fullPath = "/" + fullPath
				}

				endpoint := structure.Endpoint{
					Method:          method,
					Path:            fullPath,
					Description:     extractJavaDescription(lines, i),
					Summary:         AnalyzeMethodCode(lines, i, "java"),
					Parameters:      extractSpringParameters(lines, i, fullPath),
					RequestExamples: GenerateBasicRequestExample(method, fullPath),
				}

				if endpoint.Description == "" {
					endpoint.Description = "Endpoint detectado automaticamente (Spring Boot)"
				}

				endpoints = append(endpoints, endpoint)
				break
			}
		}
	}

	return endpoints, nil
}

func extractSpringBasePath(lines []string) string {
	requestMappingRegex := regexp.MustCompile(`@RequestMapping\s*\(\s*(?:value\s*=\s*)?["\']([^"\']+)["\']`)

	inClassDeclaration := false
	for i, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "@RestController") || strings.Contains(line, "@Controller") {
			inClassDeclaration = true
		}

		if inClassDeclaration && strings.Contains(line, "@RequestMapping") {
			matches := requestMappingRegex.FindStringSubmatch(line)
			if len(matches) > 1 {
				path := matches[1]
				if !strings.HasPrefix(path, "/") {
					path = "/" + path
				}
				return path
			}
		}

		if inClassDeclaration && strings.Contains(line, "class ") {
			break
		}

		if i > 50 {
			break
		}
	}

	return ""
}

func extractSpringPath(line string) string {
	pathRegex := regexp.MustCompile(`\(\s*(?:value\s*=\s*)?["\']([^"\']+)["\']`)
	matches := pathRegex.FindStringSubmatch(line)

	if len(matches) > 1 {
		return matches[1]
	}

	emptyParenRegex := regexp.MustCompile(`@\w+Mapping\s*\(\s*\)`)
	if emptyParenRegex.MatchString(line) {
		return ""
	}

	return ""
}

func extractRequestMethod(line string) string {
	methodRegex := regexp.MustCompile(`method\s*=\s*RequestMethod\.(\w+)`)
	matches := methodRegex.FindStringSubmatch(line)

	if len(matches) > 1 {
		return strings.ToUpper(matches[1])
	}

	return ""
}

func extractJavaDescription(lines []string, currentLine int) string {
	start := currentLine - 10
	if start < 0 {
		start = 0
	}

	for i := start; i < currentLine; i++ {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "/**") || strings.HasPrefix(line, "/*") {
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

func extractSpringParameters(lines []string, currentLine int, path string) []structure.Parameter {
	var parameters []structure.Parameter

	pathParams := ExtractBasicPathParameters(path)
	parameters = append(parameters, pathParams...)

	start := currentLine
	end := currentLine + 15
	if end > len(lines) {
		end = len(lines)
	}

	for i := start; i < end; i++ {
		line := strings.TrimSpace(lines[i])

		if strings.Contains(line, "@RequestBody") {
			bodyRegex := regexp.MustCompile(`@RequestBody\s+(\w+(?:<[^>]+>)?)\s+(\w+)`)
			if matches := bodyRegex.FindStringSubmatch(line); len(matches) > 2 {
				parameters = append(parameters, structure.Parameter{
					Name:        matches[2],
					Type:        matches[1],
					Required:    true,
					Description: "Request body",
					Location:    "body",
				})
			}
		}

		if strings.Contains(line, "@RequestParam") {
			paramRegex := regexp.MustCompile(`@RequestParam(?:\([^)]*name\s*=\s*"([^"]+)"[^)]*\))?\s+(\w+)\s+(\w+)`)
			if matches := paramRegex.FindStringSubmatch(line); len(matches) > 3 {
				paramName := matches[1]
				if paramName == "" {
					paramName = matches[3]
				}
				required := !strings.Contains(line, "required = false")
				parameters = append(parameters, structure.Parameter{
					Name:        paramName,
					Type:        matches[2],
					Required:    required,
					Description: "Query parameter",
					Location:    "query",
				})
			}
		}

		if strings.Contains(line, "@PathVariable") {
			pathVarRegex := regexp.MustCompile(`@PathVariable(?:\([^)]*name\s*=\s*"([^"]+)"[^)]*\))?\s+(\w+)\s+(\w+)`)
			if matches := pathVarRegex.FindStringSubmatch(line); len(matches) > 3 {
				paramName := matches[1]
				if paramName == "" {
					paramName = matches[3]
				}
				parameters = append(parameters, structure.Parameter{
					Name:        paramName,
					Type:        matches[2],
					Required:    true,
					Description: "Path variable",
					Location:    "path",
				})
			}
		}

		if strings.Contains(line, "public") || strings.Contains(line, "private") {
			break
		}
	}

	return parameters
}
