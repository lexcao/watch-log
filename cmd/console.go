package cmd

import (
	"github.com/lexcao/watch-log/internal/render/console"
	"github.com/lexcao/watch-log/pkg/app"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Watch logs for console output",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run(console.Renderer{})
	},
}

func init() {
	RootCmd.AddCommand(consoleCmd)
}
