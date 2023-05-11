package petitions

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
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
		Output: cliutil.OutputJSON,
	}
}

// CreateOptions specify the Petition being created.
type CreateOptions struct {
	*indentv1.CreatePetitionRequest
	ResourceNames []string
	Interactive   bool
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

			// set petitioner and prompt for missing fields
			petitioner := f.CurrentUser(cmd.Context())
			opts.Petition.Petitioners[0] = petitioner
			if len(opts.Petition.Reason) < common.MinLenReason {
				opts.Petition.Reason = promptForReason(logger, opts)
			}

			// create petition
			petition, err := client.CreatePetition(cmd.Context(), opts.CreatePetitionRequest)
			if err != nil {
				logger.Fatal("Failed to create Petition", zap.Error(err), zap.Object("petition", petition))
			}
			logger.Info("Created Petition", zap.Object("petition", petition))

			// print output
			if opts.Output == "name" {
				fmt.Println(petition.Name)
			} else if opts.Output == cliutil.OutputJSON {
				f.OutputJSON(petition)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringArrayVar(&opts.ResourceNames, "resources", opts.ResourceNames, "names of resources being requested")
	flags.BoolVar(&opts.Interactive, "interactive", opts.Interactive, "whether to prompt for missing fields")
	flags.StringVar(&opts.Output, "output", opts.Output, "format that should be output (can be 'name' or 'json')")
	flags.StringVar(&opts.Petition.Reason, "reason", opts.Output, "reason Petition is being created")
	return cmd
}

func resolveResources(ctx context.Context, f cliutil.Factory, opts *CreateOptions) (resources []*auditv1.Resource) {
	logger := f.Logger()

	// prompt for resource if interactive`
	if len(opts.ResourceNames) == 0 {
		if !opts.Interactive {
			logger.Fatal("No resources specified and not interactive")
		}
		resource := f.SelectResource(ctx, common.ViewRequestable)
		return []*auditv1.Resource{resource}
	}

	client := f.API(ctx).Resources()
	for _, resourceName := range opts.ResourceNames {
		resource, err := client.GetResource(ctx, &indentv1.GetResourceRequest{
			SpaceName:    f.Config().Space,
			ResourceName: resourceName,
		})
		if err != nil {
			logger.Fatal("Failed to resolve resource", zap.Error(err), zap.String("resourceName", resourceName))
		}
		resources = append(resources, resource)
	}
	return
}

func promptForReason(logger *zap.Logger, opts *CreateOptions) string {
	if !opts.Interactive {
		logger.Fatal("Invalid reason specified and not interactive", zap.String("reason", opts.Petition.Reason))
	}
	prompt := &promptui.Prompt{
		Label: "Reason",
		Validate: func(s string) error {
			if len(s) < common.MinLenReason {
				return fmt.Errorf("reason must be at least %d characters", common.MinLenReason)
			}
			return nil
		},
	}

	reason, err := prompt.Run()
	if err != nil {
		logger.Fatal("Failed to prompt for reason", zap.Error(err))
	}
	return reason
}
