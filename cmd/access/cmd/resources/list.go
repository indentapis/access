package resources

import (
	"github.com/spf13/cobra"

	"go.indent.com/indent-go/pkg/cliutil"
)

// NewListOptions returns a new ListOptions.
func NewListOptions() *ListOptions {
	return &ListOptions{}
}

// ListOptions are the options for listing Resources.
type ListOptions struct {
	View   string
	Output string
}

// NewCmdList returns a command allowing users to list Resources.
func NewCmdList(f cliutil.Factory) *cobra.Command {
	opts := NewListOptions()
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Resources",
		Long:  `List and filter Resources`,
		Run: func(cmd *cobra.Command, args []string) {
			if opts.Output == "json" {
				resources := f.Resources(cmd.Context(), opts.View)
				f.OutputJSON(resources)
				return
			}

			// show prompt to select resource
			resource := f.SelectResource(cmd.Context(), "")
			f.OutputJSON(resource)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.View, "view", opts.View, "View to show")
	flags.StringVarP(&opts.Output, "output", "o", opts.Output, "Output format (can be 'json')")
	return cmd
}
