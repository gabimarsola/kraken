package extractor

import (
	"kraken/src/structure"
	"regexp"
	"strings"
)

func GenerateBasicRequestExample(method string, path string) []structure.RequestExample {
	var examples []structure.RequestExample

	curlExample := "curl -X " + method + " 'http://localhost:3000" + path + "'"
	if method == "POST" || method == "PUT" || method == "PATCH" {
		curlExample += " \\\n  -H 'Content-Type: application/json' \\\n  -d '{\"key\": \"value\"}'"
	}

	examples = append(examples, structure.RequestExample{
		Language: "curl",
		Code:     curlExample,
	})

	return examples
}

func ExtractBasicPathParameters(path string) []structure.Parameter {
	var parameters []structure.Parameter
	
	paramRegex := regexp.MustCompile(`[:{]([a-zA-Z_][a-zA-Z0-9_]*)[}]?`)
	matches := paramRegex.FindAllStringSubmatch(path, -1)
	
	for _, match := range matches {
		if len(match) > 1 {
			paramName := match[1]
			if !strings.HasPrefix(paramName, ":") {
				parameters = append(parameters, structure.Parameter{
					Name:        paramName,
					Type:        "string",
					Required:    true,
					Description: "Path parameter",
					Location:    "path",
				})
			}
		}
	}
	
	return parameters
}
