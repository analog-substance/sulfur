package cmd

import (
	"bufio"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// orgRootDomainsImportCmd represents the add command
var orgRootDomainsImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import root domains",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		domains := []string{}

		if len(args) > 0 {
			domains = append(domains, args...)
		} else {
			// no args, lets read from stdin
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				scopeLine := scanner.Text()
				domains = append(domains, scopeLine)
			}

			if scanner.Err() != nil {
				log.Printf("STDIN scanner encountered an error: %s", scanner.Err())
			}
		}

		rootDomains := []sulfur.OrgRootDomain{}
		for _, domain := range domains {
			rootDomains = append(rootDomains, sulfur.OrgRootDomain{
				Domain: domain,
			})
		}

		err := sulfurAPIClient.ImportOrgRootDomains(orgId, rootDomains)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {

	orgRootDomainsCmd.AddCommand(orgRootDomainsImportCmd)

}
