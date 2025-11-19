package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// orgCertificatesCmd represents the add command
var orgCertificatesCmd = &cobra.Command{
	Use:   "certs",
	Short: "Certificates",
	Long: `get certificates for a given organization.
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		res, err := sulfurAPIClient.ListOrgCertificates(orgId)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range res.Items {
			fmt.Printf("%s\t%s\t%s\n", d.Serial, d.Issuer, d.Subject)
		}
	},
}

func init() {

	orgCmd.AddCommand(orgCertificatesCmd)

}
