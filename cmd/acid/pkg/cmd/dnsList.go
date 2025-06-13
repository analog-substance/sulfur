package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// dnsListCmd represents the add command
var dnsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list dns records",
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
