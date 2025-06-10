package cmd

import (
	"github.com/spf13/cobra"
)

// orgRootDomainsCmd represents the add command
var orgRootDomainsCmd = &cobra.Command{
	Use:   "root-domains",
	Short: "root domains",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	orgCmd.AddCommand(orgRootDomainsCmd)

}
