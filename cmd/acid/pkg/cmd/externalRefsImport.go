package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// externalRefsImportCmd represents the add command
var externalRefsImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import external references",
	Long: `For example:

	acid external-refs import --json dns.json
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		value, _ := cmd.Flags().GetString("value")
		importFile, _ := cmd.Flags().GetString("file")
		batchSize, _ := cmd.Flags().GetInt("batch-size")

		importQueue := []sulfur.ExternalReference{}

		if name != "" && value != "" {
			importQueue = append(importQueue, sulfur.ExternalReference{
				Name:  name,
				Value: value,
			})
		} else if importFile != "" {
			b, err := os.ReadFile(importFile)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(b, &importQueue)
			if err != nil {
				log.Fatal(err)
			}
		}

		importQueueLen := len(importQueue)
		if importQueueLen > 0 {

			batches := importQueueLen / batchSize

			fmt.Printf("Importing %d xternal references in %d batches\n", importQueueLen, batches)

			for len(importQueue) > batchSize {
				fmt.Printf("Sending batch %d\n", batches-(len(importQueue)/batchSize))
				batch := importQueue[:batchSize]
				importQueue = importQueue[batchSize:]

				err := sulfurAPIClient.ImportExternalReferences(batch)
				if err != nil {
					log.Println("error importing external references", err)
				}
			}

			fmt.Printf("Sending batch %d\n", batches-(len(importQueue)/batchSize))
			err := sulfurAPIClient.ImportExternalReferences(importQueue)
			if err != nil {
				log.Println(err)
			}

		}
	},
}

func init() {

	externalRefsCmd.AddCommand(externalRefsImportCmd)
	externalRefsImportCmd.Flags().StringP("name", "n", "", "name of of external reference")
	externalRefsImportCmd.Flags().StringP("value", "v", "", "value of external reference")
	externalRefsImportCmd.Flags().IntP("batch-size", "s", 500, "batch size of import")
	externalRefsImportCmd.Flags().StringP("file", "f", "", "File to read records from")

}
