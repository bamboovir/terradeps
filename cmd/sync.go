package cmd

import (
	"os"
	"runtime"

	"github.com/bamboovir/terradeps/cmd/types"
	"github.com/bamboovir/terradeps/lib/cli"
	"github.com/spf13/cobra"
)

// SyncProp define
type SyncProp struct {
	cli.Meta
	*types.SyncArgs
}

// SyncArgsCollector defination
func SyncArgsCollector(cmd *cobra.Command) *types.SyncArgs {
	props := &types.SyncArgs{}

	cmd.Flags().StringVar(&props.Arch, "arch", runtime.GOARCH, "arch")
	cmd.Flags().StringVar(&props.OS, "os", runtime.GOOS, "os")
	cmd.Flags().StringVar(&props.DepsFilePath, "deps-path", "", "deps json file path")
	cmd.Flags().StringVar(&props.TargetPath, "target-path", "", "target folder path")

	cobra.MarkFlagRequired(cmd.Flags(), "deps-path")
	cobra.MarkFlagRequired(cmd.Flags(), "target-path")

	return props
}

// SyncPropCollector defination
func SyncPropCollector(props *types.SyncArgs) (*SyncProp, error) {
	prop := &SyncProp{}
	prop.Meta = cli.DefaultMeta()
	prop.SyncArgs = props
	return prop, nil
}

// NewSyncCMD definatio
func NewSyncCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "sync",
	}

	args := SyncArgsCollector(cmd)

	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		cmd.Parent().PreRun(cmd.Parent(), args)
	}

	cmd.Run = func(cmd *cobra.Command, _ []string) {
		props, err := SyncPropCollector(args)
		if err != nil {
			os.Exit(1)
		}

		err = Sync(props)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}

	return cmd
}

// Sync defination
func Sync(props *SyncProp) error {
	_ = "sync failed"

	return nil
}
