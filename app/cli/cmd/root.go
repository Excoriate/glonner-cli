package cmd

import (
	logger "github.com/glonner/pkg/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	token        string
	organization string
)

func getLogger() logger.ILogger {
	return logger.NewLogger()
}

var rootCmd = &cobra.Command{
	Use:   "glonner",
	Short: "CLI for cloning repositories 'at scale'",
	Long: `
Use it for fetching repositories massively, particularly from a given organisation,
to quickly retrieve locally all the required repositories, using the native git clone command.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addChildCommands() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(cloneCmd)
}

func addRootFlags() {
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "",
		"GitHub private token. If not passed, it'll try to get it from the GITHUB_TOKEN environment variable.")

	rootCmd.PersistentFlags().StringVarP(&organization, "org", "g", "",
		"GitHub organization or owner to fetch required repositories from GITHUB")

	_ = rootCmd.MarkFlagRequired("token")
	_ = rootCmd.MarkFlagRequired("org")
}

func init() {
	addChildCommands()
	addRootFlags()
}
