package cmd

import (
	"fmt"
	"os"

	"github.com/bamboovir/terradeps/cmd/get"
	"github.com/bamboovir/terradeps/cmd/store"
	"github.com/bamboovir/terradeps/cmd/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewRootCMD return root command
func NewRootCMD(args []string, state *types.State, store store.Store) *cobra.Command {
	props := &types.RootArgs{}

	cmd := &cobra.Command{
		Use:   "terradeps",
		Short: "terradeps",
	}

	cmd.PreRun = func(cmd *cobra.Command, _ []string) {
		_, err := store.SetupRootLevel(state, props)
		if err != nil {
			fmt.Fprintf(os.Stderr, "root level init failed\n")
			fmt.Fprintf(os.Stderr, "%+v", err)
			os.Exit(1)
		}
		log.Infof("CLI Args: [%#v]\n", args)
	}

	cmd.PersistentFlags().BoolVar(&props.Verbose, "verbose", false, "set verbose output")
	cmd.PersistentFlags().StringVar(&props.LogLevel, "log", "FATAL", "set log level { DEBUG, ERROR, FATAL, INFO, WARN, TRACE, PANIC }")

	cmd.AddCommand(get.NewGetCMD())
	cmd.AddCommand(NewVerifyCMD())
	cmd.AddCommand(NewSyncCMD())

	cmd.SetArgs(args)
	return cmd
}
