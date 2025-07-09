package cmd

import (
	"github.com/spf13/cobra"
)

// externalRefsCmd represents the add command
var externalRefsCmd = &cobra.Command{
	Use:   "external-refs",
	Short: "external refs can be used to associate items with external systems",

	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		externalRefsListCmd.Run(cmd, args)
	},
}

func init() {

	RootCmd.AddCommand(externalRefsCmd)

}
