package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.indent.com/access/cmd/access/cmd/auth"
	configcmd "go.indent.com/access/cmd/access/cmd/config"
	"go.indent.com/indent-go/pkg/cliutil"
)

// NewCmdInit returns a command that sets up access.
func NewCmdInit(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [space name]",
		Short: "Setup access",
		Long:  "Login and set up configuration",
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()

			logger.Debug("Setting space")
			if len(args) < 1 {
				logger.Error("Expected space name")
				if err := cmd.Help(); err != nil {
					logger.Fatal("Failed to print help", zap.Error(err))
				}
				return
			}
			spaceName := args[0]
			logger = logger.With(zap.String("space", spaceName))
			configcmd.NewCmdSet(f).Run(cmd, []string{"space", spaceName})

			logger.Debug("Logging in")
			auth.NewCmdLogin(f).Run(cmd, nil)
		},
	}
	return cmd
}
