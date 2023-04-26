// Package auth is a command that performs operations related to Indent authentication.
package auth

import (
	"github.com/spf13/cobra"

	"go.indent.com/indent-go/pkg/cliutil"
)

// NewCmdAuth returns a set of commands used to configure authorization.
func NewCmdAuth(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Perform operations related to Indent authentication",
	}
	cmd.AddCommand(NewCmdLogin(f))
	cmd.AddCommand(NewCmdView(f))
	return cmd
}
