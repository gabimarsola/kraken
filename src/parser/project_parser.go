package parser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"kraken/src/detector"
	"kraken/src/structure"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

type PackageJSON struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type PomXML struct {
	XMLName     xml.Name `xml:"project"`
	GroupID     string   `xml:"groupId"`
	ArtifactID  string   `xml:"artifactId"`
	Version     string   `xml:"version"`
	Description string   `xml:"description"`
}

func ParseProject(projectPath string, projectType detector.ProjectType) (*structure.ProjectInfo, error) {
	switch projectType {
	case detector.ProjectTypeGo:
		return parseGoProject(projectPath)
	case detector.ProjectTypeNodeJS:
		return parseNodeJSProject(projectPath)
	case detector.ProjectTypeJava:
		return parseJavaProject(projectPath)
	default:
		return nil, fmt.Errorf("tipo de projeto não suportado")
	}
}

func parseGoProject(projectPath string) (*structure.ProjectInfo, error) {
	data, err := os.ReadFile(filepath.Join(projectPath, "go.mod"))
	if err != nil {
		return nil, fmt.Errorf("erro ao ler go.mod: %v", err)
	}

	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear go.mod: %v", err)
	}

	return &structure.ProjectInfo{
		Name:        f.Module.Mod.Path,
		Version:     f.Go.Version,
		Description: "API desenvolvida com Go",
		ProjectType: "Go",
	}, nil
}

func parseNodeJSProject(projectPath string) (*structure.ProjectInfo, error) {
	data, err := os.ReadFile(filepath.Join(projectPath, "package.json"))
	if err != nil {
		return nil, fmt.Errorf("erro ao ler package.json: %v", err)
	}

	var pkg PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("erro ao parsear package.json: %v", err)
	}

	description := pkg.Description
	if description == "" {
		description = "API desenvolvida com Node.js"
	}

	return &structure.ProjectInfo{
		Name:        pkg.Name,
		Version:     pkg.Version,
		Description: description,
		ProjectType: "Node.js",
	}, nil
}

func parseJavaProject(projectPath string) (*structure.ProjectInfo, error) {
	data, err := os.ReadFile(filepath.Join(projectPath, "pom.xml"))
	if err != nil {
		return nil, fmt.Errorf("erro ao ler pom.xml: %v", err)
	}

	var pom PomXML
	if err := xml.Unmarshal(data, &pom); err != nil {
		return nil, fmt.Errorf("erro ao parsear pom.xml: %v", err)
	}

	name := pom.ArtifactID
	if pom.GroupID != "" {
		name = pom.GroupID + ":" + pom.ArtifactID
	}

	description := pom.Description
	if description == "" {
		description = "API desenvolvida com Java/Maven"
	}

	return &structure.ProjectInfo{
		Name:        name,
		Version:     pom.Version,
		Description: description,
		ProjectType: "Java/Maven",
	}, nil
}
