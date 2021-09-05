package cmd

import (
	"github.com/lexcao/watch-log/internal/console/renderer"
	"github.com/lexcao/watch-log/pkg/app"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Watch logs for console output",
	Run: func(cmd *cobra.Command, args []string) {
		app.New(app.Renderer(renderer.ConsoleRenderer{})).Run()
	},
}

func init() {
	RootCmd.AddCommand(consoleCmd)
}
