package resources

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
)

// NewPullOptions returns a PullOptions with the defaults set.
func NewPullOptions() *PullOptions {
	return &PullOptions{
		PullUpdateRequest: new(indentv1.PullUpdateRequest),
	}
}

// PullOptions configures how a PullUpdate is performed.
type PullOptions struct {
	*indentv1.PullUpdateRequest
}

// NewCmdPull returns a command allowing users to pull Resources.
func NewCmdPull(f cliutil.Factory) *cobra.Command {
	opts := NewPullOptions()

	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull Resources",
		Long:  `Perform a PullUpdate and sync Resources into a space.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			client := f.API(cmd.Context()).Resources()

			opts.SpaceName = f.Config().Space
			updateResp, err := client.PullUpdate(cmd.Context(), opts.PullUpdateRequest)
			if err != nil {
				logger.Fatal("Failed to PullUpdate", zap.Error(err))
			}

			if msg := updateResp.Status.GetMessage(); msg != "" {
				logger.Fatal(msg, zap.Any("status", updateResp.Status))
			}
			logger.Info("Resources pulled successfully")
		},
	}

	flags := cmd.Flags()
	flags.StringArrayVar(&opts.Kinds, "kinds", opts.Kinds, "Kinds to be included in pull")
	flags.StringToStringVar(&opts.Flags, "flags", opts.Flags, "Flags to be set for pull")
	return cmd
}
