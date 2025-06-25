package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// orgSubdomainTakeoversCmd represents the add command
var orgSubdomainTakeoversCmd = &cobra.Command{
	Use:   "takeovers",
	Short: "list takeovers",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		res, err := sulfurAPIClient.ListOrgSubDomainTakeovers(orgId)
		if err != nil {
			log.Fatal("error communicating with API", err)
		}

		for _, takeover := range res {
			log.Println(takeover)
		}

	},
}

func init() {

	orgCmd.AddCommand(orgSubdomainTakeoversCmd)

}
