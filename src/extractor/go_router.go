package extractor

import (
	"go/ast"
	"go/parser"
	"go/token"
	"kraken/src/structure"
	"regexp"
	"strings"
)

type GoRouterExtractor struct{}

func (e *GoRouterExtractor) Supports(filePath string) bool {
	return strings.HasSuffix(filePath, ".go")
}

func (e *GoRouterExtractor) Extract(filePath string, content []byte) ([]structure.Endpoint, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, content, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var endpoints []structure.Endpoint
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	patterns := []struct {
		regex  *regexp.Regexp
		method string
	}{
		{regexp.MustCompile(`router\.HandleFunc\("([^"]+)",\s*\w+\)\.Methods\("(\w+)"\)`), ""},
		{regexp.MustCompile(`r\.Get\("([^"]+)",`), "GET"},
		{regexp.MustCompile(`r\.Post\("([^"]+)",`), "POST"},
		{regexp.MustCompile(`r\.Put\("([^"]+)",`), "PUT"},
		{regexp.MustCompile(`r\.Delete\("([^"]+)",`), "DELETE"},
		{regexp.MustCompile(`r\.Patch\("([^"]+)",`), "PATCH"},
		{regexp.MustCompile(`router\.GET\("([^"]+)",`), "GET"},
		{regexp.MustCompile(`router\.POST\("([^"]+)",`), "POST"},
		{regexp.MustCompile(`router\.PUT\("([^"]+)",`), "PUT"},
		{regexp.MustCompile(`router\.DELETE\("([^"]+)",`), "DELETE"},
		{regexp.MustCompile(`router\.PATCH\("([^"]+)",`), "PATCH"},
		{regexp.MustCompile(`e\.GET\("([^"]+)",`), "GET"},
		{regexp.MustCompile(`e\.POST\("([^"]+)",`), "POST"},
		{regexp.MustCompile(`e\.PUT\("([^"]+)",`), "PUT"},
		{regexp.MustCompile(`e\.DELETE\("([^"]+)",`), "DELETE"},
		{regexp.MustCompile(`e\.PATCH\("([^"]+)",`), "PATCH"},
	}

	for _, pattern := range patterns {
		matches := pattern.regex.FindAllStringSubmatchIndex(contentStr, -1)
		for _, match := range matches {
			if len(match) >= 4 {
				path := contentStr[match[2]:match[3]]
				method := pattern.method
				if method == "" && len(match) >= 6 {
					method = contentStr[match[4]:match[5]]
				}

				// Encontrar linha onde ocorre o match
				lineNum := strings.Count(contentStr[:match[0]], "\n")

				endpoint := structure.Endpoint{
					Method:          method,
					Path:            path,
					Description:     "Endpoint detectado automaticamente",
					Summary:         AnalyzeMethodCode(lines, lineNum, "go"),
					Parameters:      ExtractBasicPathParameters(path),
					RequestExamples: GenerateBasicRequestExample(method, path),
				}
				endpoints = append(endpoints, endpoint)
			}
		}
	}

	ast.Inspect(node, func(n ast.Node) bool {
		return true
	})

	return endpoints, nil
}
