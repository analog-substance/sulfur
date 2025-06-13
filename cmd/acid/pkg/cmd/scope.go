package cmd

import (
	"github.com/spf13/cobra"
)

// scopeCmd represents the add command
var scopeCmd = &cobra.Command{
	Use:   "scope",
	Short: "scope commands",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	RootCmd.AddCommand(scopeCmd)

}
