package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
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

		if jsonOutput {
			outBytes, err := json.MarshalIndent(domains.Items, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(outBytes))
		} else {
			for _, d := range domains.Items {
				fmt.Println(d.Domain)
			}
		}

	},
}

func init() {

	orgRootDomainsCmd.AddCommand(orgRootDomainsListCmd)

}
