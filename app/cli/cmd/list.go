package cmd

import (
	"github.com/glonner/actions"
	"github.com/spf13/cobra"
)

var (
	output string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the repositories that belong to a given organisation",
	Long: `
Only list, do not clone, all the repositories of a given GitHub organisation.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := getLogger()

		globals := actions.GlobalRequiredArgs{
			Organization: organization,
			Token:        token,
		}

		specifics := actions.ListActionArgs{} // This command, till now, does not require any specific arguments.

		actions.ListAction(globals, specifics, &logger)
	},
}

func init() {
	listCmd.Flags().StringVarP(&output,
		"output",
		"o", "table",
		"Show the results as a table, CSV or JSON format. Default is table.")
}
