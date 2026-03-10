package generator

import (
	"fmt"
	"kraken/src/structure"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GroupEndpointsByRoute(endpoints []structure.Endpoint) map[string][]structure.Endpoint {
	grouped := make(map[string][]structure.Endpoint)

	for _, endpoint := range endpoints {
		routeBase := extractRouteBase(endpoint.Path)
		grouped[routeBase] = append(grouped[routeBase], endpoint)
	}

	return grouped
}

func extractRouteBase(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) == 0 {
		return "root"
	}

	baseRoute := parts[0]

	if baseRoute == "" {
		return "root"
	}

	baseRoute = strings.ReplaceAll(baseRoute, ":", "")
	baseRoute = strings.ReplaceAll(baseRoute, "{", "")
	baseRoute = strings.ReplaceAll(baseRoute, "}", "")

	return strings.ToUpper(baseRoute)
}

func GenerateRouteDocumentation(routeName string, endpoints []structure.Endpoint, projectInfo *structure.ProjectInfo, outputDir string) error {
	fileName := fmt.Sprintf("%s-DOC.md", routeName)
	filePath := filepath.Join(outputDir, fileName)

	tmpl, err := template.New("route").Parse(structure.RouteTemplate)
	if err != nil {
		return fmt.Errorf("erro ao criar template: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo %s: %v", fileName, err)
	}
	defer file.Close()

	data := struct {
		ProjectName string
		ProjectType string
		Version     string
		RouteName   string
		Endpoints   []structure.Endpoint
	}{
		ProjectName: projectInfo.Name,
		ProjectType: projectInfo.ProjectType,
		Version:     projectInfo.Version,
		RouteName:   routeName,
		Endpoints:   endpoints,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("erro ao executar template: %v", err)
	}

	return nil
}

func GenerateAllRouteDocumentation(projectInfo *structure.ProjectInfo, outputDir string) ([]string, error) {
	docsDir := filepath.Join(outputDir, "docs", "kraken")

	err := os.MkdirAll(docsDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar diretório docs/kraken: %v", err)
	}

	grouped := GroupEndpointsByRoute(projectInfo.Endpoints)
	var generatedFiles []string

	for routeName, endpoints := range grouped {
		err := GenerateRouteDocumentation(routeName, endpoints, projectInfo, docsDir)
		if err != nil {
			return generatedFiles, err
		}
		fileName := fmt.Sprintf("%s-DOC.md", routeName)
		generatedFiles = append(generatedFiles, filepath.Join("docs/kraken", fileName))
	}

	return generatedFiles, nil
}
