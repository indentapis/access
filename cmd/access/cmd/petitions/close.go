package petitions

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
	"go.indent.com/indent-go/pkg/common"
)

// NewCmdClose returns a command allowing users to close a Petition.
func NewCmdClose(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close [petition name]",
		Short: "Close a Petition",
		Long:  `ADVANCED USE ONLY: Administratively closes a Petition'`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			client := f.API(cmd.Context()).Petitions()

			petitionName := parsePetitionArg(logger, args)
			petition, err := client.GetPetition(cmd.Context(), &indentv1.GetPetitionRequest{
				SpaceName:    f.Config().Space,
				PetitionName: petitionName,
			})
			if err != nil {
				logger.Fatal("Failed to get Petition", zap.Error(err))
			}

			// close petition
			status := petition.GetState().GetStatus()
			if status == nil {
				logger.Fatal("State isn't set, cannot close automatically")
			}
			status.Phase = common.PetitionStateClosed

			petition, err = client.UpdatePetition(cmd.Context(), &indentv1.UpdatePetitionRequest{
				SpaceName:    f.Config().Space,
				PetitionName: petitionName,
				Petition:     petition,
			})
			if err != nil {
				logger.Fatal("Failed to update Petition as closed", zap.Error(err))
			}
			logger.Info("Petition closed", zap.String("petition", petition.GetName()))
		},
	}

	return cmd
}
