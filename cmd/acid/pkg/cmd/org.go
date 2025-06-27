package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// orgCmd represents the add command
var orgCmd = &cobra.Command{
	Use:     "orgs",
	Aliases: []string{"org"},
	Short:   "org commands",
	Long: `todo
`,
	//Run: func(cmd *cobra.Command, args []string) {
	//
	//},
}

func init() {
	RootCmd.AddCommand(orgCmd)
}

func getRequiredOrgFlag(cmd *cobra.Command) string {
	orgId := viper.GetString("organization")
	if orgId == "" {
		log.Panic("must specify an organization in config or with --org")
	}

	return orgId
}
