package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// orgRootDomainsListCmd represents the add command
var orgRootDomainsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list root domains",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		domains, err := sulfurAPIClient.ListRootDomains()
		if err != nil {
			log.Fatal(err)
		}
		for _, domain := range domains.Items {
			fmt.Println(domain.Domain)
		}
	},
}

func init() {

	orgRootDomainsCmd.AddCommand(orgRootDomainsListCmd)

}
