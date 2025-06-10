package cmd

import (
	"github.com/spf13/cobra"
)

// orgCmd represents the add command
var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "org commands",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	RootCmd.AddCommand(orgCmd)

}

