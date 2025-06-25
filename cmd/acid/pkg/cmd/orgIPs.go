package cmd

import (
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
		for _, d := range res.Items {
			fmt.Println(d.Expand.IpAddress.Address)
		}
	},
}

func init() {

	orgCmd.AddCommand(orgIPsCmd)

}
