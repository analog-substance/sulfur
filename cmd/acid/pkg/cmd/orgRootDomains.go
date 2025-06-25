package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// orgRootDomainsCmd represents the add command
var orgRootDomainsCmd = &cobra.Command{
	Use:   "root-domains",
	Short: "root domains",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		res, err := sulfurAPIClient.ListOrgDomains(orgId)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range res.Items {
			fmt.Println(d.Expand.RootDomain.Domain)
		}
	},
}

func init() {

	orgCmd.AddCommand(orgRootDomainsCmd)

}
