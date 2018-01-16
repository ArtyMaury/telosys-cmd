package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var homeDir string
var configFile = ".telosys-cfg.yaml"
var config = viper.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tcmd",
	Short: "Telosys Cli in Go",
	Long:  "Telosys Cli in Go",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&homeDir, "home", ".", "home folder (default is .)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	homeDir = toAbs(homeDir)
	config.SetConfigFile(toPath(configFile))

	config.AutomaticEnv() // read in environment variables that match

}
