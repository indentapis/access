package auth

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.indent.com/indent-go/pkg/cliutil"
)

// NewCmdView returns a command allowing the current user to be viewed.
func NewCmdView(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "View currently logged in user",
		Long:  `Allows details about the currently logged in user to be introspected.'`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()

			user := f.CurrentUser(cmd.Context())
			logger.Info("Current User", zap.Object("user", user))
		},
	}

	return cmd
}
