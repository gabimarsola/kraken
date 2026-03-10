package cli

import (
	"fmt"
	"os"
)

type CLI struct {
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage: kraken <command>")
	fmt.Println("Commands:")
	fmt.Println("  generate - Generate documentation for a project")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(0)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()
}
