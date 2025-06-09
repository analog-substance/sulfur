package cmd

import (
	"bufio"
	"github.com/analog-substance/sulfur/cmd/acid/pkg/sulfur"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add items to scope",
	Long: `Add items to scope unless it has been excluded via scopious exclude. For example:

	cat customer-supplied.txt | scopious add

	scopious add -i internal 10.0.0.0/22
`,
	Run: func(cmd *cobra.Command, args []string) {
		org, _ := cmd.Flags().GetString("org")

		if len(args) > 0 {
			//s := []string{}
			//for _, arg := range args {
			//	s = append(s, arg)
			//}
			sulfur.Add(org, args...)
		} else {
			// no args, lets read from stdin
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				scopeLine := scanner.Text()
				sulfur.Add(org, scopeLine)
			}

			if scanner.Err() != nil {
				log.Printf("STDIN scanner encountered an error: %s", scanner.Err())
			}
		}
	},
}

func init() {

	RootCmd.AddCommand(AddCmd)

}
