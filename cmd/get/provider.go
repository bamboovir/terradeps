package get

import (
	"github.com/spf13/cobra"
)

// NewProviderCMD definatio
func NewProviderCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider",
		Short: "provider",
	}

	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		cmd.Parent().PreRun(cmd.Parent(), args)
	}

	cmd.Run = func(cmd *cobra.Command, _ []string) {
	}

	return cmd
}
