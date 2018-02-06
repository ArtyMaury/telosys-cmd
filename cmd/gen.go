package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates the templates",
	Long:  "Generates the templates",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if askConfirmation() == nil {
			generateFunction()
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func generateFunction() {
	fmt.Println("J'ai généré !")
}
