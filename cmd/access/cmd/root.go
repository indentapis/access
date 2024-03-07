// Package cmd is the root of access.
package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.indent.com/access/cmd/access/cmd/auth"
	configcmd "go.indent.com/access/cmd/access/cmd/config"
	"go.indent.com/access/cmd/access/cmd/petitions"
	"go.indent.com/access/cmd/access/cmd/resources"
	"go.indent.com/access/cmd/access/cmd/tokens"
	"go.indent.com/indent-go/pkg/cliutil"
)

// NewRoot returns a new root command.
func NewRoot(logger *zap.Logger) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "access",
		Short: "Manages the Indent platform",
	}
	f, config := cliutil.New(logger, rootCmd)
	rootCmd.AddCommand(NewCmdInit(f))
	rootCmd.AddCommand(NewCmdNew(f))
	rootCmd.AddCommand(auth.NewCmdAuth(f))
	rootCmd.AddCommand(configcmd.NewCmdConfig(f))
	rootCmd.AddCommand(petitions.NewCmdPetitions(f))
	rootCmd.AddCommand(resources.NewCmdResources(f))
	rootCmd.AddCommand(tokens.NewCmdTokens(f))

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&config.Space, "space", "s", config.Space, "Space to perform operations in")
	flags.BoolVar(&config.Staging, "staging", config.Staging, "Use staging environment for request")
	flags.BoolVarP(&config.Verbose, "verbose", "v", config.Verbose, "Include debug messages and additional context in logs")
	flags.BoolVar(&config.Headless, "headless", config.Headless, "Run in headless mode (no browser login prompt)")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		f.Setup()
	}
	return rootCmd
}
