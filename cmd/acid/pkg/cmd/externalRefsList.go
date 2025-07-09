package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// externalRefsListCmd represents the add command
var externalRefsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list external reference records",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		listResponse, err := sulfurAPIClient.ListExternalReferences()
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
	externalRefsCmd.AddCommand(externalRefsListCmd)
	//externalRefsListCmd.Flags()
}
