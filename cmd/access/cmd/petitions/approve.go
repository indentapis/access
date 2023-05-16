package petitions

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	"go.indent.com/indent-go/pkg/cliutil"
	"go.indent.com/indent-go/pkg/common"
)

// NewApproveOptions returns a ApproveOptions with the defaults set.
func NewApproveOptions() *ApproveOptions {
	return &ApproveOptions{}
}

// ApproveOptions specify the details of approving a Petition.
type ApproveOptions struct {
	Duration   time.Duration
	Indefinite bool
}

// NewCmdApprove returns a command allowing users to approve a Petition.
func NewCmdApprove(f cliutil.Factory) *cobra.Command {
	opts := NewApproveOptions()
	cmd := &cobra.Command{
		Use:   "approve [petition name]",
		Short: "Approve a Petition",
		Long:  `Approve a Petition for a specified amount of time`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()

			claim := &auditv1.Event{
				Event: common.EventApprove,
				Meta: &auditv1.Meta{
					Labels: approvalDurationLabels(logger, opts),
				},
			}

			petitionName := parsePetitionArg(logger, args)
			err := createClaim(cmd.Context(), f, petitionName, claim)
			if err != nil {
				logger.Fatal("Failed to approve petition", zap.Error(err), zap.Object("claim", claim))
			}
		},
	}

	flags := cmd.Flags()
	flags.DurationVarP(&opts.Duration, "duration", "d", opts.Duration, "Go formatted duration Petition should be approved"+
		" (for example '10h' or '1h10m10s')")
	flags.BoolVar(&opts.Indefinite, "indefinite", opts.Indefinite, "approve Petition indefinitely")
	return cmd
}

func approvalDurationLabels(logger *zap.Logger, opts *ApproveOptions) map[string]string {
	if opts.Indefinite && opts.Duration != 0 {
		logger.Fatal("both indefinite and duration should not be set at the same time", zap.Duration("duration", opts.Duration))
	} else if opts.Indefinite {
		return map[string]string{common.LabelIsIndefinite: strconv.FormatBool(true)}
	}
	return map[string]string{common.LabelDuration: opts.Duration.String()}
}
