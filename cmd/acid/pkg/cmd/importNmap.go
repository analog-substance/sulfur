package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// importNmapCmd represents the add command
var importNmapCmd = &cobra.Command{
	Use:   "nmap",
	Short: "import nmap xml",
	Long: `Import nmap xml files For example:

	acid import nmap nmap.xml
`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("nmap import")
	},
}

func init() {

	RootCmd.AddCommand(importNmapCmd)

}
