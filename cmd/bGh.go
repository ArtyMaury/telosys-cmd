package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// bGhCmd represents the b gh command
var bGhCmd = &cobra.Command{
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
	bCmd.AddCommand(bGhCmd)
}

func getGithubRepo() string {
	return getConfValue(cfgGithub)
}

func setGithubRepo(repo string) {
	user := getGithubUser(repo)
	setConfValue(cfgGithub, user)
}

// Finds the github user in a generic url
func getGithubUser(repo string) string {
	reg := regexp.MustCompile(`((https?:\/\/)?(www\.)?github\.com\/)?([^\/]+)`)
	user := reg.FindStringSubmatch(repo)[4]
	return user
}

func getGithubURL(user string) string {
	url := "https://github.com/" + user + "/"
	return url
}
