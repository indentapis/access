package main

import (
	"os"

	"go.indent.com/access/cmd/access/cmd"
)

func main() {
	rootCmd := cmd.NewRoot()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
