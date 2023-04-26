package petitions

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	"go.indent.com/indent-go/pkg/cliutil"
	"go.indent.com/indent-go/pkg/common"
)

// NewCmdRevoke returns a command allowing users to revoke a Petition.
func NewCmdRevoke(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke",
		Short: "Revoke a Petition and related access",
		Long:  `Revoke a Petition and any related resources`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			claim := &auditv1.Event{
				Event: common.EventRevoke,
			}

			petitionName := parsePetitionArg(logger, args)
			err := createClaim(cmd.Context(), f, petitionName, claim)
			if err != nil {
				logger.Fatal("Failed to revoke petition", zap.Error(err), zap.Object("claim", claim))
			}
		},
	}

	return cmd
}
