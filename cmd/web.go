package cmd

import "github.com/spf13/cobra"

var (
	WebPort int
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Watch logs on web ui",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO web.start
	},
}

func init() {
	webCmd.Flags().IntVarP(&WebPort, "port", "p", 7788, "Specify the port of web ui")

	RootCmd.AddCommand(webCmd)
}
