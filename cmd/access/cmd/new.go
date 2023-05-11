package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	petitionscmd "go.indent.com/access/cmd/access/cmd/petitions"
	"go.indent.com/indent-go/pkg/cliutil"
)

// NewCmdNew returns a command that interactively creates a Petition.
func NewCmdNew(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Request access",
		Long:  "Interactively create a Petition to request access to a Resource",
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			createCmd := petitionscmd.NewCmdCreate(f)
			createCmd.SetContext(cmd.Context())
			if err := createCmd.Flags().Set("interactive", "true"); err != nil {
				logger.Fatal("failed to set interactive flag", zap.Error(err))
			}
			createCmd.Run(createCmd, nil)
		},
	}
	return cmd
}
