package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	GlobalVerbose bool
)

var RootCmd = &cobra.Command{
	Use:           "wlog",
	Short:         "watch local log better",
	Long:          "Watch Log is a tool to help watch local logs better than ever",
	SilenceErrors: true,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if GlobalVerbose {
			log.SetLevel(log.DebugLevel)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&GlobalVerbose, "verbose", "v", false, "verbose output")
}

func Execute() error {
	return RootCmd.Execute()
}
