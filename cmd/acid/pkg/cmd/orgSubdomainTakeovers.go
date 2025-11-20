package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

// orgSubdomainTakeoversCmd represents the add command
var orgSubdomainTakeoversCmd = &cobra.Command{
	Use:   "takeovers",
	Short: "list takeovers",
	Long: `todo
`,
	Run: func(cmd *cobra.Command, args []string) {
		orgId := getRequiredOrgFlag(cmd)

		res, err := sulfurAPIClient.ListOrgSubDomainTakeovers(orgId)
		if err != nil {
			log.Fatal("error communicating with API", err)
		}

		if jsonOutput {
			outBytes, err := json.MarshalIndent(res, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(outBytes))
		} else {
			takeoverTable(res)
		}
	},
}

func init() {

	orgCmd.AddCommand(orgSubdomainTakeoversCmd)

}

func takeoverTable(takeovers []sulfur.SubdomainTakeover) {

	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)
	headers := []string{"Org", "Domain", "IP", "Last Resolved"}

	CapitalizeHeaders := func(data []string) []string {
		for i := range data {
			data[i] = strings.ToUpper(data[i])
		}
		return data
	}

	data := [][]string{}
	for _, takeover := range takeovers {

		data = append(data, []string{
			takeover.OrgName,
			takeover.Domain,
			takeover.IpAddress,
			takeover.LastResolved,
		})
	}

	ct := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers(CapitalizeHeaders(headers)...).
		Rows(data...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}

			even := row%2 == 0

			if even {
				return baseStyle.Foreground(lipgloss.Color("245"))
			}
			return baseStyle.Foreground(lipgloss.Color("252"))
		})
	fmt.Println(ct)

}
