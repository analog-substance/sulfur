package cmd

import (
	"github.com/spf13/cobra"
)

// domainsCmd represents the add command
var domainsCmd = &cobra.Command{
	Use:     "domains",
	Aliases: []string{"domain"},
	Short:   "import data from external jobs",
	Long: `

`,
	Run: func(cmd *cobra.Command, args []string) {
		domainsListCmd.Run(cmd, args)
	},
}

func init() {

	RootCmd.AddCommand(domainsCmd)

}
