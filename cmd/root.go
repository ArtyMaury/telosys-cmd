package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var homeDir string
var configFile = ".telosys-cfg.yaml"
// config will contain all the config variables
var config = viper.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tcmd",
	Short: "Telosys Cli in Go",
	Long:  "Telosys Cli in Go",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&homeDir, "home", ".", "home folder")
	rootCmd.PersistentFlags().BoolP("", "y", false, "Skip confirmation requests")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	homeDir = toAbsPath(homeDir)
	config.SetConfigFile(toPath(configFile))

	config.AutomaticEnv() // read in environment variables that match
}
