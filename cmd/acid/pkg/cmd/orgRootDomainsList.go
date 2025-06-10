package cmd

import (
	"github.com/spf13/cobra"
)

// orgRootDomainsListCmd represents the add command
var orgRootDomainsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list root domains",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	orgRootDomainsCmd.AddCommand(orgRootDomainsListCmd)

}
