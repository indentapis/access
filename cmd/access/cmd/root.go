// Package cmd is the root of access.
package cmd

import (
	"github.com/spf13/cobra"

	"go.indent.com/access/cmd/access/cmd/auth"
	configcmd "go.indent.com/access/cmd/access/cmd/config"
	"go.indent.com/access/cmd/access/cmd/petitions"
	"go.indent.com/access/cmd/access/cmd/resources"
	"go.indent.com/indent-go/pkg/cliutil"
)

// NewRoot returns a new root command.
func NewRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "access",
		Short: "Manages the Indent platform",
	}
	f, config := cliutil.New(rootCmd)
	rootCmd.AddCommand(NewCmdInit(f))
	rootCmd.AddCommand(NewCmdNew(f))
	rootCmd.AddCommand(auth.NewCmdAuth(f))
	rootCmd.AddCommand(configcmd.NewCmdConfig(f))
	rootCmd.AddCommand(petitions.NewCmdPetitions(f))
	rootCmd.AddCommand(resources.NewCmdResources(f))

	flags := rootCmd.PersistentFlags()
	flags.StringVar(&config.Space, "space", config.Space, "Space to perform operations in")
	flags.BoolVar(&config.Staging, "staging", config.Staging, "Use staging environment for request")
	flags.BoolVar(&config.Headless, "headless", config.Headless, "Run in headless mode (no browser login prompt)")
	f.Setup()
	return rootCmd
}
