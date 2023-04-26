// Package petitions offer subcommands for performing actions with Petitions.
package petitions

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
)

// NewCmdPetitions returns a set of commands allowing users manage Petitions.
func NewCmdPetitions(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "petitions",
		Short: "Open new Petitions and perform operations on them",
		Long: `Allows for the creation of Petitions and performing activities such
such as denying, approving, and revoking.`,
	}
	cmd.AddCommand(NewCmdList(f))
	cmd.AddCommand(NewCmdApprove(f))
	cmd.AddCommand(NewCmdClose(f))
	cmd.AddCommand(NewCmdCreate(f))
	cmd.AddCommand(NewCmdDeny(f))
	cmd.AddCommand(NewCmdRevoke(f))

	return cmd
}

func parsePetitionArg(logger *zap.Logger, args []string) (petition string) {
	switch len(args) {
	case 0:
		logger.Fatal("Petition name must be specified")
	case 1:
	default:
		logger.Fatal("only Petition name should be specified", zap.Strings("args", args))
	}
	return args[0]
}

func createClaim(ctx context.Context, f cliutil.Factory, petitionName string, claim *auditv1.Event) error {
	logger := f.Logger()
	client := f.API(ctx).Petitions()

	// set resources from petition
	petition, err := client.GetPetition(ctx, &indentv1.GetPetitionRequest{
		SpaceName:    f.Config().Space,
		PetitionName: petitionName,
	})
	if err != nil {
		logger.Fatal("Failed to get petition", zap.Error(err), zap.String("petitionName", petitionName))
	}
	claim.Resources = petition.GetResources()

	_, err = client.CreateClaim(ctx, &indentv1.CreatePetitionClaimRequest{
		SpaceName:    f.Config().Space,
		PetitionName: petitionName,
		Claim:        claim,
	})
	return err
}
