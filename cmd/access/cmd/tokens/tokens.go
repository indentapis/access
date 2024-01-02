// Package tokens is a command to manage tokens.
package tokens

import (
	"github.com/spf13/cobra"

	"go.indent.com/indent-go/pkg/cliutil"
)

// NewCmdTokens returns a set of commands used to manage tokens.
func NewCmdTokens(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tokens",
		Aliases: []string{"t"},
		Short:   "Manage tokens",
	}
	cmd.AddCommand(NewCmdCreate(f))
	return cmd
}
