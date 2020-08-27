package get

import (
	"github.com/spf13/cobra"
)

// NewGetCMD definatio
func NewGetCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get",
	}

	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		cmd.Parent().PreRun(cmd.Parent(), args)
	}

	cmd.AddCommand(NewProviderCMD())
	cmd.AddCommand(NewTerraformCMD())

	return cmd
}
