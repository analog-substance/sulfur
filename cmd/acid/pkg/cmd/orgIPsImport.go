package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/analog-substance/sulfur/pkg/sulfur"
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
		ipAddress, _ := cmd.Flags().GetString("ip")
		extRef, _ := cmd.Flags().GetString("external-ref")
		batchSize, _ := cmd.Flags().GetInt("batch-size")
		importFile, _ := cmd.Flags().GetString("file")

		importIPAddresses := []sulfur.OrgIPAddressImport{}

		if ipAddress != "" {
			ip := sulfur.OrgIPAddressImport{
				IpAddress: ipAddress,
			}
			if extRef != "" {
				ip.ExternalReference = extRef
			}
			importIPAddresses = append(importIPAddresses, ip)
		} else if importFile != "" {
			b, err := os.ReadFile(importFile)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(b, &importIPAddresses)
			if err != nil {
				log.Fatal(err)
			}
		}

		importRecordLen := len(importIPAddresses)
		if importRecordLen > 0 {
			batches := importRecordLen / batchSize
			fmt.Printf("Importing %d Org IP addresses in %d batches\n", importRecordLen, batches)

			for len(importIPAddresses) > batchSize {
				fmt.Printf("Sending batch %d\n", batches-(len(importIPAddresses)/batchSize))
				batch := importIPAddresses[:batchSize]
				importIPAddresses = importIPAddresses[batchSize:]

				err := sulfurAPIClient.ImportOrgIPAddresses(orgId, batch)
				if err != nil {
					log.Println("error importing dns records", err)
				}
			}

			fmt.Printf("Sending batch %d\n", batches-(len(importIPAddresses)/batchSize))
			err := sulfurAPIClient.ImportOrgIPAddresses(orgId, importIPAddresses)
			if err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {

	orgIPsCmd.AddCommand(orgIPsImportCmd)

	orgIPsImportCmd.Flags().StringP("ip", "i", "", "ip address")
	orgIPsImportCmd.Flags().StringP("external-ref", "r", "", "external reference id")
	orgIPsImportCmd.Flags().IntP("batch-size", "s", 500, "batch size of import")
	orgIPsImportCmd.Flags().StringP("file", "f", "", "File to data from")

}
