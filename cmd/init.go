package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfigFile()
		initProjectFiles()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initConfigFile() {
	if err := config.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", config.ConfigFileUsed())
	} else {
		config.Set("homeDir", ".")
		config.Set("model", "myModel")
		config.Set("bundle", "myBundle")
		config.WriteConfig()
		fmt.Println("Config file successfully created:", config.ConfigFileUsed())
	}
}

func initProjectFiles() {
	os.Create("projectFiles")
}
