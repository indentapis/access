package main

import (
	"os"

	"go.indent.com/access/cmd/access/cmd"
)

func main() {
	logger := newLogger()
	rootCmd := cmd.NewRoot(logger)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
