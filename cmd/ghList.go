package cmd

import (
	"github.com/spf13/cobra"
)

var githubAPI = "https://api.github.com"

// ghCmd represents the gh command
var ghListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of bundles in the github repository",
	Long:  "Get the list of bundles in the github repository",
	Run: func(cmd *cobra.Command, args []string) {
		printList(getGithubRepoList())
	},
}

func init() {
	ghCmd.AddCommand(ghListCmd)
}

func getGithubRepoList() []string {
	url := githubAPI + "/users/" + getGithubUser(getConfValue(cfgGithub)) + "/repos"
	maps := getHttpJsonValues(url, "name")
	var listRepo []string
	for _, repoMap := range maps {
		listRepo = append(listRepo, repoMap["name"].(string))
	}
	return listRepo
}
