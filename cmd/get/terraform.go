package get

import (
	"github.com/spf13/cobra"
)

// NewTerraformCMD definatio
func NewTerraformCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "terraform",
		Short: "terraform",
	}

	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		cmd.Parent().PreRun(cmd.Parent(), args)
	}

	cmd.Run = func(cmd *cobra.Command, _ []string) {
	}

	return cmd
}
