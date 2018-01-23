package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// ghCmd represents the gh command
var ghCmd = &cobra.Command{
	Use:   "gh",
	Short: "Get or set the github repository for bundles",
	Long:  "Get or set the github repository for bundles",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Le repo utilisé est \"" + getGithubRepo() + "\"")
		} else {
			setGithubRepo(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(ghCmd)
}

func getGithubRepo() string {
	err := config.ReadInConfig()
	if err == nil {
		return config.GetString(cfgGithub)
	}
	return ""
}

func setGithubRepo(repo string) {
	reg := regexp.MustCompile(`((https?:\/\/)?(www\.)?github\.com\/)?([\w-]+)`)
	cleanRepo := reg.FindStringSubmatch(repo)[4]
	fmt.Println(cleanRepo)
}
