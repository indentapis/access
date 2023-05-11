package auth

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.indent.com/indent-go/pkg/cliutil"
	"go.indent.com/indent-go/pkg/oauthutil"
)

// NewCmdLogin returns a command allowing users to login.
func NewCmdLogin(f cliutil.Factory) *cobra.Command {
	opts := oauthutil.NewLoginOptions()

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Log into an Indent account",
		Long:  `Opens a browser to authenticate access with an Indent account'`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			opts.OAuth = f.Config().Environment.OAuth

			store := f.Store()
			if err := store.Login(opts); err != nil {
				logger.Fatal("Failed to login", zap.Error(err))
			} else if err = store.UpdateUserInfo(); err != nil {
				logger.Fatal("Failed to get userinfo", zap.Error(err))
			}
			logger.Info("Login successful")
		},
	}

	// TODO: enable oauth config flags
	return cmd
}
