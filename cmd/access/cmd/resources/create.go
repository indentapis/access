package resources

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	auditv1 "go.indent.com/indent-go/api/indent/audit/v1"
	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
	"go.indent.com/indent-go/pkg/common"
)

const (
	// create requires a kind and an id
	createRequiredArgs = 2
)

// NewCreateOptions returns a CreateOptions with the defaults set.
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{
		CreateResourceRequest: &indentv1.CreateResourceRequest{
			Resource: new(auditv1.Resource),
		},
	}
}

// CreateOptions configures how a CreateUpdate is performed.
type CreateOptions struct {
	*indentv1.CreateResourceRequest
	Output string
}

// NewCmdCreate returns a command allowing users to create Resources.
func NewCmdCreate(f cliutil.Factory) *cobra.Command {
	opts := NewCreateOptions()

	cmd := &cobra.Command{
		Use:   "create kind id",
		Short: "Create Resource",
		Long:  `Create a Resource in a space.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			client := f.API(cmd.Context()).Resources()

			if len(args) != createRequiredArgs {
				logger.Fatal("A kind and ID must be specified")
			}
			opts.Resource.Kind, opts.Resource.Id = args[0], args[1]

			opts.SpaceName = f.Config().Space
			resource, err := client.CreateResource(cmd.Context(), opts.CreateResourceRequest)
			if err != nil {
				logger.Fatal("Failed to Create Resource", zap.Error(err))
			}

			logger.Info("Created Resource", zap.Object("resource", resource))

			// print output
			if opts.Output == "name" {
				fmt.Println(resource.GetLabels()[common.LabelBlockName])
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.Resource.DisplayName, "displayName", opts.Resource.DisplayName, "Display name of Resource")
	flags.StringVar(&opts.Output, "output", opts.Output, "format that should be output")
	return cmd
}
