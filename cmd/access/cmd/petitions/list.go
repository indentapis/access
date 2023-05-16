package petitions

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
)

// NewListOptions returns a new ListOptions.
func NewListOptions() *ListOptions {
	return &ListOptions{}
}

// ListOptions are the options for listing Petitions.
type ListOptions struct {
	Output string
}

// NewCmdList returns a command allowing users to list Petitions.
func NewCmdList(f cliutil.Factory) *cobra.Command {
	opts := NewListOptions()
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Petitions",
		Long:  `List and filter Petitions`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			client := f.API(cmd.Context()).Petitions()

			petitions, err := client.ListPetitions(cmd.Context(), &indentv1.ListPetitionsRequest{
				SpaceName: f.Config().Space,
			})
			if err != nil {
				logger.Fatal("Failed to list Petitions", zap.Error(err))
			}

			if opts.Output == cliutil.OutputJSON {
				f.OutputJSON(petitions)
				return
			}

			s, err := cliutil.NewSelect(petitions.GetPetitions())
			if err != nil {
				logger.Fatal("Failed to create select", zap.Error(err))
			} else if result, err := s.Run(); err != nil {
				logger.Fatal("Failed to run select", zap.Error(err))
			} else {
				f.OutputJSON(result)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.Output, "output", "o", opts.Output, "Output format (can be 'json')")
	return cmd
}
