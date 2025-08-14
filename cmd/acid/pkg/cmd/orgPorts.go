package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// orgPortsCmd represents the add command
var orgPortsCmd = &cobra.Command{
	Use:   "ports",
	Short: "ports",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)
		getDomains, _ := cmd.Flags().GetBool("domains")
		getIPs, _ := cmd.Flags().GetBool("ips")

		if !getDomains && !getIPs {
			getDomains = true
		}

		if getDomains {

			res, err := sulfurAPIClient.ListOrgDomainPorts(orgId)
			if err != nil {
				log.Fatal(err)
			}
			for _, d := range res {
				fmt.Println(d.Domain, d.Port)
			}
		}

		if getIPs {
			res, err := sulfurAPIClient.ListOrgIPPorts(orgId)
			if err != nil {
				log.Fatal(err)
			}
			for _, d := range res {
				fmt.Println(d.Address, d.Port)
			}
		}
	},
}

func init() {

	orgCmd.AddCommand(orgPortsCmd)
	orgPortsCmd.Flags().BoolP("domains", "d", false, "Get domain ports")
	orgPortsCmd.Flags().BoolP("ips", "i", false, "Get ip ports")

}
