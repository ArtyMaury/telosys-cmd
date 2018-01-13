package cmd

import (
	"github.com/spf13/cobra"
)

// nmCmd represents the nm command
var nmCmd = &cobra.Command{
	Use:     "nm",
	Aliases: []string{"newmodel"},
	Short:   "Add a new model",
	Long:    "Create a new model in the telosys project",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			newModel(arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(nmCmd)
}

func newModel(name string) {
	modelDir := name + "_model"
	modelFile := name + ".model"
	newFile(modelFile)
	newDir(modelDir)

	setModel(name)
}
