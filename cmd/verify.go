package cmd

import (
	"github.com/spf13/cobra"
)

// NewVerifyCMD defination
func NewVerifyCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "verify",
	}

	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		cmd.Parent().PreRun(cmd.Parent(), args)
	}

	cmd.Run = func(cmd *cobra.Command, _ []string) {
	}

	return cmd
}
