// Package config is a command that allows changes to access configuration.
package config

import (
	"github.com/spf13/cobra"

	"go.indent.com/indent-go/pkg/cliutil"
)

// NewCmdConfig returns a set of commands used to modify configuration.
func NewCmdConfig(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"cf"},
		Short:   "Make changes to access configuration",
	}
	cmd.AddCommand(NewCmdSet(f))
	return cmd
}
