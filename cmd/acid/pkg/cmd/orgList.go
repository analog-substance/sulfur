package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// orgListCmd represents the add command
var orgListCmd = &cobra.Command{
	Use:   "list",
	Short: "list organizations",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgsListResponse, err := sulfurAPIClient.ListOrganizations()
		if err != nil {
			log.Fatal(err)
		}

		for _, org := range orgsListResponse.Items {
			fmt.Println(org.Name, org.Id)
		}
	},
}

func init() {

	orgCmd.AddCommand(orgListCmd)

}
