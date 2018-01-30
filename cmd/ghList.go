package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var githubAPI = "https://api.github.com"
var githubMasterZipPattern = "https://github.com/${USER}/${REPO}/archive/master.zip"

// ghCmd represents the gh command
var ghListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of bundles in the github repository",
	Long:  "Get the list of bundles in the github repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(getGithubRepoList())
		} else {
			getGithubRepoList()
		}
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
