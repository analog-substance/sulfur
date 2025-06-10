package cmd

import (
	"github.com/spf13/cobra"
)

// importCmd represents the add command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import data from external jobs",
	Long: `

`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	RootCmd.AddCommand(importCmd)

}
