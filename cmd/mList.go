package cmd

import (
	"github.com/spf13/cobra"
)

// mList represents the m list command
var mListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available models",
	Long:  "List the available models",
	Run: func(cmd *cobra.Command, args []string) {
		printList(availableModels)
	},
}

func init() {
	mCmd.AddCommand(mListCmd)
}
