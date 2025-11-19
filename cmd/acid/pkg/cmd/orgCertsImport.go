package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/spf13/cobra"
)

// orgCertificatesImportCmd represents the add command
var orgCertificatesImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import org certificates",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)
		batchSize, _ := cmd.Flags().GetInt("batch-size")
		importFile, _ := cmd.Flags().GetString("file")

		importCertificates := []sulfur.OrgCertificateImport{}

		if importFile != "" {
			b, err := os.ReadFile(importFile)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(b, &importCertificates)
			if err != nil {
				log.Fatal(err)
			}
		}

		importRecordLen := len(importCertificates)
		if importRecordLen > 0 {
			batches := importRecordLen / batchSize
			fmt.Printf("Importing %d Org Certificates %d batches\n", importRecordLen, batches)

			for len(importCertificates) > batchSize {
				fmt.Printf("Sending batch %d\n", batches-(len(importCertificates)/batchSize))
				batch := importCertificates[:batchSize]
				importCertificates = importCertificates[batchSize:]

				err := sulfurAPIClient.ImportOrgCertificates(orgId, batch)
				if err != nil {
					log.Println("error importing org certs", err)
				}
			}

			fmt.Printf("Sending batch %d\n", batches-(len(importCertificates)/batchSize))
			err := sulfurAPIClient.ImportOrgCertificates(orgId, importCertificates)
			if err != nil {
				log.Println(err)
			}
		}

	},
}

func init() {

	orgCertificatesCmd.AddCommand(orgCertificatesImportCmd)
	orgCertificatesImportCmd.Flags().IntP("batch-size", "s", 500, "batch size of import")
	orgCertificatesImportCmd.Flags().StringP("file", "f", "", "File to data from")

}
