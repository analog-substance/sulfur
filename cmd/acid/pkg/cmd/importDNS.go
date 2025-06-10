package cmd

import (
	"github.com/spf13/cobra"
)

// importDNSCmd represents the add command
var importDNSCmd = &cobra.Command{
	Use:   "dns",
	Short: "import DNS reccords",
	Long: `For example:

	acid import dns --json dns.json
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	importCmd.AddCommand(importDNSCmd)

}
