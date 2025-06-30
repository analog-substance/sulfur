package cmd

import (
	"bufio"
	"github.com/analog-substance/sulfur/pkg/dns"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// importDomainCmd represents the add command
var importDomainCmd = &cobra.Command{
	Use:   "import",
	Short: "Resolve domains and import results",
	Long: `For example:

	acid import domain test.google.com
	as-crt-slurp.sh | acid import domain -s
`,
	Run: func(cmd *cobra.Command, args []string) {
		serverResolve, _ := cmd.Flags().GetBool("server-resolve")

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

		if serverResolve {
			err := sulfurAPIClient.ImportDomainAndResolve(domains)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			resolvedDNS := dns.ResolveDomains(domains)
			err := sulfurAPIClient.ImportDNSRecords(resolvedDNS)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	domainsCmd.AddCommand(importDomainCmd)
	importDomainCmd.Flags().BoolP("server-resolve", "s", false, "Send domains to server and resolve them there")
}
