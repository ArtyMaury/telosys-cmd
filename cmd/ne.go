package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// neCmd represents the ne command
var neCmd = &cobra.Command{
	Use:     "ne",
	Aliases: []string{"newentity"},
	Short:   "Creates a new entity",
	Long:    "Creates a new entity",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			newEntity(arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(neCmd)
}

func newEntity(name string) {
	if err := config.ReadInConfig(); err == nil {
		if model := config.GetString(cfgModel); model != "" {
			newFile(model+"_model", name+".entity")
		} else {
			fmt.Println("No Model selected")
		}
	}
}
