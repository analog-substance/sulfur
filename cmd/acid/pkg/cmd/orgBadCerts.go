package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// orgBadCertsCmd represents the add command
var orgBadCertsCmd = &cobra.Command{
	Use:   "bad-certs",
	Short: "Sites with certs that do not belong to the organization",
	Long: `Sites with certs that do not belong to the organization
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		res, err := sulfurAPIClient.ListOrgBadCertificates(orgId)
		if err != nil {
			log.Fatal("error communicating with API", err)
		}

		if jsonOutput {
			outBytes, err := json.MarshalIndent(res, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(outBytes))
		} else {
			for _, d := range res {
				fmt.Println(d.Domain)
			}
		}

	},
}

func init() {

	orgCmd.AddCommand(orgBadCertsCmd)

}
