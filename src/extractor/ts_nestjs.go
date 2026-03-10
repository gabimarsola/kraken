package extractor

import (
	"kraken/src/structure"
	"regexp"
	"strings"
)

type TSNestJSExtractor struct{}

func (e *TSNestJSExtractor) Supports(filePath string) bool {
	return strings.HasSuffix(filePath, ".ts") && !strings.HasSuffix(filePath, ".spec.ts")
}

func (e *TSNestJSExtractor) Extract(filePath string, content []byte) ([]structure.Endpoint, error) {
	var endpoints []structure.Endpoint
	contentStr := string(content)

	controllerRegex := regexp.MustCompile(`@Controller\s*\(\s*(?:{[^}]*path:\s*)?['"]([^'"]+)['"]`)
	controllerMatches := controllerRegex.FindStringSubmatch(contentStr)

	basePath := ""
	if len(controllerMatches) > 1 {
		basePath = controllerMatches[1]
		if !strings.HasPrefix(basePath, "/") {
			basePath = "/" + basePath
		}
	}

	methodPatterns := []struct {
		regex  *regexp.Regexp
		method string
	}{
		{regexp.MustCompile(`@Get\s*\(\s*['"]([^'"]*)['"]\s*\)`), "GET"},
		{regexp.MustCompile(`@Post\s*\(\s*['"]([^'"]*)['"]\s*\)`), "POST"},
		{regexp.MustCompile(`@Put\s*\(\s*['"]([^'"]*)['"]\s*\)`), "PUT"},
		{regexp.MustCompile(`@Delete\s*\(\s*['"]([^'"]*)['"]\s*\)`), "DELETE"},
		{regexp.MustCompile(`@Patch\s*\(\s*['"]([^'"]*)['"]\s*\)`), "PATCH"},
		{regexp.MustCompile(`@Get\s*\(\s*\)`), "GET"},
		{regexp.MustCompile(`@Post\s*\(\s*\)`), "POST"},
		{regexp.MustCompile(`@Put\s*\(\s*\)`), "PUT"},
		{regexp.MustCompile(`@Delete\s*\(\s*\)`), "DELETE"},
		{regexp.MustCompile(`@Patch\s*\(\s*\)`), "PATCH"},
	}

	lines := strings.Split(contentStr, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		for _, pattern := range methodPatterns {
			if pattern.regex.MatchString(line) {
				matches := pattern.regex.FindStringSubmatch(line)

				path := ""
				if len(matches) > 1 {
					path = matches[1]
				}

				fullPath := basePath
				if path != "" {
					if !strings.HasPrefix(path, "/") {
						fullPath = fullPath + "/" + path
					} else {
						fullPath = fullPath + path
					}
				}

				fullPath = strings.ReplaceAll(fullPath, "//", "/")

				endpoint := structure.Endpoint{
					Method:          pattern.method,
					Path:            fullPath,
					Description:     extractNestJSDescription(lines, i),
					Summary:         AnalyzeMethodCode(lines, i, "typescript"),
					Parameters:      extractNestJSParameters(lines, i, fullPath),
					RequestExamples: generateRequestExample(pattern.method, fullPath),
				}

				if endpoint.Description == "" {
					endpoint.Description = "Endpoint detectado automaticamente (NestJS)"
				}

				endpoints = append(endpoints, endpoint)
				break
			}
		}
	}

	return endpoints, nil
}

func extractNestJSDescription(lines []string, currentLine int) string {
	start := currentLine - 20
	if start < 0 {
		start = 0
	}

	apiOperationRegex := regexp.MustCompile(`@ApiOperation\s*\(\s*{\s*summary:\s*['"]([^'"]+)['"]`)
	apiTagsRegex := regexp.MustCompile(`@ApiTags\s*\(\s*['"]([^'"]+)['"]`)

	for i := start; i < currentLine; i++ {
		line := strings.TrimSpace(lines[i])

		if matches := apiOperationRegex.FindStringSubmatch(line); len(matches) > 1 {
			return matches[1]
		}

		if matches := apiTagsRegex.FindStringSubmatch(line); len(matches) > 1 {
			return matches[1]
		}
	}

	return ""
}

func extractNestJSParameters(lines []string, currentLine int, path string) []structure.Parameter {
	var parameters []structure.Parameter

	start := currentLine - 30
	if start < 0 {
		start = 0
	}

	pathParams := extractPathParameters(path)
	for _, param := range pathParams {
		parameters = append(parameters, structure.Parameter{
			Name:        param,
			Type:        "string",
			Required:    true,
			Description: "Path parameter",
			Location:    "path",
		})
	}

	for i := start; i < currentLine+10 && i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if strings.Contains(line, "@Query()") {
			queryRegex := regexp.MustCompile(`@Query\(\)\s+(\w+):\s*([\w<>\[\]]+)`)
			if matches := queryRegex.FindStringSubmatch(line); len(matches) > 2 {
				parameters = append(parameters, structure.Parameter{
					Name:        matches[1],
					Type:        matches[2],
					Required:    false,
					Description: "Query parameter",
					Location:    "query",
				})
			}
		}

		if strings.Contains(line, "@Body()") {
			bodyRegex := regexp.MustCompile(`@Body\(\)\s+(\w+):\s*([\w<>\[\]]+)`)
			if matches := bodyRegex.FindStringSubmatch(line); len(matches) > 2 {
				parameters = append(parameters, structure.Parameter{
					Name:        matches[1],
					Type:        matches[2],
					Required:    true,
					Description: "Request body",
					Location:    "body",
				})
			}
		}

		if strings.Contains(line, "@Param()") {
			paramRegex := regexp.MustCompile(`@Param\(\)\s+(\w+):\s*([\w<>\[\]]+)`)
			if matches := paramRegex.FindStringSubmatch(line); len(matches) > 2 {
				parameters = append(parameters, structure.Parameter{
					Name:        matches[1],
					Type:        matches[2],
					Required:    true,
					Description: "Path parameter",
					Location:    "path",
				})
			}
		}
	}

	return parameters
}

func extractPathParameters(path string) []string {
	var params []string
	paramRegex := regexp.MustCompile(`:([\w]+)`)
	matches := paramRegex.FindAllStringSubmatch(path, -1)
	for _, match := range matches {
		if len(match) > 1 {
			params = append(params, match[1])
		}
	}
	return params
}

func generateRequestExample(method string, path string) []structure.RequestExample {
	var examples []structure.RequestExample

	curlExample := "curl -X " + method + " 'http://localhost:3000" + path + "'"
	if method == "POST" || method == "PUT" || method == "PATCH" {
		curlExample += " \\\n  -H 'Content-Type: application/json' \\\n  -d '{\"key\": \"value\"}'"
	}

	examples = append(examples, structure.RequestExample{
		Language: "curl",
		Code:     curlExample,
	})

	tsExample := "const response = await fetch('http://localhost:3000" + path + "', {\n  method: '" + method + "'"
	if method == "POST" || method == "PUT" || method == "PATCH" {
		tsExample += ",\n  headers: { 'Content-Type': 'application/json' },\n  body: JSON.stringify({ key: 'value' })"
	}
	tsExample += "\n});"

	examples = append(examples, structure.RequestExample{
		Language: "typescript",
		Code:     tsExample,
	})

	return examples
}
