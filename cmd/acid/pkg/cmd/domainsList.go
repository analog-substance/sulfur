package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// domainsListCmd represents the add command
var domainsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list dns records",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		listResponse, err := sulfurAPIClient.ListRootDomains()
		if err != nil {
			log.Fatal(err)
		}

		if jsonOutput {
			outBytes, err := json.MarshalIndent(listResponse.Items, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(outBytes))
		} else {
			for _, org := range listResponse.Items {
				fmt.Println(org.Type, org.Name, org.Value)
			}
		}
	},
}

func init() {

	dnsCmd.AddCommand(dnsListCmd)
	dnsListCmd.Flags()

}
