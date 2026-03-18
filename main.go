package main

import (
	"fmt"
	"kraken/cli"
	"os"
)

const (
	errorMessage = "❌ Erro: %v\n"
)

func main() {
	projectPath := "."
	if len(os.Args) > 1 {
		projectPath = os.Args[1]
	}

	// Criar e executar CLI
	krakenCLI := cli.NewCLI(projectPath)
	err := krakenCLI.Run()
	if err != nil {
		fmt.Printf(errorMessage, err)
		os.Exit(1)
	}
}
