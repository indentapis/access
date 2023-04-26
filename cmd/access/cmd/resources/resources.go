// Package resources provides subcommands for managing resources.
package resources

import (
	"github.com/spf13/cobra"

	"go.indent.com/indent-go/pkg/cliutil"
)

// NewCmdResources returns a set of commands allowing management of Resources.
func NewCmdResources(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resources",
		Short: "Manage Resources within a space",
		Long: `Allows for management of Resources including activities such
such as pulling.`,
	}
	cmd.AddCommand(NewCmdCreate(f))
	cmd.AddCommand(NewCmdPull(f))
	return cmd
}
