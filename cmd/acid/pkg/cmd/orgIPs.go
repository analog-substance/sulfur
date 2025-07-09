package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// orgIPsCmd represents the add command
var orgIPsCmd = &cobra.Command{
	Use:   "ips",
	Short: "get ips",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		res, err := sulfurAPIClient.ListOrgIPAddresses(orgId)
		if err != nil {
			log.Fatal(err)
		}

		if jsonOutput {
			outBytes, err := json.MarshalIndent(res.Items, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(outBytes))
		} else {
			for _, d := range res.Items {
				fmt.Println(d.Expand.IpAddress.Address)
			}
		}

	},
}

func init() {

	orgCmd.AddCommand(orgIPsCmd)

}
