package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// orgPortsCmd represents the add command
var orgPortsCmd = &cobra.Command{
	Use:   "ports",
	Short: "ports",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		res, err := sulfurAPIClient.ListOrgPorts(orgId)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range res {
			fmt.Println(d.Address, d.Port)
		}
	},
}

func init() {

	orgCmd.AddCommand(orgPortsCmd)

}
