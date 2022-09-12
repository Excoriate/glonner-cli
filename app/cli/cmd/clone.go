package cmd

import (
	"github.com/glonner/actions"
	"github.com/spf13/cobra"
)

var (
	dir          string
	dryRun       bool
	skipIfExist  bool // Gracefully skip if the repository already exists
	forceIfExist bool // Forcefully clone the repository, even if it already exists
	storageCheck bool // Check the size of the directory before cloning
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone all the repositories that belong to a given organization",
	Long: `
Clone all the repositories that belong to a given organization.
The repositories are cloned into the directory specified by the -dir flag.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := getLogger()

		globals := actions.GlobalRequiredArgs{Organization: organization, Token: token}

		specifics := actions.CloneActionArgs{
			Dir:          dir,
			SkipIfExists: skipIfExist,
			ForceIfExist: forceIfExist,
			StorageCheck: storageCheck,
			DryRun:       dryRun,
		}

		actions.CloneAction(globals, specifics, &logger)
	},
}

func init() {
	cloneCmd.Flags().StringVarP(&dir,
		"dir",
		"d", "",
		"Directory to store the repositories.")

	cloneCmd.Flags().BoolVarP(&skipIfExist,
		"skip",
		"s", false,
		"Skip a given repository to be cloned, if it exists in the destination directory.")

	cloneCmd.Flags().BoolVarP(&forceIfExist,
		"force",
		"f", false,
		"Forcefully clone a given repository, even if it already exists in the destination directory.")

	cloneCmd.Flags().BoolVarP(&storageCheck,
		"check",
		"c", false,
		"Check the size of the destination directory ("+
			"and the total storage in the machine) before cloning.")

	cloneCmd.Flags().BoolVarP(&dryRun,
		"dry-run",
		"n",
		false,
		"Only print the commands that would be executed, but don't execute them.")

	_ = cloneCmd.MarkFlagRequired("dir")
}
