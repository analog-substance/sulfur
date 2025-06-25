package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// orgIPsImportCmd represents the add command
var orgIPsImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		ipAddresses := []string{}

		if len(args) > 0 {
			ipAddresses = append(ipAddresses, args...)
		} else {
			// no args, lets read from stdin
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				scopeLine := scanner.Text()
				ipAddresses = append(ipAddresses, scopeLine)
			}

			if scanner.Err() != nil {
				log.Printf("STDIN scanner encountered an error: %s", scanner.Err())
			}
		}

		err := sulfurAPIClient.ImportOrgIPAddresses(orgId, ipAddresses)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {

	orgIPsCmd.AddCommand(orgIPsImportCmd)

}
