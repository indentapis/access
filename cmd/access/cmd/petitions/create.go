package petitions

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
	"go.indent.com/indent-go/pkg/common"
)

// NewCreateOptions returns CreateOptions with the defaults set.
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{
		CreatePetitionRequest: &indentv1.CreatePetitionRequest{
			Petition: &indentv1.Petition{
				Petitioners: []*auditv1.Resource{nil},
				State: &indentv1.PetitionState{
					Status: &indentv1.PetitionStatus{
						Phase: common.PetitionStateOpen,
					},
				},
			},
		},
	}
}

// CreateOptions specify the Petition being created.
type CreateOptions struct {
	*indentv1.CreatePetitionRequest
	ResourceNames []string
	Output        string
}

// NewCmdCreate returns a command allowing users to create a Petition.
func NewCmdCreate(f cliutil.Factory) *cobra.Command {
	opts := NewCreateOptions()
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a Petition",
		Long:  `Open a Petition requesting access to the specified resource.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			client := f.API(cmd.Context()).Petitions()

			// setup petition
			opts.SpaceName = f.Config().Space
			opts.Petition.Resources = resolveResources(cmd.Context(), f, opts)
			opts.Petition.Meta = &auditv1.Meta{
				Labels: map[string]string{
					common.LabelAppConfigID: f.AppConfigName(cmd.Context()),
				},
			}

			// set petitioner
			petitioner := f.CurrentUser(cmd.Context())
			opts.Petition.Petitioners[0] = petitioner

			// create petition
			petition, err := client.CreatePetition(cmd.Context(), opts.CreatePetitionRequest)
			if err != nil {
				logger.Fatal("Failed to create Petition", zap.Error(err), zap.Object("petition", petition))
			}
			logger.Info("Created Petition", zap.Object("petition", petition))

			// print output
			if opts.Output == "name" {
				fmt.Println(petition.Name)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringArrayVar(&opts.ResourceNames, "resources", opts.ResourceNames, "names of resources being requested")
	flags.StringVar(&opts.Output, "output", opts.Output, "format that should be output")
	flags.StringVar(&opts.Petition.Reason, "reason", opts.Output, "reason Petition is being created")
	return cmd
}

func resolveResources(ctx context.Context, f cliutil.Factory, options *CreateOptions) (resources []*auditv1.Resource) {
	client := f.API(ctx).Resources()
	for _, resourceName := range options.ResourceNames {
		resource, err := client.GetResource(ctx, &indentv1.GetResourceRequest{
			SpaceName:    f.Config().Space,
			ResourceName: resourceName,
		})
		if err != nil {
			f.Logger().Fatal("Failed to resolve resource", zap.Error(err), zap.String("resourceName", resourceName))
		}
		resources = append(resources, resource)
	}
	return
}
