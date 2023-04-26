package petitions

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	"go.indent.com/indent-go/pkg/cliutil"
	"go.indent.com/indent-go/pkg/common"
)

// NewCmdDeny returns a command allowing users to deny a Petition.
func NewCmdDeny(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deny [petition name]",
		Short: "Deny a Petition",
		Long:  `Denies a Petition'`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			claim := &auditv1.Event{
				Event: common.EventDeny,
			}

			petitionName := parsePetitionArg(logger, args)
			err := createClaim(cmd.Context(), f, petitionName, claim)
			if err != nil {
				logger.Fatal("Failed to deny petition", zap.Error(err), zap.Object("claim", claim))
			}
		},
	}

	return cmd
}
