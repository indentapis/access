package petitions

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
)

const (
	numPetitionsList = 20
)

// NewCmdList returns a command allowing users to list Petitions.
func NewCmdList(f cliutil.Factory) *cobra.Command {
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

			for i, pet := range petitions.GetPetitions() {
				logger.Info("Petition", zap.Any("petition", pet))

				if i > numPetitionsList {
					break
				}
			}
		},
	}

	return cmd
}
