package cmd

import (
	"github.com/lexcao/watch-log/internal/web/controller"
	"github.com/lexcao/watch-log/internal/web/renderer"
	"github.com/lexcao/watch-log/pkg/app"
	"github.com/spf13/cobra"
)

var (
	WebPort int
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Watch logs on web ui",
	Run: func(cmd *cobra.Command, args []string) {
		app.New(
			app.Controller(controller.WebController{}),
			app.Renderer(renderer.WebRenderer{}),
		).Run()
	},
}

func init() {
	webCmd.Flags().IntVarP(&WebPort, "port", "p", 7788, "Specify the port of web ui")

	RootCmd.AddCommand(webCmd)
}
