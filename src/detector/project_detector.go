package detector

import (
	"os"
	"path/filepath"
)

type ProjectType string

const (
	ProjectTypeGo     ProjectType = "go"
	ProjectTypeNodeJS ProjectType = "nodejs"
	ProjectTypeJava   ProjectType = "java"
	ProjectTypeUnknown ProjectType = "unknown"
)

func DetectProjectType(projectPath string) ProjectType {
	if fileExists(filepath.Join(projectPath, "go.mod")) {
		return ProjectTypeGo
	}
	if fileExists(filepath.Join(projectPath, "package.json")) {
		return ProjectTypeNodeJS
	}
	if fileExists(filepath.Join(projectPath, "pom.xml")) {
		return ProjectTypeJava
	}
	return ProjectTypeUnknown
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
