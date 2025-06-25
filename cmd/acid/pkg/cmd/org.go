package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// orgCmd represents the add command
var orgCmd = &cobra.Command{
	Use:     "orgs",
	Aliases: []string{"org"},
	Short:   "org commands",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	orgCmd.PersistentFlags().StringP("org-id", "o", "", "ID of the organization")

	RootCmd.AddCommand(orgCmd)

}

func getRequiredOrgFlag(cmd *cobra.Command) string {
	orgId, _ := cmd.Flags().GetString("org-id")
	if orgId == "" {
		log.Panic("must specify --org-id")
	}

	return orgId
}
