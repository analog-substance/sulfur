package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// importDNSCmd represents the add command
var importDNSCmd = &cobra.Command{
	Use:   "import",
	Short: "import DNS records",
	Long: `For example:

	acid import dns --json dns.json
`,
	Run: func(cmd *cobra.Command, args []string) {
		recordName, _ := cmd.Flags().GetString("name")
		recordValue, _ := cmd.Flags().GetString("value")
		recordType, _ := cmd.Flags().GetString("type")
		recordTTL, _ := cmd.Flags().GetInt("ttl")
		importFile, _ := cmd.Flags().GetString("file")
		batchSize, _ := cmd.Flags().GetInt("batch-size")

		dnsRecords := []sulfur.DNSRecord{}

		if recordName != "" && recordValue != "" && recordType != "" && recordTTL > 0 {
			dnsRecords = append(dnsRecords, sulfur.DNSRecord{
				Name:  recordName,
				Value: recordValue,
				Type:  recordType,
				TTL:   recordTTL,
			})
		} else if importFile != "" {
			b, err := os.ReadFile(importFile)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(b, &dnsRecords)
			if err != nil {
				log.Fatal(err)
			}
		}

		dnsRecordLen := len(dnsRecords)
		if dnsRecordLen > 0 {

			batches := dnsRecordLen / batchSize

			fmt.Printf("Importing %d DNS records in %d batches\n", dnsRecordLen, batches)

			for len(dnsRecords) > batchSize {
				fmt.Printf("Sending batch %d\n", batches-(len(dnsRecords)/batchSize))
				batch := dnsRecords[:batchSize]
				dnsRecords = dnsRecords[batchSize:]

				err := sulfurAPIClient.ImportDNSRecords(batch)
				if err != nil {
					log.Println("error importing dns records", err)
				}
			}

			fmt.Printf("Sending batch %d\n", batches-(len(dnsRecords)/batchSize))
			err := sulfurAPIClient.ImportDNSRecords(dnsRecords)
			if err != nil {
				log.Println(err)
			}

		}
	},
}

func init() {

	dnsCmd.AddCommand(importDNSCmd)
	importDNSCmd.Flags().StringP("name", "n", "", "name of record")
	importDNSCmd.Flags().StringP("value", "v", "", "value of record")
	importDNSCmd.Flags().StringP("type", "t", "", "ttl of record")
	importDNSCmd.Flags().IntP("ttl", "l", 0, "ttl of record")
	importDNSCmd.Flags().IntP("batch-size", "s", 500, "batch size of import")
	importDNSCmd.Flags().StringP("file", "f", "", "File to read records from")

}
