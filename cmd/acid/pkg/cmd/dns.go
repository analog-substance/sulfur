package cmd

import (
	"github.com/spf13/cobra"
)

// dnsCmd represents the add command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "dns commands",

	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		dnsListCmd.Run(cmd, args)
	},
}

func init() {

	RootCmd.AddCommand(dnsCmd)

}
