package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the project",
	Long:  "Initiates the telosys project",
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
		config.Set(cfgHomedir, toRelPath(homeDir))
		config.Set(cfgModel, nil)
		config.Set(cfgBundle, nil)
		config.Set(cfgGithub, "https://github.com/telosys-templates-v3/")
		config.WriteConfig()
		fmt.Println("Config file successfully created:", config.ConfigFileUsed())
	}
}

func initProjectFiles() {
	newFile("initFilesOk")
	// newDir("downloads")
	// newDir("lib")
	// newDir("templates")
	// newFile("databases.dbcfg")
	// newFile("telosys-tools.cfg")
}
