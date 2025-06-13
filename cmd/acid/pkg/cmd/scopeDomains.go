package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// scopeDomainsCmd represents the add command
var scopeDomainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "domains records",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		listResponse, err := sulfurAPIClient.ListDNSRecords()
		if err != nil {
			log.Fatal(err)
		}

		for _, org := range listResponse.Items {
			fmt.Println(org.Name, org.Value)
		}
	},
}

func init() {

	dnsCmd.AddCommand(dnsListCmd)

}
