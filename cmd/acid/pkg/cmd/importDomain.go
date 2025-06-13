package cmd

import (
	"bufio"
	"github.com/analog-substance/sulfur/cmd/acid/pkg/dns"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// importDomainCmd represents the add command
var importDomainCmd = &cobra.Command{
	Use:   "domains",
	Short: "Resolve domains and import results",
	Long: `For example:


	acid import domain test.google.com
	as-crt-slurp.sh | acid import domain 
`,
	Run: func(cmd *cobra.Command, args []string) {

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

		resolvedDNS := dns.ResolveDomains(domains)

		sulfurAPIClient.ImportDNSRecords(resolvedDNS)

	},
}

func init() {

	importCmd.AddCommand(importDomainCmd)

}
